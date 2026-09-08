// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package postgres

import (
	"context"
	"errors"
	"time"

	"pos-go/internal/core/money"
	"pos-go/internal/modules/payment/cash"
	"pos-go/internal/modules/sales/domain"

	"github.com/jackc/pgx/v5"
)

func (s *SalesStore) GetSale(ctx context.Context, rootID, saleID string) (domain.Sale, error) {
	sale, err := s.getSaleHeader(ctx, rootID, saleID, false)
	if err != nil {
		return domain.Sale{}, err
	}
	rows, err := s.query(ctx, `SELECT id,catalog_item_id,item_name_snapshot,unit_price_rupiah,quantity,line_total_rupiah,created_at FROM sale_lines WHERE root_id=$1 AND sale_id=$2 ORDER BY created_at,id`, rootID, saleID)
	if err != nil {
		return domain.Sale{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var line domain.Line
		var unit, total int64
		err := rows.Scan(&line.ID, &line.CatalogItemID, &line.ItemNameSnapshot, &unit, &line.Quantity, &total, &line.CreatedAt)
		if err != nil {
			return domain.Sale{}, err
		}
		line.RootID = rootID
		line.SaleID = saleID
		line.UnitPrice = money.IDR(unit)
		line.LineTotal = money.IDR(total)
		sale.Lines = append(sale.Lines, line)
	}
	return sale, rows.Err()
}

func (s *SalesStore) getSaleHeader(ctx context.Context, rootID, saleID string, lock bool) (domain.Sale, error) {
	query := `SELECT s.id,s.status,s.total_rupiah,s.actor_account_id,s.session_id,s.posted_at,
		p.id,p.tendered_rupiah,p.applied_rupiah,p.change_rupiah,p.paid_at,
		r.id,r.actor_account_id,r.session_id,r.reason,r.reversed_at,
		f.id,f.amount_rupiah,f.refunded_at
		FROM sales s JOIN cash_payments p ON p.root_id=s.root_id AND p.sale_id=s.id
		LEFT JOIN sale_reversals r ON r.root_id=s.root_id AND r.sale_id=s.id
		LEFT JOIN cash_refunds f ON f.root_id=r.root_id AND f.reversal_id=r.id
		WHERE s.root_id=$1 AND s.id=$2`
	if lock {
		query += " FOR UPDATE OF s,p"
	}
	return scanSaleHeader(s.queryRow(ctx, query, rootID, saleID), rootID)
}

func scanSaleHeader(row pgx.Row, rootID string) (domain.Sale, error) {
	var s domain.Sale
	var total, tendered, applied, change int64
	var reversalID, reversalActor, reversalSession, reversalReason, refundID *string
	var reversedAt, refundedAt *time.Time
	var refundAmount *int64
	err := row.Scan(&s.ID, &s.Status, &total, &s.ActorAccountID, &s.SessionID, &s.PostedAt, &s.Cash.ID, &tendered, &applied, &change, &s.Cash.PaidAt, &reversalID, &reversalActor, &reversalSession, &reversalReason, &reversedAt, &refundID, &refundAmount, &refundedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return domain.Sale{}, domain.ErrSaleNotFound
	}
	if err != nil {
		return domain.Sale{}, err
	}
	s.RootID = rootID
	s.Total = money.IDR(total)
	s.Cash.RootID = rootID
	s.Cash.SaleID = s.ID
	s.Cash.Tendered = money.IDR(tendered)
	s.Cash.Applied = money.IDR(applied)
	s.Cash.Change = money.IDR(change)
	if reversalID != nil {
		if reversalActor == nil || reversalSession == nil || reversalReason == nil || reversedAt == nil || refundID == nil || refundAmount == nil || refundedAt == nil {
			return domain.Sale{}, errors.New("incomplete sale reversal persistence")
		}
		s.Reversal = &domain.Reversal{ID: *reversalID, RootID: rootID, SaleID: s.ID, ActorAccountID: *reversalActor, SessionID: *reversalSession, Reason: *reversalReason, ReversedAt: *reversedAt, Refund: cash.Refund{ID: *refundID, RootID: rootID, ReversalID: *reversalID, CashPaymentID: s.Cash.ID, Amount: money.IDR(*refundAmount), RefundedAt: *refundedAt}}
	}
	return s, nil
}
