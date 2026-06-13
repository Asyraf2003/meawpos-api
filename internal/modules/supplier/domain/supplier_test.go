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

package domain

import (
	"errors"
	"testing"
	"time"
)

func TestNewSupplierNormalizesNameAndContact(t *testing.T) {
	now := time.Date(2026, 6, 14, 10, 0, 0, 0, time.UTC)

	supplier, err := NewSupplier(" supplier-1 ", "  Bengkel   Jaya  ", SupplierContact{
		Phone:   " 0812 ",
		Email:   " owner@example.test ",
		Address: " Jalan Test ",
		Notes:   " sparepart ",
	}, now)
	if err != nil {
		t.Fatalf("NewSupplier() error = %v", err)
	}

	if supplier.ID() != "supplier-1" {
		t.Fatalf("ID() = %q", supplier.ID())
	}
	if supplier.Name() != "Bengkel Jaya" {
		t.Fatalf("Name() = %q", supplier.Name())
	}
	if supplier.NormalizedName() != "bengkel jaya" {
		t.Fatalf("NormalizedName() = %q", supplier.NormalizedName())
	}
	if supplier.Phone() != "0812" {
		t.Fatalf("Phone() = %q", supplier.Phone())
	}
	if !supplier.IsActive() {
		t.Fatal("IsActive() = false, want true")
	}
	if supplier.Status() != SupplierStatusActive {
		t.Fatalf("Status() = %q", supplier.Status())
	}
}

func TestNewSupplierRejectsBlankName(t *testing.T) {
	_, err := NewSupplier("supplier-1", " ", SupplierContact{}, time.Now())
	if !errors.Is(err, ErrBlankSupplierName) {
		t.Fatalf("error = %v, want %v", err, ErrBlankSupplierName)
	}
}

func TestSupplierUpdateAndLifecycle(t *testing.T) {
	now := time.Date(2026, 6, 14, 10, 0, 0, 0, time.UTC)
	later := now.Add(time.Hour)

	supplier, err := NewSupplier("supplier-1", "Bengkel Jaya", SupplierContact{}, now)
	if err != nil {
		t.Fatalf("NewSupplier() error = %v", err)
	}

	err = supplier.Update(" Toko   Sumber ", SupplierContact{Phone: " 555 "}, later)
	if err != nil {
		t.Fatalf("Update() error = %v", err)
	}

	if supplier.Name() != "Toko Sumber" {
		t.Fatalf("Name() = %q", supplier.Name())
	}
	if supplier.NormalizedName() != "toko sumber" {
		t.Fatalf("NormalizedName() = %q", supplier.NormalizedName())
	}
	if supplier.Phone() != "555" {
		t.Fatalf("Phone() = %q", supplier.Phone())
	}
	if !supplier.UpdatedAt().Equal(later) {
		t.Fatalf("UpdatedAt() = %v", supplier.UpdatedAt())
	}

	supplier.Deactivate(later.Add(time.Hour))
	if supplier.IsActive() {
		t.Fatal("IsActive() = true, want false")
	}
	if supplier.Status() != SupplierStatusInactive {
		t.Fatalf("Status() = %q", supplier.Status())
	}

	supplier.Activate(later.Add(2 * time.Hour))
	if !supplier.IsActive() {
		t.Fatal("IsActive() = false, want true")
	}
}
