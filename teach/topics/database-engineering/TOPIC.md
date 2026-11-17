# Database Engineering

Knowledge map for reasoning about relational data behavior and performance. Project progress is tracked only in [learning/STATE.md](../../../learning/STATE.md).

## Knowledge map

1. Relational modeling, constraints, keys, and SQL semantics
2. Transactions, isolation, atomicity, and transaction boundaries
3. MVCC, snapshots, visibility, and anomalies
4. Row/table locks, contention, deadlocks, and queue claims
5. Index design, query plans, selectivity, and measurement
6. Schema change safety, migrations, and operational rollout

Prerequisite path: SQL/modeling → transactions/MVCC → locking/indexes → workload-specific performance and migration choices.

## Cross-topic practice

Job Lease connects transaction and lock semantics to state-machine invariants, worker concurrency, recovery, and integration-test evidence. Project paths are at the repository root. Code changes are optional when reading, query reasoning, or a controlled experiment better meets the goal.

## Recommended primary resources

- [PostgreSQL: MVCC Introduction](https://www.postgresql.org/docs/current/mvcc-intro.html) — snapshots and concurrent access.
- [PostgreSQL: Transaction Isolation](https://www.postgresql.org/docs/current/transaction-iso.html) — isolation levels and anomalies.
- [PostgreSQL: Explicit Locking](https://www.postgresql.org/docs/current/explicit-locking.html) — row locks, deadlocks, and lock behavior.
- [PostgreSQL: EXPLAIN](https://www.postgresql.org/docs/current/using-explain.html) — plans, estimates, and performance investigation.
- [PostgreSQL: SELECT](https://www.postgresql.org/docs/current/sql-select.html) — locking clauses including `SKIP LOCKED`.

## Lessons

- [0001 · Why two workers do not claim the same Job](lessons/0001-postgres-lease-claim-concurrency.html) — snapshot visibility, row locking, and Lease guards; also relevant to [Concurrency & Reliability](../concurrency-reliability/TOPIC.md).

## Labs

No formal Lab has been created yet. Add one here when a scoped, valuable project or analysis task is designed.

## References

- [PostgreSQL claim concurrency](reference/postgres-claim-concurrency.html) — MVCC, locks, `SKIP LOCKED`, and Lease-token responsibilities.

## Capabilities to verify

Use [project evidence](../../../learning/STATE.md), not this map, to determine current capability. Useful checks include predicting transaction outcomes, explaining a lock interaction, selecting an index from query/workload evidence, and distinguishing an isolated test schema from production migration compatibility.
