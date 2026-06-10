package usecase

import (
	"context"

	"pos-go/internal/modules/productcatalog/ports"
)

type LookupProducts struct {
	reader ports.ProductReader
}

func NewLookupProducts(reader ports.ProductReader) *LookupProducts {
	return &LookupProducts{
		reader: reader,
	}
}

func (uc *LookupProducts) Execute(
	ctx context.Context,
	query LookupProductsQuery,
) (LookupProductsResult, error) {
	items, err := uc.reader.Lookup(ctx, ports.ProductLookupQuery{
		Query:          query.Query,
		Limit:          query.Limit,
		IncludeDeleted: query.IncludeDeleted,
	})
	if err != nil {
		return LookupProductsResult{}, err
	}

	result := LookupProductsResult{
		Items: make([]LookupProductsItem, 0, len(items)),
	}
	for _, item := range items {
		result.Items = append(result.Items, LookupProductsItem{
			ID:              item.ID,
			Code:            item.Code,
			Name:            item.Name,
			Brand:           item.Brand,
			Size:            item.Size,
			SalePriceRupiah: item.SalePriceRupiah,
			Status:          item.Status,
		})
	}

	return result, nil
}
