# Templates

This folder contains copyable documentation templates.

## Contents

- `0110_domain_scope_packet.md`: scope packet template for handing a bounded domain/API task to another AI, session, or human.
- `0120_prompt_authoring_rules.md`: copy-safe prompt writing rules, including ASCII and backtick guidance for web AI.
- `0121_codex_session_prompts.md`: terminal Codex start, implementation, review, and close-session prompts.
- `0122_web_ai_session_prompts.md`: browser AI open, continue, close, and cleanup prompts.
- `0123_analysis_and_review_prompts.md`: source-batch, domain migration, blueprint, and code review prompts.
- `0124_testing_and_proof_prompts.md`: test planning, proof running, and test-output interpretation prompts.
- `0125_data_capture_and_evidence_prompts.md`: prompts for turning external source or web AI output into organized docs.
- `0126_resume_after_pause_prompts.md`: prompts for long-pause resume, model switch, and missing-handoff recovery.

## Use This Folder When

- moving work between terminal Codex, GPT web, or another session;
- defining editable, read-only, and forbidden files;
- stating proof requirements for a scoped task.
- writing prompts that must be safe to copy from web AI;
- restarting work after a long pause or model switch;
- deciding where external analysis or source data should live in `docs/`.

Templates are starting points. Fill them with concrete files, rules, proof commands, and next steps before use.

## Copy Safety

For web AI prompts:

- keep prompts in plain text when possible;
- use ASCII;
- avoid nested Markdown code fences;
- use uppercase placeholders like `REPLACE_WITH_SCOPE`;
- put long source data below a clear `SOURCE DATA` heading;
- keep proof commands separate from claims.
