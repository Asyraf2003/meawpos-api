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

	"pos-go/internal/modules/capability/domain"
	"pos-go/internal/modules/capability/ports"
)

type CheckCapability struct {
	repository ports.CapabilityRepository
}

func NewCheckCapability(repository ports.CapabilityRepository) *CheckCapability {
	return &CheckCapability{repository: repository}
}

func (u *CheckCapability) Execute(ctx context.Context, key string) error {
	key = strings.TrimSpace(key)
	if key == "" {
		return errors.New("capability key is required")
	}

	capability, err := u.repository.Get(ctx, key)
	if err != nil {
		return err
	}

	if !capability.Enabled {
		return domain.ErrCapabilityDisabled
	}

	return nil
}
