// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package sqlite

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrInvalidUUID = errors.New("SQLite catalog identifier must be a canonical lowercase UUID")

func canonicalUUID(value string) (string, error) {
	parsed, err := uuid.Parse(value)
	if err != nil || parsed.String() != value {
		return "", ErrInvalidUUID
	}
	return value, nil
}

func unixMicro(value time.Time) int64 { return value.UTC().UnixMicro() }

func timeFromUnixMicro(value int64) time.Time { return time.UnixMicro(value).UTC() }
