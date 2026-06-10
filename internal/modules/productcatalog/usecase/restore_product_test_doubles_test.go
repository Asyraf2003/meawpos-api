package usecase

import (
	"context"

	"pos-go/internal/modules/productcatalog/domain"
)

type restoreProductRepositoryDouble struct {
	found   *domain.Product
	updated *domain.Product
	err     error
	findErr error
}

func (d *restoreProductRepositoryDouble) Create(
	_ context.Context,
	_ *domain.Product,
) error {
	return nil
}

func (d *restoreProductRepositoryDouble) Update(
	_ context.Context,
	product *domain.Product,
) error {
	d.updated = product
	return d.err
}

func (d *restoreProductRepositoryDouble) FindByID(
	_ context.Context,
	_ string,
) (*domain.Product, error) {
	if d.findErr != nil {
		return nil, d.findErr
	}

	return d.found, nil
}
