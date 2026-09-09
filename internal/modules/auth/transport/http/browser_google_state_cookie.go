// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

const browserGoogleStateCookieName = "miawpos_google_state"

type browserGoogleStateCookie struct {
	secure bool
	ttl    time.Duration
}

func (s browserGoogleStateCookie) Set(c echo.Context, state string) {
	age := max(1, int(s.ttl.Seconds()))
	s.write(c, state, age)
}

func (s browserGoogleStateCookie) Clear(c echo.Context) { s.write(c, "", -1) }

func (s browserGoogleStateCookie) write(c echo.Context, value string, age int) {
	c.SetCookie(&http.Cookie{ // #nosec G124 -- Secure is true outside accepted local HTTP mode, matching the session helper.
		Name: browserGoogleStateCookieName, Value: value, Path: browserGoogleCallbackPath,
		HttpOnly: true, Secure: s.secure, SameSite: http.SameSiteLaxMode, MaxAge: age,
	})
}
