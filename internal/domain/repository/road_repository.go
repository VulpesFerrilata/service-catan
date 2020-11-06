package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
)

type SafeRoadRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (model.Roads, error)
}

type RoadRepository interface {
	SafeRoadRepository
	InsertOrUpdate(ctx context.Context, road *model.Road) error
}
