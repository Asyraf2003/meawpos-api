// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	authdomain "pos-go/internal/modules/auth/domain"
	authusecase "pos-go/internal/modules/auth/usecase"
	httpmw "pos-go/internal/transport/http/middleware"

	"github.com/labstack/echo/v4"
)

func TestBrowserLogoutHandler_RevokesBeforeClearingCookie(t *testing.T) {
	revoker := &fakeBrowserSessionRevoker{}
	e := echo.New()
	req := browserLogoutRequest(t)
	rec := httptest.NewRecorder()

	handler := NewBrowserLogoutHandler(authusecase.NewLogoutCurrentSession(revoker), true)
	if err := handler.Logout(e.NewContext(req, rec)); err != nil {
		t.Fatalf("Logout() error = %v", err)
	}
	if rec.Code != http.StatusNoContent || revoker.sessionID != "sess-123" {
		t.Fatalf("status=%d revoked=%q", rec.Code, revoker.sessionID)
	}
	assertBrowserSessionCookie(t, rec, "", true, true)
}

func TestBrowserLogoutHandler_RevocationFailureDoesNotClearCookie(t *testing.T) {
	revoker := &fakeBrowserSessionRevoker{err: errors.New("revoke failed")}
	e := echo.New()
	rec := httptest.NewRecorder()
	handler := NewBrowserLogoutHandler(authusecase.NewLogoutCurrentSession(revoker), false)

	if err := handler.Logout(e.NewContext(browserLogoutRequest(t), rec)); err == nil {
		t.Fatal("Logout() error = nil, want error")
	}
	if got := rec.Header().Get(echo.HeaderSetCookie); got != "" {
		t.Fatalf("Set-Cookie = %q, want empty", got)
	}
}

func browserLogoutRequest(t *testing.T) *http.Request {
	t.Helper()
	req := httptest.NewRequest(http.MethodPost, "/api/auth/browser/logout", nil)
	return req.WithContext(httpmw.WithPrincipal(req.Context(), authdomain.Principal{
		AccountID: "acc-123", SessionID: "sess-123", TrustLevel: "aal1",
	}))
}
