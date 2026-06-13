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
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ErrorBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Fields  any    `json:"fields,omitempty"`
}

type ErrorEnvelope struct {
	Success bool      `json:"success"`
	Error   ErrorBody `json:"error"`
	Meta    any       `json:"meta"`
}

type HTTPError struct {
	StatusCode int
	Code       string
	Message    string
	Fields     any
}

func NewHTTPError(statusCode int, code string, message string) *HTTPError {
	return &HTTPError{
		StatusCode: statusCode,
		Code:       code,
		Message:    message,
	}
}

func NewValidationHTTPError(statusCode int, code string, message string, fields any) *HTTPError {
	return &HTTPError{
		StatusCode: statusCode,
		Code:       code,
		Message:    message,
		Fields:     fields,
	}
}

func (err *HTTPError) Error() string {
	if err == nil {
		return ""
	}

	return err.Message
}

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

func errorEnvelope(code string, message string, fields any) ErrorEnvelope {
	return ErrorEnvelope{
		Success: false,
		Error: ErrorBody{
			Code:    code,
			Message: message,
			Fields:  fields,
		},
		Meta: map[string]any{},
	}
}

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
