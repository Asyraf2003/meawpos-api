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

	"pos-go/internal/modules/servicecatalog/domain"
)

func TestUpdateServiceCatalogItemChangesNameAndPrice(t *testing.T) {
	ctx := context.Background()
	repo := newFakeServiceCatalogRepository()
	seedServiceCatalogItem(t, repo, "svc_1", "Potong Rambut", 10000)

	uc := NewUpdateServiceCatalogItem(repo, fixedClock)

	got, err := uc.Execute(ctx, UpdateServiceCatalogItemCommand{
		ID:                 "svc_1",
		Name:               "Cuci Motor",
		DefaultPriceRupiah: 15000,
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if got.Name != "Cuci Motor" {
		t.Fatalf("Name = %q, want %q", got.Name, "Cuci Motor")
	}

	if got.NormalizedName != "cuci motor" {
		t.Fatalf("NormalizedName = %q, want %q", got.NormalizedName, "cuci motor")
	}

	if got.DefaultPriceRupiah != 15000 {
		t.Fatalf("DefaultPriceRupiah = %d, want %d", got.DefaultPriceRupiah, 15000)
	}
}

func TestUpdateServiceCatalogItemRejectsDuplicateNormalizedName(t *testing.T) {
	ctx := context.Background()
	repo := newFakeServiceCatalogRepository()
	seedServiceCatalogItem(t, repo, "svc_1", "Potong Rambut", 10000)
	seedServiceCatalogItem(t, repo, "svc_2", "Cuci Motor", 15000)

	uc := NewUpdateServiceCatalogItem(repo, fixedClock)

	_, err := uc.Execute(ctx, UpdateServiceCatalogItemCommand{
		ID:                 "svc_2",
		Name:               "potong   rambut",
		DefaultPriceRupiah: 20000,
	})
	if !errors.Is(err, ErrDuplicateServiceCatalogItemNormalizedName) {
		t.Fatalf("error = %v, want duplicate normalized name", err)
	}
}

func TestUpdateServiceCatalogItemRejectsInvalidPrice(t *testing.T) {
	ctx := context.Background()
	repo := newFakeServiceCatalogRepository()
	seedServiceCatalogItem(t, repo, "svc_1", "Potong Rambut", 10000)

	uc := NewUpdateServiceCatalogItem(repo, fixedClock)

	_, err := uc.Execute(ctx, UpdateServiceCatalogItemCommand{
		ID:                 "svc_1",
		Name:               "Potong Rambut",
		DefaultPriceRupiah: 0,
	})
	if !errors.Is(err, domain.ErrInvalidServiceCatalogItemDefaultPrice) {
		t.Fatalf("error = %v, want invalid default price", err)
	}
}
