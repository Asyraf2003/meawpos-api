// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package usecase

import (
	"context"

	"pos-go/internal/core/audit"
	"pos-go/internal/modules/sales/domain"
)

func (uc *PostCashSale) persistSale(ctx context.Context, cmd PostCashSaleCommand, sale domain.Sale) error {
	if err := uc.posting.InsertSale(ctx, sale); err != nil {
		return err
	}
	for _, line := range sale.Lines {
		if err := uc.posting.InsertSaleLine(ctx, line); err != nil {
			return err
		}
	}
	if err := uc.posting.InsertCashPayment(ctx, sale.Cash); err != nil {
		return err
	}
	if err := uc.audit.WriteAuditEvent(ctx, audit.Event{ID: uc.newID(), RootID: cmd.RootID, ActorAccountID: cmd.ActorAccountID, SessionID: cmd.SessionID, RequestID: cmd.RequestID, AuthorityUsed: cmd.AuthorityUsed, Operation: postCashSaleOperation, ResourceType: "sale", ResourceID: sale.ID, OccurredAt: sale.PostedAt}); err != nil {
		return err
	}
	return uc.idempotency.BindIdempotency(ctx, cmd.RootID, postCashSaleOperation, cmd.IdempotencyKey, "sale", sale.ID)
}
