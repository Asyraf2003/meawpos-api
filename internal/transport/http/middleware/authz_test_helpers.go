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

package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"pos-go/internal/modules/auth/domain"

	"github.com/labstack/echo/v4"
)

func newAuthzTestContext(principal *domain.Principal) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	if principal != nil {
		req = req.WithContext(WithPrincipal(req.Context(), *principal))
	}

	rec := httptest.NewRecorder()
	return e.NewContext(req, rec), rec
}

func assertHTTPErrorCode(t *testing.T, err error, want int) {
	t.Helper()

	if err == nil {
		t.Fatal("handler() error = nil, want error")
	}

	httpErr, ok := err.(*echo.HTTPError)
	if !ok {
		t.Fatalf("error type = %T, want *echo.HTTPError", err)
	}
	if httpErr.Code != want {
		t.Fatalf("status = %d, want %d", httpErr.Code, want)
	}
}
