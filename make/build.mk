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

##@ Build

.PHONY: build run dev

build: ## Build the API binary
	mkdir -p .bin
	go build -o $(APP_BIN) ./cmd/api

run: build ## Run the API on HTTP_PORT
	HTTP_PORT=$(HTTP_PORT) $(APP_BIN)

dev: db-dev-setup db-migrate run ## Setup local DB, migrate, then run the API
