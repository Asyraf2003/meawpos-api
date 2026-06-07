package usecase

import (
	"context"

	"pos-go/internal/modules/servicecatalog/ports"
)

type LookupServiceCatalogItems struct {
	repo ports.ServiceCatalogRepository
}

type LookupServiceCatalogItemsCommand struct {
	Query           string
	Limit           int
	IncludeInactive bool
}

func NewLookupServiceCatalogItems(repo ports.ServiceCatalogRepository) LookupServiceCatalogItems {
	return LookupServiceCatalogItems{repo: repo}
}

func (uc LookupServiceCatalogItems) Execute(
	ctx context.Context,
	cmd LookupServiceCatalogItemsCommand,
) ([]ServiceCatalogLookupResult, error) {
	limit := cmd.Limit
	if limit == 0 {
		limit = defaultLookupLimit
	}

	if limit < 1 || limit > maxLookupLimit {
		return nil, ErrInvalidLookupLimit
	}

	items, err := uc.repo.Lookup(ctx, ports.LookupServiceCatalogItemsFilter{
		Query:      cmd.Query,
		Limit:      limit,
		ActiveOnly: !cmd.IncludeInactive,
	})
	if err != nil {
		return nil, err
	}

	return mapServiceCatalogLookupResults(items), nil
}
