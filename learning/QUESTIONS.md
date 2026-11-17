# Question Bank

Results: `✗` miss · `△` partial · `✓` good. Next due: `✗` +1 day, `△` +3 days, `✓` doubles the previous interval (3 → 7 → 14 → 30); retire after `✓` at 30. Tag questions from real interviews as `real:<company>`.

## Knowledge and behavioral

| ID | Area | Question | Last | Result | Due |
|---|---|---|---|---|---|
| K001 | Go/context | With the current `main.go`, what happens to the workers when you press Ctrl+C? | 11-16 | ✗ | 11-18 |
| K002 | Reliability | A worker is killed after claiming a job but before `Finish`. What state is the job in, and when can it run again? | 11-16 | ✗ | 11-18 |
| K003 | Go/context | During graceful shutdown, which context should `Finish` use, and why? (worker ctx / Background / WithoutCancel+timeout) | 11-16 | ✗ | 11-18 |
| K004 | Go/context | Name the cases where a context gets cancelled, and the rule for "stop now" vs "finish in-flight work". | 11-16 | △ | 11-19 |
| K005 | Go/context | Context vs transaction: what does each one guarantee? | — | — | 11-18 |
| K006 | Database | In PostgreSQL Read Committed, how long does a snapshot last: per statement or per request? What anomaly follows? | 11-17 | ✗ | 11-18 |
| K007 | Database | Does a plain `SELECT` take any lock? | 11-17 | △ | 11-20 |
| K008 | Database | What do `FOR UPDATE` and `SKIP LOCKED` each do in the job claim query? | 11-17 | ✓ | 11-20 |
| K009 | Reliability | Which `WHERE` predicates stop a stale worker's `Finish` from overwriting a takeover? | 11-17 | △ | 11-20 |
| K010 | Testing | Why didn't the DAL tests catch `ClaimNext(60)`? Which test would? | — | — | 11-18 |
| K011 | Go/testing | Why does `go build` miss errors in `_test.go` files, and what catches them? | — | — | 11-19 |
| K012 | Project (EN) | Give a 2-minute pitch of this project. | — | — | 11-18 |
| K013 | Behavioral (EN) | Tell me about a bug you found that others missed. (Use the 60ns story.) | — | — | 11-18 |

<!-- Learner: add questions from your previous interview here, tagged real:<company>. -->

## Algorithms (Python)

| ID | Problem | Pattern | Last | Result | Time | Note | Due |
|---|---|---|---|---|---|---|---|
