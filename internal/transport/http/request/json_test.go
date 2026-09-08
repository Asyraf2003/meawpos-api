// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package request

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	httpresponse "pos-go/internal/transport/http/response"

	"github.com/labstack/echo/v4"
)

func TestDecodeJSONStrictBoundary(t *testing.T) {
	tests := []struct {
		name, contentType, body string
		wantStatus              int
	}{
		{name: "valid", contentType: echo.MIMEApplicationJSON, body: `{"name":"root"}`},
		{name: "content type", contentType: echo.MIMETextPlain, body: `{}`, wantStatus: http.StatusUnsupportedMediaType},
		{name: "unknown field", contentType: echo.MIMEApplicationJSON, body: `{"name":"root","admin":true}`, wantStatus: http.StatusBadRequest},
		{name: "trailing value", contentType: echo.MIMEApplicationJSON, body: `{"name":"root"}{}`, wantStatus: http.StatusBadRequest},
		{name: "malformed", contentType: echo.MIMEApplicationJSON, body: `{"name":`, wantStatus: http.StatusBadRequest},
		{name: "too large", contentType: echo.MIMEApplicationJSON, body: `{"name":"` + strings.Repeat("x", int(MaxJSONBodyBytes)) + `"}`, wantStatus: http.StatusRequestEntityTooLarge},
		{name: "too large after value", contentType: echo.MIMEApplicationJSON, body: `{"name":"root"}` + strings.Repeat(" ", int(MaxJSONBodyBytes)), wantStatus: http.StatusRequestEntityTooLarge},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tc.body))
			req.Header.Set(echo.HeaderContentType, tc.contentType)
			ctx := e.NewContext(req, httptest.NewRecorder())
			var value struct {
				Name string `json:"name"`
			}
			err := DecodeJSON(ctx, &value)
			if tc.wantStatus == 0 {
				if err != nil || value.Name != "root" {
					t.Fatalf("DecodeJSON() = name %q, err %v", value.Name, err)
				}
				return
			}
			publicErr, ok := err.(*httpresponse.HTTPError)
			if !ok || publicErr.StatusCode != tc.wantStatus {
				t.Fatalf("DecodeJSON() error = %#v, want status %d", err, tc.wantStatus)
			}
		})
	}
}
