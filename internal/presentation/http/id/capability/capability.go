package capability

import "pos-go/internal/modules/capability/domain"

type CapabilityResponse struct {
	Key                 string `json:"key"`
	Domain              string `json:"domain"`
	Operation           string `json:"operation"`
	Method              string `json:"method"`
	Path                string `json:"path"`
	DefaultEnabled      bool   `json:"default_enabled"`
	Enabled             bool   `json:"enabled"`
	RequiredPermission  string `json:"required_permission"`
	RiskLevel           string `json:"risk_level"`
	AuditRequired       bool   `json:"audit_required"`
	IdempotencyRequired bool   `json:"idempotency_required"`
	OwnerPackage        string `json:"owner_package"`
	TestProof           string `json:"test_proof"`
	DisabledReason      string `json:"disabled_reason"`
}

func FromDomain(capability domain.Capability) CapabilityResponse {
	return CapabilityResponse{
		Key:                 capability.Key,
		Domain:              capability.Domain,
		Operation:           capability.Operation,
		Method:              capability.Method,
		Path:                capability.Path,
		DefaultEnabled:      capability.DefaultEnabled,
		Enabled:             capability.Enabled,
		RequiredPermission:  capability.RequiredPermission,
		RiskLevel:           string(capability.RiskLevel),
		AuditRequired:       capability.AuditRequired,
		IdempotencyRequired: capability.IdempotencyRequired,
		OwnerPackage:        capability.OwnerPackage,
		TestProof:           capability.TestProof,
		DisabledReason:      capability.DisabledReason,
	}
}

func FromDomainList(capabilities []domain.Capability) []CapabilityResponse {
	out := make([]CapabilityResponse, 0, len(capabilities))
	for _, capability := range capabilities {
		out = append(out, FromDomain(capability))
	}

	return out
}
