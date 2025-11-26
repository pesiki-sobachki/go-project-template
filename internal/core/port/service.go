package port

import "context"

type Service interface {
	HealthCheck(ctx context.Context) error
}
