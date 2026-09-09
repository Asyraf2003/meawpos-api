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

package http

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	browserSessionCookieName = "miawpos_session"
	browserSessionCookiePath = "/api/auth/browser"
)

var errBrowserSessionCookieMissing = errors.New("browser session cookie missing")

type browserSessionCookie struct {
	secure bool
}

func newBrowserSessionCookie(secure bool) browserSessionCookie {
	return browserSessionCookie{secure: secure}
}

func (c browserSessionCookie) Set(ctx echo.Context, refreshToken string) {
	ctx.SetCookie(&http.Cookie{
		Name:     browserSessionCookieName,
		Value:    refreshToken,
		Path:     browserSessionCookiePath,
		HttpOnly: true,
		Secure:   c.secure,
		SameSite: http.SameSiteStrictMode,
	})
}

func (c browserSessionCookie) Read(ctx echo.Context) (string, error) {
	cookie, err := ctx.Cookie(browserSessionCookieName)
	if err != nil {
		return "", errBrowserSessionCookieMissing
	}

	refreshToken := strings.TrimSpace(cookie.Value)
	if refreshToken == "" {
		return "", errBrowserSessionCookieMissing
	}

	return refreshToken, nil
}

func (c browserSessionCookie) Clear(ctx echo.Context) {
	ctx.SetCookie(&http.Cookie{
		Name:     browserSessionCookieName,
		Value:    "",
		Path:     browserSessionCookiePath,
		HttpOnly: true,
		Secure:   c.secure,
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Unix(1, 0).UTC(),
		MaxAge:   -1,
	})
}
