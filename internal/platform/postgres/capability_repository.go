// Copyright (C) 2026 Asyraf Mubarak
//
// This file is part of gopos-api.
//
// gopos-api is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, version 3 only.
//
// gopos-api is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with gopos-api. If not, see <https://www.gnu.org/licenses/>.

package postgres

import (
	"context"
	"errors"

	"pos-go/internal/modules/capability/domain"
	"pos-go/internal/modules/capability/ports"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CapabilityRepository struct {
	pool *pgxpool.Pool
}

func NewCapabilityRepository(pool *pgxpool.Pool) *CapabilityRepository {
	return &CapabilityRepository{pool: pool}
}

func (r *CapabilityRepository) List(ctx context.Context) ([]domain.Capability, error) {
	rows, err := r.query(ctx, capabilitySelectSQL()+`
		ORDER BY domain, operation, key
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	capabilities := []domain.Capability{}
	for rows.Next() {
		capability, err := scanCapability(rows)
		if err != nil {
			return nil, err
		}
		capabilities = append(capabilities, capability)
	}

	return capabilities, rows.Err()
}

func (r *CapabilityRepository) Get(ctx context.Context, key string) (domain.Capability, error) {
	row := r.queryRow(ctx, capabilitySelectSQL()+`
		WHERE key = $1
	`, key)

	capability, err := scanCapability(row)
	if errors.Is(err, pgx.ErrNoRows) {
		return domain.Capability{}, ports.ErrCapabilityNotFound
	}
	if err != nil {
		return domain.Capability{}, err
	}

	return capability, nil
}

func (r *CapabilityRepository) Save(ctx context.Context, capability domain.Capability) error {
	capability, err := domain.NewCapability(capability)
	if err != nil {
		return err
	}

	sql := `
		INSERT INTO api_capabilities (
			key, domain, operation, method, path,
			default_enabled, enabled, required_permission, risk_level,
			audit_required, idempotency_required, owner_package,
			test_proof, disabled_reason
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
		ON CONFLICT (key) DO UPDATE SET
			domain = EXCLUDED.domain,
			operation = EXCLUDED.operation,
			method = EXCLUDED.method,
			path = EXCLUDED.path,
			default_enabled = EXCLUDED.default_enabled,
			enabled = EXCLUDED.enabled,
			required_permission = EXCLUDED.required_permission,
			risk_level = EXCLUDED.risk_level,
			audit_required = EXCLUDED.audit_required,
			idempotency_required = EXCLUDED.idempotency_required,
			owner_package = EXCLUDED.owner_package,
			test_proof = EXCLUDED.test_proof,
			disabled_reason = EXCLUDED.disabled_reason,
			updated_at = now()
	`

	_, err = r.exec(ctx, sql, capabilityArgs(capability)...)
	return err
}

var _ ports.CapabilityRepository = (*CapabilityRepository)(nil)
