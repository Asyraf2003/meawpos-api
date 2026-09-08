// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package request

import (
	"encoding/json"
	"errors"
	"io"
	"mime"
	"net/http"

	httpresponse "pos-go/internal/transport/http/response"

	"github.com/labstack/echo/v4"
)

const MaxJSONBodyBytes int64 = 64 * 1024

func DecodeJSON(c echo.Context, destination any) error {
	mediaType, _, err := mime.ParseMediaType(c.Request().Header.Get(echo.HeaderContentType))
	if err != nil || mediaType != echo.MIMEApplicationJSON {
		return httpresponse.NewHTTPError(http.StatusUnsupportedMediaType, "unsupported_media_type", "content type must be application/json")
	}

	c.Request().Body = http.MaxBytesReader(c.Response(), c.Request().Body, MaxJSONBodyBytes)
	decoder := json.NewDecoder(c.Request().Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(destination); err != nil {
		return decodeError(err)
	}
	err = decoder.Decode(&struct{}{})
	if errors.Is(err, io.EOF) {
		return nil
	}
	if err != nil {
		return decodeError(err)
	}
	return invalidBody()
}

func decodeError(err error) error {
	var tooLarge *http.MaxBytesError
	if errors.As(err, &tooLarge) {
		return httpresponse.NewHTTPError(http.StatusRequestEntityTooLarge, "request_body_too_large", "request body is too large")
	}
	return invalidBody()
}

func invalidBody() error {
	return httpresponse.NewHTTPError(http.StatusBadRequest, "invalid_request_body", "invalid request body")
}
