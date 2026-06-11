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
	"testing"
	"time"

	"pos-go/internal/modules/productcatalog/domain"
)

func assertCreateProductDuplicateCandidate(t *testing.T, checker *fakeProductDuplicateChecker) {
	t.Helper()

	if !checker.createCalled {
		t.Fatalf("duplicate checker was not called")
	}
	if checker.candidate.Code == nil || *checker.candidate.Code != "PRD-001" {
		t.Fatalf("candidate.Code = %v, want PRD-001", checker.candidate.Code)
	}
	if checker.candidate.NormalizedName != "oli mesin" {
		t.Fatalf("candidate.NormalizedName = %q, want oli mesin", checker.candidate.NormalizedName)
	}
	if checker.candidate.NormalizedBrand != "yamaha genuine" {
		t.Fatalf("candidate.NormalizedBrand = %q, want yamaha genuine", checker.candidate.NormalizedBrand)
	}
	if checker.candidate.Size == nil || *checker.candidate.Size != 1000 {
		t.Fatalf("candidate.Size = %v, want 1000", checker.candidate.Size)
	}
}

func assertCreateProductResult(t *testing.T, result CreateProductResult, fixedNow time.Time) {
	t.Helper()

	if result.ID != "prod_001" {
		t.Fatalf("result.ID = %q, want prod_001", result.ID)
	}
	if result.Name != "Oli Mesin" {
		t.Fatalf("result.Name = %q, want Oli Mesin", result.Name)
	}
	if result.Brand != "Yamaha Genuine" {
		t.Fatalf("result.Brand = %q, want Yamaha Genuine", result.Brand)
	}
	if result.Status != string(domain.ProductStatusActive) {
		t.Fatalf("result.Status = %q, want active", result.Status)
	}
	if !result.CreatedAt.Equal(fixedNow) || !result.UpdatedAt.Equal(fixedNow) {
		t.Fatalf("result timestamps = %v/%v, want %v", result.CreatedAt, result.UpdatedAt, fixedNow)
	}
}

func assertCreateProductSideEffects(
	t *testing.T,
	versionRepository *fakeProductVersionRepository,
	auditRecorder *fakeProductAuditRecorder,
	fixedNow time.Time,
) {
	t.Helper()

	if len(versionRepository.appended) != 1 {
		t.Fatalf("version append count = %d, want 1", len(versionRepository.appended))
	}
	version := versionRepository.appended[0]
	if version.ProductID != "prod_001" {
		t.Fatalf("version.ProductID = %q, want prod_001", version.ProductID)
	}
	if version.RevisionNo != 1 {
		t.Fatalf("version.RevisionNo = %d, want 1", version.RevisionNo)
	}
	if version.EventName != productCreatedEventName {
		t.Fatalf("version.EventName = %q, want %q", version.EventName, productCreatedEventName)
	}
	if !version.ChangedAt.Equal(fixedNow) {
		t.Fatalf("version.ChangedAt = %v, want %v", version.ChangedAt, fixedNow)
	}

	if len(auditRecorder.records) != 1 {
		t.Fatalf("audit record count = %d, want 1", len(auditRecorder.records))
	}
	audit := auditRecorder.records[0]
	if audit.AggregateID != "prod_001" {
		t.Fatalf("audit.AggregateID = %q, want prod_001", audit.AggregateID)
	}
	if audit.EventName != productCreatedEventName {
		t.Fatalf("audit.EventName = %q, want %q", audit.EventName, productCreatedEventName)
	}
	if audit.Operation != "create" {
		t.Fatalf("audit.Operation = %q, want create", audit.Operation)
	}
	if audit.RevisionNo != 1 {
		t.Fatalf("audit.RevisionNo = %d, want 1", audit.RevisionNo)
	}
	if !audit.OccurredAt.Equal(fixedNow) {
		t.Fatalf("audit.OccurredAt = %v, want %v", audit.OccurredAt, fixedNow)
	}
}
