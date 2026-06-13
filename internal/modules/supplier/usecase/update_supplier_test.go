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
// along with gopos-api. If not, see https://www.gnu.org/licenses/.

package usecase

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestUpdateSupplierUpdatesExistingSupplier(t *testing.T) {
	repo := newFakeSupplierRepository()
	repo.byID["supplier-1"] = mustSupplier(t, "supplier-1", "Bengkel Jaya", true)
	uc := NewUpdateSupplier(repo, func() time.Time {
		return time.Date(2026, 6, 14, 12, 0, 0, 0, time.UTC)
	})

	result, err := uc.Execute(context.Background(), UpdateSupplierCommand{
		ID:   " supplier-1 ",
		Name: " Toko   Sumber ",
		Contact: SupplierContactInput{
			Phone: " 555 ",
		},
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if result.Name != "Toko Sumber" {
		t.Fatalf("Name = %q", result.Name)
	}
	if result.Phone != "555" {
		t.Fatalf("Phone = %q", result.Phone)
	}
	if repo.updateCalls != 1 {
		t.Fatalf("updateCalls = %d", repo.updateCalls)
	}
}

func TestUpdateSupplierReturnsNotFound(t *testing.T) {
	repo := newFakeSupplierRepository()
	uc := NewUpdateSupplier(repo, time.Now)

	_, err := uc.Execute(context.Background(), UpdateSupplierCommand{
		ID:   "missing",
		Name: "Supplier",
	})
	if !errors.Is(err, ErrSupplierNotFound) {
		t.Fatalf("error = %v, want %v", err, ErrSupplierNotFound)
	}
}

func TestUpdateSupplierRejectsDuplicateActiveName(t *testing.T) {
	repo := newFakeSupplierRepository()
	repo.byID["supplier-1"] = mustSupplier(t, "supplier-1", "Bengkel Jaya", true)
	repo.byID["supplier-2"] = mustSupplier(t, "supplier-2", "Toko Sumber", true)

	uc := NewUpdateSupplier(repo, time.Now)

	_, err := uc.Execute(context.Background(), UpdateSupplierCommand{
		ID:   "supplier-1",
		Name: "Toko Sumber",
	})
	if !errors.Is(err, ErrDuplicateSupplierActiveName) {
		t.Fatalf("error = %v, want %v", err, ErrDuplicateSupplierActiveName)
	}
}

func TestUpdateInactiveSupplierAllowsDuplicateActiveName(t *testing.T) {
	repo := newFakeSupplierRepository()
	repo.byID["supplier-1"] = mustSupplier(t, "supplier-1", "Bengkel Jaya", false)
	repo.byID["supplier-2"] = mustSupplier(t, "supplier-2", "Toko Sumber", true)

	uc := NewUpdateSupplier(repo, time.Now)

	_, err := uc.Execute(context.Background(), UpdateSupplierCommand{
		ID:   "supplier-1",
		Name: "Toko Sumber",
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}
}
