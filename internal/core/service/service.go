package service

import (
	"context"

	"github.com/shanth1/gotools/log"
	"github.com/shanth1/template/internal/core/port"
)

type Service struct {
	logger log.Logger
}

var _ port.Service = (*Service)(nil)

func New(
	logger log.Logger,
) port.Service {
	return &Service{
		logger: logger,
	}
}

func (s *Service) HealthCheck(_ context.Context) error {
	s.logger.Info().Msg("starting health check...")
	return nil
}
