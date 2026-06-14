# Copyright (C) 2026 Asyraf Mubarak
#
# This file is part of gopos-api.
#
# gopos-api is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as published by
# the Free Software Foundation, version 3 only.
#
# gopos-api is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
# GNU Affero General Public License for more details.
#
# You should have received a copy of the GNU Affero General Public License
# along with gopos-api. If not, see <https://www.gnu.org/licenses/>.

##@ Test

.PHONY: test test-unit test-api test-db test-db-integration vet lint check verify ci screening

test: ## Run all Go tests
	$(GO_TEST) ./...

test-unit: ## Run module-focused tests
	$(GO_TEST) ./internal/modules/... ./internal/platform/state/... ./internal/platform/token/... ./internal/config

test-api: ## Run HTTP transport and presentation tests
	$(GO_TEST) ./internal/modules/*/transport/http ./internal/transport/http/... ./internal/presentation/http/...

test-db: ## Run PostgreSQL adapter tests
	$(GO_TEST) ./internal/platform/postgres/...

test-db-integration: ## Load .env and run DB-backed PostgreSQL integration tests, defaulting to Supplier
	@set -a; if [[ -f .env ]]; then source .env; fi; set +a; $(GO_TEST) -tags integration ./internal/platform/postgres/... -run "$${RUN:-Supplier}" -count=1 -v

vet: ## Run go vet audit
	bash scripts/audit_go_vet.sh

lint: vet ## Run static analysis currently wired as go vet

check: audit-format vet audit-file-size audit-hex audit-ai-rules ## Run local doc and structure checks

verify: audit-all ## Run the aggregate local quality gate

ci: verify ## Alias to verify

screening: audit-all ## Alias to aggregate audit
