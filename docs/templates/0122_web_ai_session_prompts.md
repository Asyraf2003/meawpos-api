# Web AI Session Prompts

Use these prompts when moving work to GPT web or another browser-based AI.

## Open A Web AI Session

Copy this as plain text. Do not wrap it in Markdown fences.

```text
You are helping with a Go Echo API + PostgreSQL migration project.

IMPORTANT
Read and obey the provided docs excerpts. If a file is not provided, mark it as GAP instead of guessing.

CONTEXT
Repository: /home/asyraf/Code/go/pos-go
Active scope: REPLACE_WITH_SCOPE
Current blueprint: docs/blueprints/REPLACE_WITH_BLUEPRINT.md

DOCS PROVIDED
REPLACE_WITH_DOCS_OR_SUMMARIES

SOURCE DATA PROVIDED
REPLACE_WITH_SOURCE_DATA

TASK
REPLACE_WITH_TASK

RULES
- Answer in English unless the requested user-facing app text must be Indonesian.
- Do not invent files, tests, schema, or repo state.
- Keep domain logic out of HTTP handlers.
- Keep SQL inside persistence adapters.
- Do not propose an endpoint without a capability key.
- Do not claim implementation completion because you cannot access the repo directly.

EXPECTED OUTPUT
- FACT
- GAP
- DECISION
- BLUEPRINT or PATCH PLAN
- RISKS
- PROOF THE TERMINAL AGENT MUST RUN
- HANDOFF TEXT FOR CODEX
```

## Continue An Existing Web AI Problem

```text
Continue the same active scope.

PREVIOUS HANDOFF
REPLACE_WITH_HANDOFF

NEW DATA
REPLACE_WITH_NEW_DATA

TASK NOW
REPLACE_WITH_TASK

RULES
- Preserve existing decisions unless new evidence contradicts them.
- If new evidence changes the plan, say exactly what changed and why.
- Keep output structured so it can be pasted into docs/handoffs or docs/evidence.
```

## Close A Web AI Session

```text
Prepare a handoff for terminal Codex.

OUTPUT FORMAT
Use plain text headings.

INCLUDE
Active scope:
Blueprint referenced:
Files that Codex should read:
Files Codex may edit:
Files Codex must not edit:
Decisions made:
Facts proven from provided data:
Gaps:
Recommended next active step:
Proof commands Codex should run:

Do not include claims about commands being run unless command output was provided.
```

## Web AI Output Cleanup Prompt

```text
Rewrite your previous answer so it is safe to paste into a repository Markdown file.

RULES
- English only.
- ASCII only.
- No nested Markdown code fences.
- Use file paths exactly.
- Separate facts from recommendations.
- Remove speculative claims.
- Add a GAP section for unknowns.
```
