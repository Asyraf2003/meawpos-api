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

	"github.com/google/uuid"
)

type manualLoginSeed struct {
	now              time.Time
	sessionID        string
	refreshToken     string
	refreshExp       time.Time
	refreshTokenHash string
}

func newManualLoginSeed(sessionTTL time.Duration) (manualLoginSeed, error) {
	now := time.Now()
	refreshToken, err := randB64(48)
	if err != nil {
		return manualLoginSeed{}, err
	}

	return manualLoginSeed{
		now:              now,
		sessionID:        uuid.NewString(),
		refreshToken:     refreshToken,
		refreshExp:       now.Add(sessionTTL),
		refreshTokenHash: sha256Hex(refreshToken),
	}, nil
}
