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
