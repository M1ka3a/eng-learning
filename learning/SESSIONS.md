# Learning Session Log

Append one concise entry after each session. Record evidence, not assumptions.

## Prior progress (date not recorded)

- **Topics:** Go HTTP/JSON API, PostgreSQL and `database/sql`, layered structure, job lifecycle, workers, ticker/context/goroutines, `FOR UPDATE SKIP LOCKED`.
- **Evidence:** implementation in the repository; manual logs showed three tasks overlapping across multiple workers.
- **Implemented:** the listed functionality exists; worker execution remains a five-second Sleep placeholder.
- **Understood:** not established by repository contents alone; continue checking through learner explanations and design reasoning.
- **Open:** crash recovery, leases, retry, timeout, idempotency, graceful shutdown, robust PostgreSQL integration tests, fault injection, deployment, observability.

## 2026-11-12 — Lease implementation and test review

- **Topic:** Lease ownership, expiry/reclaim, renewal/completion guards, and PostgreSQL integration-test design.
- **Implemented:** existing lease implementation reviewed; DAL unit tests use a fake SQL driver. Lease integration tests use a random schema, require database name `backend_lab_test`, assert database state after stale-token attempts, and exercise concurrent claiming of one job.
- **Verified:** `env -u LEASE_TEST_DATABASE_URL GOCACHE=/private/tmp/backend-lab-go-gocache go test ./...` passed. Integration tests were not run; no PostgreSQL connection was made.
- **Understood:** learner identified the missing lease fields in a prior `GetJobByID` implementation and clarified that SQL lease behavior should remain covered by integration tests rather than mock-only tests. Broader independent explanation of PostgreSQL concurrency guarantees remains unassessed.
- **To reinforce:** run integration tests against a dedicated PostgreSQL test database; connect lease renewal to long-running execution; continue with retry and crash-recovery semantics.

## 2026-11-12 — Executor direction checkpoint

- **Encountered:** compared test failure triage, support ticket triage, and document extraction against current Software Engineering and AI Engineering learning goals.
- **Decision:** learner selected text-only customer support ticket triage as the first real Job Executor; first version has no external support-system integration or LLM call.
- **Implemented/Verified:** no executor code or behavior yet.
- **Understood/Transferable:** selection and scope are explicit; reasoning about the input/output contract and layer responsibilities is the next exercise, so understanding and transfer are not yet established.
- **Design checkpoint:** learner proposed description, order ID, and submitter UID as inputs, and independently distinguished accepting an incomplete ticket from allowing a worker to claim it. Exact readiness representation and output contract remain open.

## 2026-11-12 — Requirements-to-design learning checkpoint

- **Encountered:** learner questioned project sequencing and identified difficulty translating requirements into software design.
- **Decision:** make requirements-to-design translation the primary capability goal for the current ticket-triage session; define acceptance examples and contracts before coding. Keep the first implementation within the existing layers, with processing in Biz.
- **Implemented/Verified:** no ticket-triage code or behavior yet.
- **Understood/Transferable:** learner identified the gap independently; ability to derive a design from a new requirement remains to be demonstrated.

## 2026-11-12 — Ticket create/supplement session end

- **Implemented:** partial text-ticket flow: `TicketDetails`, initial `needs_input`/`pending` selection, PATCH route and Handler, Biz UID comparison, conditional DAL `jsonb_set` update to add order ID and move to `pending`; nullable lease fields in `GetJobByID` are handled.
- **Verified:** `go build ./...` passed after the latest Biz changes. No tests or PostgreSQL calls were made for the ticket flow. Earlier Lease integration tests remain unrun.
- **Understood:** learner distinguished submission from execution readiness, noticed the need to recover a job ID for later operations, and recognized that persisted ticket data should not be named only as input. Broader requirements-to-design independence remains in progress.
- **Transferable:** not yet demonstrated in a different use case.
- **Open:** validate UID/description on create; define triage output and rules; replace Worker Sleep, persist result, and expose it through GET. The request UID check is not authentication; the `jsonb_set` database behavior has not been verified.
