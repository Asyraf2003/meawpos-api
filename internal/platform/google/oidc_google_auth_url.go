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
	"pos-go/internal/modules/auth/ports"

	"golang.org/x/oauth2"
)

func (o *OIDC) AuthCodeURL(p ports.OIDCAuthURLParams) string {
	cfg := o.oauthConfigFor(p.RedirectURL)

	opts := []oauth2.AuthCodeOption{
		oauth2.SetAuthURLParam("nonce", p.Nonce),
		oauth2.SetAuthURLParam("code_challenge", p.CodeChallenge),
		oauth2.SetAuthURLParam("code_challenge_method", "S256"),
		oauth2.SetAuthURLParam("prompt", "select_account"),
	}

	return cfg.AuthCodeURL(p.State, opts...)
}
