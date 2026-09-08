// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package postgres

import (
	"context"

	"pos-go/internal/modules/payment/cash"
	"pos-go/internal/modules/sales/domain"
)

func (s *SalesStore) InsertSale(ctx context.Context, sale domain.Sale) error {
	_, err := s.exec(ctx, `INSERT INTO sales (id,root_id,status,total_rupiah,actor_account_id,session_id,posted_at) VALUES ($1,$2,$3,$4,$5,$6,$7)`, sale.ID, sale.RootID, sale.Status, sale.Total.AmountRupiah(), sale.ActorAccountID, sale.SessionID, sale.PostedAt)
	return err
}
func (s *SalesStore) InsertSaleLine(ctx context.Context, line domain.Line) error {
	_, err := s.exec(ctx, `INSERT INTO sale_lines (id,root_id,sale_id,catalog_item_id,item_name_snapshot,unit_price_rupiah,quantity,line_total_rupiah,created_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`, line.ID, line.RootID, line.SaleID, line.CatalogItemID, line.ItemNameSnapshot, line.UnitPrice.AmountRupiah(), line.Quantity, line.LineTotal.AmountRupiah(), line.CreatedAt)
	return err
}
func (s *SalesStore) InsertCashPayment(ctx context.Context, payment cash.Payment) error {
	_, err := s.exec(ctx, `INSERT INTO cash_payments (id,root_id,sale_id,tendered_rupiah,applied_rupiah,change_rupiah,paid_at) VALUES ($1,$2,$3,$4,$5,$6,$7)`, payment.ID, payment.RootID, payment.SaleID, payment.Tendered.AmountRupiah(), payment.Applied.AmountRupiah(), payment.Change.AmountRupiah(), payment.PaidAt)
	return err
}
