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

type GoogleFlow struct {
	oidc         ports.OIDCProvider
	states       ports.AuthStateStore
	accounts     ports.AccountIdentityRepository
	sessions     ports.SessionStore
	tokens       ports.TokenIssuer
	tx           ports.Transactor
	roleAssigner ports.AccountRoleAssigner
	stateTTL     time.Duration
	sessionTTL   time.Duration
}

func NewGoogleFlow(
	oidc ports.OIDCProvider,
	states ports.AuthStateStore,
	accounts ports.AccountIdentityRepository,
	sessions ports.SessionStore,
	tokens ports.TokenIssuer,
	tx ports.Transactor,
	stateTTL time.Duration,
	sessionTTL time.Duration,
) *GoogleFlow {
	return &GoogleFlow{
		oidc:       oidc,
		states:     states,
		accounts:   accounts,
		sessions:   sessions,
		tokens:     tokens,
		tx:         tx,
		stateTTL:   stateTTL,
		sessionTTL: sessionTTL,
	}
}

func (u *GoogleFlow) WithRoleAssigner(roleAssigner ports.AccountRoleAssigner) *GoogleFlow {
	u.roleAssigner = roleAssigner
	return u
}
