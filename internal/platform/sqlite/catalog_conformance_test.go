// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package sqlite

import (
	"context"
	"database/sql"
	"path/filepath"
	"testing"

	"pos-go/internal/testsupport/catalogconformance"

	"github.com/google/uuid"
)

func TestCatalogConformance(t *testing.T) {
	catalogconformance.Run(t, newCatalogConformanceFixture)
}

func newCatalogConformanceFixture(t *testing.T) catalogconformance.Fixture {
	t.Helper()
	path := filepath.Join(t.TempDir(), "catalog.sqlite")
	db, err := Open(context.Background(), filepath.Dir(path), filepath.Base(path))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = db.Close() })
	if err := ApplyMigrations(context.Background(), db, sqliteMigrationDirectory(t)); err != nil {
		t.Fatal(err)
	}
	rootID, otherRootID := uuid.NewString(), uuid.NewString()
	seedSQLiteRoots(t, db, rootID, otherRootID)
	store := NewCatalogStore(db)
	return catalogconformance.Fixture{
		Store: store, Transactor: NewTransactor(db), RootID: rootID, OtherRootID: otherRootID,
		Verify: func(ctx context.Context) error { return verifySQLiteFixture(ctx, db, path) },
		Counts: func(ctx context.Context, id string) (int, int, error) { return sqliteCounts(ctx, db, id) },
	}
}

func sqliteCounts(ctx context.Context, db *sql.DB, id string) (int, int, error) {
	var items, prices int
	if err := db.QueryRowContext(ctx, `SELECT count(*) FROM catalog_items WHERE id=?`, id).Scan(&items); err != nil {
		return 0, 0, err
	}
	if err := db.QueryRowContext(ctx, `SELECT count(*) FROM catalog_item_prices WHERE catalog_item_id=?`, id).Scan(&prices); err != nil {
		return 0, 0, err
	}
	return items, prices, nil
}
