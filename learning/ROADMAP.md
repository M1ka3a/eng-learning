# Roadmap: project finish and interview prep

The project is a portfolio piece: **a reliable PostgreSQL-backed job queue in Go, plus the test system that proves its guarantees.** Finish it, freeze it, and be able to defend it. Do not add product features unless an interview story needs them.

## A. Project finish list

Each item counts only after the learner answers its "defend it" questions.

| # | Item | Acceptance | Defend-it focus |
|---|---|---|---|
| A1 | `docker-compose` (PG + service) and README | A fresh clone runs with one command; README has a diagram, guarantees (at-least-once, lease takeover), known limits, and how to test | Why a lease; what is and isn't guaranteed |
| A2 | Python pytest API suite (`tests/api/`) | Covers create/get/list/supplement order ID, validation errors, 404/409/403, and job reaching `succeeded`; fixtures manage data; runs against compose | Test design choices; how async completion is waited on without flakiness |
| A3 | GitHub Actions CI | Go unit tests, Go PG integration tests (service container), and pytest all green on push | Test pyramid; why each layer exists |
| A4 | Context everywhere + graceful shutdown | ctx flows through all DAL/Biz calls; `signal.NotifyContext` (SIGINT+SIGTERM); one shared worker ctx (fix the extra 4th worker); `http.Server` with timeouts + `Shutdown`; `WaitGroup`; bounded `Finish` ctx | What happens on SIGTERM mid-job |
| A5 | Fixes from the 11-16 review | Category handler status codes; DAL maps `sql.ErrNoRows` to domain errors; `%w` wrapping; `go mod tidy`; claim index; base schema migration | Error handling and indexing rationale |
| A6 | (Optional) attempts / max retries | Poison job stops after N attempts | Retry vs idempotency |

Suggested order: A1 → A2 → A3 → A4 → A5 → A6. A1–A3 make the project presentable; A4 also completes current Go learning.

## B. Algorithm route (Python, by pattern)

1. Arrays and hashing
2. Two pointers
3. Sliding window
4. Stack
5. Binary search
6. Linked list
7. Trees (DFS/BFS)
8. Heap
9. Intervals
10. Graphs (BFS/DFS, topological sort)
11. 1D dynamic programming
12. Backtracking

For each pattern, do about 5–8 classic problems, then move on and interleave re-solves. Target: common Mediums in ~30 minutes.

## C. Knowledge route (ordered by SDET weight)

1. Test design techniques; how to test a feature, an API, a system
2. Automation framework design, pytest depth, CI/CD
3. Python core
4. HTTP and networking basics
5. Database: transactions, isolation, locks, indexes
6. Flaky tests, test data, environments, observability for testers
7. Go and concurrency (for backend roles and project defense)
8. Light system design

## D. Behavioral (English)

Prepare 5–6 STAR stories that can be adapted to common prompts:
- conflict
- failure
- bug found
- improved a process
- learned fast
- why move toward engineering

## Later (after an offer)

The original long-term direction (AI-enabled systems, MCP, agent platforms) is parked. Revisit it once the job search is done.
