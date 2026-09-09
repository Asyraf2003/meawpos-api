// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestBrowserGoogleCallback_FailsClosed(t *testing.T) {
	for _, tc := range []struct {
		name, cookie, query string
		flowErr             error
		called              bool
	}{
		{name: "missing cookie", query: "state=s&code=c"},
		{name: "mismatched state", cookie: "other", query: "state=s&code=c"},
		{name: "missing state", cookie: "s", query: "code=c"},
		{name: "provider denied", cookie: "s", query: "state=s&error=access_denied"},
		{name: "invalid code", cookie: "s", query: "state=s&code=bad", flowErr: errors.New("sensitive provider error"), called: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			flow := &fakeGoogleFlow{callbackErr: tc.flowErr}
			h := browserGoogleTestHandler(t, flow, true)
			req := httptest.NewRequest(http.MethodGet, "https://ui.example"+browserGoogleCallbackPath+"?"+tc.query, nil)
			if tc.cookie != "" {
				req.AddCookie(&http.Cookie{Name: browserGoogleStateCookieName, Value: tc.cookie})
			}
			rec := httptest.NewRecorder()
			if err := h.Callback(echo.New().NewContext(req, rec)); err != nil {
				t.Fatal(err)
			}
			if rec.Header().Get("Location") != "/login?auth=failed" || rec.Body.Len() != 0 {
				t.Fatal("unsafe error redirect")
			}
			if (flow.callbackInput.State != "") != tc.called {
				t.Fatal("unexpected GoogleFlow execution")
			}
			for _, cookie := range rec.Result().Cookies() {
				if cookie.Name == browserSessionCookieName {
					t.Fatal("failed callback created session cookie")
				}
			}
		})
	}
}

func TestBrowserGoogleConfig_RejectsUnsafeOrigins(t *testing.T) {
	for _, raw := range []string{"", "/relative", "//evil.example", "http://ui.example/callback", "https://user@ui.example/callback", "https://ui.example/callback?next=bad", "https://ui.example/#fragment", "javascript:alert(1)"} {
		if _, err := browserGoogleCallbackURL(raw, true); err == nil {
			t.Errorf("accepted invalid origin %q", raw)
		}
	}
	if got, err := browserGoogleCallbackURL("http://localhost:5173/api/auth/google/callback", false); err != nil || !strings.HasPrefix(got, "http://localhost:5173/") {
		t.Fatal("local HTTP unsupported")
	}
	h := browserGoogleTestHandler(t, &fakeGoogleFlow{}, true)
	for _, handle := range []echo.HandlerFunc{h.Start, h.Callback} {
		rec := httptest.NewRecorder()
		err := handle(echo.New().NewContext(httptest.NewRequest(http.MethodGet, "https://evil.example/", nil), rec))
		assertHTTPErrorStatus(t, err, http.StatusBadRequest)
		if len(rec.Result().Cookies()) != 0 {
			t.Fatal("foreign host received cookie")
		}
	}
}
