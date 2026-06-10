package usecase

import (
	"context"
	"testing"

	"pos-go/internal/modules/productcatalog/ports"
)

func TestLookupProductsMapsItems(t *testing.T) {
	code := "FLT-001"
	size := 12
	reader := &lookupProductsReaderDouble{
		lookupItems: []ports.ProductLookupItem{
			{
				ID:              "product-1",
				Code:            &code,
				Name:            "Filter Udara",
				Brand:           "Aspira",
				Size:            &size,
				SalePriceRupiah: 30000,
				Status:          "active",
			},
		},
	}
	usecase := NewLookupProducts(reader)

	result, err := usecase.Execute(context.Background(), LookupProductsQuery{})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if len(result.Items) != 1 {
		t.Fatalf("Items length = %d, want 1", len(result.Items))
	}

	item := result.Items[0]
	if item.ID != "product-1" {
		t.Fatalf("ID = %q, want product-1", item.ID)
	}
	if item.Code == nil || *item.Code != "FLT-001" {
		t.Fatalf("Code = %v, want FLT-001", item.Code)
	}
	if item.Name != "Filter Udara" {
		t.Fatalf("Name = %q, want Filter Udara", item.Name)
	}
	if item.Brand != "Aspira" {
		t.Fatalf("Brand = %q, want Aspira", item.Brand)
	}
	if item.Size == nil || *item.Size != 12 {
		t.Fatalf("Size = %v, want 12", item.Size)
	}
	if item.SalePriceRupiah != 30000 {
		t.Fatalf("SalePriceRupiah = %d, want 30000", item.SalePriceRupiah)
	}
	if item.Status != "active" {
		t.Fatalf("Status = %q, want active", item.Status)
	}
}
