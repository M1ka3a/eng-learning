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
- Lease integration tests have not been run against PostgreSQL. The database environment was unavailable/inconsistent, and per the learner's boundary no service or database changes were made.
- Earlier `go test ./...` passed with `LEASE_TEST_DATABASE_URL` explicitly unset; this predates the ticket changes and checked compilation and non-database tests only. PostgreSQL SQL/locking/time semantics remain unverified.
- Lease renewal exists in the DAL but is not wired into the worker execution loop. Retry, idempotency, and timeout remain on the roadmap and are deferred until a real job creates a concrete need. Graceful shutdown, fault injection, and production readiness remain incomplete.
- Prior code and logs do not establish the learner's mastery of Go syntax, Context, concurrency, or database semantics.

## Knowledge to reinforce

- Go syntax and control flow in the existing layers.
- Context cancellation and its limits when work is blocking; goroutine lifecycle and graceful shutdown.
- Translating a real business use case into a clear job input/output contract and execution lifecycle.
- Translating business statements into concrete scenarios, acceptance criteria, data contracts, state transitions, and layer responsibilities before coding.
- Go execution design, Context cancellation, and the boundary between Worker orchestration and business execution.
- PostgreSQL transaction boundaries, row locks, and `SKIP LOCKED`; Lease database behavior remains unverified.

## Next session

**Next session:** Review the revised user-selected category contract and ask whether routing should be synchronous. Select a real asynchronous workload only if it has meaningful processing needs. Study the coach-authored example against the revised requirements before re-implementing a bounded part independently. Database migration and Lease behavior remain unverified; database checks require separate authorization and a safe test environment.
