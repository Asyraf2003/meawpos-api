package ports

import (
	"context"

	"pos-go/internal/modules/servicecatalog/domain"
)

type ListStatusFilter string

const (
	ListStatusActive   ListStatusFilter = "active"
	ListStatusInactive ListStatusFilter = "inactive"
	ListStatusAll      ListStatusFilter = "all"
)

type ListServiceCatalogItemsFilter struct {
	Query   string
	Status  ListStatusFilter
	Page    int
	PerPage int
}

type LookupServiceCatalogItemsFilter struct {
	Query      string
	Limit      int
	ActiveOnly bool
}

type ServiceCatalogRepository interface {
	Create(ctx context.Context, item domain.ServiceCatalogItem) error
	Update(ctx context.Context, item domain.ServiceCatalogItem) error
	FindByID(ctx context.Context, id domain.ServiceCatalogItemID) (domain.ServiceCatalogItem, bool, error)
	FindByNormalizedName(ctx context.Context, normalizedName domain.NormalizedName) (domain.ServiceCatalogItem, bool, error)
	List(ctx context.Context, filter ListServiceCatalogItemsFilter) ([]domain.ServiceCatalogItem, error)
	Lookup(ctx context.Context, filter LookupServiceCatalogItemsFilter) ([]domain.ServiceCatalogItem, error)
	SetActive(ctx context.Context, id domain.ServiceCatalogItemID, active bool) (domain.ServiceCatalogItem, bool, error)
}
