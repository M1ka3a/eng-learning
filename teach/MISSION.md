# Mission: Land a Test Infra / Developer Productivity role in Singapore, with backend as a stretch

## Why

The learner has ~6 years of QA/SDET experience (ByteDance/TikTok, Garena) and wants engineering-heavy work. The near-term goal is an offer from a mid-size or larger company in Singapore. Target order, in the learner's own priority:

1. **Developer Tooling / Test Infrastructure / Developer Productivity / AI Quality** (primary).
2. A QA or SDET title is acceptable only if roughly 60–70% of the work is framework, platform, or tool development.
3. **Backend / general SWE** as a stretch. It is what the learner most wants long-term.
4. Plain QA automation only as a last fallback. Traditional QA roles are not wanted.

There is **no offer in hand** (the Ola offer was declined), so there is no fallback.

Background, existing interview prep, and real interview history: `learning/BACKGROUND.md` (private).

## Timeline and budget

- Latest acceptable start date: **January 2027** (learner, 2026-11-17). The employment gap (since 2026-09-19) is the main concern.
- Applications start **2026-11-19**. Interviews are expected to follow within a few weeks, so learning runs in parallel with interviewing.
- Study time: **4+ hours per day**.
- Resume: done. Deep project stories already exist in Lark; the gaps are listed in BACKGROUND.md.

## What interviews test (weights; revise after each real interview)

| Area | Test Infra / DevProd / AI Quality | Backend (stretch) |
|---|---|---|
| Past-project deep-dive (ByteDance stories, English) | High | High |
| Test framework / platform design (execution, scheduling, data model, retries, failure attribution) | High | Medium |
| AI combined with testing (where AI helps, limits, evaluating AI output) | High | Low |
| Algorithms (LeetCode Easy–Medium, Python) | Medium–High | High |
| Behavioral (English), including the employment-gap answer | High | High |
| Python (language, pytest, tooling) | High | Medium |
| Backend fundamentals: database, concurrency, HTTP/networking | Medium | High |
| Go (language, concurrency, context) | Medium | High |
| System design (platform-style for Test Infra; general for backend) | Medium | Medium–High |
| CI/CD, flaky tests, test strategy | Medium–High | Low |

Real interviews so far (Ola) asked about projects, AI for QA, and framework design. They did not ask algorithms. Larger companies usually have a coding round anyway.

## The role of this repository

This repo is the learner's **hands-on practice ground** for writing real backend code and practicing framework and system design. It is not a portfolio piece to polish. Choose project work by learning value for the target roles, such as concurrency, scheduling, reliability, test architecture, and service design. Polish is a non-goal. Byproducts that help the job search:

- concrete hands-on stories (for example the 60ns lease bug)
- a credible answer to "what have you been doing since September"
- practice explaining design decisions out loud

## Language homes (to avoid mixing Python and Go)

- **Python:** all algorithm practice and all test automation.
- **Go:** the service in this repository.
- Do not practice the same skill in both languages. When a Python/Go difference matters (`None`/`nil`, exceptions vs error returns, slicing, `=` vs `:=`), point it out and add it to the question bank if the learner trips on it.

## Success looks like

- Every resume claim (including MCP/agent work) and every known real question has a rehearsed English answer that holds up under follow-ups.
- Common Medium problems, by pattern, solved in Python within ~30 minutes.
- The learner can explain and defend the design and code they practiced in this repo.
- An offer in the target order above, starting no later than January 2027.

## Out of scope for now

- AI platform, MCP, or agent-orchestration features in this repo. AI-for-testing remains a high-weight **interview** topic.
- Polishing this repo for presentation.
- Polished bilingual HTML lessons by default.
