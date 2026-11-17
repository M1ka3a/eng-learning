# Concurrency & Reliability

Knowledge map for coordinating work safely under cancellation, contention, crashes, and partial failure. Capability evidence is in [learning/STATE.md](../../../learning/STATE.md).

## Knowledge map

1. Goroutines, shared state, ownership, and synchronization
2. Cancellation, deadlines, Context propagation, and shutdown
3. Work queues, worker pools, backpressure, and scheduling
4. Leases, stale owners, fencing, and completion guards
5. Retries, idempotency, timeouts, and delivery guarantees
6. Failure injection, recovery, and operational limits

Prerequisite path: goroutine lifecycle → shared ownership → cancellation/bounded work → coordination and recovery for identified failure modes.

## Recommended primary resources

- [Go Concurrency Patterns: Context](https://go.dev/blog/context) — cancellation, deadlines, and propagation.
- [Go context package](https://pkg.go.dev/context) — current API reference.
- [PostgreSQL SELECT](https://www.postgresql.org/docs/current/sql-select.html) — `SKIP LOCKED` semantics for queue-like use.
- [Google SRE Book: Addressing Cascading Failures](https://sre.google/sre-book/addressing-cascading-failures/) — overload and retry amplification.

## Cross-topic lessons

- [Database Engineering · 0001 · Job claim concurrency](../database-engineering/lessons/0001-postgres-lease-claim-concurrency.html) — row locks and Lease ownership; canonical lesson is stored in Database Engineering.

## Labs and references

Create topic-local material only when a concrete learning need exists. For the shared Lease quick reference, use [PostgreSQL claim concurrency](../database-engineering/reference/postgres-claim-concurrency.html).

## Capabilities to verify

Check [learning/STATE.md](../../../learning/STATE.md): Context vs transaction, bounded graceful shutdown, Lease takeover and stale completion, and actual test guarantees remain evidence-led questions.
