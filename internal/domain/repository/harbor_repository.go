package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
)

type SafeHarborRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (model.Harbors, error)
}

type HarborRepository interface {
	SafeHarborRepository
	InsertOrUpdate(ctx context.Context, harbor *model.Harbor) error
}
