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
	"net/http"
	"net/http/httptest"
	"testing"

	"pos-go/internal/modules/productcatalog/ports"

	"github.com/labstack/echo/v4"
)

func TestProductCatalogHandlerShowMapsNotFound(t *testing.T) {
	show := &fakeGetProductDetail{err: ports.ErrProductNotFound}
	handler := NewProductCatalogHandler(nil, nil, show, nil, nil, nil, nil, nil)

	req := httptest.NewRequest(http.MethodGet, "/products/missing", nil)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("missing")

	err := handler.Show(c)
	httpErr, ok := err.(*echo.HTTPError)
	if !ok {
		t.Fatalf("Show() error = %T, want *echo.HTTPError", err)
	}
	if httpErr.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want %d", httpErr.Code, http.StatusNotFound)
	}
}
