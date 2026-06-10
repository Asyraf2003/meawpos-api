package usecase

import (
	"context"
	"testing"
	"time"

	"pos-go/internal/modules/productcatalog/domain"
)

func TestRestoreProductRecordsAudit(t *testing.T) {
	product, err := domain.NewProduct(domain.ProductInput{
		ID:              "product-1",
		Name:            "Filter Udara",
		Brand:           "Aspira",
		SalePriceRupiah: 30000,
	})
	if err != nil {
		t.Fatalf("NewProduct() error = %v", err)
	}

	if err := product.SoftDelete(domain.DeleteInput{
		DeletedAt:        time.Date(2026, 6, 10, 9, 0, 0, 0, time.UTC),
		DeletedByActorID: "actor-1",
		Reason:           "setup deleted product",
	}); err != nil {
		t.Fatalf("SoftDelete() setup error = %v", err)
	}

	auditRecorder := &fakeProductAuditRecorder{}
	restoredAt := time.Date(2026, 6, 10, 11, 0, 0, 0, time.UTC)
	usecase := NewRestoreProduct(
		&restoreProductRepositoryDouble{found: product},
		&fakeProductVersionRepository{},
		auditRecorder,
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

	if len(auditRecorder.records) != 1 {
		t.Fatalf("audit record count = %d, want 1", len(auditRecorder.records))
	}

	record := auditRecorder.records[0]
	if record.AggregateID != "product-1" {
		t.Fatalf("AggregateID = %q, want product-1", record.AggregateID)
	}
	if record.EventName != "product_restored" {
		t.Fatalf("EventName = %q, want product_restored", record.EventName)
	}
	if record.Operation != "restore" {
		t.Fatalf("Operation = %q, want restore", record.Operation)
	}
	if record.ActorID != "actor-2" {
		t.Fatalf("ActorID = %q, want actor-2", record.ActorID)
	}
	if record.Reason != "restore product" {
		t.Fatalf("Reason = %q, want restore product", record.Reason)
	}
	if !record.OccurredAt.Equal(restoredAt) {
		t.Fatalf("OccurredAt = %v, want %v", record.OccurredAt, restoredAt)
	}
	if record.RevisionNo != 1 {
		t.Fatalf("RevisionNo = %d, want 1", record.RevisionNo)
	}
}
