package http

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"pos-go/internal/modules/capability/domain"
	"pos-go/internal/modules/capability/ports"
	capabilitypresenter "pos-go/internal/presentation/http/id/capability"

	"github.com/labstack/echo/v4"
)

type ListCapabilitiesUsecase interface {
	Execute(ctx context.Context) ([]domain.Capability, error)
}

type ShowCapabilityUsecase interface {
	Execute(ctx context.Context, key string) (domain.Capability, error)
}

type EnableCapabilityUsecase interface {
	Execute(ctx context.Context, key string) error
}

type DisableCapabilityUsecase interface {
	Execute(ctx context.Context, key string, reason string) error
}

type CapabilityHandler struct {
	listUsecase    ListCapabilitiesUsecase
	showUsecase    ShowCapabilityUsecase
	enableUsecase  EnableCapabilityUsecase
	disableUsecase DisableCapabilityUsecase
}

func NewCapabilityHandler(
	listUsecase ListCapabilitiesUsecase,
	showUsecase ShowCapabilityUsecase,
	enableUsecase EnableCapabilityUsecase,
	disableUsecase DisableCapabilityUsecase,
) *CapabilityHandler {
	return &CapabilityHandler{
		listUsecase:    listUsecase,
		showUsecase:    showUsecase,
		enableUsecase:  enableUsecase,
		disableUsecase: disableUsecase,
	}
}

func (h *CapabilityHandler) Register(group *echo.Group) {
	group.GET("/capabilities", h.List)
	group.GET("/capabilities/:key", h.Show)
	group.POST("/capabilities/:key/enable", h.Enable)
	group.POST("/capabilities/:key/disable", h.Disable)
}

type responseEnvelope struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
}

type disableCapabilityRequest struct {
	Reason string `json:"reason"`
}

func (h *CapabilityHandler) List(c echo.Context) error {
	capabilities, err := h.listUsecase.Execute(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, responseEnvelope{
		Success: true,
		Data:    capabilitypresenter.FromDomainList(capabilities),
	})
}

func (h *CapabilityHandler) Show(c echo.Context) error {
	key := strings.TrimSpace(c.Param("key"))
	if key == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "capability key is required")
	}

	capability, err := h.showUsecase.Execute(c.Request().Context(), key)
	if err != nil {
		return capabilityHTTPError(err)
	}

	return c.JSON(http.StatusOK, responseEnvelope{
		Success: true,
		Data:    capabilitypresenter.FromDomain(capability),
	})
}

func (h *CapabilityHandler) Enable(c echo.Context) error {
	key := strings.TrimSpace(c.Param("key"))
	if key == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "capability key is required")
	}

	if err := h.enableUsecase.Execute(c.Request().Context(), key); err != nil {
		return capabilityHTTPError(err)
	}

	capability, err := h.showUsecase.Execute(c.Request().Context(), key)
	if err != nil {
		return capabilityHTTPError(err)
	}

	return c.JSON(http.StatusOK, responseEnvelope{
		Success: true,
		Data:    capabilitypresenter.FromDomain(capability),
	})
}

func (h *CapabilityHandler) Disable(c echo.Context) error {
	key := strings.TrimSpace(c.Param("key"))
	if key == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "capability key is required")
	}

	var req disableCapabilityRequest
	if c.Request().Body != nil && c.Request().ContentLength != 0 {
		if err := c.Bind(&req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
		}
	}

	if err := h.disableUsecase.Execute(c.Request().Context(), key, req.Reason); err != nil {
		return capabilityHTTPError(err)
	}

	capability, err := h.showUsecase.Execute(c.Request().Context(), key)
	if err != nil {
		return capabilityHTTPError(err)
	}

	return c.JSON(http.StatusOK, responseEnvelope{
		Success: true,
		Data:    capabilitypresenter.FromDomain(capability),
	})
}

func capabilityHTTPError(err error) error {
	if errors.Is(err, ports.ErrCapabilityNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, "capability not found")
	}

	return err
}
