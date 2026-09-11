// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"
)

func sqliteMigrationDirectory(t *testing.T) string {
	t.Helper()
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("locate SQLite conformance test")
	}
	return filepath.Join(filepath.Dir(filename), "..", "..", "..", "migrations", "sqlite")
}

func seedSQLiteRoots(t *testing.T, db *sql.DB, rootIDs ...string) {
	t.Helper()
	for _, id := range rootIDs {
		now := time.Now().UTC().UnixMicro()
		if _, err := db.Exec(`INSERT INTO roots (id,name,created_at,updated_at) VALUES (?,?,?,?)`,
			id, "Conformance Root", now, now); err != nil {
			t.Fatalf("seed SQLite root: %v", err)
		}
	}
}

func verifySQLiteFixture(ctx context.Context, db *sql.DB, path string) error {
	checks := []struct {
		query string
		want  any
	}{
		{"PRAGMA foreign_keys", 1},
		{"PRAGMA busy_timeout", 5000},
		{"PRAGMA journal_mode", "delete"},
		{"PRAGMA synchronous", 2},
	}
	for _, check := range checks {
		var got any
		if err := db.QueryRowContext(ctx, check.query).Scan(&got); err != nil {
			return err
		}
		if fmt.Sprint(got) != fmt.Sprint(check.want) {
			return fmt.Errorf("%s=%v, want %v", check.query, got, check.want)
		}
	}
	var strictTables int
	if err := db.QueryRowContext(ctx, `SELECT count(*) FROM pragma_table_list
		WHERE name IN ('roots','catalog_items','catalog_item_prices') AND strict=1`).Scan(&strictTables); err != nil {
		return err
	}
	if strictTables != 3 {
		return fmt.Errorf("strict selected tables=%d, want 3", strictTables)
	}
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	if info.Mode().Perm()&0o077 != 0 {
		return fmt.Errorf("database permissions=%#o, want owner-only", info.Mode().Perm())
	}
	return nil
}
