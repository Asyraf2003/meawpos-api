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
	"time"

	"pos-go/internal/modules/productcatalog/domain"
	"pos-go/internal/modules/productcatalog/ports"
)

type ProductIDGenerator interface {
	NewProductID() (string, error)
}

type CreateProduct struct {
	repository        ports.ProductRepository
	duplicateChecker  ports.ProductDuplicateChecker
	versionRepository ports.ProductVersionRepository
	auditRecorder     ports.ProductAuditRecorder
	idGenerator       ProductIDGenerator
	now               func() time.Time
}

func NewCreateProduct(
	repository ports.ProductRepository,
	duplicateChecker ports.ProductDuplicateChecker,
	versionRepository ports.ProductVersionRepository,
	auditRecorder ports.ProductAuditRecorder,
	idGenerator ProductIDGenerator,
	now func() time.Time,
) *CreateProduct {
	return &CreateProduct{
		repository:        repository,
		duplicateChecker:  duplicateChecker,
		versionRepository: versionRepository,
		auditRecorder:     auditRecorder,
		idGenerator:       idGenerator,
		now:               now,
	}
}

func (uc *CreateProduct) Execute(
	ctx context.Context,
	cmd CreateProductCommand,
) (CreateProductResult, error) {
	id, err := uc.idGenerator.NewProductID()
	if err != nil {
		return CreateProductResult{}, err
	}

	product, err := domain.NewProduct(domain.ProductInput{
		ID:                   id,
		Code:                 cmd.Code,
		Name:                 cmd.Name,
		Brand:                cmd.Brand,
		Size:                 cmd.Size,
		SalePriceRupiah:      cmd.SalePriceRupiah,
		ReorderPointQty:      cmd.ReorderPointQty,
		CriticalThresholdQty: cmd.CriticalThresholdQty,
	})
	if err != nil {
		return CreateProductResult{}, err
	}

	if err := uc.duplicateChecker.CheckCreateDuplicate(ctx, ports.ProductDuplicateCandidate{
		Code:            product.Code(),
		NormalizedName:  product.NormalizedName(),
		NormalizedBrand: product.NormalizedBrand(),
		Size:            product.Size(),
	}); err != nil {
		return CreateProductResult{}, err
	}

	if err := uc.repository.Create(ctx, product); err != nil {
		return CreateProductResult{}, err
	}

	now := uc.now()
	if err := uc.recordCreateProductSideEffects(ctx, product, cmd, now); err != nil {
		return CreateProductResult{}, err
	}

	return createProductResultFromDomain(product, now), nil
}
