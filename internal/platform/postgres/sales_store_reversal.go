// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package postgres

import (
	"context"
	"time"

	"pos-go/internal/modules/payment/cash"
	"pos-go/internal/modules/sales/domain"
)

func (s *SalesStore) LockSaleForReversal(ctx context.Context, rootID, saleID string) (domain.Sale, error) {
	return s.getSaleHeader(ctx, rootID, saleID, true)
}

func (s *SalesStore) MarkSaleReversed(ctx context.Context, rootID, saleID string, at time.Time) error {
	_, err := s.exec(ctx, `UPDATE sales SET status='REVERSED',reversed_at=$3 WHERE root_id=$1 AND id=$2`, rootID, saleID, at)
	return err
}
func (s *SalesStore) InsertReversal(ctx context.Context, reversal domain.Reversal) error {
	_, err := s.exec(ctx, `INSERT INTO sale_reversals (id,root_id,sale_id,actor_account_id,session_id,reason,reversed_at) VALUES ($1,$2,$3,$4,$5,$6,$7)`, reversal.ID, reversal.RootID, reversal.SaleID, reversal.ActorAccountID, reversal.SessionID, reversal.Reason, reversal.ReversedAt)
	return err
}
func (s *SalesStore) InsertCashRefund(ctx context.Context, refund cash.Refund) error {
	_, err := s.exec(ctx, `INSERT INTO cash_refunds (id,root_id,reversal_id,cash_payment_id,amount_rupiah,refunded_at) VALUES ($1,$2,$3,$4,$5,$6)`, refund.ID, refund.RootID, refund.ReversalID, refund.CashPaymentID, refund.Amount.AmountRupiah(), refund.RefundedAt)
	return err
}
