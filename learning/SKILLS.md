# Skills Matrix

Single source of capability level. Levels: `0` not seen · `1` can recall · `2` explains and survives follow-ups · `3` applies in new scenario / under interview conditions · `?` not assessed · `?(self)` self-reported, not yet checked.

Weight is for the primary SDET/QA target (H/M/L); see `teach/MISSION.md`. Target is the level needed for interviews.

## Algorithms (Python)

| Skill | W | Now | Target | Evidence | Last | Next |
|---|---|---|---|---|---|---|
| Overall Medium problems | H | 1 | 3 | Self: did one round before; Medium feels hard | — | Start the pattern route in ROADMAP |
| Arrays / hashing | H | ? | 3 | — | — | First pattern |
| Two pointers / sliding window | H | ? | 3 | — | — | — |
| Stack / queue | H | ? | 2 | — | — | — |
| Binary search | H | ? | 2 | — | — | — |
| Linked list | M | ? | 2 | — | — | — |
| Trees / BFS / DFS | H | ? | 3 | — | — | — |
| Graphs | M | ? | 2 | — | — | — |
| Heap / intervals | M | ? | 2 | — | — | — |
| Dynamic programming (1D) | M | ? | 2 | — | — | — |

## Python

| Skill | W | Now | Target | Evidence | Last | Next |
|---|---|---|---|---|---|---|
| Core language (data types, comprehensions, iterators/generators, exceptions, OOP) | H | ?(self) | 2 | — | — | Quick drill |
| pytest (fixtures, parametrize, markers, conftest) | H | ?(self) | 3 | — | — | Will be exercised by the project's API test suite |
| requests / API automation | H | ?(self) | 3 | — | — | Same |
| UI automation (Selenium/Playwright) concepts | M | ?(self) | 2 | — | — | — |

## SDET / Testing

| Skill | W | Now | Target | Evidence | Last | Next |
|---|---|---|---|---|---|---|
| Test case design (equivalence, boundary, state transition, pairwise) | H | ?(self) | 3 | — | — | Mock: "how would you test X" |
| Automation framework design (layers, data, reporting) | H | ?(self) | 3 | — | — | — |
| Test strategy and pyramid (unit/integration/E2E tradeoffs) | H | 2 | 3 | 2026-11-16: proposed reproducing the 60ns bug at the Biz layer first; existing DAL tests could not catch a caller unit bug | 11-16 | Explain in English as a project story |
| CI/CD and test pipelines | H | ?(self) | 2 | — | — | Project: GitHub Actions |
| Flaky tests: causes and handling | M | ? | 2 | — | — | — |
| Concurrency / DB integration testing | M | 1 | 2 | Repo has isolated-schema PG tests (coach-authored) | 11-16 | Explain how the isolation works |

## Project (this repo)

| Skill | W | Now | Target | Evidence | Last | Next |
|---|---|---|---|---|---|---|
| 2-minute English pitch | H | ? | 3 | — | — | Write and rehearse |
| Defend architecture and every line | H | 1 | 3 | Many parts coach-authored; not yet explained by learner | — | Defend-it questions per finish item |
| 60ns bug story (STAR) | H | 1 | 3 | Lived it on 11-16; not yet told | 11-16 | Mock interview |

## Database

| Skill | W | Now | Target | Evidence | Last | Next |
|---|---|---|---|---|---|---|
| SQL basics / joins / aggregation | M | ?(self) | 2 | — | — | — |
| Transactions and isolation levels, MVCC | M | 1 | 2 | 11-17: said MVCC gives a request-level view; Read Committed is per statement | 11-17 | Review Q in bank |
| Row locks, `FOR UPDATE`, `SKIP LOCKED` | M | 2 | 2 | 11-17: correctly explained `FOR UPDATE` and predicted the `SKIP LOCKED` empty result | 11-17 | Keep via review |
| Indexes and query plans | M | ? | 2 | — | — | — |
| MySQL vs PostgreSQL differences (common SG stacks) | L | ? | 1 | — | — | Only if interviews ask |

## Go

| Skill | W | Now | Target | Evidence | Last | Next |
|---|---|---|---|---|---|---|
| Syntax basics, exported names, error returns | M | 1 | 2 | 11-16: `t.logF`, discarded `Scan` error | 11-16 | Learn through project work |
| Error handling (`errors.Is`, `%w`) | M | 1 | 2 | Encountered in review and experiment | 11-16 | — |
| context: cancellation, deadline, propagation | M | 1 | 2 | 11-16: threaded ctx handler→DAL with coaching; could not answer shutdown/lease questions unaided | 11-16 | Finish graceful shutdown |
| Goroutines, channels, graceful shutdown | M | 1 | 2 | Observed signal demo; not implemented | 11-16 | Project item |
| Interfaces and layering | L | ? | 2 | — | — | — |

## Other

| Skill | W | Now | Target | Evidence | Last | Next |
|---|---|---|---|---|---|---|
| HTTP basics (methods, status codes, idempotency, headers) | M | ? | 2 | — | — | Drill |
| Networking (TCP/UDP, DNS, TLS, what happens on a URL) | M | ? | 2 | — | — | Drill |
| Linux / shell basics for testers | M | ?(self) | 2 | — | — | — |
| System design: "how would you test this system" | M | ? | 2 | — | — | Mock |
| Behavioral (English, STAR) | H | ? | 3 | Has prior interview prep | — | Build 5–6 stories |
