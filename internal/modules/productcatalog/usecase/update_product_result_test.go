package usecase

import (
	"context"
	"testing"
	"time"

	"pos-go/internal/modules/productcatalog/domain"
	"pos-go/internal/modules/productcatalog/ports"
)

func TestUpdateProductReturnsMappedResult(t *testing.T) {
	fixedNow := time.Date(2026, 6, 10, 12, 0, 0, 0, time.UTC)
	existing, err := domain.NewProduct(domain.ProductInput{
		ID:              "prod_009",
		Name:            "Busi Lama",
		Brand:           "NGK",
		SalePriceRupiah: 20000,
	})
	if err != nil {
		t.Fatalf("NewProduct() error = %v", err)
	}

	uc := NewUpdateProduct(
		&fakeUpdateProductRepository{product: existing},
		&fakeUpdateProductDuplicateChecker{},
		&fakeUpdateProductVersionRepository{
			existing: []ports.ProductVersionRecord{{RevisionNo: 1}},
		},
		&fakeProductAuditRecorder{},
		func() time.Time { return fixedNow },
	)

	result, err := uc.Execute(context.Background(), UpdateProductCommand{
		ID:              "prod_009",
		Code:            "  PRD-009  ",
		Name:            "  Oli   Mesin ",
		Brand:           " Yamaha   Genuine ",
		Size:            domain.IntPtr(1000),
		SalePriceRupiah: 55000,
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if result.ID != "prod_009" {
		t.Fatalf("result.ID = %q, want prod_009", result.ID)
	}
	if result.Code == nil || *result.Code != "PRD-009" {
		t.Fatalf("result.Code = %v, want PRD-009", result.Code)
	}
	if result.Name != "Oli Mesin" {
		t.Fatalf("result.Name = %q, want Oli Mesin", result.Name)
	}
	if result.NormalizedBrand != "yamaha genuine" {
		t.Fatalf("result.NormalizedBrand = %q, want yamaha genuine", result.NormalizedBrand)
	}
	if result.RevisionNo != 2 {
		t.Fatalf("result.RevisionNo = %d, want 2", result.RevisionNo)
	}
	if !result.UpdatedAt.Equal(fixedNow) {
		t.Fatalf("result.UpdatedAt = %v, want %v", result.UpdatedAt, fixedNow)
	}
}
