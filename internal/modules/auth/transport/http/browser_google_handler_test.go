// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"pos-go/internal/modules/auth/usecase"
)

func browserGoogleTestHandler(t *testing.T, flow *fakeGoogleFlow, secure bool) *BrowserGoogleHandler {
	t.Helper()
	scheme := "https"
	if !secure {
		scheme = "http"
	}
	h, err := NewBrowserGoogleHandler(flow, scheme+"://ui.example/api/auth/google/callback", secure, 10*time.Minute)
	if err != nil {
		t.Fatal(err)
	}
	return h
}

func TestBrowserGoogleStart_UsesConfiguredOriginAndBindsBrowser(t *testing.T) {
	flow := &fakeGoogleFlow{startOutput: usecase.GoogleStartOutput{State: "random-state", RedirectTo: "https://accounts.google.com/authorize"}}
	h := browserGoogleTestHandler(t, flow, true)
	req := httptest.NewRequest(http.MethodGet, "https://ui.example/api/auth/browser/google/start?redirect_url=https://evil.example&purpose=other", nil)
	req.Header.Set("X-Forwarded-Host", "evil.example")
	rec := httptest.NewRecorder()
	if err := h.Start(echo.New().NewContext(req, rec)); err != nil {
		t.Fatal(err)
	}
	if flow.startInput.RedirectURL != "https://ui.example"+browserGoogleCallbackPath || flow.startInput.Purpose != "login" {
		t.Fatal("request overrode configured authority")
	}
	if rec.Code != 303 || rec.Header().Get("Location") != flow.startOutput.RedirectTo {
		t.Fatal("missing provider redirect")
	}
	cookie := rec.Result().Cookies()[0]
	if cookie.Value != flow.startOutput.State || cookie.Name != browserGoogleStateCookieName || !cookie.HttpOnly || !cookie.Secure || cookie.SameSite != http.SameSiteLaxMode || cookie.MaxAge != 600 || cookie.Path != browserGoogleCallbackPath || cookie.Domain != "" {
		t.Fatal("invalid state cookie posture")
	}
	if rec.Header().Get("Cache-Control") != "no-store" || rec.Header().Get("Referrer-Policy") != "no-referrer" {
		t.Fatal("missing privacy headers")
	}
}

func TestBrowserGoogleCallback_SetsOnlySessionCookieAndFixedRedirect(t *testing.T) {
	for _, secure := range []bool{true, false} {
		flow := &fakeGoogleFlow{callbackOut: usecase.GoogleCallbackOutput{AccessToken: "private-access", RefreshToken: "private-refresh"}}
		h := browserGoogleTestHandler(t, flow, secure)
		req := httptest.NewRequest(http.MethodGet, "https://ui.example"+browserGoogleCallbackPath+"?state=bound-state&code=valid&return_to=//evil.example", nil)
		req.AddCookie(&http.Cookie{Name: browserGoogleStateCookieName, Value: "bound-state"})
		rec := httptest.NewRecorder()
		if err := h.Callback(echo.New().NewContext(req, rec)); err != nil {
			t.Fatal(err)
		}
		if rec.Code != 303 || rec.Header().Get("Location") != "/account" || rec.Body.Len() != 0 {
			t.Fatal("callback must redirect without a body")
		}
		cookies := rec.Result().Cookies()
		if len(cookies) != 2 || cookies[0].MaxAge != -1 {
			t.Fatal("state cookie not consumed")
		}
		session := cookies[1]
		if session.Name != browserSessionCookieName || session.Value != "private-refresh" || !session.HttpOnly || session.Secure != secure || session.SameSite != http.SameSiteStrictMode || session.Path != browserSessionCookiePath || session.Domain != "" || session.MaxAge != 0 || !session.Expires.IsZero() {
			t.Fatal("session cookie posture changed")
		}
		if flow.callbackInput.State != "bound-state" || flow.callbackInput.Code != "valid" || flow.callbackInput.RedirectURL != h.callbackURL {
			t.Fatal("GoogleFlow input mismatch")
		}
	}
}
