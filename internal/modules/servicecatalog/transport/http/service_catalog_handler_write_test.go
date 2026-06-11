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
	"net/http"
	"testing"

	uc "pos-go/internal/modules/servicecatalog/usecase"
)

func TestServiceCatalogHandler_CreateRejectsInvalidBodyBeforeUsecase(t *testing.T) {
	h := newTestServiceCatalogHandler()

	called := false
	h.create = createFn(func(ctx context.Context, cmd uc.CreateServiceCatalogItemCommand) (uc.ServiceCatalogItemResult, error) {
		_ = ctx
		_ = cmd
		called = true

		return uc.ServiceCatalogItemResult{}, nil
	})

	c, _ := newServiceCatalogTestContext(http.MethodPost, "/items", `{"name":`)

	assertHTTPErrorCode(t, h.Create(c), http.StatusBadRequest)

	if called {
		t.Fatal("create usecase was called for invalid body")
	}
}

func TestServiceCatalogHandler_UpdateMapsIDAndBody(t *testing.T) {
	h := newTestServiceCatalogHandler()

	var got uc.UpdateServiceCatalogItemCommand
	h.update = updateFn(func(ctx context.Context, cmd uc.UpdateServiceCatalogItemCommand) (uc.ServiceCatalogItemResult, error) {
		_ = ctx
		got = cmd

		result := testServiceCatalogItemResult()
		result.ID = cmd.ID
		result.Name = cmd.Name
		result.DefaultPriceRupiah = cmd.DefaultPriceRupiah

		return result, nil
	})

	c, rec := newServiceCatalogTestContext(
		http.MethodPut,
		"/items/svc_1",
		`{"name":"Premium Wash","default_price_rupiah":25000}`,
	)
	c.SetParamNames("id")
	c.SetParamValues("svc_1")

	if err := h.Update(c); err != nil {
		t.Fatalf("Update() error = %v", err)
	}

	if got.ID != "svc_1" || got.Name != "Premium Wash" || got.DefaultPriceRupiah != 25000 {
		t.Fatalf("command = %+v, want id/name/default_price_rupiah mapped", got)
	}

	assertJSONStatus(t, rec, http.StatusOK)
	assertSuccessEnvelope(t, rec)
}
