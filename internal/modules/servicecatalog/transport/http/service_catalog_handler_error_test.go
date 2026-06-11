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

func TestServiceCatalogHandler_ShowMapsNotFoundTo404(t *testing.T) {
	h := newTestServiceCatalogHandler()

	h.show = showFn(func(ctx context.Context, cmd uc.ShowServiceCatalogItemCommand) (uc.ServiceCatalogItemResult, error) {
		_ = ctx

		if cmd.ID != "svc_missing" {
			t.Fatalf("id = %q, want svc_missing", cmd.ID)
		}

		return uc.ServiceCatalogItemResult{}, uc.ErrServiceCatalogItemNotFound
	})

	c, _ := newServiceCatalogTestContext(http.MethodGet, "/items/svc_missing", "")
	c.SetParamNames("id")
	c.SetParamValues("svc_missing")

	assertHTTPErrorCode(t, h.Show(c), http.StatusNotFound)
}
