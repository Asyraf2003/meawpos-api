// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	authusecase "pos-go/internal/modules/auth/usecase"

	"github.com/labstack/echo/v4"
)

func TestBrowserManualLoginHandler_SetsCookieAndHidesRefreshToken(t *testing.T) {
	sessionExp := time.Unix(1779277000, 0).UTC()
	usecase := &fakeManualLoginUsecase{output: authusecase.ManualLoginOutput{
		AccessToken: "access-token", AccessExp: time.Unix(1776685000, 0).UTC(),
		RefreshToken: "refresh-token", RefreshExp: sessionExp, TrustLevel: "aal1",
	}}
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/auth/browser/manual/login",
		strings.NewReader(`{"email":"admin@example.com","password":"12345678"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	if err := NewBrowserManualLoginHandler(usecase, false).Login(e.NewContext(req, rec)); err != nil {
		t.Fatalf("Login() error = %v", err)
	}
	assertBrowserSessionCookie(t, rec, "refresh-token", false, false)

	var body map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatal(err)
	}
	if _, exists := body["refresh_token"]; exists {
		t.Fatalf("browser response exposed refresh_token: %v", body)
	}
	if body["access_token"] != "access-token" || body["session_exp"] != sessionExp.Format(time.RFC3339) {
		t.Fatalf("browser login body=%v", body)
	}
}

func TestBrowserManualLoginHandler_InvalidCredentialsDoesNotSetCookie(t *testing.T) {
	usecase := &fakeManualLoginUsecase{err: authusecase.ErrManualLoginInvalidCredentials}
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/auth/browser/manual/login",
		strings.NewReader(`{"email":"admin@example.com","password":"wrong"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	err := NewBrowserManualLoginHandler(usecase, false).Login(e.NewContext(req, rec))
	assertHTTPErrorStatus(t, err, http.StatusUnauthorized)
	if got := rec.Header().Get(echo.HeaderSetCookie); got != "" {
		t.Fatalf("Set-Cookie = %q, want empty", got)
	}
}
