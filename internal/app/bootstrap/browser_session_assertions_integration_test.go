// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

//go:build integration

package bootstrap

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"testing"
	"time"
)

func browserSessionExpiryByRefreshToken(t *testing.T, app *App, refreshToken string) time.Time {
	t.Helper()
	sum := sha256.Sum256([]byte(refreshToken))
	var expiresAt time.Time
	if err := app.DB.QueryRow(t.Context(),
		`SELECT expires_at FROM auth_sessions WHERE refresh_token_hash = $1`,
		hex.EncodeToString(sum[:]),
	).Scan(&expiresAt); err != nil {
		t.Fatalf("query session expiry: %v", err)
	}
	return expiresAt
}

func parseBrowserSessionExpiry(t *testing.T, body map[string]any) time.Time {
	t.Helper()
	raw, _ := body["session_exp"].(string)
	if raw == "" {
		t.Fatalf("session_exp missing from body=%v", body)
	}
	expiresAt, err := time.Parse(time.RFC3339Nano, raw)
	if err != nil {
		t.Fatalf("parse session_exp %q: %v", raw, err)
	}
	return expiresAt
}

func assertBrowserResponseHasNoRefreshToken(t *testing.T, body map[string]any) {
	t.Helper()
	if _, exists := body["refresh_token"]; exists {
		t.Fatalf("browser response exposed refresh_token: %v", body)
	}
}

func assertBrowserCookieSecurity(t *testing.T, cookie *http.Cookie) {
	t.Helper()
	if cookie == nil || cookie.Name != "miawpos_session" || cookie.Path != "/api/auth/browser" {
		t.Fatalf("browser cookie identity/scope=%+v", cookie)
	}
	if !cookie.HttpOnly || !cookie.Secure || cookie.SameSite != http.SameSiteStrictMode || cookie.Domain != "" {
		t.Fatalf("browser cookie security=%+v", cookie)
	}
}

func assertDeletedBrowserCookie(t *testing.T, cookie *http.Cookie) {
	t.Helper()
	assertBrowserCookieSecurity(t, cookie)
	if cookie.MaxAge >= 0 || cookie.Expires.IsZero() || !cookie.Expires.Before(time.Now()) {
		t.Fatalf("deleted cookie is not expired=%+v", cookie)
	}
}

func assertSameInstant(t *testing.T, label string, got, want time.Time) {
	t.Helper()
	if !got.Equal(want) {
		t.Fatalf("%s=%s, want %s", label, got, want)
	}
}
