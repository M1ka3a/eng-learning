# Topic Router

This index routes `$teach-se` to one topic workspace. Each topic has its own lesson numbering and `lessons/`, `labs/`, and `reference/` collections. Cross-topic work has one canonical home and is linked from every relevant topic; do not duplicate it.

Project progress and verified capability evidence live only in [learning/STATE.md](../learning/STATE.md); chronological session evidence lives in [learning/SESSIONS.md](../learning/SESSIONS.md). The [Roadmap](../learning/ROADMAP.md) is a source of options, not a fixed topic sequence.

## Topics

| Topic | Workspace | Focus |
|---|---|---|
| Database Engineering | [database-engineering](topics/database-engineering/TOPIC.md) | SQL, transactions, indexes, MVCC, locks, performance |
| State Machine & Domain Modeling | [state-machine-domain-modeling](topics/state-machine-domain-modeling/TOPIC.md) | Business contracts, states, invariants, domain boundaries |
| System Design | [system-design](topics/system-design/TOPIC.md) | Architecture, caching, queues, distributed systems, reliability |
| Software Implementation | [software-implementation](topics/software-implementation/TOPIC.md) | Go, interfaces, modularity, code quality, refactoring |
| Concurrency & Reliability | [concurrency-reliability](topics/concurrency-reliability/TOPIC.md) | Context, coordination, workers, leases, recovery |
| Testing & Delivery | [testing-delivery](topics/testing-delivery/TOPIC.md) | Test strategy, observability, CI/CD, quality and delivery evidence |

## Routing and shared material

Invoke `$teach-se` from the repository root or any subdirectory and name a topic, a learning goal, or a Lab to resume. The skill resolves the repository root and then uses the selected `TOPIC.md` and its local materials. Read [the shared mission](MISSION.md), [shared resources](RESOURCES.md), and [teaching preferences](NOTES.md) as relevant. A user may switch topics without migrating or resetting the prior Lab checkpoint.

Courses are numbered within each topic. A lesson may connect to multiple topics; list it once in its canonical topic and link it from the others. Shared assets live in `assets/`. Topic knowledge maps are navigation, not progress or mastery state.
