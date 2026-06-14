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
	"database/sql"
	"fmt"
	"time"

	"pos-go/internal/modules/supplier/domain"
)

type supplierScanner interface {
	Scan(dest ...any) error
}

func supplierSelectSQL() string {
	return `
		SELECT
			id,
			name,
			name_normalized,
			phone,
			email,
			address,
			notes,
			is_active,
			created_at,
			updated_at
		FROM suppliers
	`
}

func scanSupplier(row supplierScanner) (domain.Supplier, error) {
	var id string
	var name string
	var normalizedName string
	var phone sql.NullString
	var email sql.NullString
	var address sql.NullString
	var notes sql.NullString
	var isActive bool
	var createdAt time.Time
	var updatedAt time.Time

	err := row.Scan(
		&id,
		&name,
		&normalizedName,
		&phone,
		&email,
		&address,
		&notes,
		&isActive,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return domain.Supplier{}, err
	}

	supplier, err := domain.NewSupplier(
		domain.SupplierID(id),
		name,
		domain.SupplierContact{
			Phone:   supplierStringValue(phone),
			Email:   supplierStringValue(email),
			Address: supplierStringValue(address),
			Notes:   supplierStringValue(notes),
		},
		createdAt,
	)
	if err != nil {
		return domain.Supplier{}, err
	}

	if supplier.NormalizedName() != domain.NormalizedName(normalizedName) {
		return domain.Supplier{}, fmt.Errorf("supplier normalized name mismatch for id %q", id)
	}

	if isActive {
		supplier.Activate(updatedAt)
	} else {
		supplier.Deactivate(updatedAt)
	}

	return supplier, nil
}

func supplierArgs(supplier domain.Supplier) []any {
	return []any{
		string(supplier.ID()),
		supplier.Name(),
		string(supplier.NormalizedName()),
		supplierNullableText(supplier.Phone()),
		supplierNullableText(supplier.Email()),
		supplierNullableText(supplier.Address()),
		supplierNullableText(supplier.Notes()),
		supplier.IsActive(),
		supplier.CreatedAt(),
		supplier.UpdatedAt(),
	}
}

func supplierStringValue(value sql.NullString) string {
	if !value.Valid {
		return ""
	}

	return value.String
}

func supplierNullableText(value string) any {
	if value == "" {
		return nil
	}

	return value
}
