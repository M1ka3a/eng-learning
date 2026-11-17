# Teach SE workflow reference

## Frame the unit

State:

- Primary topic and any cross-topic links.
- Learning goal and why it follows from current capability evidence.
- Assumed knowledge and uncertainties.
- Chosen unit(s): lesson, source study, Quiz, Lab, or a smaller subset.
- Acceptance evidence and a bounded stopping point.
- Project files and authorized actions, if any.

Prefer high-trust primary material. Do not teach every prerequisite from scratch; check it briefly and fill only the gaps needed for the selected goal.

## Quiz design

For a full learning loop, normally ask one concept question, one concrete scenario question, and one decision/tradeoff question. Scale count and complexity to the topic. Questions should elicit reasoning, not trivia. Give feedback promptly after the learner answers. Keep Quiz assessment separate from Lab delivery.

## Lab brief (before any implementation)

Every substantive Lab brief contains:

1. Stable topic-prefixed Lab ID and title; relevant topic maps.
2. Real problem/risk and intended capability.
3. Allowed scope, paths, constraints, and implementation attribution plan.
4. Observable acceptance criteria, including relevant failure/edge cases.
5. Exact verification methods/commands and what each can prove; environment safety boundary.
6. Evidence requested from learner: diff/files, actual outputs, design rationale, review findings, and interpretation.
7. Stop condition and unmet criteria to resume later.

The brief describes the target, not the solution. Do not edit the repository or run tests merely because a Lab was described; use the authorization already present for each action. When no code change adds value, a reviewed design, test plan, or controlled read-only investigation can be a valid Lab.

## Assessment rubric

Report each dimension in plain language, with evidence and limits:

| Dimension | Evidence |
|---|---|
| Knowledge | Learner explanation and Quiz reasoning; distinguish correct/partial/incorrect/not assessed. |
| Application | Whether the specified behavior or analysis was achieved; attribute learner vs Agent work. |
| Verification | Tests/observations actually run, environment, result, and claims they support. |
| Judgment | Learner's explanation of requirements, alternatives, design, risks, and review. |
| Transfer | Application to a distinct scenario, only if checked. |

Use `demonstrated`, `partial`, `not demonstrated`, or `not assessed` per dimension as appropriate. Do not create an aggregate score that hides a mismatch, such as a successful Agent implementation with unverified learner understanding.

## Record and resume

At meaningful checkpoints update `learning/STATE.md`, the only progress source. Preserve separate evidence for Encountered, Implemented, Verified, Understood, and Transferable. Record ownership of implementation and what remains unknown. At substantive session end append a concise entry to `learning/SESSIONS.md`; never overwrite prior entries. A durable Teach learning record may capture a demonstrated insight, but not project progress or a Lab status table.

To resume, identify the existing Lab ID, inspect its open criteria and next checkpoint in `learning/STATE.md`, and inspect only relevant source files/tests. Do not restart completed acceptance criteria or claim completion until the stop conditions are met. To switch topics, route through `teach/TOPICS.md`; keep the former checkpoint intact.
