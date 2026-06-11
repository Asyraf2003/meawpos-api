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

import "strings"

type Capability struct {
	Key                 string
	Domain              string
	Operation           string
	Method              string
	Path                string
	DefaultEnabled      bool
	Enabled             bool
	RequiredPermission  string
	RiskLevel           RiskLevel
	AuditRequired       bool
	IdempotencyRequired bool
	OwnerPackage        string
	TestProof           string
	DisabledReason      string
}

func NewCapability(input Capability) (Capability, error) {
	capability := Capability{
		Key:                 normalize(input.Key),
		Domain:              normalize(input.Domain),
		Operation:           normalize(input.Operation),
		Method:              strings.ToUpper(normalize(input.Method)),
		Path:                normalize(input.Path),
		DefaultEnabled:      input.DefaultEnabled,
		Enabled:             input.Enabled,
		RequiredPermission:  normalize(input.RequiredPermission),
		RiskLevel:           input.RiskLevel,
		AuditRequired:       input.AuditRequired,
		IdempotencyRequired: input.IdempotencyRequired,
		OwnerPackage:        normalize(input.OwnerPackage),
		TestProof:           normalize(input.TestProof),
		DisabledReason:      normalize(input.DisabledReason),
	}

	if err := capability.Validate(); err != nil {
		return Capability{}, err
	}

	if capability.Enabled {
		capability.DisabledReason = ""
	}

	return capability, nil
}

func (c Capability) Disable(reason string) Capability {
	c.Enabled = false
	c.DisabledReason = normalize(reason)

	return c
}

func (c Capability) Enable() Capability {
	c.Enabled = true
	c.DisabledReason = ""

	return c
}
