# Learning State

## Current milestone

Build a reliable asynchronous job execution service. Current focus: Crash Recovery and Lease.

## Evidence-backed progress

### Implemented and observed

- Go HTTP API with JSON encoding/decoding.
- PostgreSQL access through `database/sql`.
- Handler, Biz, DAL, and Model layers.
- Job creation, lookup, and state transitions.
- Background workers using ticker, context, and goroutines.
- Atomic job claiming with PostgreSQL `FOR UPDATE SKIP LOCKED`.
- Three workers processed jobs with overlapping execution, verified through manual logs.

### Important limits

- `internal/worker/worker.go` still uses a roughly five-second `Sleep` placeholder; it is not a real business executor.
- Crash recovery, leases, retry, timeout, and idempotency are not implemented.
- PostgreSQL integration coverage, fault injection, and production readiness are incomplete.
- Prior code and logs do not establish the learner's mastery of Go syntax, Context, concurrency, or database semantics.

## Knowledge to reinforce

- Go syntax and control flow in the existing layers.
- Context cancellation and its limits when work is blocking.
- Goroutine lifecycle and graceful shutdown.
- PostgreSQL transaction boundaries, row locks, and `SKIP LOCKED` behavior.
- Crash windows and the guarantees a lease/retry design can honestly provide.

## Next session

**Suggested topic:** Crash recovery and lease design.

Before changing code, draw the current job state transitions and identify what happens if a worker exits after claiming a job. Discuss lease ownership, expiry, renewal, and races; compare a minimal recovery approach with alternatives. The learner should then implement one agreed increment independently.

**Acceptance evidence:** a written state/failure model; learner can explain the chosen lease semantics and at least one race/failure case; a focused implementation and repeatable PostgreSQL test or manual failure experiment. Record implementation, verification, and understanding separately.
