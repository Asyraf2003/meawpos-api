// Copyright (C) 2026 Asyraf Mubarak
//
// This file is part of gopos-api.
//
// gopos-api is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, version 3 only.
//
// gopos-api is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with gopos-api. If not, see <https://www.gnu.org/licenses/>.

package postgres

import (
	"pos-go/internal/modules/capability/domain"
)

type capabilityScanner interface {
	Scan(dest ...any) error
}

func capabilitySelectSQL() string {
	return `
		SELECT
			key, domain, operation, method, path,
			default_enabled, enabled, required_permission, risk_level,
			audit_required, idempotency_required, owner_package,
			test_proof, disabled_reason
		FROM api_capabilities
	`
}

func scanCapability(row capabilityScanner) (domain.Capability, error) {
	var capability domain.Capability
	var riskLevel string
	var disabledReason *string

	err := row.Scan(
		&capability.Key,
		&capability.Domain,
		&capability.Operation,
		&capability.Method,
		&capability.Path,
		&capability.DefaultEnabled,
		&capability.Enabled,
		&capability.RequiredPermission,
		&riskLevel,
		&capability.AuditRequired,
		&capability.IdempotencyRequired,
		&capability.OwnerPackage,
		&capability.TestProof,
		&disabledReason,
	)
	if err != nil {
		return domain.Capability{}, err
	}

	capability.RiskLevel = domain.RiskLevel(riskLevel)
	if disabledReason != nil {
		capability.DisabledReason = *disabledReason
	}

	return domain.NewCapability(capability)
}

func capabilityArgs(capability domain.Capability) []any {
	var disabledReason any
	if capability.DisabledReason != "" {
		disabledReason = capability.DisabledReason
	}

	return []any{
		capability.Key,
		capability.Domain,
		capability.Operation,
		capability.Method,
		capability.Path,
		capability.DefaultEnabled,
		capability.Enabled,
		capability.RequiredPermission,
		string(capability.RiskLevel),
		capability.AuditRequired,
		capability.IdempotencyRequired,
		capability.OwnerPackage,
		capability.TestProof,
		disabledReason,
	}
}
