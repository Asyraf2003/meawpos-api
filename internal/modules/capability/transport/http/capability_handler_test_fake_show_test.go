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

	"pos-go/internal/modules/capability/domain"
	"pos-go/internal/modules/capability/ports"
)

type fakeShowCapabilityUsecase struct {
	fake *fakeCapabilityUsecases
}

func (f fakeShowCapabilityUsecase) Execute(ctx context.Context, key string) (domain.Capability, error) {
	_ = ctx
	f.fake.showCalls++
	if f.fake.err != nil {
		return domain.Capability{}, f.fake.err
	}

	capability, ok := f.fake.capabilities[key]
	if !ok {
		return domain.Capability{}, ports.ErrCapabilityNotFound
	}

	return capability, nil
}

var _ ShowCapabilityUsecase = fakeShowCapabilityUsecase{}
