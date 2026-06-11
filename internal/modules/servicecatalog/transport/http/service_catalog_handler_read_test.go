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

func TestServiceCatalogHandler_ListParsesQueryAndReturnsEnvelope(t *testing.T) {
	h := newTestServiceCatalogHandler()

	var got uc.ListServiceCatalogItemsCommand
	h.list = listFn(func(ctx context.Context, cmd uc.ListServiceCatalogItemsCommand) ([]uc.ServiceCatalogItemResult, error) {
		_ = ctx
		got = cmd

		return []uc.ServiceCatalogItemResult{testServiceCatalogItemResult()}, nil
	})

	c, rec := newServiceCatalogTestContext(http.MethodGet, "/items?q=wash&status=active&page=2&per_page=5", "")

	if err := h.List(c); err != nil {
		t.Fatalf("List() error = %v", err)
	}

	if got.Query != "wash" || string(got.Status) != "active" || got.Page != 2 || got.PerPage != 5 {
		t.Fatalf("command = %+v, want query/status/page/per_page mapped", got)
	}

	assertJSONStatus(t, rec, http.StatusOK)
	assertSuccessEnvelope(t, rec)
}

func TestServiceCatalogHandler_LookupParsesActiveOnly(t *testing.T) {
	h := newTestServiceCatalogHandler()

	var got uc.LookupServiceCatalogItemsCommand
	h.lookup = lookupFn(func(ctx context.Context, cmd uc.LookupServiceCatalogItemsCommand) ([]uc.ServiceCatalogLookupResult, error) {
		_ = ctx
		got = cmd

		return []uc.ServiceCatalogLookupResult{{
			ID: "svc_1", Name: "Express Wash", DefaultPriceRupiah: 15000,
		}}, nil
	})

	c, rec := newServiceCatalogTestContext(http.MethodGet, "/items/lookup?q=wash&limit=7&active_only=true", "")

	if err := h.Lookup(c); err != nil {
		t.Fatalf("Lookup() error = %v", err)
	}

	if got.Query != "wash" || got.Limit != 7 || got.IncludeInactive {
		t.Fatalf("command = %+v, want query/limit mapped and include_inactive false", got)
	}

	assertJSONStatus(t, rec, http.StatusOK)
	assertSuccessEnvelope(t, rec)
}
