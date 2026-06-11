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

	"pos-go/internal/modules/auth/domain"
	"pos-go/internal/modules/auth/ports"
)

type fakeCallbackOIDCProvider struct {
	claims ports.OIDCClaims
}

func (f *fakeCallbackOIDCProvider) AuthCodeURL(p ports.OIDCAuthURLParams) string {
	_ = p
	return "https://example.com/oauth"
}

func (f *fakeCallbackOIDCProvider) ExchangeAndVerify(ctx context.Context, code, codeVerifier, redirectURL, nonce string) (ports.OIDCClaims, error) {
	_ = ctx

	if code == "" {
		return ports.OIDCClaims{}, errString("code empty")
	}
	if codeVerifier == "" {
		return ports.OIDCClaims{}, errString("code verifier empty")
	}
	if redirectURL == "" {
		return ports.OIDCClaims{}, errString("redirect url empty")
	}
	if nonce == "" {
		return ports.OIDCClaims{}, errString("nonce empty")
	}

	return f.claims, nil
}

type fakeAccountIdentityRepository struct {
	accountID       string
	resolveCalls    int
	upsertCalls     int
	lastClaims      ports.OIDCClaims
	lastIdentity    domain.Identity
	lastUpsertAccID string
}

func (f *fakeAccountIdentityRepository) ResolveOrCreateAccountByGoogle(ctx context.Context, claims ports.OIDCClaims) (string, error) {
	_ = ctx
	f.resolveCalls++
	f.lastClaims = claims
	return f.accountID, nil
}

func (f *fakeAccountIdentityRepository) UpsertIdentity(ctx context.Context, accountID string, identity domain.Identity) error {
	_ = ctx
	f.upsertCalls++
	f.lastUpsertAccID = accountID
	f.lastIdentity = identity
	return nil
}

type errString string

func (e errString) Error() string { return string(e) }
