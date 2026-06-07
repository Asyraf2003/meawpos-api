package usecase

import (
	"time"

	"pos-go/internal/modules/servicecatalog/domain"
)

const (
	defaultListPage    = 1
	defaultListPerPage = 10
	maxListPerPage     = 50
	defaultLookupLimit = 20
	maxLookupLimit     = 50
)

type ServiceCatalogItemIDGenerator func() (domain.ServiceCatalogItemID, error)

type Clock func() time.Time

type ServiceCatalogItemResult struct {
	ID                 string
	Name               string
	NormalizedName     string
	DefaultPriceRupiah int64
	IsActive           bool
	Status             string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type ServiceCatalogLookupResult struct {
	ID                 string
	Name               string
	DefaultPriceRupiah int64
}

func mapServiceCatalogItemResult(item domain.ServiceCatalogItem) ServiceCatalogItemResult {
	return ServiceCatalogItemResult{
		ID:                 string(item.ID()),
		Name:               item.Name(),
		NormalizedName:     string(item.NormalizedName()),
		DefaultPriceRupiah: int64(item.DefaultPriceRupiah()),
		IsActive:           item.IsActive(),
		Status:             string(item.Status()),
		CreatedAt:          item.CreatedAt(),
		UpdatedAt:          item.UpdatedAt(),
	}
}

func mapServiceCatalogItemResults(items []domain.ServiceCatalogItem) []ServiceCatalogItemResult {
	results := make([]ServiceCatalogItemResult, 0, len(items))

	for _, item := range items {
		results = append(results, mapServiceCatalogItemResult(item))
	}

	return results
}

func mapServiceCatalogLookupResults(items []domain.ServiceCatalogItem) []ServiceCatalogLookupResult {
	results := make([]ServiceCatalogLookupResult, 0, len(items))

	for _, item := range items {
		results = append(results, ServiceCatalogLookupResult{
			ID:                 string(item.ID()),
			Name:               item.Name(),
			DefaultPriceRupiah: int64(item.DefaultPriceRupiah()),
		})
	}

	return results
}

func ensureClock(clock Clock) Clock {
	if clock != nil {
		return clock
	}

	return time.Now
}
