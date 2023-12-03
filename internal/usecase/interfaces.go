package usecase

import (
	"context"
)

type Shortener interface {
	Ping(ctx context.Context) error
}
