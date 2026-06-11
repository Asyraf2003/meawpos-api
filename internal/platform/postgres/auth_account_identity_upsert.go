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
	"encoding/json"

	"pos-go/internal/modules/auth/domain"
)

func (r *AccountIdentityRepository) UpsertIdentity(ctx context.Context, accountID string, identity domain.Identity) error {
	metaJSON, err := json.Marshal(identity.Meta)
	if err != nil {
		return err
	}

	return r.exec(ctx, `
		INSERT INTO auth_identities (
			account_id, provider, subject, email, email_verified, meta_json
		) VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (provider, subject)
		DO UPDATE SET
			account_id = EXCLUDED.account_id,
			email = EXCLUDED.email,
			email_verified = EXCLUDED.email_verified,
			meta_json = EXCLUDED.meta_json,
			updated_at = now()
	`,
		accountID,
		string(identity.Provider),
		identity.Subject,
		identity.Email,
		identity.EmailVerified,
		metaJSON,
	)
}
