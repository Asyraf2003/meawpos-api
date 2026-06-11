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
	"time"

	"pos-go/internal/modules/auth/ports"
)

type fakeRefreshSessionRepository struct {
	session          ports.RefreshSession
	findErr          error
	rotateErr        error
	findCalls        int
	rotateCalls      int
	lastLookupHash   string
	lastSessionID    string
	lastNewHash      string
	lastNewExpiresAt time.Time
}

func (f *fakeRefreshSessionRepository) FindActiveByRefreshTokenHash(ctx context.Context, refreshTokenHash string) (ports.RefreshSession, error) {
	_ = ctx
	f.findCalls++
	f.lastLookupHash = refreshTokenHash
	return f.session, f.findErr
}

func (f *fakeRefreshSessionRepository) RotateRefreshToken(ctx context.Context, sessionID string, newRefreshTokenHash string, newExpiresAt time.Time) error {
	_ = ctx
	f.rotateCalls++
	f.lastSessionID = sessionID
	f.lastNewHash = newRefreshTokenHash
	f.lastNewExpiresAt = newExpiresAt
	return f.rotateErr
}
