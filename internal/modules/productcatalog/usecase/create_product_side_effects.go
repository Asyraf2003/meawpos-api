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
