package http

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/shanth1/gotools/log"
	httpMw "github.com/shanth1/template/internal/adapter/inbound/http/middleware"
	v1 "github.com/shanth1/template/internal/adapter/inbound/http/v1"
	"github.com/shanth1/template/internal/core/port"
)

func NewRouter(service port.Service, logger log.Logger) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(httpMw.Logger(logger))
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	handlerV1 := v1.NewHandler(service, logger)

	r.Get("/health", handlerV1.HealthCheck)
	r.Route("/api/v1", func(_ chi.Router) {
		// r.Post("/users", handlerV1.CreateUser)
	})

	return r
}
