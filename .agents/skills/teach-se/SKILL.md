---
name: teach-se
description: Teach software engineering and AI engineering through topic-routed lessons, quizzes, and evidence-based project labs in this repository.
metadata:
  short-description: "Topic-based software engineering lessons and project labs"
---

# Teach SE

Use this project-specific extension of Matt Pocock's Teach Skill for long-running Software Engineering and AI Engineering learning in `backend-lab-go`. Preserve Teach's mission-grounded teaching, trusted sources, short focused lessons, retrieval practice, fast feedback, reference materials, and evidence-based learning records while following this repository's topic and project-state rules.

## Resolve the workspace and topic

The invocation directory is not the teaching workspace. Find the repository root (prefer `git rev-parse --show-toplevel`; otherwise walk upward) and verify that it contains `AGENTS.md`, `learning/STATE.md`, and `teach/TOPICS.md`. The Teach workspace is always `<repo>/teach/`; the real Go project is `<repo>/`.

Read `teach/TOPICS.md` and resolve the requested topic to its folder `teach/topics/<topic-slug>/TOPIC.md`. If the user names an unlisted topic, determine whether it belongs under an existing topic or merits adding a new one; update the index and create only needed topic files. Read only the selected topic's material and connected prerequisites.

Before planning a learning unit, consult `teach/MISSION.md`, the selected topic map, `learning/STATE.md`, relevant recent `learning/SESSIONS.md` entries, and `learning/ROADMAP.md` as an option set rather than a queue. Resolve open Quiz/Lab checkpoints from `learning/STATE.md` and the conversation. Never create a parallel progress state, topic mastery tracker, gradebook, or per-topic mission.

## Topic workspace layout

```text
teach/
  MISSION.md                 # one shared mission
  TOPICS.md                  # topic router
  RESOURCES.md               # shared curated sources
  assets/                    # shared lesson components
  topics/<topic-slug>/
    TOPIC.md                 # knowledge map, topic resource links, cross-links
    lessons/                 # topic-local 0001, 0002, ... numbering
    labs/                    # topic-local Lab briefs/artifacts
    reference/               # topic-specific quick references
  learning-records/          # verified durable insights only; create lazily
  NOTES.md                   # teaching preferences
```

Create topic subdirectories only when needed; do not add empty placeholders. Keep a cross-topic lesson or Lab in one canonical topic folder and link it from other relevant topic maps. Do not duplicate artifacts to make them appear in several topics. Lesson filenames increment within their topic, e.g. `topics/database-engineering/lessons/0001-lease-claims.html`. Lab identifiers include the topic slug and local sequence, e.g. `database-engineering-LAB-0001`. Use relative links and verify them after moving content. Shared CSS and reusable interactive components live in `teach/assets/`; topical content can link upward to them.

## Learning loop

For consequential engineering capabilities, use a proportional **Lesson → Quiz → Project Lab → Assessment → Learning State** loop. For a small concept, use a smaller unit without forcing code changes or HTML. Before teaching, tell the learner the goal, why now, scope/constraints, and observable acceptance evidence. Do not reveal Quiz answers or implement a Lab before the learner has had a chance to reason or act.

1. **Lesson:** teach only the knowledge needed for the goal. Prefer official/primary sources; record useful sources in shared `teach/RESOURCES.md` or the topic map and cite claims. Use a short, polished standalone HTML lesson when it adds durable value; preserve the Teach HTML conventions and feedback mechanisms. Otherwise teach conversationally or use a design review, source-reading, incident analysis, or independent experiment. **Bilingual correspondence is a teaching principle for lessons:** default to Chinese-English paired content for the core explanations, concepts, and takeaways, keeping each English passage aligned with its Chinese counterpart rather than collecting an unrelated translation at the end. Preserve standard English technical terms and source titles. Keep conversational coaching in Chinese by default unless the learner asks for more English.
2. **Quiz:** use a suitable mix of concept explanation, scenario prediction, and technical decision questions. Ask before giving solutions. Use answers to distinguish demonstrated reasoning from exposure and to tune the Lab scope.
3. **Project Lab:** prefer valuable work against actual code under the repository root, not a toy copy under `teach/`. A Lab may be read-only analysis, review, test design, failure reproduction, independent experiment, refactor, or implementation; add functionality only when the learning goal and product need justify it. Before solving, specify the Lab ID, connected topics, problem, allowed paths/actions, acceptance criteria, verification method and safety boundary, evidence to return, and stop condition. The learner may code or delegate implementation to an Agent; attribute the work and require learner review/explanation where understanding is being assessed. Follow repository authorization rules separately for edits, tests, database access, and environment changes.
4. **Assessment:** assess knowledge/Quiz, engineering application, verification, and technical judgment separately. Record who implemented (learner, Agent, or jointly). A passing test proves only the behavior it actually exercises; a present test is not a run result. Agent-authored implementation can be delivered work but is not proof of the learner's understanding. Mark unsupported understanding or transfer as unknown/not assessed.
5. **Learning State:** update only `learning/STATE.md` for meaningful capability/project evidence, tracking Encountered, Implemented, Verified, Understood, and Transferable separately. Append concise substantive session evidence to `learning/SESSIONS.md` at session end. Keep unfinished Lab acceptance criteria and the next checkpoint in the same state source when meaningful. Do not copy state into topic maps. Create a Teach learning record only for a durable insight supported by learner explanation/reasoning, following the copied format; it must not become a second progress record.

## Source-derived formats and upstream compatibility

- Use `references/MISSION-FORMAT.md` for the single mission, `references/RESOURCES-FORMAT.md` for curated trustworthy sources, `references/LEARNING-RECORD-FORMAT.md` for durable insights, and `references/GLOSSARY-FORMAT.md` only when the learner has demonstrated command of a term.
- Reuse assets before creating new ones. A standalone lesson uses topic-local numbering and links to its topic map, connected lessons/references, a primary source, and a reminder to ask follow-up questions.
- `UPSTREAM-NOTES.md` records intentional differences from the upstream Teach Skill. When upstream Teach changes, compare its SKILL and FORMAT files with the local copies; merge useful teaching improvements without restoring incompatible workspace, numbering, or state behavior. Do not edit the globally installed `teach` skill.

For the detailed planning, Lab, assessment, and resume checklist, read [`references/workflow.md`](references/workflow.md). For the rationale and mapping to upstream behavior, read [`UPSTREAM-NOTES.md`](UPSTREAM-NOTES.md).
