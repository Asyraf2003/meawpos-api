package usecase

import (
	"context"
	"testing"
	"time"

	"pos-go/internal/modules/productcatalog/domain"
	"pos-go/internal/modules/productcatalog/ports"
)

func TestUpdateProductRecordsAudit(t *testing.T) {
	fixedNow := time.Date(2026, 6, 10, 11, 0, 0, 0, time.UTC)
	existing, err := domain.NewProduct(domain.ProductInput{
		ID:              "prod_007",
		Name:            "Busi Lama",
		Brand:           "NGK",
		SalePriceRupiah: 20000,
	})
	if err != nil {
		t.Fatalf("NewProduct() error = %v", err)
	}

	auditRecorder := &fakeProductAuditRecorder{}
	uc := NewUpdateProduct(
		&fakeUpdateProductRepository{product: existing},
		&fakeUpdateProductDuplicateChecker{},
		&fakeUpdateProductVersionRepository{
			existing: []ports.ProductVersionRecord{{RevisionNo: 1}},
		},
		auditRecorder,
		func() time.Time { return fixedNow },
	)

	_, err = uc.Execute(context.Background(), UpdateProductCommand{
		ID:              "prod_007",
		Name:            "Busi Baru",
		Brand:           "Denso",
		SalePriceRupiah: 30000,
		ActorID:         "actor_002",
		Reason:          "catalog correction",
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if len(auditRecorder.records) != 1 {
		t.Fatalf("audit record count = %d, want 1", len(auditRecorder.records))
	}
	audit := auditRecorder.records[0]
	if audit.AggregateID != "prod_007" {
		t.Fatalf("audit.AggregateID = %q, want prod_007", audit.AggregateID)
	}
	if audit.EventName != productUpdatedEventName {
		t.Fatalf("audit.EventName = %q, want %q", audit.EventName, productUpdatedEventName)
	}
	if audit.Operation != "update" {
		t.Fatalf("audit.Operation = %q, want update", audit.Operation)
	}
	if audit.ActorID != "actor_002" {
		t.Fatalf("audit.ActorID = %q, want actor_002", audit.ActorID)
	}
	if audit.Reason != "catalog correction" {
		t.Fatalf("audit.Reason = %q, want catalog correction", audit.Reason)
	}
	if audit.RevisionNo != 2 {
		t.Fatalf("audit.RevisionNo = %d, want 2", audit.RevisionNo)
	}
	if !audit.OccurredAt.Equal(fixedNow) {
		t.Fatalf("audit.OccurredAt = %v, want %v", audit.OccurredAt, fixedNow)
	}
}
