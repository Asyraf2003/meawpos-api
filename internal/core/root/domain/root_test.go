// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package domain

import (
	"errors"
	"strings"
	"testing"
	"time"
)

func TestNewRootRejectsOversizedName(t *testing.T) {
	_, err := NewRoot("root", strings.Repeat("a", MaxRootNameLength+1), "owner", time.Now())
	if !errors.Is(err, ErrRootNameTooLong) {
		t.Fatalf("NewRoot() error = %v, want %v", err, ErrRootNameTooLong)
	}
}
