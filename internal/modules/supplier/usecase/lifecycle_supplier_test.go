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
)

func TestActivateSupplier(t *testing.T) {
	repo := newFakeSupplierRepository()
	repo.byID["supplier-1"] = mustSupplier(t, "supplier-1", "Bengkel Jaya", false)
	uc := NewActivateSupplier(repo)

	result, err := uc.Execute(context.Background(), ActivateSupplierCommand{ID: "supplier-1"})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}
	if !result.IsActive {
		t.Fatal("IsActive = false, want true")
	}
}

func TestActivateSupplierRejectsDuplicateActiveName(t *testing.T) {
	repo := newFakeSupplierRepository()
	repo.byID["supplier-1"] = mustSupplier(t, "supplier-1", "Bengkel Jaya", false)
	repo.byID["supplier-2"] = mustSupplier(t, "supplier-2", "Bengkel Jaya", true)
	uc := NewActivateSupplier(repo)

	_, err := uc.Execute(context.Background(), ActivateSupplierCommand{ID: "supplier-1"})
	if !errors.Is(err, ErrDuplicateSupplierActiveName) {
		t.Fatalf("error = %v, want %v", err, ErrDuplicateSupplierActiveName)
	}
}

func TestDeactivateSupplier(t *testing.T) {
	repo := newFakeSupplierRepository()
	repo.byID["supplier-1"] = mustSupplier(t, "supplier-1", "Bengkel Jaya", true)
	uc := NewDeactivateSupplier(repo)

	result, err := uc.Execute(context.Background(), DeactivateSupplierCommand{
		ID:     "supplier-1",
		Reason: "temporary",
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}
	if result.IsActive {
		t.Fatal("IsActive = true, want false")
	}
}
