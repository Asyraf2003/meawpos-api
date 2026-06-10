package usecase

import (
	"time"

	"pos-go/internal/modules/productcatalog/ports"
)

type RestoreProduct struct {
	repository        ports.ProductRepository
	versionRepository ports.ProductVersionRepository
	auditRecorder     ports.ProductAuditRecorder
	now               func() time.Time
}

func NewRestoreProduct(
	repository ports.ProductRepository,
	versionRepository ports.ProductVersionRepository,
	auditRecorder ports.ProductAuditRecorder,
	now func() time.Time,
) *RestoreProduct {
	return &RestoreProduct{
		repository:        repository,
		versionRepository: versionRepository,
		auditRecorder:     auditRecorder,
		now:               now,
	}
}
