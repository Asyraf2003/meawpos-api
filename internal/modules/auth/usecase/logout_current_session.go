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

	"pos-go/internal/modules/auth/domain"
	"pos-go/internal/modules/auth/ports"
)

type LogoutCurrentSession struct {
	revoker ports.SessionRevoker
}

func NewLogoutCurrentSession(revoker ports.SessionRevoker) *LogoutCurrentSession {
	return &LogoutCurrentSession{revoker: revoker}
}

func (u *LogoutCurrentSession) Execute(ctx context.Context, principal domain.Principal) error {
	if principal.SessionID == "" {
		return errors.New("session id is required")
	}

	return u.revoker.RevokeSession(ctx, principal.SessionID)
}
