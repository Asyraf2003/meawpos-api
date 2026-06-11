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
	"context"
	"errors"
	"fmt"
	"time"

	"pos-go/internal/modules/auth/ports"
)

func (i *Issuer) IssueAccessToken(ctx context.Context, req ports.AccessTokenRequest) (string, time.Time, error) {
	_ = ctx

	if req.AccountID == "" || req.SessionID == "" {
		return "", time.Time{}, errors.New("missing account/session")
	}

	now := time.Now()
	exp := now.Add(i.ttl)

	hs, err := encodeJSON(header{
		Alg: "HS256",
		Typ: "JWT",
		Kid: i.kid,
	})
	if err != nil {
		return "", time.Time{}, err
	}

	ps, err := encodeJSON(payload{
		Iss: i.issuer,
		Aud: i.aud,
		Sub: req.AccountID,
		Sid: req.SessionID,
		AAL: req.TrustLevel,
		IAT: now.Unix(),
		EXP: exp.Unix(),
	})
	if err != nil {
		return "", time.Time{}, err
	}

	input := hs + "." + ps
	sig := signHS256(i.secret, input)

	return fmt.Sprintf("%s.%s", input, sig), exp, nil
}
