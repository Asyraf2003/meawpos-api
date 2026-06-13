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

	"pos-go/internal/modules/supplier/ports"
)

func TestListSuppliersAppliesDefaults(t *testing.T) {
	repo := newFakeSupplierRepository()
	repo.byID["supplier-1"] = mustSupplier(t, "supplier-1", "Bengkel Jaya", true)
	uc := NewListSuppliers(repo)

	results, err := uc.Execute(context.Background(), ListSuppliersCommand{Query: "jay"})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if len(results) != 1 {
		t.Fatalf("len(results) = %d", len(results))
	}
	if repo.listFilter.Status != ports.ListStatusActive {
		t.Fatalf("Status = %q", repo.listFilter.Status)
	}
	if repo.listFilter.Page != 1 {
		t.Fatalf("Page = %d", repo.listFilter.Page)
	}
	if repo.listFilter.PerPage != 10 {
		t.Fatalf("PerPage = %d", repo.listFilter.PerPage)
	}
}

func TestLookupSuppliersDefaultsActiveOnly(t *testing.T) {
	repo := newFakeSupplierRepository()
	repo.byID["supplier-1"] = mustSupplier(t, "supplier-1", "Bengkel Jaya", true)
	uc := NewLookupSuppliers(repo)

	results, err := uc.Execute(context.Background(), LookupSuppliersCommand{})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if len(results) != 1 {
		t.Fatalf("len(results) = %d", len(results))
	}
	if repo.lookupFilter.Limit != 20 {
		t.Fatalf("Limit = %d", repo.lookupFilter.Limit)
	}
	if !repo.lookupFilter.ActiveOnly {
		t.Fatal("ActiveOnly = false, want true")
	}
}

func TestLookupSuppliersRejectsInvalidLimit(t *testing.T) {
	repo := newFakeSupplierRepository()
	uc := NewLookupSuppliers(repo)

	_, err := uc.Execute(context.Background(), LookupSuppliersCommand{Limit: 51})
	if !errors.Is(err, ErrInvalidSupplierLookupLimit) {
		t.Fatalf("error = %v, want %v", err, ErrInvalidSupplierLookupLimit)
	}
}

func TestShowSupplierReturnsNotFound(t *testing.T) {
	repo := newFakeSupplierRepository()
	uc := NewShowSupplier(repo)

	_, err := uc.Execute(context.Background(), ShowSupplierCommand{ID: "missing"})
	if !errors.Is(err, ErrSupplierNotFound) {
		t.Fatalf("error = %v, want %v", err, ErrSupplierNotFound)
	}
}
