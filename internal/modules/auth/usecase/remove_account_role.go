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

	"pos-go/internal/modules/auth/ports"
)

var ErrBaseRoleRemovalNotAllowed = errors.New("base role cannot be removed")

type RemoveAccountRole struct {
	remover ports.AccountRoleRemover
}

func NewRemoveAccountRole(remover ports.AccountRoleRemover) *RemoveAccountRole {
	return &RemoveAccountRole{remover: remover}
}

func (u *RemoveAccountRole) Execute(ctx context.Context, accountID string, roleKey string) error {
	accountID = strings.TrimSpace(accountID)
	roleKey = strings.TrimSpace(roleKey)

	if accountID == "" {
		return errors.New("account id is required")
	}
	if roleKey == "" {
		return errors.New("role key is required")
	}
	if roleKey == "base" {
		return ErrBaseRoleRemovalNotAllowed
	}

	return u.remover.RemoveRole(ctx, accountID, roleKey)
}
