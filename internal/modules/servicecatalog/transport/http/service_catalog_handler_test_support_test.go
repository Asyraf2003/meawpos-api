package http

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	uc "pos-go/internal/modules/servicecatalog/usecase"

	"github.com/labstack/echo/v4"
)

func newServiceCatalogTestContext(method, target, body string) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(method, target, strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	return e.NewContext(req, rec), rec
}

func testServiceCatalogItemResult() uc.ServiceCatalogItemResult {
	now := time.Date(2026, 6, 8, 10, 30, 0, 0, time.UTC)

	return uc.ServiceCatalogItemResult{
		ID:                 "svc_1",
		Name:               "Express Wash",
		NormalizedName:     "express wash",
		DefaultPriceRupiah: 15000,
		IsActive:           true,
		Status:             "active",
		CreatedAt:          now,
		UpdatedAt:          now,
	}
}

func assertJSONStatus(t *testing.T, rec *httptest.ResponseRecorder, want int) {
	t.Helper()

	if rec.Code != want {
		t.Fatalf("status = %d, want %d, body = %s", rec.Code, want, rec.Body.String())
	}
	if got := rec.Header().Get(echo.HeaderContentType); !strings.Contains(got, echo.MIMEApplicationJSON) {
		t.Fatalf("content-type = %q, want JSON", got)
	}
}

func assertSuccessEnvelope(t *testing.T, rec *httptest.ResponseRecorder) {
	t.Helper()

	var body struct {
		Success bool            `json:"success"`
		Data    json.RawMessage `json:"data"`
		Meta    json.RawMessage `json:"meta"`
	}

	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("response is not JSON: %v; body = %s", err, rec.Body.String())
	}
	if !body.Success {
		t.Fatalf("success = false, want true; body = %s", rec.Body.String())
	}
	if len(body.Data) == 0 || string(body.Data) == "null" {
		t.Fatalf("data is empty; body = %s", rec.Body.String())
	}
	if len(body.Meta) == 0 || string(body.Meta) == "null" {
		t.Fatalf("meta is empty; body = %s", rec.Body.String())
	}
}

func assertHTTPErrorCode(t *testing.T, err error, want int) {
	t.Helper()

	if err == nil {
		t.Fatalf("error = nil, want HTTP %d", want)
	}

	httpErr, ok := err.(*echo.HTTPError)
	if !ok {
		t.Fatalf("error type = %T, want *echo.HTTPError", err)
	}
	if httpErr.Code != want {
		t.Fatalf("status = %d, want %d", httpErr.Code, want)
	}
}
