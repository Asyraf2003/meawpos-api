// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	_ "modernc.org/sqlite"
)

var ErrInvalidDatabasePath = errors.New("invalid SQLite database path")

func Open(ctx context.Context, directory, name string) (*sql.DB, error) {
	if !filepath.IsAbs(directory) || name == "" || name == "." || name == ":memory:" ||
		strings.HasPrefix(strings.ToLower(name), "file:") || filepath.Base(name) != name {
		return nil, ErrInvalidDatabasePath
	}
	root, err := os.OpenRoot(directory)
	if err != nil {
		return nil, fmt.Errorf("open SQLite database directory: %w", err)
	}
	file, err := root.OpenFile(name, os.O_CREATE|os.O_EXCL|os.O_RDWR, 0o600)
	if err != nil {
		_ = root.Close()
		return nil, fmt.Errorf("create SQLite database: %w", err)
	}
	if err := file.Close(); err != nil {
		_ = root.Close()
		return nil, fmt.Errorf("close SQLite database: %w", err)
	}
	if err := root.Close(); err != nil {
		return nil, fmt.Errorf("close SQLite database directory: %w", err)
	}

	path := filepath.Join(directory, name)
	dsn := (&url.URL{Scheme: "file", Path: path, RawQuery: url.Values{
		"_pragma": []string{
			"foreign_keys(1)",
			"busy_timeout(5000)",
			"journal_mode(DELETE)",
			"synchronous(FULL)",
		},
		"_txlock": []string{"immediate"},
	}.Encode()}).String()
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, fmt.Errorf("open SQLite database: %w", err)
	}
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	if err := db.PingContext(ctx); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("ping SQLite database: %w", err)
	}
	return db, nil
}
