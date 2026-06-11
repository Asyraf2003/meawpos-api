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

type SoftDeleteProduct struct {
	repository        ports.ProductRepository
	versionRepository ports.ProductVersionRepository
	auditRecorder     ports.ProductAuditRecorder
	now               func() time.Time
}

func NewSoftDeleteProduct(
	repository ports.ProductRepository,
	versionRepository ports.ProductVersionRepository,
	auditRecorder ports.ProductAuditRecorder,
	now func() time.Time,
) *SoftDeleteProduct {
	return &SoftDeleteProduct{
		repository:        repository,
		versionRepository: versionRepository,
		auditRecorder:     auditRecorder,
		now:               now,
	}
}

func (uc *SoftDeleteProduct) Execute(
	ctx context.Context,
	cmd SoftDeleteProductCommand,
) (SoftDeleteProductResult, error) {
	product, err := uc.repository.FindByID(ctx, cmd.ID)
	if err != nil {
		return SoftDeleteProductResult{}, err
	}

	deletedAt := uc.now()
	if err := product.SoftDelete(domain.DeleteInput{
		DeletedAt:        deletedAt,
		DeletedByActorID: cmd.ActorID,
		Reason:           cmd.Reason,
	}); err != nil {
		return SoftDeleteProductResult{}, err
	}

	if err := uc.repository.Update(ctx, product); err != nil {
		return SoftDeleteProductResult{}, err
	}

	revisionNo, err := uc.recordSoftDeleteProductVersion(ctx, product.ID(), cmd, deletedAt)
	if err != nil {
		return SoftDeleteProductResult{}, err
	}

	if err := uc.recordSoftDeleteProductAudit(ctx, product.ID(), cmd, deletedAt, revisionNo); err != nil {
		return SoftDeleteProductResult{}, err
	}

	return SoftDeleteProductResult{
		ID:         product.ID(),
		Status:     string(product.Status()),
		DeletedAt:  deletedAt,
		RevisionNo: revisionNo,
	}, nil
}
