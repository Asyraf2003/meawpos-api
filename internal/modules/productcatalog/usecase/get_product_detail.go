package usecase

import (
	"context"

	"pos-go/internal/modules/productcatalog/ports"
)

type GetProductDetail struct {
	reader ports.ProductReader
}

func NewGetProductDetail(reader ports.ProductReader) *GetProductDetail {
	return &GetProductDetail{
		reader: reader,
	}
}

func (uc *GetProductDetail) Execute(
	ctx context.Context,
	query GetProductDetailQuery,
) (GetProductDetailResult, error) {
	product, err := uc.reader.GetByID(ctx, query.ID)
	if err != nil {
		return GetProductDetailResult{}, err
	}

	return GetProductDetailResult{
		ID:                   product.ID(),
		Code:                 product.Code(),
		Name:                 product.Name(),
		NormalizedName:       product.NormalizedName(),
		Brand:                product.Brand(),
		NormalizedBrand:      product.NormalizedBrand(),
		Size:                 product.Size(),
		SalePriceRupiah:      product.SalePriceRupiah(),
		ReorderPointQty:      product.ReorderPointQty(),
		CriticalThresholdQty: product.CriticalThresholdQty(),
		Status:               string(product.Status()),
	}, nil
}
