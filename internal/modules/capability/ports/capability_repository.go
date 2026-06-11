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

package ports

import (
	"context"
	"errors"

	"pos-go/internal/modules/capability/domain"
)

var ErrCapabilityNotFound = errors.New("capability not found")

type CapabilityRepository interface {
	List(ctx context.Context) ([]domain.Capability, error)
	Get(ctx context.Context, key string) (domain.Capability, error)
	Save(ctx context.Context, capability domain.Capability) error
}
