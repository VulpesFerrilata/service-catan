package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
)

type SafeDevelopmentCardRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (model.DevelopmentCards, error)
}

type DevelopmentCardRepository interface {
	SafeDevelopmentCardRepository
	InsertOrUpdate(ctx context.Context, developmentCard *model.DevelopmentCard) error
}
