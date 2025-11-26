package v1

import (
	"net/http"

	"github.com/shanth1/gotools/log"
	"github.com/shanth1/template/internal/core/port"
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

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if err := h.service.HealthCheck(ctx); err != nil {
		h.logger.Error().Err(err).Msg("health check failed")
		w.WriteHeader(http.StatusServiceUnavailable)
		_, _ = w.Write([]byte("Service Unavailable"))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK"))
}
