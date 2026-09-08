// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package usecase

import (
	"context"
	"time"

	"pos-go/internal/core/audit"
	coretransaction "pos-go/internal/core/transaction"
	"pos-go/internal/modules/payment/cash"
	"pos-go/internal/modules/sales/domain"
	"pos-go/internal/modules/sales/ports"
)

const reverseCashSaleOperation = "sale.cash.reverse"

type ReverseSale struct {
	store ports.ReversalStore
	audit audit.Writer
	tx    coretransaction.Transactor
	newID func() string
	now   func() time.Time
}

func NewReverseSale(store ports.ReversalStore, writer audit.Writer, tx coretransaction.Transactor, newID func() string, now func() time.Time) *ReverseSale {
	return &ReverseSale{store: store, audit: writer, tx: tx, newID: newID, now: now}
}

func (uc *ReverseSale) Execute(ctx context.Context, cmd ReverseSaleCommand) (domain.Reversal, error) {
	reason, err := domain.NormalizeReason(cmd.Reason)
	if err != nil {
		return domain.Reversal{}, err
	}
	var result domain.Reversal
	err = uc.tx.RunInTx(ctx, func(txCtx context.Context) error {
		sale, err := uc.store.LockSaleForReversal(txCtx, cmd.RootID, cmd.SaleID)
		if err != nil {
			return err
		}
		if sale.Status == domain.StatusReversed {
			return domain.ErrSaleAlreadyReversed
		}
		now := uc.now().UTC()
		result = domain.Reversal{ID: uc.newID(), RootID: cmd.RootID, SaleID: cmd.SaleID, ActorAccountID: cmd.ActorAccountID, SessionID: cmd.SessionID, Reason: reason, ReversedAt: now}
		result.Refund = cash.Refund{ID: uc.newID(), RootID: cmd.RootID, ReversalID: result.ID, CashPaymentID: sale.Cash.ID, Amount: sale.Cash.Applied, RefundedAt: now}
		if err := uc.store.MarkSaleReversed(txCtx, cmd.RootID, cmd.SaleID, now); err != nil {
			return err
		}
		if err := uc.store.InsertReversal(txCtx, result); err != nil {
			return err
		}
		if err := uc.store.InsertCashRefund(txCtx, result.Refund); err != nil {
			return err
		}
		return uc.audit.WriteAuditEvent(txCtx, audit.Event{ID: uc.newID(), RootID: cmd.RootID, ActorAccountID: cmd.ActorAccountID, SessionID: cmd.SessionID, RequestID: cmd.RequestID, AuthorityUsed: cmd.AuthorityUsed, Operation: reverseCashSaleOperation, ResourceType: "sale_reversal", ResourceID: result.ID, Reason: reason, OccurredAt: now})
	})
	return result, err
}
