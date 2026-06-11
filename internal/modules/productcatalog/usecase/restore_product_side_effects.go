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

	"pos-go/internal/modules/productcatalog/ports"
)

const productRestoredEventName = "product_restored"

func (uc *RestoreProduct) recordRestoreProductVersion(
	ctx context.Context,
	productID string,
	cmd RestoreProductCommand,
	occurredAt time.Time,
) (int, error) {
	versions, err := uc.versionRepository.ListByProductID(ctx, productID)
	if err != nil {
		return 0, err
	}

	revisionNo := len(versions) + 1
	version := ports.ProductVersionRecord{
		ProductID:        productID,
		RevisionNo:       revisionNo,
		EventName:        productRestoredEventName,
		ChangedByActorID: cmd.ActorID,
		ChangeReason:     cmd.Reason,
		ChangedAt:        occurredAt,
	}

	return revisionNo, uc.versionRepository.Append(ctx, version)
}

func (uc *RestoreProduct) recordRestoreProductAudit(
	ctx context.Context,
	productID string,
	cmd RestoreProductCommand,
	occurredAt time.Time,
	revisionNo int,
) error {
	return uc.auditRecorder.RecordProductAudit(ctx, ports.ProductAuditRecord{
		AggregateID: productID,
		EventName:   productRestoredEventName,
		Operation:   "restore",
		ActorID:     cmd.ActorID,
		Reason:      cmd.Reason,
		OccurredAt:  occurredAt,
		RevisionNo:  revisionNo,
	})
}
