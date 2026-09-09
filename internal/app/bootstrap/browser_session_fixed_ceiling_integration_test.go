// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

//go:build integration

package bootstrap

import (
	"net/http"
	"testing"
)

func TestWalkingSkeletonHTTP_BrowserSessionFixedCeilingAcrossRefreshes(t *testing.T) {
	app := newR4HTTPApp(t)
	login := map[string]any{"email": "admin@example.com", "password": "1234" + "5678"}
	status, loginBody, loginCookie := browserSessionRequest(t, app, http.MethodPost,
		"/api/auth/browser/manual/login", "", nil, login)
	if status != http.StatusOK {
		t.Fatalf("browser login status=%d body=%v", status, loginBody)
	}
	assertBrowserResponseHasNoRefreshToken(t, loginBody)
	assertBrowserCookieSecurity(t, loginCookie)
	originalExpiry := browserSessionExpiryByRefreshToken(t, app, loginCookie.Value)

	status, firstBody, firstCookie := browserSessionRequest(t, app, http.MethodPost,
		"/api/auth/browser/refresh", "", loginCookie, nil)
	if status != http.StatusOK {
		t.Fatalf("first refresh status=%d body=%v", status, firstBody)
	}
	assertBrowserResponseHasNoRefreshToken(t, firstBody)
	assertBrowserCookieSecurity(t, firstCookie)
	if firstCookie.Value == loginCookie.Value {
		t.Fatal("first refresh cookie was not rotated")
	}
	assertSameInstant(t, "first response expiry", parseBrowserSessionExpiry(t, firstBody), originalExpiry)
	assertSameInstant(t, "first stored expiry", browserSessionExpiryByRefreshToken(t, app, firstCookie.Value), originalExpiry)

	status, staleBody, staleCookie := browserSessionRequest(t, app, http.MethodPost,
		"/api/auth/browser/refresh", "", loginCookie, nil)
	if status != http.StatusUnauthorized {
		t.Fatalf("stale cookie status=%d body=%v", status, staleBody)
	}
	assertDeletedBrowserCookie(t, staleCookie)

	status, secondBody, secondCookie := browserSessionRequest(t, app, http.MethodPost,
		"/api/auth/browser/refresh", "", firstCookie, nil)
	if status != http.StatusOK {
		t.Fatalf("second refresh status=%d body=%v", status, secondBody)
	}
	assertBrowserResponseHasNoRefreshToken(t, secondBody)
	if secondCookie == nil || secondCookie.Value == firstCookie.Value {
		t.Fatal("second refresh cookie was not rotated")
	}
	assertSameInstant(t, "second response expiry", parseBrowserSessionExpiry(t, secondBody), originalExpiry)
	assertSameInstant(t, "second stored expiry", browserSessionExpiryByRefreshToken(t, app, secondCookie.Value), originalExpiry)
}
