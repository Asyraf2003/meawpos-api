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
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ErrorResponse(err error) (int, ErrorEnvelope) {
	var publicErr *HTTPError
	if errors.As(err, &publicErr) {
		statusCode := normalizeStatusCode(publicErr.StatusCode)
		code := normalizeCode(publicErr.Code, defaultCodeForStatus(statusCode))
		message := normalizeMessage(publicErr.Message, http.StatusText(statusCode))

		return statusCode, errorEnvelope(code, message, publicErr.Fields)
	}

	var echoErr *echo.HTTPError
	if errors.As(err, &echoErr) {
		statusCode := normalizeStatusCode(echoErr.Code)
		message := echoErrorMessage(echoErr.Message, http.StatusText(statusCode))
		code := codeForEchoHTTPError(statusCode, message)

		return statusCode, errorEnvelope(code, message, nil)
	}

	return http.StatusInternalServerError, errorEnvelope(
		"internal_server_error",
		"internal server error",
		nil,
	)
}
