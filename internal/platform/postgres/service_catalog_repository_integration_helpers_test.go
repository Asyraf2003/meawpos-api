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

	"pos-go/internal/modules/servicecatalog/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func mustOpenServiceCatalogRepoTx(
	t *testing.T,
	ctx context.Context,
) (*pgxpool.Pool, context.Context) {
	t.Helper()

	pool := mustOpenIntegrationPool(t, ctx)
	tx := mustBeginIntegrationTx(t, ctx, pool)
	t.Cleanup(func() {
		_ = tx.Rollback(ctx)
		pool.Close()
	})

	return pool, contextWithTx(ctx, tx)
}

func newServiceCatalogTestItem(t *testing.T, name string) domain.ServiceCatalogItem {
	t.Helper()

	item, err := domain.NewServiceCatalogItem(
		domain.ServiceCatalogItemID(uuid.NewString()),
		name,
		domain.MoneyRupiah(15000),
		time.Now().UTC(),
	)
	if err != nil {
		t.Fatalf("NewServiceCatalogItem() error = %v", err)
	}

	return item
}
