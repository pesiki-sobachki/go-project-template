package v1

import (
	"net/http"

	"github.com/shanth1/gotools/log"
	"github.com/shanth1/template/internal/core/port"
	"github.com/shanth1/template/internal/pkg/response"
)

type Handler struct {
	service port.Service
	logger  log.Logger
}

func NewHandler(service port.Service, logger log.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

type HealthResponse struct {
	Status string `json:"status" example:"OK"`
}

// HealthCheck godoc
// @Summary      Health Check
// @Tags         system
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.Response[HealthResponse]
// @Failure      503  {object}  response.ErrorResponse
// @Router       /health [get]
func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if err := h.service.HealthCheck(ctx); err != nil {
		h.logger.Error().Err(err).Msg("health check failed")
		response.WithError(w, http.StatusServiceUnavailable, err)
		return
	}

	response.JSON(w, http.StatusOK, HealthResponse{Status: "OK"})
}
