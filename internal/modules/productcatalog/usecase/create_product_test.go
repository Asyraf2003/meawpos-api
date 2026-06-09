package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"pos-go/internal/modules/productcatalog/domain"
)

func TestCreateProductSuccessPersistsProductAndChecksDuplicateCandidate(t *testing.T) {
	fixedNow := time.Date(2026, 6, 9, 12, 0, 0, 0, time.UTC)
	repository := &fakeProductRepository{}
	duplicateChecker := &fakeProductDuplicateChecker{}
	versionRepository := &fakeProductVersionRepository{}
	auditRecorder := &fakeProductAuditRecorder{}

	uc := NewCreateProduct(
		repository,
		duplicateChecker,
		versionRepository,
		auditRecorder,
		fakeProductIDGenerator{id: "prod_001"},
		func() time.Time { return fixedNow },
	)

	result, err := uc.Execute(context.Background(), CreateProductCommand{
		Code:            "  PRD-001  ",
		Name:            "  Oli   Mesin ",
		Brand:           " Yamaha   Genuine ",
		Size:            domain.IntPtr(1000),
		SalePriceRupiah: 55000,
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if repository.created == nil {
		t.Fatalf("repository.Create was not called")
	}

	assertCreateProductDuplicateCandidate(t, duplicateChecker)
	assertCreateProductResult(t, result, fixedNow)
	assertCreateProductSideEffects(t, versionRepository, auditRecorder, fixedNow)

}

func TestCreateProductReturnsDuplicateCheckerError(t *testing.T) {
	duplicateErr := errors.New("duplicate failure")

	uc := NewCreateProduct(
		&fakeProductRepository{},
		&fakeProductDuplicateChecker{err: duplicateErr},
		&fakeProductVersionRepository{},
		&fakeProductAuditRecorder{},
		fakeProductIDGenerator{id: "prod_002"},
		time.Now,
	)

	_, err := uc.Execute(context.Background(), CreateProductCommand{
		Name:            "Busi",
		Brand:           "NGK",
		SalePriceRupiah: 25000,
	})
	if !errors.Is(err, duplicateErr) {
		t.Fatalf("Execute() error = %v, want %v", err, duplicateErr)
	}
}
