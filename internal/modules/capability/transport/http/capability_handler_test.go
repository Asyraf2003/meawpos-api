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

package http

import (
	"os"
	"strings"
	"testing"
)

func TestCapabilityHandler_DoesNotImportPostgresAdapter(t *testing.T) {
	files := []string{
		"capability_handler.go",
		"capability_handler_read.go",
		"capability_handler_write.go",
		"capability_handler_response.go",
	}

	for _, file := range files {
		source, err := os.ReadFile(file)
		if err != nil {
			t.Fatalf("ReadFile(%s) error = %v", file, err)
		}
		if strings.Contains(string(source), "internal/platform/postgres") {
			t.Fatalf("%s imports PostgreSQL adapter", file)
		}
	}
}
