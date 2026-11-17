# Upstream Teach compatibility notes

Upstream basis: the installed Matt Pocock Teach Skill at `~/.agents/skills/teach/` as inspected on 2026-11-17. The global installation is left unchanged. The copied format references in `references/` preserve the upstream MISSION, RESOURCES, LEARNING-RECORD, and GLOSSARY conventions.

## Preserved

- One mission grounds teaching choices.
- Trustworthy knowledge sources, with annotated resources and citations.
- Short, focused lessons designed for retention and an appropriate challenge level.
- Retrieval practice, spacing/interleaving when useful, and fast feedback.
- HTML lessons with reusable assets, links to related learning materials, primary sources, and an invitation to ask follow-up questions.
- Reference documents as compressed, reusable knowledge.
- Learning records as concise ADR-like durable insights, created only when evidence supports them.

## Intentional project extensions

- Renamed the skill to `teach-se`; it is explicitly invoked and does not shadow or modify global `$teach`.
- Workspace resolution locates the Git repository and always routes teaching files to `<repo>/teach/`, independent of invocation CWD. The code Lab root is the repository itself.
- Replaced one flat lesson/reference collection with topic-local `topics/<slug>/{lessons,labs,reference}/` collections and per-topic lesson numbering; cross-topic artifacts have one canonical home and are linked from other topic maps.
- Added a Lesson → Quiz → Project Lab → Assessment → Learning State loop that evaluates understanding, engineering application, verification, and judgment separately, and attributes Agent implementation.
- `learning/STATE.md` and `learning/SESSIONS.md` remain the project learning record. Teach learning records can preserve only verified durable insights and are not a mastery/progress ledger.
- Lessons use paired Chinese-English explanations by default, with aligned passages and standard English technical terms. Conversational coaching remains Chinese-first unless requested otherwise.

## Future merge guidance

When updating from upstream, compare the original `SKILL.md` and its `*-FORMAT.md` files with these copied references. Merge new teaching practices and format improvements, then preserve the project-specific workspace resolver, topic routing, local numbering, Lab/assessment process, and single-state rule. Record material changes here. Do not copy upstream instructions that assume the invocation CWD is a flat teaching workspace.
