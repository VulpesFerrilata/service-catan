package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
)

type SafeFieldRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (model.Fields, error)
}

type FieldRepository interface {
	SafeFieldRepository
	InsertOrUpdate(ctx context.Context, field *model.Field) error
}
