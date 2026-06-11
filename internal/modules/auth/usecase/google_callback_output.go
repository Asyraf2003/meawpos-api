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

	"pos-go/internal/modules/auth/ports"
)

func (u *GoogleFlow) issueGoogleCallbackOutput(
	ctx context.Context,
	accountID string,
	seed googleCallbackSeed,
) (GoogleCallbackOutput, error) {
	accessToken, accessExp, err := u.tokens.IssueAccessToken(ctx, ports.AccessTokenRequest{
		AccountID:  accountID,
		SessionID:  seed.sessionID,
		TrustLevel: "aal1",
	})
	if err != nil {
		return GoogleCallbackOutput{}, err
	}

	return GoogleCallbackOutput{
		AccessToken:    accessToken,
		AccessExp:      accessExp,
		RefreshToken:   seed.refreshToken,
		RefreshExp:     seed.refreshExp,
		TrustLevel:     "aal1",
		StepUpRequired: false,
	}, nil
}
