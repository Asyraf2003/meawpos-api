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
	_, err := uc.reader.GetByID(ctx, query.ID)
	if err != nil {
		return GetProductDetailResult{}, err
	}

	return GetProductDetailResult{}, nil
}
