// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package sqlite

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestOpenRejectsNonFileAndExistingTargets(t *testing.T) {
	for _, target := range []struct{ directory, name string }{
		{"", "catalog.sqlite"},
		{"relative", "catalog.sqlite"},
		{t.TempDir(), ""},
		{t.TempDir(), ":memory:"},
		{t.TempDir(), "file:catalog.sqlite"},
		{t.TempDir(), "../catalog.sqlite"},
	} {
		if _, err := Open(context.Background(), target.directory, target.name); !errors.Is(err, ErrInvalidDatabasePath) {
			t.Fatalf("Open(%q, %q) error=%v, want invalid path", target.directory, target.name, err)
		}
	}

	path := filepath.Join(t.TempDir(), "existing.sqlite")
	if err := os.WriteFile(path, []byte("preserve"), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := Open(context.Background(), filepath.Dir(path), filepath.Base(path)); err == nil {
		t.Fatal("Open existing database unexpectedly succeeded")
	}
	contents, err := os.ReadFile(path)
	if err != nil || string(contents) != "preserve" {
		t.Fatalf("existing file changed: contents=%q err=%v", contents, err)
	}
}
