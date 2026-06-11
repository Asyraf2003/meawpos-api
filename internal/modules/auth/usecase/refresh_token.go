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
	"strings"
	"time"

	"pos-go/internal/modules/auth/ports"
)

var ErrInvalidRefreshToken = errors.New("invalid refresh token")

type RefreshToken struct {
	repo       ports.RefreshSessionRepository
	tokens     ports.TokenIssuer
	sessionTTL time.Duration
}

type RefreshTokenInput struct {
	RefreshToken string
}

type RefreshTokenOutput struct {
	AccessToken    string    `json:"access_token"`
	AccessExp      time.Time `json:"access_exp"`
	RefreshToken   string    `json:"refresh_token"`
	RefreshExp     time.Time `json:"refresh_exp"`
	TrustLevel     string    `json:"trust_level"`
	StepUpRequired bool      `json:"step_up_required"`
}

func NewRefreshToken(
	repo ports.RefreshSessionRepository,
	tokens ports.TokenIssuer,
	sessionTTL time.Duration,
) *RefreshToken {
	return &RefreshToken{
		repo:       repo,
		tokens:     tokens,
		sessionTTL: sessionTTL,
	}
}

func (u *RefreshToken) Execute(ctx context.Context, in RefreshTokenInput) (RefreshTokenOutput, error) {
	in.RefreshToken = strings.TrimSpace(in.RefreshToken)
	if in.RefreshToken == "" {
		return RefreshTokenOutput{}, ErrInvalidRefreshToken
	}

	refreshTokenHash := sha256Hex(in.RefreshToken)

	session, err := u.repo.FindActiveByRefreshTokenHash(ctx, refreshTokenHash)
	if err != nil {
		return RefreshTokenOutput{}, ErrInvalidRefreshToken
	}

	now := time.Now()
	if session.RevokedAt != nil {
		return RefreshTokenOutput{}, ErrInvalidRefreshToken
	}
	if !session.ExpiresAt.After(now) {
		return RefreshTokenOutput{}, ErrInvalidRefreshToken
	}

	newRefreshToken, err := randB64(48)
	if err != nil {
		return RefreshTokenOutput{}, err
	}

	newRefreshExp := now.Add(u.sessionTTL)
	newRefreshTokenHash := sha256Hex(newRefreshToken)

	accessToken, accessExp, err := u.tokens.IssueAccessToken(ctx, ports.AccessTokenRequest{
		AccountID:  session.AccountID,
		SessionID:  session.SessionID,
		TrustLevel: "aal1",
	})
	if err != nil {
		return RefreshTokenOutput{}, err
	}

	if err := u.repo.RotateRefreshToken(ctx, session.SessionID, newRefreshTokenHash, newRefreshExp); err != nil {
		return RefreshTokenOutput{}, err
	}

	return RefreshTokenOutput{
		AccessToken:    accessToken,
		AccessExp:      accessExp,
		RefreshToken:   newRefreshToken,
		RefreshExp:     newRefreshExp,
		TrustLevel:     "aal1",
		StepUpRequired: false,
	}, nil
}
