package port

import "context"

//go:generate mockgen -source=service.go -destination=mocks/service_mock.go -package=mocks

type Service interface {
	HealthCheck(ctx context.Context) error
}
