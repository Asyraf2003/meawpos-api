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

import "net/http"

func codeForEchoHTTPError(statusCode int, message string) string {
	switch message {
	case "authentication required",
		"missing bearer token",
		"invalid bearer token",
		"invalid access token",
		"session check failed",
		"inactive session",
		"principal resolve failed":
		return "authentication_required"
	case "forbidden":
		return "forbidden"
	case "capability disabled":
		return "capability_disabled"
	case "invalid request body":
		return "invalid_request_body"
	default:
		return defaultCodeForStatus(statusCode)
	}
}

func defaultCodeForStatus(statusCode int) string {
	switch statusCode {
	case http.StatusBadRequest:
		return "bad_request"
	case http.StatusUnauthorized:
		return "authentication_required"
	case http.StatusForbidden:
		return "forbidden"
	case http.StatusNotFound:
		return "not_found"
	case http.StatusConflict:
		return "conflict"
	default:
		return "internal_server_error"
	}
}
