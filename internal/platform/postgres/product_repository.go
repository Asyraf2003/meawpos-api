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

package postgres

import (
	"errors"

	"pos-go/internal/modules/productcatalog/ports"

	"github.com/jackc/pgx/v5/pgxpool"
)

var errProductRepositoryNotImplemented = errors.New("product postgres repository behavior not implemented")

type ProductRepository struct {
	pool *pgxpool.Pool
}

func NewProductRepository(pool *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{pool: pool}
}

var _ ports.ProductRepository = (*ProductRepository)(nil)
var _ ports.ProductReader = (*ProductRepository)(nil)
var _ ports.ProductVersionRepository = (*ProductRepository)(nil)
var _ ports.ProductDuplicateChecker = (*ProductRepository)(nil)
