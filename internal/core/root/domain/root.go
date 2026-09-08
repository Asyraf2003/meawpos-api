// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package domain

import (
	"errors"
	"strings"
	"time"
)

const MaxRootNameLength = 200

var (
	ErrRootNameRequired = errors.New("root name is required")
	ErrRootNameTooLong  = errors.New("root name is too long")
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
	if len([]rune(name)) > MaxRootNameLength {
		return Root{}, ErrRootNameTooLong
	}
	return Root{
		ID: id, Name: name, PrimaryOwnerAccountID: ownerAccountID,
		CreatedAt: now, UpdatedAt: now,
	}, nil
}
