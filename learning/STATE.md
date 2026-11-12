# Learning State

## Current milestone

Build a real Job Executor that replaces the Sleep placeholder and can evolve toward AI Engineering. Current work is at the use-case and contract design stage; the job type and result contract are not yet selected.

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

### Important limits

- `internal/worker/worker.go` still uses a roughly five-second `Sleep` placeholder; implementing a real executor is now the next priority.
- Lease integration tests have not been run against PostgreSQL. The database environment was unavailable/inconsistent, and per the learner's boundary no service or database changes were made.
- `go test ./...` passed with `LEASE_TEST_DATABASE_URL` explicitly unset; this checked compilation and non-database tests only. PostgreSQL SQL/locking/time semantics remain unverified.
- Lease renewal exists in the DAL but is not wired into the worker execution loop. Retry, idempotency, and timeout remain on the roadmap and are deferred until a real job creates a concrete need. Graceful shutdown, fault injection, and production readiness remain incomplete.
- Prior code and logs do not establish the learner's mastery of Go syntax, Context, concurrency, or database semantics.

## Knowledge to reinforce

- Go syntax and control flow in the existing layers.
- Context cancellation and its limits when work is blocking; goroutine lifecycle and graceful shutdown.
- PostgreSQL transaction boundaries, row locks, and `SKIP LOCKED`; lease correctness has test code but no live database evidence yet.
- Translating a real business use case into a clear job input/output contract and execution lifecycle.
- Go execution design, Context cancellation, and the boundary between Worker orchestration and business execution.
- PostgreSQL transaction boundaries, row locks, and `SKIP LOCKED`; Lease database behavior remains unverified.

## Next session

**Suggested topic:** Design and implement the first real Job Executor. First select a useful deterministic job that can later gain an AI step; define its versioned input, output, errors, and execution flow. Review whether the current Handler/Biz/DAL/Worker boundaries fit before making changes. Keep the Lease PostgreSQL integration test as an explicit unverified item, and revisit retry/idempotency/timeout only when required by the selected job.
