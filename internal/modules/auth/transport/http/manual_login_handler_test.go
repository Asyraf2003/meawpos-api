// Copyright (C) 2026 Asyraf Mubarak
//
// This file is part of gopos-api.
//
// gopos-api is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, version 3 only.
//
// gopos-api is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with gopos-api. If not, see <https://www.gnu.org/licenses/>.

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

func TestManualLoginHandler_Success(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(
		http.MethodPost,
		"/api/auth/manual/login",
		strings.NewReader(`{"email":"admin@example.com","password":"12345678"}`),
	)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	usecase := &fakeManualLoginUsecase{
		output: authusecase.ManualLoginOutput{
			AccessToken:    "access-token",
			AccessExp:      time.Unix(1776685000, 0),
			RefreshToken:   "refresh-token",
			RefreshExp:     time.Unix(1779277000, 0),
			TrustLevel:     "aal1",
			StepUpRequired: false,
		},
	}
	handler := NewManualLoginHandler(usecase)

	if err := handler.Login(c); err != nil {
		t.Fatalf("Login() error = %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if usecase.lastInput.Email != "admin@example.com" {
		t.Fatalf("email input = %q", usecase.lastInput.Email)
	}
	if usecase.lastInput.Password != "12345678" {
		t.Fatalf("password input = %q", usecase.lastInput.Password)
	}

	var body authusecase.ManualLoginOutput
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	if body.AccessToken != "access-token" {
		t.Fatalf("access token = %q", body.AccessToken)
	}
}
