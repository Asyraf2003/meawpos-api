// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package core

import (
	"errors"
	"strings"
	"testing"
	"time"
)

func TestNewItemRejectsOversizedName(t *testing.T) {
	_, err := NewItem("item", "root", strings.Repeat("a", MaxNameLength+1), time.Now())
	if !errors.Is(err, ErrNameTooLong) {
		t.Fatalf("NewItem() error = %v, want %v", err, ErrNameTooLong)
	}
}
