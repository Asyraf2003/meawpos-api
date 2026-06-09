package usecase

import (
	"time"

	"pos-go/internal/modules/productcatalog/ports"
)

type UpdateProduct struct {
	repository        ports.ProductRepository
	duplicateChecker  ports.ProductDuplicateChecker
	versionRepository ports.ProductVersionRepository
	auditRecorder     ports.ProductAuditRecorder
	now               func() time.Time
}

func NewUpdateProduct(
	repository ports.ProductRepository,
	duplicateChecker ports.ProductDuplicateChecker,
	versionRepository ports.ProductVersionRepository,
	auditRecorder ports.ProductAuditRecorder,
	now func() time.Time,
) *UpdateProduct {
	return &UpdateProduct{
		repository:        repository,
		duplicateChecker:  duplicateChecker,
		versionRepository: versionRepository,
		auditRecorder:     auditRecorder,
		now:               now,
	}
}
