# Testing & Delivery

Knowledge map for building confidence from requirements through operation and release. Implementation and verification evidence belongs in [learning/SKILLS.md](../../../learning/SKILLS.md).

## Knowledge map

1. Acceptance evidence and testability
2. Unit, component, integration, contract, and end-to-end test boundaries
3. Deterministic tests, fixtures, isolation, and safe environments
4. Concurrency tests, fault injection, and failure assertions
5. Logs, metrics, traces, health checks, and diagnosis
6. CI/CD, migrations, deployment, rollback, and release confidence

Prerequisite path: define observable behavior → choose a test boundary → make evidence repeatable → connect runtime signals to delivery decisions.

## Recommended primary resources

- [Go: Add a test](https://go.dev/doc/tutorial/add-a-test) — official Go testing introduction.
- [Google SRE Book: Monitoring Distributed Systems](https://sre.google/sre-book/monitoring-distributed-systems/) — monitoring signals and actionable observability.
- [DORA capabilities](https://dora.dev/capabilities/) — research-based delivery capability catalog; use to frame questions, not as a recipe.

## Cross-topic practice

Lease tests illustrate why mocks cannot prove PostgreSQL locking semantics and why isolated test schemas matter. Production schema/migration verification is a separate claim.

## Lessons, Labs, and references

No topic-local artifacts yet. Link cross-topic material from its canonical location.

## Capabilities to verify

Consult [learning/SKILLS.md](../../../learning/SKILLS.md). Check that tests support the stated claim, failures are reproducible, integration environments are isolated, and release risks are visible.
