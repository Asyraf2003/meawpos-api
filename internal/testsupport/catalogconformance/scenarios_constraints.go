// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package catalogconformance

import (
	"context"
	"testing"
)

func scenarioCrossRootRead(t *testing.T, f Fixture) {
	id := newID()
	mustCreate(t, f, id, "Private", nil)
	assertNotFound(t, f, f.OtherRootID, id)
}

func scenarioCrossRootPrice(t *testing.T, f Fixture) {
	id := newID()
	mustCreate(t, f, id, "Owned", nil)
	if err := f.Store.CreatePrice(context.Background(), price(f.OtherRootID, id, 1)); err == nil {
		t.Fatal("cross-root price unexpectedly accepted")
	}
	mustCounts(t, f, id, 1, 0)
}

func scenarioOrphan(t *testing.T, f Fixture) {
	id := newID()
	if err := f.Store.CreateItem(context.Background(), item(id, newID(), "Orphan")); err == nil {
		t.Fatal("orphan item unexpectedly accepted")
	}
	mustCounts(t, f, id, 0, 0)
}

func scenarioDuplicateItem(t *testing.T, f Fixture) {
	id := newID()
	if err := f.Store.CreateItem(context.Background(), item(id, f.RootID, "First")); err != nil {
		t.Fatal(err)
	}
	if err := f.Store.CreateItem(context.Background(), item(id, f.RootID, "Second")); err == nil {
		t.Fatal("duplicate item id unexpectedly accepted")
	}
	mustCounts(t, f, id, 1, 0)
}

func scenarioDuplicatePrice(t *testing.T, f Fixture) {
	id := newID()
	mustCreate(t, f, id, "Priced", nil)
	if err := f.Store.CreatePrice(context.Background(), price(f.RootID, id, 10)); err != nil {
		t.Fatal(err)
	}
	if err := f.Store.CreatePrice(context.Background(), price(f.RootID, id, 20)); err == nil {
		t.Fatal("duplicate price unexpectedly accepted")
	}
	mustCounts(t, f, id, 1, 1)
}

func scenarioDuplicateNames(t *testing.T, f Fixture) {
	first, second := newID(), newID()
	mustCreate(t, f, first, "Same name", nil)
	mustCreate(t, f, second, "Same name", nil)
	mustCounts(t, f, first, 1, 0)
	mustCounts(t, f, second, 1, 0)
}
