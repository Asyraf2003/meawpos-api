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

package domain

import "errors"

func (c Capability) Validate() error {
	if c.Key == "" {
		return errors.New("capability key is required")
	}
	if c.Domain == "" {
		return errors.New("capability domain is required")
	}
	if c.Operation == "" {
		return errors.New("capability operation is required")
	}
	if c.Method == "" {
		return errors.New("capability method is required")
	}
	if c.Path == "" {
		return errors.New("capability path is required")
	}
	if c.RequiredPermission == "" {
		return errors.New("capability required permission is required")
	}
	if c.OwnerPackage == "" {
		return errors.New("capability owner package is required")
	}
	if c.TestProof == "" {
		return errors.New("capability test proof is required")
	}
	if !c.RiskLevel.Valid() {
		return errors.New("capability risk level is invalid")
	}

	return nil
}
