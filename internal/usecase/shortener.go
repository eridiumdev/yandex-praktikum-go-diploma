package usecase

import (
	"context"

	"github.com/eridiumdev/yandex-praktikum-go-diploma/pkg/logger"
)

type ShortenerUC struct {
	log *logger.Logger
}

func NewShortener(log *logger.Logger) *ShortenerUC {
	return &ShortenerUC{
		log: log,
	}
}

func (uc *ShortenerUC) Ping(ctx context.Context) error {
	return nil
}
