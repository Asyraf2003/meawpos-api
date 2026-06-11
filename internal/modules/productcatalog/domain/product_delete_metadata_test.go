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

package domain

import (
	"testing"
	"time"
)

func TestProductDeleteMetadataAccessors(t *testing.T) {
	product, err := NewProduct(ProductInput{
		ID:              "prod_delete_metadata",
		Name:            "Kampas Rem",
		Brand:           "Honda",
		SalePriceRupiah: 40000,
	})
	if err != nil {
		t.Fatalf("NewProduct() error = %v", err)
	}

	deletedAt := time.Date(2026, 6, 10, 12, 0, 0, 0, time.UTC)
	if err := product.SoftDelete(DeleteInput{
		DeletedAt:        deletedAt,
		DeletedByActorID: " actor-1 ",
		Reason:           " duplicate product ",
	}); err != nil {
		t.Fatalf("SoftDelete() error = %v", err)
	}

	if product.DeletedByActorID() != "actor-1" {
		t.Fatalf("DeletedByActorID() = %q, want actor-1", product.DeletedByActorID())
	}
	if product.DeleteReason() != "duplicate product" {
		t.Fatalf("DeleteReason() = %q, want duplicate product", product.DeleteReason())
	}

	if err := product.Restore(); err != nil {
		t.Fatalf("Restore() error = %v", err)
	}

	if product.DeletedByActorID() != "" {
		t.Fatalf("DeletedByActorID() after restore = %q, want empty", product.DeletedByActorID())
	}
	if product.DeleteReason() != "" {
		t.Fatalf("DeleteReason() after restore = %q, want empty", product.DeleteReason())
	}
}
