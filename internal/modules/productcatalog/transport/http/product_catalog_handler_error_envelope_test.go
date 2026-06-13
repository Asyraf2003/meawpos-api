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
	"context"
	"encoding/json"
	stdhttp "net/http"
	"net/http/httptest"
	"testing"

	"pos-go/internal/modules/productcatalog/ports"
	productcatalogusecase "pos-go/internal/modules/productcatalog/usecase"
	httpresponse "pos-go/internal/transport/http/response"

	"github.com/labstack/echo/v4"
)

type getProductDetailFunc func(context.Context, productcatalogusecase.GetProductDetailQuery) (productcatalogusecase.GetProductDetailResult, error)

func (f getProductDetailFunc) Execute(
	ctx context.Context,
	query productcatalogusecase.GetProductDetailQuery,
) (productcatalogusecase.GetProductDetailResult, error) {
	return f(ctx, query)
}

func TestProductCatalogShowWritesErrorEnvelope(t *testing.T) {
	e := echo.New()
	e.HTTPErrorHandler = httpresponse.HTTPErrorHandler

	handler := NewProductCatalogHandler(
		nil,
		nil,
		getProductDetailFunc(func(
			ctx context.Context,
			query productcatalogusecase.GetProductDetailQuery,
		) (productcatalogusecase.GetProductDetailResult, error) {
			_ = ctx

			if query.ID != "product-1" {
				t.Fatalf("query.ID = %q, want product-1", query.ID)
			}

			return productcatalogusecase.GetProductDetailResult{}, ports.ErrProductNotFound
		}),
		nil,
		nil,
		nil,
		nil,
		nil,
	)

	handler.RegisterShow(e.Group("/api"))

	req := httptest.NewRequest(stdhttp.MethodGet, "/api/products/product-1", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	if rec.Code != stdhttp.StatusNotFound {
		t.Fatalf("status = %d, want %d", rec.Code, stdhttp.StatusNotFound)
	}

	var body httpresponse.ErrorEnvelope
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode response: %v", err)
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
