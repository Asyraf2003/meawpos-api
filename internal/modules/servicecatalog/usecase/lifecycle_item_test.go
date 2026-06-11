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
	"context"
	"errors"
	"testing"
)

func TestActivateServiceCatalogItemMarksInactiveItemActive(t *testing.T) {
	ctx := context.Background()
	repo := newFakeServiceCatalogRepository()
	seedServiceCatalogItem(t, repo, "svc_1", "Potong Rambut", 10000)

	if _, _, err := repo.SetActive(ctx, "svc_1", false); err != nil {
		t.Fatalf("SetActive() error = %v", err)
	}

	uc := NewActivateServiceCatalogItem(repo)

	got, err := uc.Execute(ctx, ActivateServiceCatalogItemCommand{ID: "svc_1"})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if !got.IsActive {
		t.Fatal("activated item should be active")
	}
}

func TestDeactivateServiceCatalogItemMarksActiveItemInactive(t *testing.T) {
	ctx := context.Background()
	repo := newFakeServiceCatalogRepository()
	seedServiceCatalogItem(t, repo, "svc_1", "Potong Rambut", 10000)

	uc := NewDeactivateServiceCatalogItem(repo)

	got, err := uc.Execute(ctx, DeactivateServiceCatalogItemCommand{ID: "svc_1"})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if got.IsActive {
		t.Fatal("deactivated item should be inactive")
	}
}

func TestShowServiceCatalogItemMissingItemReturnsNotFound(t *testing.T) {
	ctx := context.Background()
	repo := newFakeServiceCatalogRepository()
	uc := NewShowServiceCatalogItem(repo)

	_, err := uc.Execute(ctx, ShowServiceCatalogItemCommand{ID: "svc_missing"})
	if !errors.Is(err, ErrServiceCatalogItemNotFound) {
		t.Fatalf("error = %v, want not found", err)
	}
}
