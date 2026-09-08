// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package bootstrap

import "testing"

func TestComponentSetFailsFastForMissingDependency(t *testing.T) {
	if _, err := newComponentSet([]string{"sales", "payment.cash"}); err == nil {
		t.Fatal("expected missing dependency error")
	}
}

func TestComponentSetUsesExplicitConfiguredSet(t *testing.T) {
	set, err := newComponentSet([]string{"catalog.core"})
	if err != nil || !set["catalog.core"] || set["sales"] {
		t.Fatalf("set=%v err=%v", set, err)
	}
}
