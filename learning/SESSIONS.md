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

## 2026-11-14 — Ticket triage example implementation checkpoint

- **Decision:** learner requested the coach to commit the existing worktree, then implement the discussed ticket features as a reference to study and re-create.
- **Implemented:** baseline commit `db07412`; coach-authored create validation, Biz phrase-based classification, Worker execution, result/error persistence under lease-token guard, GET result, UID list, and review-only SQL migration.
- **Verified:** `go build ./...` and `git diff --check` passed. No tests were run, no migration was applied, and no database/service was accessed.
- **Understood/Transferable:** learner has not yet explained or independently re-implemented this example; do not infer mastery from the coach-authored code.
- **Open:** actual jobs schema compatibility, PostgreSQL behavior, classification quality, and real authentication for UID-based access.

## 2026-11-14 — Classification requirement correction

- **Understood:** learner challenged keyword-based automatic classification as an unjustified business rule and suggested customer or support-agent selection; this is direct evidence of requirements judgment.
- **Decision/Implemented:** coach revised the example to require a customer-selected category at creation and to produce a routing result marked `user_selected`. Phrase matching was removed. Existing Job execution and result persistence remain in place for now.
- **Verified:** non-database `go build ./...` and `git diff --check` passed. No tests, service run, migration, or PostgreSQL call.
- **Open:** human review/editing is not implemented; the simple routing mapping may not justify an asynchronous Job Executor. Actual database behavior and authentication are unverified.

## 2026-11-16 — Code review and lease-duration bug fix checkpoint

- **Encountered:** read-only review of layering, error handling, context propagation, resource management, SQL/migrations, and go.mod. Main findings: `ClaimNext(60)` was 60ns; no `context.Context` passed below handlers; `main.go` starts 4 workers, only one of them cancellable, with no graceful shutdown or server timeouts; handlers depend on `sql.ErrNoRows`; category handler returns wrong status codes; claim/UID queries have no indexes; base `jobs`/`categories` schema missing from migrations; `go mod tidy` needed (direct deps marked indirect).
- **Implemented:** learner directed a test-first fix. Coach wrote `internal/biz/job_integration_test.go`, then changed Biz to use `claimLeaseDuration = 60 * time.Second`.
- **Verified:** before the fix, the test failed at `Finish: job 1 finish rejected: status, lease, or token no longer matches`. After the fix, `go test ./... -count=1` against `backend_lab_test` passed, including the DAL Lease integration tests on real PostgreSQL for the first time. Test schemas were dropped afterward. Changes are not committed.
- **Understood/Transferable:** learner chose the reproduce-then-fix order and gated each step. Their explanation of why DAL tests missed the bug, and of the remaining review items, has not been assessed.
- **Open:** prioritized next fixes: worker lifecycle/graceful shutdown in `main.go`, context propagation plus domain errors with `%w` wrapping, then indexes and base-schema migration.

## 2026-11-16 — Go context session end

- **Topic:** first exposure to `context.Context`: cancellation, deadlines, propagation, signals, and its relationship to leases.
- **Implemented:** learner, with step-by-step coaching, propagated ctx from `r.Context()` to `QueryRowContext` for `GetJobByID` and `AddOrderID`, going one layer further than instructed; fixed compile/vet errors layer by layer, including test call sites. The coach completed the experiment test at the learner's request after they got stuck on `t.Logf`, the ignored `Scan` error, and the test name.
- **Verified:** build/vet/full tests pass against `backend_lab_test`. Observed: deadline ctx cancels `pg_sleep(2)` at ~201ms (`errors.Is` sees `DeadlineExceeded` through pgx's wrapper), and no-ctx query runs 2s. Signal demo: unhandled SIGTERM gives exit 143 with no cleanup; `signal.NotifyContext` gives a clean exit 0. A first SIGINT attempt was invalid because background processes in non-interactive shells ignore SIGINT; this was replaced with SIGTERM.
- **Understood:** partial. The learner's question "what is context for?" led to the purpose table (caller gone, timeout, shutdown). Learner generalized "usually finish current work", which the coach refined to a result-still-needed rule. They did not independently answer the Ctrl+C, lease-takeover, or `Finish` ctx questions; the answers were given. They expressed early linkage of context and transactions; the distinction is unchecked.
- **Transferable:** not demonstrated.
- **Open:** learner's own explanation of the cancellation model; graceful shutdown implementation; ctx for the remaining DAL methods and the Worker; at-least-once/idempotency implications of lease takeover.
