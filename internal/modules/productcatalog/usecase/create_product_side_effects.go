// Copyright (C) 2026 Asyraf Mubarak
//
// This file is part of gopos-api.
//
// gopos-api is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, version 3 only.
//
// gopos-api is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with gopos-api. If not, see <https://www.gnu.org/licenses/>.

package usecase

import (
	"context"
	"time"

	"pos-go/internal/modules/productcatalog/domain"
	"pos-go/internal/modules/productcatalog/ports"
)

const productCreatedEventName = "product_created"

func (uc *CreateProduct) recordCreateProductSideEffects(
	ctx context.Context,
	product *domain.Product,
	cmd CreateProductCommand,
	occurredAt time.Time,
) error {
	version := ports.ProductVersionRecord{
		ProductID:        product.ID(),
		RevisionNo:       1,
		EventName:        productCreatedEventName,
		ChangedByActorID: cmd.ActorID,
		ChangeReason:     cmd.Reason,
		ChangedAt:        occurredAt,
	}
	if err := uc.versionRepository.Append(ctx, version); err != nil {
		return err
	}

	return uc.auditRecorder.RecordProductAudit(ctx, ports.ProductAuditRecord{
		AggregateID: product.ID(),
		EventName:   productCreatedEventName,
		Operation:   "create",
		ActorID:     cmd.ActorID,
		Reason:      cmd.Reason,
		OccurredAt:  occurredAt,
		RevisionNo:  1,
	})
}
