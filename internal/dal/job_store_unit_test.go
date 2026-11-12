package dal

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"strings"
	"testing"

	"backend-lab-go/internal/model"
)

// fakeSQLDriver captures Exec calls so DAL predicates and result handling can
// be tested without starting PostgreSQL. It does not emulate PostgreSQL SQL.
type fakeSQLState struct {
	query    string
	args     []driver.NamedValue
	affected int64
}

type fakeConnector struct{ state *fakeSQLState }

func (c fakeConnector) Connect(context.Context) (driver.Conn, error) { return fakeConn{c.state}, nil }
func (c fakeConnector) Driver() driver.Driver                        { return fakeDriver{} }

type fakeDriver struct{}

func (fakeDriver) Open(string) (driver.Conn, error) { return nil, errors.New("use connector") }

type fakeConn struct{ state *fakeSQLState }

func (c fakeConn) Prepare(string) (driver.Stmt, error) { return nil, errors.New("not implemented") }
func (c fakeConn) Close() error                        { return nil }
func (c fakeConn) Begin() (driver.Tx, error)           { return nil, errors.New("not implemented") }
func (c fakeConn) ExecContext(_ context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	c.state.query = query
	c.state.args = append([]driver.NamedValue(nil), args...)
	return driver.RowsAffected(c.state.affected), nil
}

func newFakeStore(t *testing.T, affected int64) (*JobStore, *fakeSQLState) {
	t.Helper()
	state := &fakeSQLState{affected: affected}
	db := sql.OpenDB(fakeConnector{state})
	t.Cleanup(func() { db.Close() })
	return NewJobStore(db), state
}

func assertLeasePredicate(t *testing.T, state *fakeSQLState) {
	t.Helper()
	query := strings.ToLower(state.query)
	for _, alternatives := range [][]string{
		{"id = $1", "id = $2"},
		{"status = 'running'", "status = $3"},
		{"lease_token = $2", "lease_token = $4"},
		{"lease_expires_at > now()"},
	} {
		found := false
		for _, predicate := range alternatives {
			if strings.Contains(query, predicate) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("query missing one of %q: %s", alternatives, state.query)
		}
	}
}

func TestExtendLeaseUsesCurrentUnexpiredOwner(t *testing.T) {
	store, state := newFakeStore(t, 1)
	if err := store.ExtendLease(42, "worker-token"); err != nil {
		t.Fatal(err)
	}
	assertLeasePredicate(t, state)
	if len(state.args) != 2 || state.args[0].Value != int64(42) || state.args[1].Value != "worker-token" {
		t.Fatalf("unexpected SQL args: %#v", state.args)
	}
}

func TestExtendLeaseRejectsNoMatchingLease(t *testing.T) {
	store, _ := newFakeStore(t, 0)
	if err := store.ExtendLease(42, "stale-token"); err == nil {
		t.Fatal("expected stale lease to be rejected")
	}
}

func TestUpdateStatusUsesCurrentUnexpiredOwner(t *testing.T) {
	store, state := newFakeStore(t, 1)
	if err := store.UpdateStatus(42, model.JobRunning, model.JobSucceeded, "worker-token"); err != nil {
		t.Fatal(err)
	}
	assertLeasePredicate(t, state)
	if len(state.args) != 4 || state.args[0].Value != string(model.JobSucceeded) || state.args[1].Value != int64(42) || state.args[2].Value != string(model.JobRunning) || state.args[3].Value != "worker-token" {
		t.Fatalf("unexpected SQL args: %#v", state.args)
	}
}

func TestUpdateStatusRejectsNoMatchingLease(t *testing.T) {
	store, _ := newFakeStore(t, 0)
	if err := store.UpdateStatus(42, model.JobRunning, model.JobSucceeded, "stale-token"); err == nil {
		t.Fatal("expected stale lease completion to be rejected")
	}
}
