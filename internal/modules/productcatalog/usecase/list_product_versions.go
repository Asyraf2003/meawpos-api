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

	"pos-go/internal/modules/productcatalog/ports"
)

type ListProductVersions struct {
	versions ports.ProductVersionRepository
}

func NewListProductVersions(versions ports.ProductVersionRepository) *ListProductVersions {
	return &ListProductVersions{
		versions: versions,
	}
}

func (uc *ListProductVersions) Execute(
	ctx context.Context,
	query ListProductVersionsQuery,
) (ListProductVersionsResult, error) {
	records, err := uc.versions.ListByProductID(ctx, query.ProductID)
	if err != nil {
		return ListProductVersionsResult{}, err
	}

	result := ListProductVersionsResult{
		Items: make([]ListProductVersionItem, 0, len(records)),
	}
	for _, record := range records {
		result.Items = append(result.Items, ListProductVersionItem{
			ProductID:        record.ProductID,
			RevisionNo:       record.RevisionNo,
			EventName:        record.EventName,
			ChangedByActorID: record.ChangedByActorID,
			ChangeReason:     record.ChangeReason,
			ChangedAt:        record.ChangedAt,
		})
	}

	return result, nil
}
