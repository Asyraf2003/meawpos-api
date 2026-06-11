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

	"pos-go/internal/modules/productcatalog/domain"
)

func softDeleteProductForListTest(
	t *testing.T,
	repo *ProductRepository,
	txCtx context.Context,
	product *domain.Product,
) {
	t.Helper()

	if err := product.SoftDelete(domain.DeleteInput{
		DeletedAt: time.Now().UTC(),
		Reason:    "integration test",
	}); err != nil {
		t.Fatalf("SoftDelete() error = %v", err)
	}
	if err := repo.Update(txCtx, product); err != nil {
		t.Fatalf("Update() deleted error = %v", err)
	}
}
