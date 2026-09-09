// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	authusecase "pos-go/internal/modules/auth/usecase"

	"github.com/labstack/echo/v4"
)

func TestBrowserRefreshHandler_MissingOrInvalidCredentialClearsCookie(t *testing.T) {
	tests := []struct {
		name       string
		addCookie  bool
		usecaseErr error
	}{
		{name: "missing cookie"},
		{name: "invalid token", addCookie: true, usecaseErr: authusecase.ErrInvalidRefreshToken},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := &fakeRefreshTokenUsecase{err: tt.usecaseErr}
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/api/auth/browser/refresh", nil)
			if tt.addCookie {
				req.AddCookie(&http.Cookie{Name: browserSessionCookieName, Value: "dead-refresh-token"})
			}
			rec := httptest.NewRecorder()

			err := NewBrowserRefreshHandler(usecase, false).Refresh(e.NewContext(req, rec))
			assertHTTPErrorStatus(t, err, http.StatusUnauthorized)
			assertBrowserSessionCookie(t, rec, "", false, true)
		})
	}
}
