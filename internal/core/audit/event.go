// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package audit

import (
	"context"
	"time"
)

type Event struct {
	ID             string
	RootID         string
	ActorAccountID string
	SessionID      string
	RequestID      string
	AuthorityUsed  string
	Operation      string
	ResourceType   string
	ResourceID     string
	Reason         string
	OccurredAt     time.Time
}

type Writer interface {
	WriteAuditEvent(ctx context.Context, event Event) error
}
