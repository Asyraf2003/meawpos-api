#!/usr/bin/env bash
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

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

ENV_HTTP_PORT="${HTTP_PORT:-}"
ENV_API_BASE_URL="${API_BASE_URL:-}"
ENV_AUTH_DEBUG_ENABLED="${AUTH_DEBUG_ENABLED:-}"

if [[ -f ".env" ]]; then
  set -a
  source .env
  set +a
fi

if [[ -n "$ENV_HTTP_PORT" ]]; then
  HTTP_PORT="$ENV_HTTP_PORT"
fi

if [[ -n "$ENV_API_BASE_URL" ]]; then
  API_BASE_URL="$ENV_API_BASE_URL"
fi

if [[ -n "$ENV_AUTH_DEBUG_ENABLED" ]]; then
  AUTH_DEBUG_ENABLED="$ENV_AUTH_DEBUG_ENABLED"
fi

HTTP_PORT="${HTTP_PORT:-8081}"
API_BASE_URL="${API_BASE_URL:-http://127.0.0.1:${HTTP_PORT}}"
AUTH_DEBUG_ENABLED="${AUTH_DEBUG_ENABLED:-true}"
APP_BIN="${APP_BIN:-.bin/pos-go-api}"
LOG_FILE="${API_SMOKE_LOG_FILE:-/tmp/pos-go-api-${HTTP_PORT}-smoke.log}"

if [[ ! -x "$APP_BIN" ]]; then
  echo "[FAIL] API binary not found or not executable: ${APP_BIN}"
  echo "Run: make build"
  exit 1
fi

port_open() {
  timeout 1 bash -c "cat < /dev/null > /dev/tcp/127.0.0.1/${HTTP_PORT}" >/dev/null 2>&1
}

started_server=false
api_pid=""

cleanup() {
  if [[ "$started_server" == "true" && -n "$api_pid" ]]; then
    echo "== stopping temporary API server =="
    kill "$api_pid" >/dev/null 2>&1 || true
    wait "$api_pid" >/dev/null 2>&1 || true
  fi
}
trap cleanup EXIT

if port_open; then
  echo "== using existing API server on port ${HTTP_PORT} =="
else
  echo "== starting temporary API server =="
  echo "log: ${LOG_FILE}"

  HTTP_PORT="$HTTP_PORT"   AUTH_DEBUG_ENABLED="$AUTH_DEBUG_ENABLED"   "$APP_BIN" > "$LOG_FILE" 2>&1 &

  api_pid="$!"
  started_server=true

  for _ in $(seq 1 30); do
    if curl -sS "${API_BASE_URL}/api/health" >/dev/null 2>&1; then
      break
    fi

    if ! kill -0 "$api_pid" >/dev/null 2>&1; then
      echo "[FAIL] API server exited before becoming ready"
      tail -n 120 "$LOG_FILE" || true
      exit 1
    fi

    sleep 1
  done

  if ! curl -sS "${API_BASE_URL}/api/health" >/dev/null 2>&1; then
    echo "[FAIL] API server did not become ready"
    tail -n 120 "$LOG_FILE" || true
    exit 1
  fi
fi

HTTP_PORT="$HTTP_PORT" API_BASE_URL="$API_BASE_URL" AUTH_DEBUG_ENABLED="$AUTH_DEBUG_ENABLED" bash scripts/dev_api_smoke_admin.sh
