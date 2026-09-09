// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
)

type fakeBrowserSessionRevoker struct {
	sessionID string
	err       error
}

func (f *fakeBrowserSessionRevoker) RevokeSession(ctx context.Context, sessionID string) error {
	_ = ctx
	f.sessionID = sessionID
	return f.err
}

func assertHTTPErrorStatus(t *testing.T, err error, want int) {
	t.Helper()
	httpErr, ok := err.(*echo.HTTPError)
	if !ok {
		t.Fatalf("error = %T %v, want *echo.HTTPError", err, err)
	}
	if httpErr.Code != want {
		t.Fatalf("HTTP status = %d, want %d", httpErr.Code, want)
	}
}

func assertBrowserSessionCookie(
	t *testing.T,
	rec *httptest.ResponseRecorder,
	wantValue string,
	secure bool,
	deleted bool,
) {
	t.Helper()
	cookies := rec.Result().Cookies()
	if len(cookies) != 1 {
		t.Fatalf("cookies = %d, want 1", len(cookies))
	}
	cookie := cookies[0]
	if cookie.Name != browserSessionCookieName || cookie.Value != wantValue {
		t.Fatalf("cookie identity/value = %+v", cookie)
	}
	if cookie.Path != browserSessionCookiePath || cookie.Domain != "" {
		t.Fatalf("cookie scope = %+v", cookie)
	}
	if !cookie.HttpOnly || cookie.Secure != secure || cookie.SameSite != http.SameSiteStrictMode {
		t.Fatalf("cookie security = %+v", cookie)
	}
	if deleted {
		if cookie.MaxAge >= 0 || cookie.Expires.IsZero() || !cookie.Expires.Before(time.Now()) {
			t.Fatalf("cookie deletion = %+v", cookie)
		}
		return
	}
	if cookie.MaxAge != 0 || !cookie.Expires.IsZero() {
		t.Fatalf("session cookie lifetime = %+v", cookie)
	}
}
