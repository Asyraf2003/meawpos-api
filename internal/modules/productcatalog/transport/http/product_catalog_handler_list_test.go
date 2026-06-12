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

	"github.com/labstack/echo/v4"
)

func TestProductCatalogHandlerListParsesQuery(t *testing.T) {
	list := &fakeListProducts{}
	handler := NewProductCatalogHandler(list, nil, nil, nil, nil, nil, nil, nil)

	req := httptest.NewRequest(http.MethodGet, "/products?q=kampas&status=deleted&page=2&per_page=10", nil)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)

	if err := handler.List(c); err != nil {
		t.Fatalf("List() error = %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	if list.got.Search != "kampas" || list.got.Status != "deleted" {
		t.Fatalf("query identity = %+v", list.got)
	}
	if list.got.Page != 2 || list.got.PerPage != 10 {
		t.Fatalf("pagination = %d/%d", list.got.Page, list.got.PerPage)
	}
}
