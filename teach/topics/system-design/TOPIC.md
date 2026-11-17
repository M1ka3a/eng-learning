# System Design

Knowledge map for choosing system boundaries and evaluating whole-system tradeoffs. Project milestones are tracked in [learning/STATE.md](../../../learning/STATE.md).

## Knowledge map

1. Workload, requirements, constraints, and quality attributes
2. Components, ownership boundaries, APIs, and data flow
3. Synchronous calls, queues, caches, and persistence choices
4. Scaling, consistency, contention, backpressure, and capacity
5. Failure modes, redundancy, recovery, and operations
6. Architecture diagrams, decision records, and design review

Prerequisite path: define workload/quality attributes → map boundaries and flows → compare alternatives → analyze failures/capacity → validate with evidence.

## Recommended primary resources

- [Google SRE Book: Addressing Cascading Failures](https://sre.google/sre-book/addressing-cascading-failures/) — overload, queues, retries, cascading failures.
- [AWS Well-Architected Framework](https://docs.aws.amazon.com/wellarchitected/latest/framework/welcome.html) — review prompts across architecture qualities; use as a checklist, not a universal recipe.

## Cross-topic practice

Evaluate the job platform as a queue-backed executor: whether async work is justified, what delivery guarantee is required, and where persistence/worker ownership belong. Architecture analysis outside this repository is also valid.

## Lessons, Labs, and references

No topic-local artifacts yet. Link cross-topic lessons with one canonical path when created.

## Capabilities to verify

Use [learning/STATE.md](../../../learning/STATE.md) and decisions the learner explained. Check assumptions, alternatives, failure propagation, and whether a queue/cache/service boundary earns its complexity.
