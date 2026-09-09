// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package http

import (
	"crypto/subtle"
	"net/http"
	"net/url"
	"time"

	"github.com/labstack/echo/v4"
	"pos-go/internal/modules/auth/usecase"
)

type BrowserGoogleHandler struct {
	flow        GoogleFlow
	callbackURL string
	host        string
	cookie      browserSessionCookie
	stateCookie browserGoogleStateCookie
}

func NewBrowserGoogleHandler(flow GoogleFlow, configuredRedirect string, secure bool, stateTTL time.Duration) (*BrowserGoogleHandler, error) {
	callback, err := browserGoogleCallbackURL(configuredRedirect, secure)
	if err != nil {
		return nil, err
	}
	parsed, _ := url.Parse(callback)
	return &BrowserGoogleHandler{flow: flow, callbackURL: callback, host: parsed.Host,
		cookie: newBrowserSessionCookie(secure), stateCookie: browserGoogleStateCookie{secure: secure, ttl: stateTTL}}, nil
}

func (h *BrowserGoogleHandler) Register(group *echo.Group) {
	group.GET("/browser/google/start", h.Start)
	group.GET("/browser/google/callback", h.Callback)
}

func (h *BrowserGoogleHandler) Start(c echo.Context) error {
	browserGoogleHeaders(c)
	if c.Request().Host != h.host {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid browser origin")
	}
	out, err := h.flow.GoogleStart(c, usecase.GoogleStartInput{Purpose: "login", RedirectURL: h.callbackURL})
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/login?auth=failed")
	}
	h.stateCookie.Set(c, out.State)
	return c.Redirect(http.StatusSeeOther, out.RedirectTo)
}

func (h *BrowserGoogleHandler) Callback(c echo.Context) error {
	browserGoogleHeaders(c)
	if c.Request().Host != h.host {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid browser origin")
	}
	cookie, err := c.Cookie(browserGoogleStateCookieName)
	state := c.QueryParam("state")
	if err != nil || state == "" || subtle.ConstantTimeCompare([]byte(cookie.Value), []byte(state)) != 1 {
		return c.Redirect(http.StatusSeeOther, "/login?auth=failed")
	}
	h.stateCookie.Clear(c)
	if c.QueryParam("error") != "" {
		return c.Redirect(http.StatusSeeOther, "/login?auth=failed")
	}
	out, err := h.flow.GoogleCallback(c, usecase.GoogleCallbackInput{
		Code: c.QueryParam("code"), State: state, RedirectURL: h.callbackURL,
		Client: usecase.ClientInfo{UserAgent: c.Request().UserAgent(), IP: c.RealIP()},
	})
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/login?auth=failed")
	}
	h.cookie.Set(c, out.RefreshToken)
	// The UI's existing bootstrap refresh obtains its memory-only bearer.
	// Neither credential is serialized in the callback body or redirect URL.
	return c.Redirect(http.StatusSeeOther, "/account")
}

func browserGoogleHeaders(c echo.Context) {
	c.Response().Header().Set("Cache-Control", "no-store")
	c.Response().Header().Set("Referrer-Policy", "no-referrer")
}
