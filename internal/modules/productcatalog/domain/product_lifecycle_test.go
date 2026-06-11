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

func TestProductSoftDeleteAndRestoreLifecycle(t *testing.T) {
	product, err := NewProduct(ProductInput{
		ID:              "prod_003",
		Name:            "Filter Udara",
		Brand:           "Aspira",
		SalePriceRupiah: 30000,
	})
	if err != nil {
		t.Fatalf("NewProduct() error = %v", err)
	}

	deletedAt := time.Date(2026, 6, 9, 10, 0, 0, 0, time.UTC)
	if err := product.SoftDelete(DeleteInput{
		DeletedAt:        deletedAt,
		DeletedByActorID: "actor_admin",
		Reason:           "duplicate test product",
	}); err != nil {
		t.Fatalf("SoftDelete() error = %v", err)
	}

	if product.Status() != ProductStatusDeleted {
		t.Fatalf("Status() = %v, want %v", product.Status(), ProductStatusDeleted)
	}
	if product.DeletedAt() == nil || !product.DeletedAt().Equal(deletedAt) {
		t.Fatalf("DeletedAt() = %v, want %v", product.DeletedAt(), deletedAt)
	}

	if err := product.Restore(); err != nil {
		t.Fatalf("Restore() error = %v", err)
	}

	if product.Status() != ProductStatusActive {
		t.Fatalf("Status() = %v, want %v", product.Status(), ProductStatusActive)
	}
	if product.DeletedAt() != nil {
		t.Fatalf("DeletedAt() = %v, want nil", product.DeletedAt())
	}
}
