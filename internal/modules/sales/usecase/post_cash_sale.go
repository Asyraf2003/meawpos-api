// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package usecase

import (
	"context"
	"strings"
	"time"

	"pos-go/internal/core/audit"
	coretransaction "pos-go/internal/core/transaction"
	"pos-go/internal/modules/sales/domain"
	"pos-go/internal/modules/sales/ports"
)

const postCashSaleOperation = "sale.cash.post"

type PostCashSale struct {
	pricing     ports.PricingReader
	posting     ports.PostingStore
	idempotency ports.IdempotencyStore
	reader      ports.SaleReader
	audit       audit.Writer
	tx          coretransaction.Transactor
	newID       func() string
	now         func() time.Time
}

func NewPostCashSale(pricing ports.PricingReader, posting ports.PostingStore, idem ports.IdempotencyStore, reader ports.SaleReader, writer audit.Writer, tx coretransaction.Transactor, newID func() string, now func() time.Time) *PostCashSale {
	return &PostCashSale{pricing: pricing, posting: posting, idempotency: idem, reader: reader, audit: writer, tx: tx, newID: newID, now: now}
}

func (uc *PostCashSale) Execute(ctx context.Context, cmd PostCashSaleCommand) (PostCashSaleResult, error) {
	cmd.IdempotencyKey = strings.TrimSpace(cmd.IdempotencyKey)
	if cmd.IdempotencyKey == "" {
		return PostCashSaleResult{}, domain.ErrIdempotencyKeyRequired
	}
	if len(cmd.Items) == 0 {
		return PostCashSaleResult{}, domain.ErrItemsRequired
	}
	if cmd.PaymentType != "cash" {
		return PostCashSaleResult{}, domain.ErrUnsupportedPayment
	}
	for _, item := range cmd.Items {
		if item.Quantity <= 0 {
			return PostCashSaleResult{}, domain.ErrInvalidQuantity
		}
	}
	fingerprint, err := requestFingerprint(cmd)
	if err != nil {
		return PostCashSaleResult{}, err
	}
	var result PostCashSaleResult
	err = uc.tx.RunInTx(ctx, func(txCtx context.Context) error {
		claim, err := uc.idempotency.ClaimIdempotency(txCtx, cmd.RootID, postCashSaleOperation, cmd.IdempotencyKey, fingerprint, uc.now().UTC())
		if err != nil {
			return err
		}
		if claim.Existing {
			if claim.Fingerprint != fingerprint {
				return domain.ErrIdempotencyConflict
			}
			if _, err := uc.reader.GetSale(txCtx, cmd.RootID, claim.ResourceID); err != nil {
				return err
			}
			result = PostCashSaleResult{SaleID: claim.ResourceID, Replayed: true}
			return nil
		}
		sale, err := uc.buildSale(txCtx, cmd)
		if err != nil {
			return err
		}
		if err := uc.persistSale(txCtx, cmd, sale); err != nil {
			return err
		}
		result = PostCashSaleResult{SaleID: sale.ID}
		return nil
	})
	return result, err
}
