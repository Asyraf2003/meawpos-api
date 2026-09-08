// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package usecase

import (
	"context"
	"time"

	"pos-go/internal/core/audit"
	rootdomain "pos-go/internal/core/root/domain"
	rootports "pos-go/internal/core/root/ports"
	coretransaction "pos-go/internal/core/transaction"
)

type CreateRootCommand struct {
	Name, AccountID, SessionID, RequestID string
}

type CreateRoot struct {
	store rootports.Store
	audit audit.Writer
	tx    coretransaction.Transactor
	newID func() string
	now   func() time.Time
}

func NewCreateRoot(store rootports.Store, writer audit.Writer, tx coretransaction.Transactor, newID func() string, now func() time.Time) *CreateRoot {
	return &CreateRoot{store: store, audit: writer, tx: tx, newID: newID, now: now}
}

func (uc *CreateRoot) Execute(ctx context.Context, cmd CreateRootCommand) (rootdomain.Root, error) {
	now := uc.now().UTC()
	root, err := rootdomain.NewRoot(uc.newID(), cmd.Name, cmd.AccountID, now)
	if err != nil {
		return rootdomain.Root{}, err
	}
	err = uc.tx.RunInTx(ctx, func(txCtx context.Context) error {
		if err := uc.store.CreateRoot(txCtx, root, uc.newID()); err != nil {
			return err
		}
		return uc.audit.WriteAuditEvent(txCtx, audit.Event{
			ID: uc.newID(), RootID: root.ID, ActorAccountID: cmd.AccountID,
			SessionID: cmd.SessionID, RequestID: cmd.RequestID,
			AuthorityUsed: "authenticated_account", Operation: "root.create",
			ResourceType: "root", ResourceID: root.ID, OccurredAt: now,
		})
	})
	return root, err
}
