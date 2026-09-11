// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package sqlite

import (
	"context"
	"database/sql"

	coretransaction "pos-go/internal/core/transaction"
)

type txContextKey struct{}

type Transactor struct{ db *sql.DB }

func NewTransactor(db *sql.DB) *Transactor { return &Transactor{db: db} }

func (t *Transactor) RunInTx(ctx context.Context, fn func(context.Context) error) error {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	committed := false
	defer func() {
		if !committed {
			_ = tx.Rollback()
		}
	}()
	if err := fn(context.WithValue(ctx, txContextKey{}, tx)); err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	committed = true
	return nil
}

func txFromContext(ctx context.Context) (*sql.Tx, bool) {
	tx, ok := ctx.Value(txContextKey{}).(*sql.Tx)
	return tx, ok
}

var _ coretransaction.Transactor = (*Transactor)(nil)
