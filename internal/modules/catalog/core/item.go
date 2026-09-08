// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package core

import (
	"errors"
	"strings"
	"time"
)

const MaxNameLength = 200

var (
	ErrNameRequired = errors.New("catalog item name is required")
	ErrNameTooLong  = errors.New("catalog item name is too long")
)

type Item struct {
	ID, RootID, Name     string
	CreatedAt, UpdatedAt time.Time
}

func NewItem(id, rootID, name string, now time.Time) (Item, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return Item{}, ErrNameRequired
	}
	if len([]rune(name)) > MaxNameLength {
		return Item{}, ErrNameTooLong
	}
	return Item{ID: id, RootID: rootID, Name: name, CreatedAt: now, UpdatedAt: now}, nil
}
