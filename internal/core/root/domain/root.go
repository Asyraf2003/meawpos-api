// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package domain

import (
	"errors"
	"strings"
	"time"
)

var (
	ErrRootNameRequired = errors.New("root name is required")
	ErrRootNotFound     = errors.New("root not found")
)

type Root struct {
	ID                    string
	Name                  string
	PrimaryOwnerAccountID string
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

func NewRoot(id, name, ownerAccountID string, now time.Time) (Root, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return Root{}, ErrRootNameRequired
	}
	return Root{
		ID: id, Name: name, PrimaryOwnerAccountID: ownerAccountID,
		CreatedAt: now, UpdatedAt: now,
	}, nil
}
