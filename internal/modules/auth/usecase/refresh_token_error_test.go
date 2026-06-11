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
	"testing"
	"time"

	"pos-go/internal/modules/auth/ports"
)

func TestRefreshToken_RejectsEmptyRefreshToken(t *testing.T) {
	usecase := NewRefreshToken(&fakeRefreshSessionRepository{}, &fakeTokenIssuer{}, 24*time.Hour)

	_, err := usecase.Execute(context.Background(), RefreshTokenInput{})
	if err == nil {
		t.Fatal("Execute() error = nil, want error")
	}
}

func TestRefreshToken_RejectsExpiredRefreshToken(t *testing.T) {
	repo := &fakeRefreshSessionRepository{
		session: ports.RefreshSession{
			SessionID: "sess-123",
			AccountID: "acc-123",
			ExpiresAt: time.Now().Add(-1 * time.Minute),
		},
	}

	usecase := NewRefreshToken(repo, &fakeTokenIssuer{}, 24*time.Hour)

	_, err := usecase.Execute(context.Background(), RefreshTokenInput{
		RefreshToken: "old-refresh-token",
	})
	if err == nil {
		t.Fatal("Execute() error = nil, want error")
	}
}
