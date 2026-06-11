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

package google

import (
	"context"
	"errors"
	"strings"
	"time"

	"pos-go/internal/modules/auth/ports"

	"golang.org/x/oauth2"
)

func (o *OIDC) ExchangeAndVerify(
	ctx context.Context,
	code string,
	codeVerifier string,
	redirectURL string,
	nonce string,
) (ports.OIDCClaims, error) {
	cfg := o.oauthConfigFor(redirectURL)

	token, err := cfg.Exchange(
		ctx,
		code,
		oauth2.SetAuthURLParam("code_verifier", codeVerifier),
	)
	if err != nil {
		return ports.OIDCClaims{}, err
	}

	rawIDToken, _ := token.Extra("id_token").(string)
	if rawIDToken == "" {
		return ports.OIDCClaims{}, errors.New("missing id_token")
	}

	idToken, err := o.verifier.Verify(ctx, rawIDToken)
	if err != nil {
		return ports.OIDCClaims{}, err
	}

	var claims struct {
		Sub           string `json:"sub"`
		Email         string `json:"email"`
		EmailVerified bool   `json:"email_verified"`
		Nonce         string `json:"nonce"`
		AuthTime      int64  `json:"auth_time"`
	}

	if err := idToken.Claims(&claims); err != nil {
		return ports.OIDCClaims{}, err
	}

	if nonce != "" && claims.Nonce != nonce {
		return ports.OIDCClaims{}, errors.New("nonce mismatch")
	}

	authTime := time.Unix(claims.AuthTime, 0)

	return ports.OIDCClaims{
		Provider:      "google",
		Subject:       claims.Sub,
		Email:         claims.Email,
		EmailVerified: claims.EmailVerified,
		AuthTime:      authTime,
	}, nil
}

func (o *OIDC) oauthConfigFor(redirectURL string) oauth2.Config {
	cfg := o.oauth
	if strings.TrimSpace(redirectURL) != "" {
		cfg.RedirectURL = redirectURL
	}
	return cfg
}
