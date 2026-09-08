// Copyright (C) 2026 Asyraf Mubarak
// This file is part of gopos-api and licensed under GNU AGPLv3.

package postgres

import (
	"context"
	"time"

	"pos-go/internal/modules/sales/ports"
)

func (s *SalesStore) ClaimIdempotency(ctx context.Context, rootID, operation, key, fingerprint string, createdAt time.Time) (ports.Claim, error) {
	tag, err := s.exec(ctx, `INSERT INTO financial_idempotency (root_id,operation,idempotency_key,request_fingerprint,created_at) VALUES ($1,$2,$3,$4,$5) ON CONFLICT DO NOTHING`, rootID, operation, key, fingerprint, createdAt)
	if err != nil {
		return ports.Claim{}, err
	}
	if tag.RowsAffected() == 1 {
		return ports.Claim{}, nil
	}
	var stored string
	var resourceID *string
	err = s.queryRow(ctx, `SELECT request_fingerprint,resource_id FROM financial_idempotency WHERE root_id=$1 AND operation=$2 AND idempotency_key=$3 FOR UPDATE`, rootID, operation, key).Scan(&stored, &resourceID)
	if err != nil {
		return ports.Claim{}, err
	}
	claim := ports.Claim{Existing: true, Fingerprint: stored}
	if resourceID != nil {
		claim.ResourceID = *resourceID
	}
	return claim, nil
}

func (s *SalesStore) BindIdempotency(ctx context.Context, rootID, operation, key, resourceType, resourceID string) error {
	_, err := s.exec(ctx, `UPDATE financial_idempotency SET resource_type=$4,resource_id=$5 WHERE root_id=$1 AND operation=$2 AND idempotency_key=$3`, rootID, operation, key, resourceType, resourceID)
	return err
}
