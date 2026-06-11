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

package middleware

import (
	"context"

	"pos-go/internal/modules/auth/domain"
	"pos-go/internal/modules/auth/ports"
)

type fakeAccessTokenVerifier struct {
	claims ports.AccessTokenClaims
	err    error
}

func (f *fakeAccessTokenVerifier) VerifyAccessToken(ctx context.Context, token string) (ports.AccessTokenClaims, error) {
	_ = ctx
	_ = token
	return f.claims, f.err
}

type fakePrincipalResolver struct {
	principal domain.Principal
	err       error
	lastInput ports.ResolvePrincipalInput
}

func (f *fakePrincipalResolver) Resolve(ctx context.Context, in ports.ResolvePrincipalInput) (domain.Principal, error) {
	_ = ctx
	f.lastInput = in
	return f.principal, f.err
}

type fakeSessionStatusChecker struct {
	active        bool
	err           error
	lastSessionID string
}

func (f *fakeSessionStatusChecker) IsSessionActive(ctx context.Context, sessionID string) (bool, error) {
	_ = ctx
	f.lastSessionID = sessionID
	return f.active, f.err
}
