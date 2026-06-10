package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"pos-go/internal/modules/productcatalog/domain"
)

func TestUpdateProductReturnsVersionAppendError(t *testing.T) {
	versionErr := errors.New("version append failure")
	existing, err := domain.NewProduct(domain.ProductInput{
		ID:              "prod_006",
		Name:            "Busi Lama",
		Brand:           "NGK",
		SalePriceRupiah: 20000,
	})
	if err != nil {
		t.Fatalf("NewProduct() error = %v", err)
	}

	versionRepository := &fakeUpdateProductVersionRepository{err: versionErr}
	uc := NewUpdateProduct(
		&fakeUpdateProductRepository{product: existing},
		&fakeUpdateProductDuplicateChecker{},
		versionRepository,
		&fakeProductAuditRecorder{},
		time.Now,
	)

	_, err = uc.Execute(context.Background(), UpdateProductCommand{
		ID:              "prod_006",
		Name:            "Busi Baru",
		Brand:           "Denso",
		SalePriceRupiah: 30000,
	})

	if !errors.Is(err, versionErr) {
		t.Fatalf("Execute() error = %v, want %v", err, versionErr)
	}
	if len(versionRepository.appended) != 1 {
		t.Fatalf("version append count = %d, want 1", len(versionRepository.appended))
	}
}
