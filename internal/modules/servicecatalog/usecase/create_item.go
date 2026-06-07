package usecase

import (
	"context"

	"pos-go/internal/modules/servicecatalog/domain"
	"pos-go/internal/modules/servicecatalog/ports"
)

type CreateServiceCatalogItem struct {
	repo  ports.ServiceCatalogRepository
	newID ServiceCatalogItemIDGenerator
	now   Clock
}

type CreateServiceCatalogItemCommand struct {
	Name               string
	DefaultPriceRupiah int64
}

func NewCreateServiceCatalogItem(
	repo ports.ServiceCatalogRepository,
	newID ServiceCatalogItemIDGenerator,
	now Clock,
) CreateServiceCatalogItem {
	return CreateServiceCatalogItem{
		repo:  repo,
		newID: newID,
		now:   ensureClock(now),
	}
}

func (uc CreateServiceCatalogItem) Execute(
	ctx context.Context,
	cmd CreateServiceCatalogItemCommand,
) (ServiceCatalogItemResult, error) {
	if uc.newID == nil {
		return ServiceCatalogItemResult{}, ErrMissingServiceCatalogItemIDGenerator
	}

	normalizedName := domain.NormalizeName(cmd.Name)
	existing, found, err := uc.repo.FindByNormalizedName(ctx, normalizedName)
	if err != nil {
		return ServiceCatalogItemResult{}, err
	}

	if found && existing.ID() != "" {
		return ServiceCatalogItemResult{}, ErrDuplicateServiceCatalogItemNormalizedName
	}

	id, err := uc.newID()
	if err != nil {
		return ServiceCatalogItemResult{}, err
	}

	item, err := domain.NewServiceCatalogItem(
		id,
		cmd.Name,
		domain.MoneyRupiah(cmd.DefaultPriceRupiah),
		uc.now(),
	)
	if err != nil {
		return ServiceCatalogItemResult{}, err
	}

	if err := uc.repo.Create(ctx, item); err != nil {
		return ServiceCatalogItemResult{}, err
	}

	return mapServiceCatalogItemResult(item), nil
}
