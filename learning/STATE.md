# Project Technical State

Technical state of the repository only. Capability levels live in `SKILLS.md`; history lives in `SESSIONS.md`. (Pre-redesign narrative state: commit `b00ce9a`.)

## What exists

- Go HTTP API (Handler → Biz → DAL → Model) on PostgreSQL via `database/sql` + pgx.
- Job state machine: `needs_input` → `pending` → `running` → `succeeded` | `failed`.
- Ticket flow: create with UID, description, user-selected category, and optional order ID; `PATCH /jobs/{id}/input` supplies the order ID; the worker routes by category and persists the result. List by UID.
- Worker pool: ticker-based; claims with `FOR UPDATE SKIP LOCKED`; lease token and expiry guard `Finish` and `ExtendLease`.
- Most ticket and lease code was coach-authored. The ctx chain for `GetJobByID` and `AddOrderID` was written by the learner with coaching.

## Verified

- 2026-11-16: `go build`, `go vet`, and `go test ./... -count=1` pass against local DB `backend_lab_test` (`LEASE_TEST_DATABASE_URL='host=/tmp dbname=backend_lab_test sslmode=disable'`). This includes the DAL lease integration tests (claim/renew/takeover/finish, expired-lease rejection, concurrent claim) and the Biz Create→Claim→Route→Finish test.
- Tests create their own tables in a random schema. Compatibility with the real `backend_lab` schema and the migration is **not** verified.
- `ClaimNext(60)` bug (60ns lease) fixed with `claimLeaseDuration = 60 * time.Second`; the Biz test reproduced it before the fix.
- `internal/dal/context_experiment_test.go` shows that a deadline cancels `pg_sleep(2)` at about 200ms. It adds about 2s to the test run; keep or delete it.

## Known issues (from the 11-16 review; see ROADMAP A4/A5)

- ctx is missing from the remaining DAL/Biz methods and from worker calls.
- `main.go` starts 4 workers, and only one is cancellable. There is no signal handling, no `http.Server` timeouts, no graceful shutdown, and the DSN is hard-coded.
- Category handler returns 400 for a DB error and 502 for a validation error, and leaks `err.Error()`.
- Handlers and worker depend on `sql.ErrNoRows`. DAL errors carry no context, and lease rejection has no sentinel error.
- No index for the claim query or `ctx->>'uid'`. No base-schema migration for `jobs`/`categories`. No attempts limit (poison jobs).
- `go.mod` needs `go mod tidy` (direct deps are marked indirect).
- `ExtendLease` is not wired into the worker.
- UID matching is not authentication.
