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

func (u *GoogleFlow) GoogleStart(ctx context.Context, in GoogleStartInput) (GoogleStartOutput, error) {
	in.Purpose = strings.TrimSpace(in.Purpose)
	if in.Purpose == "" {
		in.Purpose = "login"
	}
	if in.Purpose != "login" {
		return GoogleStartOutput{}, errors.New("invalid auth purpose")
	}

	in.RedirectURL = strings.TrimSpace(in.RedirectURL)
	if in.RedirectURL == "" {
		return GoogleStartOutput{}, errors.New("redirect url is required")
	}

	state, err := randB64(32)
	if err != nil {
		return GoogleStartOutput{}, err
	}

	nonce, err := randB64(32)
	if err != nil {
		return GoogleStartOutput{}, err
	}

	verifier, err := randB64(32)
	if err != nil {
		return GoogleStartOutput{}, err
	}

	authState := ports.AuthState{
		Nonce:        nonce,
		CodeVerifier: verifier,
		Purpose:      in.Purpose,
		CreatedAt:    time.Now(),
	}

	if err := u.states.Put(ctx, state, authState, u.stateTTL); err != nil {
		return GoogleStartOutput{}, err
	}

	redirectTo := u.oidc.AuthCodeURL(ports.OIDCAuthURLParams{
		State:         state,
		Nonce:         nonce,
		CodeChallenge: pkceChallenge(verifier),
		RedirectURL:   in.RedirectURL,
		Purpose:       in.Purpose,
	})

	return GoogleStartOutput{
		RedirectTo: redirectTo,
		State:      state,
	}, nil
}
