# State Machine & Domain Modeling

Knowledge map for turning business needs into explicit behavior, contracts, and enforceable invariants. Progress evidence belongs in [learning/STATE.md](../../../learning/STATE.md).

## Knowledge map

1. Actors, goals, scenarios, and domain vocabulary
2. Commands, events, inputs, outputs, and business contracts
3. Entity identity, lifecycle, legal states, and transitions
4. Invariants and enforcement boundaries: API, domain, database, or integration
5. Readiness, authorization, and failure semantics
6. Acceptance examples and traceability to verification

Prerequisite path: identify actors/outcomes → define scenarios/contracts → model lifecycle/invariants → assign enforcement → derive tests.

## Recommended primary resources

- [Microsoft: Domain analysis for microservices](https://learn.microsoft.com/en-us/azure/architecture/microservices/model/domain-analysis) — domain analysis and bounded contexts.
- [Martin Fowler: Domain-Driven Design](https://martinfowler.com/bliki/DomainDrivenDesign.html) — overview and pointers to core DDD concepts.

## Cross-topic practice

Job submission/readiness spans state transitions, conditional SQL updates, worker claiming, and reliability. Lessons may model a business flow without implementing it.

## Lessons, Labs, and references

No topic-specific artifacts yet. Cross-topic work should have one canonical home and be linked here rather than copied.

## Capabilities to verify

Consult [learning/STATE.md](../../../learning/STATE.md). Check translating ambiguous statements into acceptance criteria, identifying invalid states, assigning transition ownership, and separating evidence from product assumptions.
