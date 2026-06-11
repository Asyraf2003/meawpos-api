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

package jwt

import (
	"errors"
	"strings"
	"time"

	"pos-go/internal/modules/auth/ports"
)

type Issuer struct {
	issuer string
	aud    string
	kid    string
	ttl    time.Duration
	secret []byte
}

type header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
	Kid string `json:"kid,omitempty"`
}

type payload struct {
	Iss string `json:"iss"`
	Aud string `json:"aud"`
	Sub string `json:"sub"`
	Sid string `json:"sid"`
	AAL string `json:"aal"`
	IAT int64  `json:"iat"`
	EXP int64  `json:"exp"`
}

func NewHMACIssuer(issuer, aud, kid, secret string, ttl time.Duration) (*Issuer, error) {
	if strings.TrimSpace(secret) == "" {
		return nil, errors.New("jwt secret empty")
	}
	if ttl <= 0 {
		return nil, errors.New("jwt ttl invalid")
	}

	return &Issuer{
		issuer: issuer,
		aud:    aud,
		kid:    kid,
		ttl:    ttl,
		secret: []byte(secret),
	}, nil
}

var _ ports.TokenIssuer = (*Issuer)(nil)
