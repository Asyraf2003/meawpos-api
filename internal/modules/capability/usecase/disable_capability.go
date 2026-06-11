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

package usecase

import (
	"context"
	"errors"
	"strings"

	"pos-go/internal/modules/capability/ports"
)

type DisableCapability struct {
	repository ports.CapabilityRepository
}

func NewDisableCapability(repository ports.CapabilityRepository) *DisableCapability {
	return &DisableCapability{repository: repository}
}

func (u *DisableCapability) Execute(ctx context.Context, key string, reason string) error {
	key = strings.TrimSpace(key)
	if key == "" {
		return errors.New("capability key is required")
	}

	capability, err := u.repository.Get(ctx, key)
	if err != nil {
		return err
	}

	return u.repository.Save(ctx, capability.Disable(reason))
}
