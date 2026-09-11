// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"io/fs"
	"os"
	"sort"
	"strings"
)

func ApplyMigrations(ctx context.Context, db *sql.DB, directory string) (returnErr error) {
	root, err := os.OpenRoot(directory)
	if err != nil {
		return fmt.Errorf("open SQLite migrations: %w", err)
	}
	defer func() {
		if err := root.Close(); returnErr == nil && err != nil {
			returnErr = fmt.Errorf("close SQLite migrations: %w", err)
		}
	}()
	entries, err := fs.ReadDir(root.FS(), ".")
	if err != nil {
		return fmt.Errorf("read SQLite migrations: %w", err)
	}
	names := make([]string, 0, len(entries))
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".up.sql") {
			names = append(names, entry.Name())
		}
	}
	sort.Strings(names)
	if len(names) == 0 {
		return errorsNoMigrations(directory)
	}
	for _, name := range names {
		migration, err := root.ReadFile(name)
		if err != nil {
			return fmt.Errorf("read SQLite migration %s: %w", name, err)
		}
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			return fmt.Errorf("begin SQLite migration %s: %w", name, err)
		}
		if _, err := tx.ExecContext(ctx, string(migration)); err != nil {
			_ = tx.Rollback()
			return fmt.Errorf("apply SQLite migration %s: %w", name, err)
		}
		if err := tx.Commit(); err != nil {
			return fmt.Errorf("commit SQLite migration %s: %w", name, err)
		}
	}
	return nil
}

func errorsNoMigrations(directory string) error {
	return fmt.Errorf("no SQLite migrations in %s", directory)
}
