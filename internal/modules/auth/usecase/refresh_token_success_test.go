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
	"strings"
	"testing"
	"time"

	"pos-go/internal/modules/auth/ports"
)

func TestRefreshToken_Success(t *testing.T) {
	originalExpiry := time.Now().Add(24 * time.Hour).Truncate(time.Second)
	repo := &fakeRefreshSessionRepository{session: ports.RefreshSession{
		SessionID: "sess-123", AccountID: "acc-123",
		RefreshTokenHash: sha256Hex("old-refresh-token"), ExpiresAt: originalExpiry,
	}}
	tokenIssuer := &fakeTokenIssuer{token: "new-access-token", exp: time.Now().Add(15 * time.Minute)}
	usecase := NewRefreshToken(repo, tokenIssuer)

	out, err := usecase.Execute(context.Background(), RefreshTokenInput{RefreshToken: "old-refresh-token"})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}
	if out.AccessToken != "new-access-token" || strings.TrimSpace(out.RefreshToken) == "" {
		t.Fatalf("unexpected output = %+v", out)
	}
	if out.RefreshToken == "old-refresh-token" {
		t.Fatal("refresh token was not rotated")
	}
	if !out.RefreshExp.Equal(originalExpiry) || !repo.lastNewExpiresAt.Equal(originalExpiry) {
		t.Fatalf("refresh expiry moved: output=%s stored=%s want=%s", out.RefreshExp, repo.lastNewExpiresAt, originalExpiry)
	}
	if out.TrustLevel != "aal1" || out.StepUpRequired {
		t.Fatalf("trust output = %+v", out)
	}
	if repo.findCalls != 1 || repo.rotateCalls != 1 || repo.lastSessionID != "sess-123" {
		t.Fatalf("repository calls/state = %+v", repo)
	}
	if repo.lastNewHash == "" || repo.lastNewHash == sha256Hex("old-refresh-token") {
		t.Fatalf("refresh hash was not rotated: %q", repo.lastNewHash)
	}
	if tokenIssuer.issueCalls != 1 || tokenIssuer.lastReq.AccountID != "acc-123" || tokenIssuer.lastReq.SessionID != "sess-123" {
		t.Fatalf("token issuer state = %+v", tokenIssuer)
	}
	if tokenIssuer.lastReq.TrustLevel != "aal1" {
		t.Fatalf("token issuer trust level = %q", tokenIssuer.lastReq.TrustLevel)
	}
}
