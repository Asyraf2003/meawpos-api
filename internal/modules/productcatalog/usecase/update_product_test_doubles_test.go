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
	"pos-go/internal/modules/productcatalog/ports"
)

type fakeUpdateProductRepository struct {
	product *domain.Product
	updated *domain.Product
	err     error
}

func (f *fakeUpdateProductRepository) Create(_ context.Context, _ *domain.Product) error {
	return nil
}

func (f *fakeUpdateProductRepository) Update(_ context.Context, product *domain.Product) error {
	f.updated = product

	return f.err
}

func (f *fakeUpdateProductRepository) FindByID(
	_ context.Context,
	_ string,
) (*domain.Product, error) {
	return f.product, nil
}

type fakeUpdateProductDuplicateChecker struct {
	updateCalled bool
	productID    string
	candidate    ports.ProductDuplicateCandidate
	err          error
}

func (f *fakeUpdateProductDuplicateChecker) CheckCreateDuplicate(
	_ context.Context,
	_ ports.ProductDuplicateCandidate,
) error {
	return nil
}

func (f *fakeUpdateProductDuplicateChecker) CheckUpdateDuplicate(
	_ context.Context,
	productID string,
	candidate ports.ProductDuplicateCandidate,
) error {
	f.updateCalled = true
	f.productID = productID
	f.candidate = candidate

	return f.err
}
