// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package catalogconformance

import (
	"context"
	"testing"

	coretransaction "pos-go/internal/core/transaction"
	catalogports "pos-go/internal/modules/catalog/ports"
)

type Fixture struct {
	Store       catalogports.Store
	Transactor  coretransaction.Transactor
	RootID      string
	OtherRootID string
	Verify      func(context.Context) error
	Counts      func(context.Context, string) (items, prices int, err error)
}

type Factory func(*testing.T) Fixture

func Run(t *testing.T, factory Factory) {
	t.Helper()
	t.Run("schema and connection policy", func(t *testing.T) { scenarioSchema(t, factory(t)) })
	t.Run("create and read without price", func(t *testing.T) { scenarioWithoutPrice(t, factory(t)) })
	t.Run("create and read exact price", func(t *testing.T) { scenarioExactPrice(t, factory(t)) })
	t.Run("invalid input leaves no rows", func(t *testing.T) { scenarioInvalidInput(t, factory(t)) })
	t.Run("price failure rolls back item", func(t *testing.T) { scenarioRollback(t, factory(t)) })
	t.Run("cross root read is redacted", func(t *testing.T) { scenarioCrossRootRead(t, factory(t)) })
	t.Run("cross root price is rejected", func(t *testing.T) { scenarioCrossRootPrice(t, factory(t)) })
	t.Run("orphan item is rejected", func(t *testing.T) { scenarioOrphan(t, factory(t)) })
	t.Run("duplicate item id is rejected", func(t *testing.T) { scenarioDuplicateItem(t, factory(t)) })
	t.Run("duplicate price is rejected", func(t *testing.T) { scenarioDuplicatePrice(t, factory(t)) })
	t.Run("duplicate display names are accepted", func(t *testing.T) { scenarioDuplicateNames(t, factory(t)) })
	t.Run("maximum int64 price round trips", func(t *testing.T) { scenarioMaximumPrice(t, factory(t)) })
	t.Run("concurrent duplicate id stays atomic", func(t *testing.T) { scenarioConcurrentID(t, factory(t)) })
}
