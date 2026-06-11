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
	"time"

	"pos-go/internal/modules/auth/ports"
)

func NewManualLogin(
	accounts ports.ManualAccountRepository,
	roles ports.AccountRoleAssigner,
	sessions ports.SessionStore,
	tokens ports.TokenIssuer,
	tx ports.Transactor,
	sessionTTL time.Duration,
) *ManualLogin {
	return &ManualLogin{
		accounts:   accounts,
		roles:      roles,
		sessions:   sessions,
		tokens:     tokens,
		tx:         tx,
		sessionTTL: sessionTTL,
		allowedRoles: map[string]string{
			"admin@example.com": "admin",
			"kasir@example.com": "cashier",
		},
	}
}
