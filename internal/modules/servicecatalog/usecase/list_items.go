package usecase

import (
	"context"

	"pos-go/internal/modules/servicecatalog/ports"
)

type ListServiceCatalogItems struct {
	repo ports.ServiceCatalogRepository
}

type ListServiceCatalogItemsCommand struct {
	Query   string
	Status  ports.ListStatusFilter
	Page    int
	PerPage int
}

func NewListServiceCatalogItems(repo ports.ServiceCatalogRepository) ListServiceCatalogItems {
	return ListServiceCatalogItems{repo: repo}
}

func (uc ListServiceCatalogItems) Execute(
	ctx context.Context,
	cmd ListServiceCatalogItemsCommand,
) ([]ServiceCatalogItemResult, error) {
	status := cmd.Status
	if status == "" {
		status = ports.ListStatusActive
	}

	page := cmd.Page
	if page <= 0 {
		page = defaultListPage
	}

	perPage := cmd.PerPage
	if perPage <= 0 {
		perPage = defaultListPerPage
	}

	if perPage > maxListPerPage {
		perPage = maxListPerPage
	}

	items, err := uc.repo.List(ctx, ports.ListServiceCatalogItemsFilter{
		Query:   cmd.Query,
		Status:  status,
		Page:    page,
		PerPage: perPage,
	})
	if err != nil {
		return nil, err
	}

	return mapServiceCatalogItemResults(items), nil
}
