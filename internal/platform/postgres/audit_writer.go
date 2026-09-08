// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package postgres

import (
	"context"
	"strings"

	"pos-go/internal/core/audit"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AuditWriter struct{ pool *pgxpool.Pool }

func NewAuditWriter(pool *pgxpool.Pool) *AuditWriter { return &AuditWriter{pool: pool} }

func (w *AuditWriter) WriteAuditEvent(ctx context.Context, event audit.Event) error {
	exec := w.pool.Exec
	if tx, ok := TxFromContext(ctx); ok {
		exec = tx.Exec
	}
	var reason any
	if strings.TrimSpace(event.Reason) != "" {
		reason = strings.TrimSpace(event.Reason)
	}
	_, err := exec(ctx, `INSERT INTO audit_events
		(id,root_id,actor_account_id,session_id,request_id,authority_used,operation,resource_type,resource_id,reason,occurred_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`,
		event.ID, event.RootID, event.ActorAccountID, event.SessionID, event.RequestID, event.AuthorityUsed,
		event.Operation, event.ResourceType, event.ResourceID, reason, event.OccurredAt)
	return err
}

var _ audit.Writer = (*AuditWriter)(nil)
