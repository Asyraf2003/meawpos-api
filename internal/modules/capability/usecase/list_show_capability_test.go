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
	"testing"
)

func TestShowCapabilityReturnsCapability(t *testing.T) {
	repository := fakeRepository(t, true)
	usecase := NewShowCapability(repository)

	capability, err := usecase.Execute(context.Background(), "account.role.assign")
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if capability.Key != "account.role.assign" {
		t.Fatalf("key = %q", capability.Key)
	}
}

func TestListCapabilitiesReturnsAllCapabilities(t *testing.T) {
	repository := fakeRepository(t, true)
	usecase := NewListCapabilities(repository)

	capabilities, err := usecase.Execute(context.Background())
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if len(capabilities) != 1 {
		t.Fatalf("capabilities len = %d, want 1", len(capabilities))
	}
}
