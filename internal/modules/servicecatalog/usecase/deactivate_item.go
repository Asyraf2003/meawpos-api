package usecase

import (
	"context"
	"strings"

	"pos-go/internal/modules/servicecatalog/domain"
	"pos-go/internal/modules/servicecatalog/ports"
)

type DeactivateServiceCatalogItem struct {
	repo ports.ServiceCatalogRepository
}

type DeactivateServiceCatalogItemCommand struct {
	ID string
}

func NewDeactivateServiceCatalogItem(repo ports.ServiceCatalogRepository) DeactivateServiceCatalogItem {
	return DeactivateServiceCatalogItem{repo: repo}
}

func (uc DeactivateServiceCatalogItem) Execute(
	ctx context.Context,
	cmd DeactivateServiceCatalogItemCommand,
) (ServiceCatalogItemResult, error) {
	id := domain.ServiceCatalogItemID(strings.TrimSpace(cmd.ID))
	if err := domain.ValidateServiceCatalogItemID(id); err != nil {
		return ServiceCatalogItemResult{}, err
	}

	item, found, err := uc.repo.SetActive(ctx, id, false)
	if err != nil {
		return ServiceCatalogItemResult{}, err
	}

	if !found {
		return ServiceCatalogItemResult{}, ErrServiceCatalogItemNotFound
	}

	return mapServiceCatalogItemResult(item), nil
}
