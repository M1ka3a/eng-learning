# Learning State

## Current milestone

Primary learning goal: translate business requirements into a small, testable software design, using text-only customer support ticket triage as the case. A partial create/supplement workflow now exists; the actual triage execution and result contract are still missing. Processing remains in Biz for the first version.

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

### Important limits

- `internal/worker/worker.go` still uses a roughly five-second `Sleep` placeholder; implementing the text-ticket executor is now the next priority.
- No triage result schema, classification rules, persisted result, or user-facing result retrieval exists yet. `Create` does not validate UID/description as required by the proposed contract; matching a request UID is not authentication.
- The ticket create/supplement code compiled with `go build ./...` after the latest changes. No tests or database behavior for this workflow have been run; the `jsonb_set` update assumes `jobs.ctx` is `jsonb`.
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

**Next session:** Review and finalize a small set of ticket acceptance examples and the `TriageResult` contract, then implement one complete input → Biz classification → persisted result → GET path. Review the partial create/supplement code against the contract before adding more behavior. Do not treat `go build` as database verification; Lease PostgreSQL behavior also remains unverified. Retry/idempotency/timeout stay deferred until needed by real execution.
