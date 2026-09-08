// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package core

import (
	"errors"
	"strings"
	"time"
)

var ErrNameRequired = errors.New("catalog item name is required")

type Item struct {
	ID, RootID, Name     string
	CreatedAt, UpdatedAt time.Time
}

func NewItem(id, rootID, name string, now time.Time) (Item, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return Item{}, ErrNameRequired
	}
	return Item{ID: id, RootID: rootID, Name: name, CreatedAt: now, UpdatedAt: now}, nil
}
