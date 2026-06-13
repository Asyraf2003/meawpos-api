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
// along with gopos-api. If not, see https://www.gnu.org/licenses/.

package domain

import (
	"strings"
	"time"
)

func NewSupplier(
	id SupplierID,
	name string,
	contact SupplierContact,
	now time.Time,
) (Supplier, error) {
	id = SupplierID(strings.TrimSpace(string(id)))
	if err := ValidateSupplierID(id); err != nil {
		return Supplier{}, err
	}
	if err := ValidateSupplierName(name); err != nil {
		return Supplier{}, err
	}

	return Supplier{
		id:             id,
		name:           normalizeDisplayName(name),
		normalizedName: NormalizeName(name),
		phone:          normalizeDisplayText(contact.Phone),
		email:          normalizeDisplayText(contact.Email),
		address:        normalizeDisplayText(contact.Address),
		notes:          normalizeDisplayText(contact.Notes),
		isActive:       true,
		createdAt:      now,
		updatedAt:      now,
	}, nil
}

func (s *Supplier) Update(
	name string,
	contact SupplierContact,
	now time.Time,
) error {
	if err := ValidateSupplierName(name); err != nil {
		return err
	}

	s.name = normalizeDisplayName(name)
	s.normalizedName = NormalizeName(name)
	s.phone = normalizeDisplayText(contact.Phone)
	s.email = normalizeDisplayText(contact.Email)
	s.address = normalizeDisplayText(contact.Address)
	s.notes = normalizeDisplayText(contact.Notes)
	s.updatedAt = now

	return nil
}

func (s *Supplier) Activate(now time.Time) {
	s.isActive = true
	s.updatedAt = now
}

func (s *Supplier) Deactivate(now time.Time) {
	s.isActive = false
	s.updatedAt = now
}
