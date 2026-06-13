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

import "time"

type SupplierID string

type SupplierStatus string

const (
	SupplierStatusActive   SupplierStatus = "active"
	SupplierStatusInactive SupplierStatus = "inactive"
)

type NormalizedName string

type SupplierContact struct {
	Phone   string
	Email   string
	Address string
	Notes   string
}

type Supplier struct {
	id             SupplierID
	name           string
	normalizedName NormalizedName
	phone          string
	email          string
	address        string
	notes          string
	isActive       bool
	createdAt      time.Time
	updatedAt      time.Time
}

func (s Supplier) ID() SupplierID {
	return s.id
}

func (s Supplier) Name() string {
	return s.name
}

func (s Supplier) NormalizedName() NormalizedName {
	return s.normalizedName
}

func (s Supplier) Phone() string {
	return s.phone
}

func (s Supplier) Email() string {
	return s.email
}

func (s Supplier) Address() string {
	return s.address
}

func (s Supplier) Notes() string {
	return s.notes
}

func (s Supplier) IsActive() bool {
	return s.isActive
}

func (s Supplier) Status() SupplierStatus {
	if s.isActive {
		return SupplierStatusActive
	}

	return SupplierStatusInactive
}

func (s Supplier) CreatedAt() time.Time {
	return s.createdAt
}

func (s Supplier) UpdatedAt() time.Time {
	return s.updatedAt
}
