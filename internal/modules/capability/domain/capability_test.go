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

package domain

import "testing"

func TestNewCapabilityNormalizesAndValidates(t *testing.T) {
	capability, err := NewCapability(Capability{
		Key:                " account.role.assign ",
		Domain:             " account ",
		Operation:          " assign-role ",
		Method:             " post ",
		Path:               " /api/admin/accounts/:account_id/roles ",
		DefaultEnabled:     true,
		Enabled:            true,
		RequiredPermission: " account.role.assign ",
		RiskLevel:          RiskHigh,
		AuditRequired:      true,
		OwnerPackage:       " internal/modules/auth ",
		TestProof:          " internal/modules/auth/transport/http/account_role_handler_assign_test.go ",
		DisabledReason:     " stale reason ",
	})
	if err != nil {
		t.Fatalf("NewCapability() error = %v", err)
	}

	if capability.Key != "account.role.assign" {
		t.Fatalf("key = %q", capability.Key)
	}
	if capability.Method != "POST" {
		t.Fatalf("method = %q", capability.Method)
	}
	if capability.DisabledReason != "" {
		t.Fatalf("disabled reason = %q, want empty", capability.DisabledReason)
	}
}

func TestNewCapabilityRejectsInvalidRisk(t *testing.T) {
	_, err := NewCapability(validCapability(func(c *Capability) {
		c.RiskLevel = "critical"
	}))
	if err == nil {
		t.Fatal("NewCapability() error = nil, want error")
	}
}

func TestCapabilityEnableAndDisable(t *testing.T) {
	capability, err := NewCapability(validCapability(func(c *Capability) {
		c.Enabled = true
	}))
	if err != nil {
		t.Fatalf("NewCapability() error = %v", err)
	}

	disabled := capability.Disable("maintenance")
	if disabled.Enabled {
		t.Fatal("disabled capability is enabled")
	}
	if disabled.DisabledReason != "maintenance" {
		t.Fatalf("disabled reason = %q", disabled.DisabledReason)
	}

	enabled := disabled.Enable()
	if !enabled.Enabled {
		t.Fatal("enabled capability is disabled")
	}
	if enabled.DisabledReason != "" {
		t.Fatalf("enabled disabled reason = %q", enabled.DisabledReason)
	}
}

func validCapability(mutate func(*Capability)) Capability {
	capability := Capability{
		Key:                "account.role.assign",
		Domain:             "account",
		Operation:          "assign-role",
		Method:             "POST",
		Path:               "/api/admin/accounts/:account_id/roles",
		DefaultEnabled:     true,
		Enabled:            true,
		RequiredPermission: "account.role.assign",
		RiskLevel:          RiskHigh,
		AuditRequired:      true,
		OwnerPackage:       "internal/modules/auth",
		TestProof:          "internal/modules/auth/transport/http/account_role_handler_assign_test.go",
	}

	if mutate != nil {
		mutate(&capability)
	}

	return capability
}
