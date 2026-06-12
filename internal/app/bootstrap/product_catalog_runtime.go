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

package bootstrap

import (
	"context"

	"pos-go/internal/modules/productcatalog/ports"

	"github.com/google/uuid"
)

type productCatalogUUIDGenerator struct{}

func (productCatalogUUIDGenerator) NewProductID() (string, error) {
	return uuid.NewString(), nil
}

type productCatalogNoopAuditRecorder struct{}

func (productCatalogNoopAuditRecorder) RecordProductAudit(
	_ context.Context,
	_ ports.ProductAuditRecord,
) error {
	return nil
}
