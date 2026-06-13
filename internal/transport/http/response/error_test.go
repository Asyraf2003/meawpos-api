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

package response

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestErrorResponseMapsPublicHTTPError(t *testing.T) {
	statusCode, body := ErrorResponse(NewHTTPError(
		http.StatusNotFound,
		"product_not_found",
		"product not found",
	))

	if statusCode != http.StatusNotFound {
		t.Fatalf("statusCode = %d, want %d", statusCode, http.StatusNotFound)
	}
	if body.Success {
		t.Fatal("success = true, want false")
	}
	if body.Error.Code != "product_not_found" {
		t.Fatalf("error.code = %q, want product_not_found", body.Error.Code)
	}
	if body.Error.Message != "product not found" {
		t.Fatalf("error.message = %q, want product not found", body.Error.Message)
	}

	meta, ok := body.Meta.(map[string]any)
	if !ok {
		t.Fatalf("meta type = %T, want map[string]any", body.Meta)
	}
	if len(meta) != 0 {
		t.Fatalf("meta len = %d, want 0", len(meta))
	}
}

func TestErrorResponseMapsEchoCapabilityDisabled(t *testing.T) {
	statusCode, body := ErrorResponse(echo.NewHTTPError(
		http.StatusForbidden,
		"capability disabled",
	))

	if statusCode != http.StatusForbidden {
		t.Fatalf("statusCode = %d, want %d", statusCode, http.StatusForbidden)
	}
	if body.Error.Code != "capability_disabled" {
		t.Fatalf("error.code = %q, want capability_disabled", body.Error.Code)
	}
	if body.Error.Message != "capability disabled" {
		t.Fatalf("error.message = %q, want capability disabled", body.Error.Message)
	}
}

func TestHTTPErrorHandlerWritesEnvelope(t *testing.T) {
	e := echo.New()
	e.HTTPErrorHandler = HTTPErrorHandler
	e.GET("/boom", func(c echo.Context) error {
		return NewHTTPError(http.StatusConflict, "product_code_already_exists", "product code already exists")
	})

	req := httptest.NewRequest(http.MethodGet, "/boom", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusConflict {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusConflict)
	}

	var body ErrorEnvelope
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if body.Success {
		t.Fatal("success = true, want false")
	}
	if body.Error.Code != "product_code_already_exists" {
		t.Fatalf("error.code = %q, want product_code_already_exists", body.Error.Code)
	}
	if body.Error.Message != "product code already exists" {
		t.Fatalf("error.message = %q, want product code already exists", body.Error.Message)
	}
}
