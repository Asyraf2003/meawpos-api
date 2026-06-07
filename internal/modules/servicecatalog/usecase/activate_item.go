package usecase

import (
	"context"
	"strings"

	"pos-go/internal/modules/servicecatalog/domain"
	"pos-go/internal/modules/servicecatalog/ports"
)

type ActivateServiceCatalogItem struct {
	repo ports.ServiceCatalogRepository
}

type ActivateServiceCatalogItemCommand struct {
	ID string
}

func NewActivateServiceCatalogItem(repo ports.ServiceCatalogRepository) ActivateServiceCatalogItem {
	return ActivateServiceCatalogItem{repo: repo}
}

func (uc ActivateServiceCatalogItem) Execute(
	ctx context.Context,
	cmd ActivateServiceCatalogItemCommand,
) (ServiceCatalogItemResult, error) {
	id := domain.ServiceCatalogItemID(strings.TrimSpace(cmd.ID))
	if err := domain.ValidateServiceCatalogItemID(id); err != nil {
		return ServiceCatalogItemResult{}, err
	}

	item, found, err := uc.repo.SetActive(ctx, id, true)
	if err != nil {
		return ServiceCatalogItemResult{}, err
	}

	if !found {
		return ServiceCatalogItemResult{}, ErrServiceCatalogItemNotFound
	}

	return mapServiceCatalogItemResult(item), nil
}
