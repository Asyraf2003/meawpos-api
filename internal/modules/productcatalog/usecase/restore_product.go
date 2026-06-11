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

	"pos-go/internal/modules/productcatalog/ports"
)

type RestoreProduct struct {
	repository        ports.ProductRepository
	versionRepository ports.ProductVersionRepository
	auditRecorder     ports.ProductAuditRecorder
	now               func() time.Time
}

func NewRestoreProduct(
	repository ports.ProductRepository,
	versionRepository ports.ProductVersionRepository,
	auditRecorder ports.ProductAuditRecorder,
	now func() time.Time,
) *RestoreProduct {
	return &RestoreProduct{
		repository:        repository,
		versionRepository: versionRepository,
		auditRecorder:     auditRecorder,
		now:               now,
	}
}

func (uc *RestoreProduct) Execute(
	ctx context.Context,
	cmd RestoreProductCommand,
) (RestoreProductResult, error) {
	product, err := uc.repository.FindByID(ctx, cmd.ID)
	if err != nil {
		return RestoreProductResult{}, err
	}

	if err := product.Restore(); err != nil {
		return RestoreProductResult{}, err
	}

	if err := uc.repository.Update(ctx, product); err != nil {
		return RestoreProductResult{}, err
	}

	restoredAt := uc.now()
	revisionNo, err := uc.recordRestoreProductVersion(ctx, product.ID(), cmd, restoredAt)
	if err != nil {
		return RestoreProductResult{}, err
	}

	if err := uc.recordRestoreProductAudit(ctx, product.ID(), cmd, restoredAt, revisionNo); err != nil {
		return RestoreProductResult{}, err
	}

	return RestoreProductResult{
		ID:         product.ID(),
		Status:     string(product.Status()),
		RestoredAt: restoredAt,
		RevisionNo: revisionNo,
	}, nil
}
