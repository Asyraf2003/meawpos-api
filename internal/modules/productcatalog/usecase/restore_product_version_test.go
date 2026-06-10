package usecase

import (
	"context"
	"testing"
	"time"

	"pos-go/internal/modules/productcatalog/domain"
)

func TestRestoreProductRecordsVersion(t *testing.T) {
	product, err := domain.NewProduct(domain.ProductInput{
		ID:              "product-1",
		Name:            "Filter Udara",
		Brand:           "Aspira",
		SalePriceRupiah: 30000,
	})
	if err != nil {
		t.Fatalf("NewProduct() error = %v", err)
	}

	deletedAt := time.Date(2026, 6, 10, 9, 0, 0, 0, time.UTC)
	if err := product.SoftDelete(domain.DeleteInput{
		DeletedAt:        deletedAt,
		DeletedByActorID: "actor-1",
		Reason:           "setup deleted product",
	}); err != nil {
		t.Fatalf("SoftDelete() setup error = %v", err)
	}

	versionRepository := &fakeProductVersionRepository{}
	restoredAt := time.Date(2026, 6, 10, 11, 0, 0, 0, time.UTC)
	usecase := NewRestoreProduct(
		&restoreProductRepositoryDouble{found: product},
		versionRepository,
		&fakeProductAuditRecorder{},
		func() time.Time { return restoredAt },
	)

	_, err = usecase.Execute(context.Background(), RestoreProductCommand{
		ID:      "product-1",
		ActorID: "actor-2",
		Reason:  "restore product",
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if len(versionRepository.appended) != 1 {
		t.Fatalf("version append count = %d, want 1", len(versionRepository.appended))
	}

	version := versionRepository.appended[0]
	if version.ProductID != "product-1" {
		t.Fatalf("ProductID = %q, want product-1", version.ProductID)
	}
	if version.EventName != "product_restored" {
		t.Fatalf("EventName = %q, want product_restored", version.EventName)
	}
	if version.ChangedByActorID != "actor-2" {
		t.Fatalf("ChangedByActorID = %q, want actor-2", version.ChangedByActorID)
	}
	if version.ChangeReason != "restore product" {
		t.Fatalf("ChangeReason = %q, want restore product", version.ChangeReason)
	}
	if !version.ChangedAt.Equal(restoredAt) {
		t.Fatalf("ChangedAt = %v, want %v", version.ChangedAt, restoredAt)
	}
	if version.RevisionNo != 1 {
		t.Fatalf("RevisionNo = %d, want 1", version.RevisionNo)
	}
}
