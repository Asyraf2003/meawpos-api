package usecase

import (
	"context"
	"time"

	"pos-go/internal/modules/productcatalog/ports"
)

const productUpdatedEventName = "product_updated"

func (uc *UpdateProduct) recordUpdateProductVersion(
	ctx context.Context,
	productID string,
	cmd UpdateProductCommand,
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
		EventName:        productUpdatedEventName,
		ChangedByActorID: cmd.ActorID,
		ChangeReason:     cmd.Reason,
		ChangedAt:        occurredAt,
	}

	return revisionNo, uc.versionRepository.Append(ctx, version)
}

func (uc *UpdateProduct) recordUpdateProductAudit(
	ctx context.Context,
	productID string,
	cmd UpdateProductCommand,
	occurredAt time.Time,
	revisionNo int,
) error {
	return uc.auditRecorder.RecordProductAudit(ctx, ports.ProductAuditRecord{
		AggregateID: productID,
		EventName:   productUpdatedEventName,
		Operation:   "update",
		ActorID:     cmd.ActorID,
		Reason:      cmd.Reason,
		OccurredAt:  occurredAt,
		RevisionNo:  revisionNo,
	})
}
