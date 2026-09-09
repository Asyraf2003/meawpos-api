#!/usr/bin/env bash
# Copyright (C) 2026 Asyraf Mubarak

set -euo pipefail

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$REPO_ROOT"

required_files=(
  "AGENTS.md"
  "README.md"
  "docs/README.md"
  "docs/AGENTS.md"
  "docs/0000_index.md"
  "docs/0001_north_star.md"
  "docs/0002_decision_policy.md"
  "docs/0003_document_authority_register.md"
  "docs/architecture/0005_foundation_redesign_baseline.md"
  "docs/architecture/0006_client_ui_form_factors_and_localization.md"
  "docs/architecture/0007_outcome_feedback_and_presentation_contract.md"
  "docs/transition/0000_master_execution_workflow.md"
  "docs/transition/0002_foundation_redesign_roadmap.md"
  "docs/transition/0003_phase_execution_ledger.md"
  "docs/engineering/AGENTS.md"
  "docs/engineering/0001_index.md"
  "docs/engineering/blueprints/0040_first_walking_skeleton_contract.md"
  "docs/engineering/blueprints/0041_r5_reproducible_security_release_gate.md"
  "docs/engineering/blueprints/0042_r6_human_end_to_end_client_foundation.md"
  "docs/evidence/0004_r4_postgresql_walking_skeleton_implementation.md"
  "docs/evidence/0005_r5_reproducible_security_release_gate.md"
  "docs/handoffs/2026-09-09_r6_implementation_handoff.md"
  "docs/legal/0001_repository_license_presentation.md"
  "LICENSE"
)

check_file() {
  local path="$1"
  [[ -f "$path" ]] || { echo "[FAIL] missing file: $path"; exit 1; }
  echo "[OK] file exists: $path"
}

check_contains() {
  local path="$1"
  local text="$2"
  rg -Fq "$text" "$path" || { echo "[FAIL] missing text in $path: $text"; exit 1; }
  echo "[OK] text found in $path: $text"
}

echo "== canonical file existence =="
for path in "${required_files[@]}"; do
  check_file "$path"
done

echo
echo "== documentation folder entrypoints =="
while IFS= read -r directory; do
  check_file "${directory%/}/README.md"
done < <(fd --type d . docs)

echo
echo "== canonical direction markers =="
check_contains "docs/README.md" "Canonical Read Order"
check_contains "docs/README.md" "R6 IMPLEMENTATION AUTHORIZED; LOGIN VERTICAL SLICE EXACT NEXT"
check_contains "docs/README.md" "handoffs/2026-09-09_r6_implementation_handoff.md"
check_contains "docs/README.md" "engineering/blueprints/0041_r5_reproducible_security_release_gate.md"
check_contains "docs/README.md" "evidence/0005_r5_reproducible_security_release_gate.md"
check_contains "docs/AGENTS.md" "Current owner direction and the North Star"
check_contains "docs/0003_document_authority_register.md" "HISTORICAL EVIDENCE"
check_contains "docs/architecture/0005_foundation_redesign_baseline.md" "Financial Mutation Boundary"
check_contains "docs/architecture/0005_foundation_redesign_baseline.md" "R4 uses additive PostgreSQL migrations only"
check_contains "docs/transition/0002_foundation_redesign_roadmap.md" 'Status: `IMPLEMENTATION AUTHORIZED; LOGIN/SESSION VERTICAL SLICE EXACT NEXT`'
check_contains "docs/transition/0002_foundation_redesign_roadmap.md" 'Status: `PLANNED AFTER R6; NOT STARTED`'
check_contains "docs/transition/0002_foundation_redesign_roadmap.md" "bounded meawpos-api fixed-ceiling refresh + browser auth transport"
check_contains "docs/transition/0003_phase_execution_ledger.md" "R6 IMPLEMENTATION AUTHORIZED; LOGIN VERTICAL SLICE EXACT NEXT"
check_contains "docs/transition/0003_phase_execution_ledger.md" "ACCEPTED: DEFER FROM R6"
check_contains "docs/handoffs/2026-09-09_r6_implementation_handoff.md" "CANONICAL IMPLEMENTATION HANDOFF; R6 IMPLEMENTATION AUTHORIZED"
check_contains "docs/engineering/domain/0030_domain_contracts.md" "There is no universal CRUD requirement."
check_contains "docs/engineering/architecture/0022_api_capability_control.md" "Business component activation"
check_contains "docs/engineering/security/0080_security_baseline.md" "make release-gate"
check_contains "docs/legal/0001_repository_license_presentation.md" 'canonical `LICENSE` file'
check_contains "LICENSE" "GNU AFFERO GENERAL PUBLIC LICENSE"
check_contains "LICENSE" "END OF TERMS AND CONDITIONS"

echo
echo "[PASS] canonical AI/documentation rules audit passed"
