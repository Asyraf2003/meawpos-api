package usecase

import (
	"context"
	"strings"

	"pos-go/internal/modules/servicecatalog/domain"
	"pos-go/internal/modules/servicecatalog/ports"
)

type ShowServiceCatalogItem struct {
	repo ports.ServiceCatalogRepository
}

type ShowServiceCatalogItemCommand struct {
	ID string
}

func NewShowServiceCatalogItem(repo ports.ServiceCatalogRepository) ShowServiceCatalogItem {
	return ShowServiceCatalogItem{repo: repo}
}

func (uc ShowServiceCatalogItem) Execute(
	ctx context.Context,
	cmd ShowServiceCatalogItemCommand,
) (ServiceCatalogItemResult, error) {
	id := domain.ServiceCatalogItemID(strings.TrimSpace(cmd.ID))
	if err := domain.ValidateServiceCatalogItemID(id); err != nil {
		return ServiceCatalogItemResult{}, err
	}

	item, found, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return ServiceCatalogItemResult{}, err
	}

	if !found {
		return ServiceCatalogItemResult{}, ErrServiceCatalogItemNotFound
	}

	return mapServiceCatalogItemResult(item), nil
}
