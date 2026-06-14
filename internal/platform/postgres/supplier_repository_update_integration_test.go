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
	"context"
	"testing"
	"time"

	"pos-go/internal/modules/supplier/domain"
)

func TestSupplierRepository_UpdateChangesFields(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenSupplierRepoTx(t, ctx)
	repo := NewSupplierRepository(pool)
	supplier := newSupplierRepositoryTestSupplier(t, "Bengkel Lama")
	updatedAt := supplierRepositoryTestTime.Add(2 * time.Hour)
	if err := repo.Create(txCtx, supplier); err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	if err := supplier.Update("Bengkel Baru", domain.SupplierContact{
		Phone:   "08999999999",
		Email:   "baru@example.com",
		Address: "Jalan Baru 2",
		Notes:   "updated supplier",
	}, updatedAt); err != nil {
		t.Fatalf("Supplier.Update() error = %v", err)
	}
	supplier.Deactivate(updatedAt.Add(time.Minute))
	if err := repo.Update(txCtx, supplier); err != nil {
		t.Fatalf("Update() error = %v", err)
	}
	got, found, err := repo.FindByID(txCtx, supplier.ID())
	if err != nil {
		t.Fatalf("FindByID() error = %v", err)
	}
	if !found {
		t.Fatal("FindByID() found = false, want true")
	}
	assertSupplierFields(t, got, supplier)
}

func TestSupplierRepository_UpdateStoresNormalizedNameFromDomain(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenSupplierRepoTx(t, ctx)
	repo := NewSupplierRepository(pool)
	supplier := newSupplierRepositoryTestSupplier(t, "Nama Awal")
	if err := repo.Create(txCtx, supplier); err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	if err := supplier.Update("  Nama   Baru  ", domain.SupplierContact{}, supplierRepositoryTestTime.Add(time.Hour)); err != nil {
		t.Fatalf("Supplier.Update() error = %v", err)
	}
	if err := repo.Update(txCtx, supplier); err != nil {
		t.Fatalf("Update() error = %v", err)
	}
	stored := mustReadSupplierNormalizedName(t, txCtx, repo, supplier.ID())
	if stored != string(supplier.NormalizedName()) {
		t.Fatalf("stored normalized name = %q, want %q", stored, supplier.NormalizedName())
	}
}

func TestSupplierRepository_UpdateMissingSupplierUsesLocalConvention(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenSupplierRepoTx(t, ctx)
	repo := NewSupplierRepository(pool)
	supplier := newSupplierRepositoryTestSupplier(t, "Missing Supplier")
	if err := repo.Update(txCtx, supplier); err != nil {
		t.Fatalf("Update() error = %v, want nil for missing id convention", err)
	}
}

func TestSupplierRepository_UpdateRejectsDuplicateActiveNormalizedName(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenSupplierRepoTx(t, ctx)
	repo := NewSupplierRepository(pool)
	existing := newSupplierRepositoryTestSupplier(t, "Nama Sama")
	candidate := newSupplierRepositoryTestSupplier(t, "Nama Beda")
	if err := repo.Create(txCtx, existing); err != nil {
		t.Fatalf("Create() existing error = %v", err)
	}
	if err := repo.Create(txCtx, candidate); err != nil {
		t.Fatalf("Create() candidate error = %v", err)
	}
	if err := candidate.Update(" nama   sama ", domain.SupplierContact{}, supplierRepositoryTestTime.Add(time.Hour)); err != nil {
		t.Fatalf("Supplier.Update() error = %v", err)
	}
	err := repo.Update(txCtx, candidate)
	if !isSupplierActiveNameUniqueViolation(err) {
		t.Fatalf("Update() error = %v, want active name unique violation", err)
	}
}
