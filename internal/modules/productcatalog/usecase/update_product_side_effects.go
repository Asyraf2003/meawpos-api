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
