// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

//go:build integration

package bootstrap

import (
	"net/http"
	"testing"
)

func TestWalkingSkeletonHTTP_BrowserLogoutPreservesAuthorityAndRevokesSession(t *testing.T) {
	app := newR4HTTPApp(t)
	login := map[string]any{"email": "admin@example.com", "password": "1234" + "5678"}
	status, loginBody, loginCookie := browserSessionRequest(t, app, http.MethodPost,
		"/api/auth/browser/manual/login", "", nil, login)
	if status != http.StatusOK {
		t.Fatalf("browser login status=%d body=%v", status, loginBody)
	}

	status, unauthorizedBody, unauthorizedCookie := browserSessionRequest(t, app, http.MethodPost,
		"/api/auth/browser/logout", "", loginCookie, nil)
	if status != http.StatusUnauthorized {
		t.Fatalf("unauthorized logout status=%d body=%v", status, unauthorizedBody)
	}
	if unauthorizedCookie != nil {
		t.Fatalf("unauthorized logout changed cookie=%+v", unauthorizedCookie)
	}

	status, refreshBody, refreshCookie := browserSessionRequest(t, app, http.MethodPost,
		"/api/auth/browser/refresh", "", loginCookie, nil)
	if status != http.StatusOK {
		t.Fatalf("session revoked without authority status=%d body=%v", status, refreshBody)
	}
	accessToken, _ := refreshBody["access_token"].(string)
	if accessToken == "" {
		t.Fatalf("refresh body=%v", refreshBody)
	}

	status, logoutBody, clearCookie := browserSessionRequest(t, app, http.MethodPost,
		"/api/auth/browser/logout", accessToken, refreshCookie, nil)
	if status != http.StatusNoContent {
		t.Fatalf("authorized logout status=%d body=%v", status, logoutBody)
	}
	assertDeletedBrowserCookie(t, clearCookie)

	status, revokedBody, revokedCookie := browserSessionRequest(t, app, http.MethodPost,
		"/api/auth/browser/refresh", "", refreshCookie, nil)
	if status != http.StatusUnauthorized {
		t.Fatalf("revoked refresh status=%d body=%v", status, revokedBody)
	}
	assertDeletedBrowserCookie(t, revokedCookie)
}
