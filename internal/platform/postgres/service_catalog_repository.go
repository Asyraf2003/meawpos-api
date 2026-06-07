package postgres

import (
	"pos-go/internal/modules/servicecatalog/ports"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ServiceCatalogRepository struct {
	pool *pgxpool.Pool
}

func NewServiceCatalogRepository(pool *pgxpool.Pool) *ServiceCatalogRepository {
	return &ServiceCatalogRepository{pool: pool}
}

var _ ports.ServiceCatalogRepository = (*ServiceCatalogRepository)(nil)
