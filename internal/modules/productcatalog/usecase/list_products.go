package usecase

import (
	"context"

	"pos-go/internal/modules/productcatalog/ports"
)

type ListProducts struct {
	reader ports.ProductReader
}

func NewListProducts(reader ports.ProductReader) *ListProducts {
	return &ListProducts{
		reader: reader,
	}
}

func (uc *ListProducts) Execute(
	ctx context.Context,
	query ListProductsQuery,
) (ListProductsResult, error) {
	_, err := uc.reader.List(ctx, ports.ProductListQuery{
		Search:  query.Search,
		Status:  query.Status,
		Page:    query.Page,
		PerPage: query.PerPage,
	})
	if err != nil {
		return ListProductsResult{}, err
	}

	return ListProductsResult{}, nil
}
