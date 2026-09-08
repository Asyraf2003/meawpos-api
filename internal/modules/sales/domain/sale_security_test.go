// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package domain

import (
	"errors"
	"strings"
	"testing"
)

func TestNormalizeReasonRejectsOversizedReason(t *testing.T) {
	_, err := NormalizeReason(strings.Repeat("a", MaxReversalReasonLength+1))
	if !errors.Is(err, ErrReversalReasonTooLong) {
		t.Fatalf("NormalizeReason() error = %v, want %v", err, ErrReversalReasonTooLong)
	}
}
