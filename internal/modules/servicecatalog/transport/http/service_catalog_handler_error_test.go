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
