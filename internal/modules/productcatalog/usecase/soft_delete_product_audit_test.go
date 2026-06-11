// Copyright (C) 2026 Asyraf Mubarak
//
// This file is part of gopos-api.
//
// gopos-api is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, version 3 only.
//
// gopos-api is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with gopos-api. If not, see <https://www.gnu.org/licenses/>.

package usecase

import (
	"context"
	"testing"
	"time"

	"pos-go/internal/modules/productcatalog/domain"
)

func TestSoftDeleteProductRecordsAudit(t *testing.T) {
	product, err := domain.NewProduct(domain.ProductInput{
		ID:              "product-1",
		Name:            "Filter Udara",
		Brand:           "Aspira",
		SalePriceRupiah: 30000,
	})
	if err != nil {
		t.Fatalf("NewProduct() error = %v", err)
	}

	auditRecorder := &fakeProductAuditRecorder{}
	deletedAt := time.Date(2026, 6, 10, 10, 0, 0, 0, time.UTC)
	usecase := NewSoftDeleteProduct(
		&softDeleteProductRepositoryDouble{found: product},
		&fakeProductVersionRepository{},
		auditRecorder,
		func() time.Time { return deletedAt },
	)

	_, err = usecase.Execute(context.Background(), SoftDeleteProductCommand{
		ID:      "product-1",
		ActorID: "actor-1",
		Reason:  "obsolete product",
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
	if record.EventName != "product_deleted" {
		t.Fatalf("EventName = %q, want product_deleted", record.EventName)
	}
	if record.Operation != "delete" {
		t.Fatalf("Operation = %q, want delete", record.Operation)
	}
	if record.ActorID != "actor-1" {
		t.Fatalf("ActorID = %q, want actor-1", record.ActorID)
	}
	if record.Reason != "obsolete product" {
		t.Fatalf("Reason = %q, want obsolete product", record.Reason)
	}
	if !record.OccurredAt.Equal(deletedAt) {
		t.Fatalf("OccurredAt = %v, want %v", record.OccurredAt, deletedAt)
	}
	if record.RevisionNo != 1 {
		t.Fatalf("RevisionNo = %d, want 1", record.RevisionNo)
	}
}
