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

package jwt

import (
	"context"
	"testing"
	"time"

	"pos-go/internal/modules/auth/ports"
)

func mustIssueTestAccessToken(
	t *testing.T,
	ttl time.Duration,
	secret string,
	accountID string,
	sessionID string,
	trustLevel string,
) string {
	t.Helper()

	issuer, err := NewHMACIssuer(
		"pos-go",
		"pos-go-client",
		"local-dev-key",
		secret,
		ttl,
	)
	if err != nil {
		t.Fatalf("NewHMACIssuer() error = %v", err)
	}

	token, _, err := issuer.IssueAccessToken(context.Background(), ports.AccessTokenRequest{
		AccountID:  accountID,
		SessionID:  sessionID,
		TrustLevel: trustLevel,
	})
	if err != nil {
		t.Fatalf("IssueAccessToken() error = %v", err)
	}

	return token
}

func mustNewTestVerifier(t *testing.T, secret string) *Verifier {
	t.Helper()

	verifier, err := NewHMACVerifier("pos-go", "pos-go-client", secret)
	if err != nil {
		t.Fatalf("NewHMACVerifier() error = %v", err)
	}

	return verifier
}
