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

	"pos-go/internal/modules/capability/ports"
)

type fakeEnableCapabilityUsecase struct {
	fake *fakeCapabilityUsecases
}

func (f fakeEnableCapabilityUsecase) Execute(ctx context.Context, key string) error {
	_ = ctx
	f.fake.enableCalls++
	capability, ok := f.fake.capabilities[key]
	if !ok {
		return ports.ErrCapabilityNotFound
	}
	f.fake.capabilities[key] = capability.Enable()

	return f.fake.err
}

type fakeDisableCapabilityUsecase struct {
	fake *fakeCapabilityUsecases
}

func (f fakeDisableCapabilityUsecase) Execute(ctx context.Context, key string, reason string) error {
	_ = ctx
	f.fake.disableCalls++
	f.fake.lastDisableReason = reason
	capability, ok := f.fake.capabilities[key]
	if !ok {
		return ports.ErrCapabilityNotFound
	}
	f.fake.capabilities[key] = capability.Disable(reason)

	return f.fake.err
}

var _ EnableCapabilityUsecase = fakeEnableCapabilityUsecase{}
var _ DisableCapabilityUsecase = fakeDisableCapabilityUsecase{}
