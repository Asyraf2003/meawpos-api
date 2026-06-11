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

//go:build integration

package postgres

import (
	"testing"
	"time"

	"pos-go/internal/modules/productcatalog/ports"
)

func newProductVersionRecord(
	productID string,
	revisionNo int,
	eventName string,
	changedAt time.Time,
) ports.ProductVersionRecord {
	return ports.ProductVersionRecord{
		ProductID:        productID,
		RevisionNo:       revisionNo,
		EventName:        eventName,
		ChangedByActorID: "actor-1",
		ChangeReason:     "integration test",
		ChangedAt:        changedAt,
	}
}

func assertProductVersionRecord(
	t *testing.T,
	got ports.ProductVersionRecord,
	want ports.ProductVersionRecord,
) {
	t.Helper()

	if got.ProductID != want.ProductID {
		t.Fatalf("ProductID = %q, want %q", got.ProductID, want.ProductID)
	}
	if got.RevisionNo != want.RevisionNo {
		t.Fatalf("RevisionNo = %d, want %d", got.RevisionNo, want.RevisionNo)
	}
	if got.EventName != want.EventName {
		t.Fatalf("EventName = %q, want %q", got.EventName, want.EventName)
	}
	if got.ChangedByActorID != want.ChangedByActorID {
		t.Fatalf("ChangedByActorID = %q, want %q", got.ChangedByActorID, want.ChangedByActorID)
	}
	if got.ChangeReason != want.ChangeReason {
		t.Fatalf("ChangeReason = %q, want %q", got.ChangeReason, want.ChangeReason)
	}
	if !got.ChangedAt.Equal(want.ChangedAt) {
		t.Fatalf("ChangedAt = %v, want %v", got.ChangedAt, want.ChangedAt)
	}
}
