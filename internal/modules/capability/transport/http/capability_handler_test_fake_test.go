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

package http

import (
	"context"
	"testing"

	"pos-go/internal/modules/capability/domain"
	"pos-go/internal/modules/capability/ports"
)

type fakeCapabilityUsecases struct {
	capabilities      map[string]domain.Capability
	listCalls         int
	showCalls         int
	enableCalls       int
	disableCalls      int
	lastDisableReason string
	err               error
}

func (f *fakeCapabilityUsecases) Execute(ctx context.Context) ([]domain.Capability, error) {
	_ = ctx
	f.listCalls++
	if f.err != nil {
		return nil, f.err
	}

	out := make([]domain.Capability, 0, len(f.capabilities))
	for _, capability := range f.capabilities {
		out = append(out, capability)
	}

	return out, nil
}

func newCapabilityHandlerForTest(t *testing.T) (*CapabilityHandler, *fakeCapabilityUsecases) {
	t.Helper()

	capability, err := domain.NewCapability(domain.Capability{
		Key:                 "capability.manage",
		Domain:              "capability",
		Operation:           "manage",
		Method:              "*",
		Path:                "/api/admin/capabilities",
		DefaultEnabled:      true,
		Enabled:             true,
		RequiredPermission:  "capability.manage",
		RiskLevel:           domain.RiskHigh,
		AuditRequired:       true,
		IdempotencyRequired: false,
		OwnerPackage:        "internal/modules/capability/transport/http",
		TestProof:           "handler tests and SQL proof placeholders",
	})
	if err != nil {
		t.Fatalf("NewCapability() error = %v", err)
	}

	fake := &fakeCapabilityUsecases{
		capabilities: map[string]domain.Capability{capability.Key: capability},
	}
	handler := NewCapabilityHandler(
		fake,
		fakeShowCapabilityUsecase{fake: fake},
		fakeEnableCapabilityUsecase{fake: fake},
		fakeDisableCapabilityUsecase{fake: fake},
	)

	return handler, fake
}

var _ ListCapabilitiesUsecase = (*fakeCapabilityUsecases)(nil)
var _ = ports.ErrCapabilityNotFound
