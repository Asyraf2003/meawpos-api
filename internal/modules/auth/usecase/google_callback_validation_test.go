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
	"time"

	"pos-go/internal/platform/state/memory"
)

func TestGoogleCallback_RejectsMissingFields(t *testing.T) {
	flow := NewGoogleFlow(
		&fakeCallbackOIDCProvider{},
		memory.NewAuthStateStore(),
		nil,
		nil,
		nil,
		nil,
		10*time.Minute,
		24*time.Hour,
	)

	_, err := flow.GoogleCallback(context.Background(), GoogleCallbackInput{})
	if err == nil {
		t.Fatal("GoogleCallback() error = nil, want error")
	}
}
