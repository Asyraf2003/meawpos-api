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
	"errors"
	stdhttp "net/http"
	"testing"

	"pos-go/internal/modules/productcatalog/domain"
	"pos-go/internal/modules/productcatalog/ports"
	httpresponse "pos-go/internal/transport/http/response"
)

func TestMapProductCatalogErrorReturnsPublicErrorCodes(t *testing.T) {
	tests := []struct {
		name       string
		err        error
		wantStatus int
		wantCode   string
		wantMsg    string
	}{
		{
			name:       "not found",
			err:        ports.ErrProductNotFound,
			wantStatus: stdhttp.StatusNotFound,
			wantCode:   "product_not_found",
			wantMsg:    "product not found",
		},
		{
			name:       "duplicate code",
			err:        ports.ErrDuplicateProductCode,
			wantStatus: stdhttp.StatusConflict,
			wantCode:   "product_code_already_exists",
			wantMsg:    "product code already exists",
		},
		{
			name:       "duplicate identity",
			err:        ports.ErrDuplicateProductIdentity,
			wantStatus: stdhttp.StatusConflict,
			wantCode:   "product_identity_already_exists",
			wantMsg:    "product identity already exists",
		},
		{
			name:       "domain validation",
			err:        domain.ErrProductNameRequired,
			wantStatus: stdhttp.StatusBadRequest,
			wantCode:   "product_validation_failed",
			wantMsg:    domain.ErrProductNameRequired.Error(),
		},
		{
			name:       "unknown",
			err:        errors.New("database exploded because apparently computers are fragile"),
			wantStatus: stdhttp.StatusInternalServerError,
			wantCode:   "product_catalog_request_failed",
			wantMsg:    "product catalog request failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			statusCode, body := httpresponse.ErrorResponse(mapProductCatalogError(tt.err))

			if statusCode != tt.wantStatus {
				t.Fatalf("statusCode = %d, want %d", statusCode, tt.wantStatus)
			}
			if body.Success {
				t.Fatal("success = true, want false")
			}
			if body.Error.Code != tt.wantCode {
				t.Fatalf("error.code = %q, want %q", body.Error.Code, tt.wantCode)
			}
			if body.Error.Message != tt.wantMsg {
				t.Fatalf("error.message = %q, want %q", body.Error.Message, tt.wantMsg)
			}

			meta, ok := body.Meta.(map[string]any)
			if !ok {
				t.Fatalf("meta type = %T, want map[string]any", body.Meta)
			}
			if len(meta) != 0 {
				t.Fatalf("meta len = %d, want 0", len(meta))
			}
		})
	}
}
