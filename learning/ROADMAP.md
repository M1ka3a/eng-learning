# Engineering Learning Roadmap

## Long-term capability goals

Build the ability to independently take a software idea from requirements through design, implementation, testing, deployment, operation, and iteration. Develop sound judgment about failure modes, tradeoffs, and technology choices, then apply those skills to AI-enabled systems.

The learner brings substantial QA/SDET experience, Go/Python development, automation, and AI/MCP tool exposure. The development focus is independent backend delivery, system design, operations, and deeper understanding of concurrency and data semantics.

## Phases and acceptance criteria

### 1. Build a real Job Executor with an AI Engineering path

Topics: choose a meaningful job type; define its input/output contract and execution lifecycle; implement a real executor in place of the Sleep placeholder; decide how results and failures are represented; evolve from deterministic processing toward a bounded AI capability.

Acceptance: defend the selected job use case and its contract; execute a real job end to end through the existing API and worker; persist or return a useful result; demonstrate success and failure behavior; explain the architecture and remaining limits.

### 2. Reliable asynchronous execution and recovery

Topics: crash recovery, leases and lease expiry, retry policy, idempotency, timeout/cancellation, graceful shutdown, PostgreSQL transaction and locking semantics. Prioritize these when the real executor creates a concrete need; do not expand scheduler mechanics without clear learning value.

Acceptance: explain the job state machine and failure windows; implement only the recovery behavior required by the workload; demonstrate key guarantees with PostgreSQL integration tests and targeted failure scenarios; explain what guarantees are and are not provided.

### 3. Testing and production operation

Topics: unit versus integration tests, fault injection, configuration and secrets, deployment, migrations, logs/metrics/traces, health checks, operational runbooks.

Acceptance: reproduce key correctness claims in repeatable tests; deploy the service in a documented environment; diagnose a seeded failure from observable evidence; explain operational tradeoffs and remaining risks.

### 4. Distributed systems and system design

Topics: delivery guarantees, consistency, contention, backpressure, scaling, coordination, data ownership, capacity and failure analysis.

Acceptance: produce and defend a design for a stated workload; identify bottlenecks and failure modes; compare alternatives with explicit requirements and costs; revise the design when assumptions change.

### 5. AI engineering foundations

Topics: LLM APIs, structured output, tool calling, prompt and context design, safety boundaries, evaluation, latency/cost/reliability.

Acceptance: build a bounded AI capability with explicit inputs, outputs, failure handling, and evaluation cases; measure quality and operational behavior; explain where deterministic code is preferable.

### 6. MCP, agents, and AI task execution platform

Topics: MCP, tool contracts, agent loops, persistence, permissions, observability, evaluations, human approval, orchestration.

Acceptance: evolve the job platform through documented increments into an AI task execution platform; demonstrate tool execution and recovery under failures; evaluate behavior against repeatable scenarios; defend the system boundaries and tradeoffs.

## Working method

Use real project work in 90–120 minute deep-work sessions. Discuss the problem and design before implementation. The learner implements; the coach provides questions, hints, and review. Record implementation, verification, and understanding separately. Feature count is not a measure of learning progress.

Plan sessions by expected learning value and relevance to the long-term Software Engineering and AI Engineering goals, not only by roadmap order. Prefer real product behavior over continued development of generic scheduler features when the scheduler already supports the learning objective.
