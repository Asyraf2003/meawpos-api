package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"pos-go/internal/modules/productcatalog/domain"
)

func TestUpdateProductReturnsRepositoryUpdateError(t *testing.T) {
	updateErr := errors.New("product update failure")
	existing, err := domain.NewProduct(domain.ProductInput{
		ID:              "prod_004",
		Name:            "Busi Lama",
		Brand:           "NGK",
		SalePriceRupiah: 20000,
	})
	if err != nil {
		t.Fatalf("NewProduct() error = %v", err)
	}

	repository := &fakeUpdateProductRepository{
		product: existing,
		err:     updateErr,
	}
	uc := NewUpdateProduct(
		repository,
		&fakeUpdateProductDuplicateChecker{},
		&fakeProductVersionRepository{},
		&fakeProductAuditRecorder{},
		time.Now,
	)

	_, err = uc.Execute(context.Background(), UpdateProductCommand{
		ID:              "prod_004",
		Name:            "Busi Baru",
		Brand:           "Denso",
		SalePriceRupiah: 30000,
	})

	if !errors.Is(err, updateErr) {
		t.Fatalf("Execute() error = %v, want %v", err, updateErr)
	}
	if repository.updated == nil {
		t.Fatalf("repository.Update was not called")
	}
}
