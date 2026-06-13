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

	"pos-go/internal/modules/supplier/domain"
)

func TestCreateSupplierCreatesActiveSupplier(t *testing.T) {
	repo := newFakeSupplierRepository()
	uc := NewCreateSupplier(
		repo,
		func() (domain.SupplierID, error) { return "supplier-1", nil },
		func() time.Time { return time.Date(2026, 6, 14, 10, 0, 0, 0, time.UTC) },
	)

	result, err := uc.Execute(context.Background(), CreateSupplierCommand{
		Name: " Bengkel   Jaya ",
		Contact: SupplierContactInput{
			Phone: " 0812 ",
			Email: " owner@example.test ",
		},
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if result.ID != "supplier-1" {
		t.Fatalf("ID = %q", result.ID)
	}
	if result.Name != "Bengkel Jaya" {
		t.Fatalf("Name = %q", result.Name)
	}
	if result.NormalizedName != "bengkel jaya" {
		t.Fatalf("NormalizedName = %q", result.NormalizedName)
	}
	if !result.IsActive {
		t.Fatal("IsActive = false, want true")
	}
	if repo.createCalls != 1 {
		t.Fatalf("createCalls = %d", repo.createCalls)
	}
}

func TestCreateSupplierRejectsDuplicateActiveName(t *testing.T) {
	repo := newFakeSupplierRepository()
	repo.byID["supplier-existing"] = mustSupplier(t, "supplier-existing", "Bengkel Jaya", true)

	uc := NewCreateSupplier(
		repo,
		func() (domain.SupplierID, error) { return "supplier-1", nil },
		func() time.Time { return time.Now() },
	)

	_, err := uc.Execute(context.Background(), CreateSupplierCommand{Name: "bengkel jaya"})
	if !errors.Is(err, ErrDuplicateSupplierActiveName) {
		t.Fatalf("error = %v, want %v", err, ErrDuplicateSupplierActiveName)
	}
}

func TestCreateSupplierAllowsInactiveNameReuse(t *testing.T) {
	repo := newFakeSupplierRepository()
	repo.byID["supplier-existing"] = mustSupplier(t, "supplier-existing", "Bengkel Jaya", false)

	uc := NewCreateSupplier(
		repo,
		func() (domain.SupplierID, error) { return "supplier-1", nil },
		func() time.Time { return time.Now() },
	)

	_, err := uc.Execute(context.Background(), CreateSupplierCommand{Name: "bengkel jaya"})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}
}
