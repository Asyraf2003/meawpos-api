// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package authorization

import (
	"context"
	"errors"
	"sort"
	"strings"
)

var ErrRootAccessDenied = errors.New("root access denied")

type Authority struct {
	RootID       string
	AccountID    string
	PrimaryOwner bool
	Roles        []string
	Permissions  []string
}

func (a Authority) HasPermissions(required ...string) bool {
	if a.PrimaryOwner {
		return true
	}
	for _, want := range required {
		found := false
		for _, permission := range a.Permissions {
			if permission == want {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func (a Authority) Snapshot() string {
	if a.PrimaryOwner {
		return "primary_owner"
	}
	roles := append([]string(nil), a.Roles...)
	sort.Strings(roles)
	return "roles:" + strings.Join(roles, ",")
}

type Resolver interface {
	ResolveRootAuthority(ctx context.Context, rootID, accountID string) (Authority, error)
}
