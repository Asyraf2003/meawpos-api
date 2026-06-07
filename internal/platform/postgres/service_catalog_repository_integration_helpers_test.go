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
