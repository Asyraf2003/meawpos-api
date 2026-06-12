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
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	authdomain "pos-go/internal/modules/auth/domain"
	httpmw "pos-go/internal/transport/http/middleware"

	"github.com/labstack/echo/v4"
)

func TestProductCatalogHandlerCreateMapsPublicFieldsAndActor(t *testing.T) {
	create := &fakeCreateProduct{}
	handler := NewProductCatalogHandler(nil, nil, nil, create, nil, nil, nil, nil)

	body := []byte(`{
		"kode_barang":"SKU-001",
		"nama_barang":"Kampas Rem",
		"merek":"Honda",
		"ukuran":14,
		"harga_jual":40000,
		"reorder_point_qty":5,
		"critical_threshold_qty":2,
		"reason":"created by API"
	}`)

	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req = req.WithContext(httpmw.WithPrincipal(req.Context(), authdomain.Principal{AccountID: "actor-1"}))
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)

	if err := handler.Create(c); err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	assertCreateCommand(t, create)
}

func assertCreateCommand(t *testing.T, create *fakeCreateProduct) {
	t.Helper()

	if create.got.Code != "SKU-001" || create.got.Name != "Kampas Rem" || create.got.Brand != "Honda" {
		t.Fatalf("unexpected identity command: %+v", create.got)
	}
	if create.got.Size == nil || *create.got.Size != 14 {
		t.Fatalf("Size = %v", create.got.Size)
	}
	if create.got.SalePriceRupiah != 40000 {
		t.Fatalf("SalePriceRupiah = %d", create.got.SalePriceRupiah)
	}
	if create.got.ReorderPointQty == nil || *create.got.ReorderPointQty != 5 {
		t.Fatalf("ReorderPointQty = %v", create.got.ReorderPointQty)
	}
	if create.got.CriticalThresholdQty == nil || *create.got.CriticalThresholdQty != 2 {
		t.Fatalf("CriticalThresholdQty = %v", create.got.CriticalThresholdQty)
	}
	if create.got.ActorID != "actor-1" || create.got.Reason != "created by API" {
		t.Fatalf("actor/reason = %q/%q", create.got.ActorID, create.got.Reason)
	}
}
