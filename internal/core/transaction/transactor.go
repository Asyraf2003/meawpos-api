// Copyright (C) 2026 Asyraf Mubarak
//
// This file is part of gopos-api.
//
// gopos-api is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, version 3 only.

package transaction

import "context"

type Transactor interface {
	RunInTx(ctx context.Context, fn func(context.Context) error) error
}
