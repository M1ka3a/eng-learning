package dal

import (
	"context"
	"errors"
	"testing"
	"time"
)

// Experiment: does cancelling a context actually stop a running SQL query?
func TestContextCancelsSlowQuery(t *testing.T) {
	store := leaseTestStore(t)

	// A: with context, deadline 200ms.
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	start := time.Now()
	var x any
	err := store.db.QueryRowContext(ctx, "SELECT pg_sleep(2)").Scan(&x)
	t.Logf("A (QueryRowContext, 200ms deadline): elapsed=%v err=%v", time.Since(start), err)

	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("A: want context.DeadlineExceeded, got %v", err)
	}

	// B: without context. Nothing can stop it, so it runs the full 2 seconds.
	start = time.Now()
	err = store.db.QueryRow("SELECT pg_sleep(2)").Scan(&x)
	t.Logf("B (QueryRow, no context): elapsed=%v err=%v", time.Since(start), err)
}
