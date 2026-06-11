// Copyright (C) 2026 Asyraf Mubarak
//
// This file is part of gopos-api.
//
// gopos-api is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, version 3 only.
//
// gopos-api is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with gopos-api. If not, see <https://www.gnu.org/licenses/>.

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
