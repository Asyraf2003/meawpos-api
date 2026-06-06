# Data Capture And Evidence Prompts

Use this file when collecting source data, command output, or AI analysis so the repository stays organized.

## Where To Put Data

Use these targets:

```text
docs/evidence/     source facts, command output summaries, extracted inventories
docs/blueprints/   planned work and implementation order
docs/adr/          accepted architecture or policy decisions
docs/handoffs/     continuation notes between sessions
docs/templates/    reusable prompts
docs/archive/      obsolete historical material
```

Never create loose migration notes at the repository root.

## Capture A Laravel Source Batch

```text
Convert this Laravel source batch into a Go migration evidence document.

OUTPUT TARGET
docs/evidence/REPLACE_WITH_NUMBERED_NAME.md

SOURCE BATCH
REPLACE_WITH_SOURCE_BATCH

RULES
- Do not lose table names, field names, route names, class names, or test behavior.
- Do not invent missing files.
- Preserve exact operation behavior when visible.
- Mark unknown source as GAP.
- Extract API candidates, capability candidates, DB constraints, and tests to preserve.
- Do not implement code.

EXPECTED OUTPUT
- evidence document content;
- related blueprint update notes;
- smallest next source batch request.
```

## Capture Web AI Analysis Into Docs

```text
Turn the analysis below into a repository document.

DOCUMENT CLASS
Choose one: evidence, blueprint, ADR, handoff, archive.

TARGET FILE
docs/REPLACE_WITH_FOLDER/REPLACE_WITH_FILE.md

ANALYSIS
REPLACE_WITH_ANALYSIS

RULES
- English only.
- ASCII only.
- Separate proven facts from recommendations.
- Move guesses into GAP.
- Do not duplicate another active document's purpose.
- Include references to source data or command output when available.
```

## Create A Handoff From Current Chat

```text
Create a handoff from this chat.

TARGET
docs/handoffs/REPLACE_WITH_DATE_SCOPE.md

INCLUDE
- date;
- active scope;
- repository path;
- current blueprint;
- files changed;
- evidence created;
- commands run;
- proof result;
- decisions made;
- gaps;
- next valid active step;
- prompt to start the next session.

Do not invent command output.
```

## Data Quality Checklist

Before committing evidence:

```text
Does it say where the data came from?
Does it separate FACT and GAP?
Does it avoid implementation claims?
Does it preserve names exactly?
Does it point to the next blueprint or decision?
Does it avoid duplicate active docs?
```
