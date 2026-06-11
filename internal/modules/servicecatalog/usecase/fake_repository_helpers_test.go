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
	"testing"
	"time"

	"pos-go/internal/modules/servicecatalog/domain"
)

func seedServiceCatalogItem(
	t *testing.T,
	repo *fakeServiceCatalogRepository,
	id domain.ServiceCatalogItemID,
	name string,
	price domain.MoneyRupiah,
) domain.ServiceCatalogItem {
	t.Helper()

	item, err := domain.NewServiceCatalogItem(id, name, price, fixedNow())
	if err != nil {
		t.Fatalf("NewServiceCatalogItem() error = %v", err)
	}

	if err := repo.Create(context.Background(), item); err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	return item
}

func fixedIDGenerator(id domain.ServiceCatalogItemID) ServiceCatalogItemIDGenerator {
	return func() (domain.ServiceCatalogItemID, error) {
		return id, nil
	}
}

func fixedClock() time.Time {
	return time.Date(2026, 6, 8, 10, 0, 0, 0, time.UTC)
}

func fixedNow() time.Time {
	return fixedClock()
}
