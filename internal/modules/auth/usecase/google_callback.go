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

func (u *GoogleFlow) GoogleCallback(ctx context.Context, in GoogleCallbackInput) (GoogleCallbackOutput, error) {
	if err := normalizeGoogleCallbackInput(&in); err != nil {
		return GoogleCallbackOutput{}, err
	}

	storedState, err := u.states.GetDel(ctx, in.State)
	if err != nil {
		return GoogleCallbackOutput{}, err
	}

	claims, err := u.oidc.ExchangeAndVerify(
		ctx,
		in.Code,
		storedState.CodeVerifier,
		in.RedirectURL,
		storedState.Nonce,
	)
	if err != nil {
		return GoogleCallbackOutput{}, err
	}

	seed, err := newGoogleCallbackSeed(u.sessionTTL)
	if err != nil {
		return GoogleCallbackOutput{}, err
	}

	accountID, err := u.persistGoogleCallback(ctx, claims, seed)
	if err != nil {
		return GoogleCallbackOutput{}, err
	}

	return u.issueGoogleCallbackOutput(ctx, accountID, seed)
}

func (u *GoogleFlow) persistGoogleCallback(
	ctx context.Context,
	claims ports.OIDCClaims,
	seed googleCallbackSeed,
) (string, error) {
	var accountID string

	err := u.tx.RunInTx(ctx, func(txCtx context.Context) error {
		resolvedAccountID, err := u.resolveGoogleCallbackAccount(txCtx, claims, seed.now)
		if err != nil {
			return err
		}

		accountID = resolvedAccountID
		return u.createGoogleCallbackSession(txCtx, accountID, seed)
	})
	if err != nil {
		return "", err
	}

	return accountID, nil
}
