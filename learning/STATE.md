# Learning State

## Current milestone

Primary learning goal: translate business requirements into a small, testable software design. The learner challenged the coach-authored keyword classifier as an unsupported product assumption. The current ticket example uses a user-selected category and deterministic routing; it is a small routing prototype, not yet a compelling real asynchronous executor. Baseline commit `db07412` preserves the pre-example worktree. Independent learner implementation and understanding remain unestablished.

## Evidence-backed progress

### Implemented and observed

- Go HTTP API with JSON encoding/decoding.
- PostgreSQL access through `database/sql`.
- Handler, Biz, DAL, and Model layers.
- Job creation, lookup, and state transitions.
- Background workers using ticker, context, and goroutines.
- Atomic job claiming with PostgreSQL `FOR UPDATE SKIP LOCKED`.
- Three workers processed jobs with overlapping execution, verified through manual logs.
- Lease implementation in `internal/dal/job_store.go`: claims pending or expired jobs, assigns a fresh token and expiry, gates renewal and completion on running status, matching token, and unexpired lease.
- Lease test code: PostgreSQL integration scenarios use an isolated random schema and now require database name `backend_lab_test`; a concurrent claim scenario and database row assertions are included.
- DAL unit tests use a fake `database/sql` driver to inspect SQL predicates and affected-row handling without PostgreSQL.
- Learner has practiced the job state machine, worker pool, lease, and basic failure-recovery design. This records practice, not full mastery.
- Learner independently distinguished accepting an incomplete ticket from making it eligible for worker claiming; this is evidence of reasoning about input completeness and execution readiness.
- Learner identified requirements-to-design translation as a current skill gap and chose to practice it explicitly before implementation.
- Ticket progress: `TicketDetails` (UID, order ID, description); create selects `needs_input` or `pending`; `PATCH /jobs/{id}/input` calls Biz `AddOrderID`, which compares the stored UID and updates the order ID through a conditional DAL update that also moves the job to `pending`.
- Learner independently noticed that clients need a way to recover job IDs and that `TicketInput` was misleading for data read from stored `ctx`.
- Coach-authored example includes create validation, user-selected ticket category (`payment_issue`, `order_issue`, `other`), deterministic category-to-queue routing in Biz, Worker execution, atomic lease-guarded result/status persistence, GET result, and UID-based list endpoint. A review-only SQL migration file adds result/error columns and a status check.
- Learner independently identified the phrase-based classification as a poor product assumption and proposed human or user selection instead. This is evidence of critical requirements review.

### Important limits

- Worker no longer uses the Sleep placeholder in the coach-authored example. The current category-to-queue mapping is simple enough that its need for asynchronous execution is unresolved; a genuinely useful Executor workload remains to be selected.
- Matching a request UID is not authentication. The UID-based list endpoint is therefore only a prototype until authorization is designed.
- The SQL migration has not been applied; the actual local jobs schema was not available in the repository and must be reviewed before use.
- `go build ./...` and `git diff --check` passed after the coach-authored example. No tests, server run, or PostgreSQL behavior for this workflow have been verified; `jsonb_set` requires `jobs.ctx` to be `jsonb`.
- 2026-11-16: with learner authorization, empty local database `backend_lab_test` was created; `go test ./... -count=1` with `LEASE_TEST_DATABASE_URL='host=/tmp dbname=backend_lab_test sslmode=disable'` passed, including the three DAL Lease integration tests (claim/renew/takeover/finish, expired lease rejection, concurrent `SKIP LOCKED` claim) and a new Biz Create→ClaimNext→RouteTicket→Finish test. These tests use their own `CREATE TABLE`, not the real `backend_lab` schema; migration compatibility and the worker/HTTP path remain unverified.
- Fixed bug: Biz called `ClaimNext(60)` (60ns as `time.Duration`), so every `Finish` was rejected and jobs were reclaimed forever. The new Biz integration test reproduced the failure before the fix (`claimLeaseDuration = 60 * time.Second`). DAL tests passed `time.Minute` and could not catch the caller's unit mistake.
- Lease renewal exists in the DAL but is not wired into the worker execution loop. Retry, idempotency, and timeout remain on the roadmap and are deferred until a real job creates a concrete need. Graceful shutdown, fault injection, and production readiness remain incomplete.
- Prior code and logs do not establish the learner's mastery of Go syntax, Context, concurrency, or database semantics.

### Context (2026-11-16, first exposure)

- **Implemented (learner, coached):** ctx added from `r.Context()` through Handler → Biz → DAL `QueryRowContext` for `GetJobByID` and `AddOrderID` (both handlers, both SQL calls); test call sites use `context.Background()`. Learner chose to give `AddOrderID` a real ctx parameter instead of `context.TODO()`. Remaining DAL/Biz methods, Worker calls, and `main.go` lifecycle still have no ctx.
- **Verified:** `go build`, `go vet`, and full `go test` against `backend_lab_test` pass. Experiment test `internal/dal/context_experiment_test.go` (coach-completed at learner request): `QueryRowContext` with a 200ms deadline stopped `pg_sleep(2)` at ~201ms with an error matching `errors.Is(err, context.DeadlineExceeded)`; plain `QueryRow` ran the full 2s. Scratchpad signal demo: without `signal.NotifyContext`, SIGTERM killed the process (exit 143, no defer/ctx.Done); with it, ctx was cancelled and cleanup ran (exit 0).
- **Understood (partial):** learner asked for context's purpose and correctly connected it to stopping work during shutdown; after explanation, independently generalized that in-flight work should usually be finished. The coach refined this to "stop when no one needs the result; finish valuable in-flight work with a bounded deadline". Learner could not answer the Ctrl+C or lease-takeover questions before explanation, and did not answer the `Finish` ctx choice (a/b/c) independently; the coach gave the answer (c). Learner says they are starting to relate context to transactions. The distinction (context = when to stop; transaction = consistent state after stopping) has been explained but not checked.
- Struggled with Go basics in the experiment: unexported `t.logF` vs `t.Logf`, a discarded `Scan` error, a test-name typo.

## Knowledge to reinforce

- Go syntax and control flow in the existing layers.
- Context cancellation and its limits when work is blocking; goroutine lifecycle and graceful shutdown.
- Translating a real business use case into a clear job input/output contract and execution lifecycle.
- Translating business statements into concrete scenarios, acceptance criteria, data contracts, state transitions, and layer responsibilities before coding.
- Go execution design, Context cancellation, and the boundary between Worker orchestration and business execution.
- PostgreSQL transaction boundaries, row locks, and `SKIP LOCKED`; Lease database behavior remains unverified.

## Next session

**Next session (Context continuation, recommended):** Learner first explains in their own words the three cancellation cases (caller gone, timeout, shutdown), the "does anyone still need the result" rule, and context vs transaction. Then they implement graceful shutdown in `main.go`: `signal.NotifyContext` with SIGINT+SIGTERM, a single shared worker ctx, no extra fourth worker, `http.Server.Shutdown`, a `WaitGroup`, and a `Finish` ctx decided by the learner. Uncommitted work: the ctx chain changes and the experiment test. Decide whether to keep the experiment test.

**Earlier pending direction:** Review the revised user-selected category contract and ask whether routing should be synchronous. Select a real asynchronous workload only if it has meaningful processing needs. Study the coach-authored example against the revised requirements before re-implementing a bounded part independently. Database migration and Lease behavior remain unverified; database checks require separate authorization and a safe test environment.
