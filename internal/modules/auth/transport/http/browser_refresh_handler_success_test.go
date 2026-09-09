// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	authusecase "pos-go/internal/modules/auth/usecase"

	"github.com/labstack/echo/v4"
)

func TestBrowserRefreshHandler_ReadsCookieRotatesItAndHidesRefreshToken(t *testing.T) {
	sessionExp := time.Unix(1779277000, 0).UTC()
	usecase := &fakeRefreshTokenUsecase{output: authusecase.RefreshTokenOutput{
		AccessToken: "new-access-token", AccessExp: time.Unix(1776685000, 0).UTC(),
		RefreshToken: "new-refresh-token", RefreshExp: sessionExp, TrustLevel: "aal1",
	}}
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/auth/browser/refresh", nil)
	req.AddCookie(&http.Cookie{Name: browserSessionCookieName, Value: "old-refresh-token"})
	rec := httptest.NewRecorder()

	if err := NewBrowserRefreshHandler(usecase, true).Refresh(e.NewContext(req, rec)); err != nil {
		t.Fatalf("Refresh() error = %v", err)
	}
	if usecase.lastInput.RefreshToken != "old-refresh-token" {
		t.Fatalf("refresh input = %q", usecase.lastInput.RefreshToken)
	}
	assertBrowserSessionCookie(t, rec, "new-refresh-token", true, false)

	var body map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatal(err)
	}
	if _, exists := body["refresh_token"]; exists {
		t.Fatalf("browser response exposed refresh_token: %v", body)
	}
	if body["session_exp"] != sessionExp.Format(time.RFC3339) {
		t.Fatalf("browser refresh body=%v", body)
	}
}
