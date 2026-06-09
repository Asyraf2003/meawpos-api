package usecase

import (
	"context"
	"testing"
	"time"

	"pos-go/internal/modules/productcatalog/domain"
)

func TestUpdateProductPersistsUpdatedProduct(t *testing.T) {
	existing, err := domain.NewProduct(domain.ProductInput{
		ID:              "prod_003",
		Name:            "Busi Lama",
		Brand:           "NGK",
		SalePriceRupiah: 20000,
	})
	if err != nil {
		t.Fatalf("NewProduct() error = %v", err)
	}

	repository := &fakeUpdateProductRepository{product: existing}
	uc := NewUpdateProduct(
		repository,
		&fakeUpdateProductDuplicateChecker{},
		&fakeProductVersionRepository{},
		&fakeProductAuditRecorder{},
		time.Now,
	)

	_, err = uc.Execute(context.Background(), UpdateProductCommand{
		ID:              "prod_003",
		Code:            "PRD-003",
		Name:            "Busi Baru",
		Brand:           "Denso",
		SalePriceRupiah: 30000,
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if repository.updated == nil {
		t.Fatalf("repository.Update was not called")
	}
	if repository.updated.ID() != "prod_003" {
		t.Fatalf("updated.ID = %q, want prod_003", repository.updated.ID())
	}
	if repository.updated.Name() != "Busi Baru" {
		t.Fatalf("updated.Name = %q, want Busi Baru", repository.updated.Name())
	}
	if repository.updated.Brand() != "Denso" {
		t.Fatalf("updated.Brand = %q, want Denso", repository.updated.Brand())
	}
}
