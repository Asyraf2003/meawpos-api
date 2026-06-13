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

package http

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
)

type testEnvelope struct {
	Success bool            `json:"success"`
	Data    json.RawMessage `json:"data"`
	Meta    map[string]any  `json:"meta"`
}

type capabilityResponseForTest struct {
	Key            string `json:"key"`
	RiskLevel      string `json:"risk_level"`
	Enabled        bool   `json:"enabled"`
	DisabledReason string `json:"disabled_reason"`
}

func decodeEnvelope(t *testing.T, rec *httptest.ResponseRecorder) testEnvelope {
	t.Helper()

	var envelope testEnvelope
	if err := json.Unmarshal(rec.Body.Bytes(), &envelope); err != nil {
		t.Fatalf("json.Unmarshal(envelope) error = %v", err)
	}
	if envelope.Meta == nil {
		t.Fatal("meta is nil, want empty object")
	}
	if len(envelope.Meta) != 0 {
		t.Fatalf("meta = %#v, want empty object", envelope.Meta)
	}

	return envelope
}

func decodeCapability(t *testing.T, raw json.RawMessage) capabilityResponseForTest {
	t.Helper()

	var data capabilityResponseForTest
	decodeRawData(t, raw, &data)

	return data
}

func decodeCapabilityList(t *testing.T, raw json.RawMessage) []capabilityResponseForTest {
	t.Helper()

	var data []capabilityResponseForTest
	decodeRawData(t, raw, &data)

	return data
}

func decodeRawData(t *testing.T, raw json.RawMessage, out any) {
	t.Helper()

	if len(raw) == 0 {
		t.Fatal("data is empty")
	}
	if err := json.Unmarshal(raw, out); err != nil {
		t.Fatalf("json.Unmarshal(data) error = %v", err)
	}
}
