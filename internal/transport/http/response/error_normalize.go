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

package response

import (
	"fmt"
	"net/http"
)

func normalizeStatusCode(statusCode int) int {
	if statusCode < 400 || statusCode > 599 {
		return http.StatusInternalServerError
	}

	return statusCode
}

func normalizeCode(code string, fallback string) string {
	if code == "" {
		return fallback
	}

	return code
}

func normalizeMessage(message string, fallback string) string {
	if message == "" {
		return fallback
	}

	return message
}

func echoErrorMessage(message any, fallback string) string {
	if message == nil {
		return fallback
	}

	switch value := message.(type) {
	case string:
		return normalizeMessage(value, fallback)
	case error:
		return normalizeMessage(value.Error(), fallback)
	default:
		return normalizeMessage(fmt.Sprint(value), fallback)
	}
}
