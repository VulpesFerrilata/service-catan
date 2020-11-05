package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
)

type SafeResourceCardRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (model.ResourceCards, error)
}

type ResourceCardRepository interface {
	SafeResourceCardRepository
	InsertOrUpdate(ctx context.Context, resourceCard *model.ResourceCard) error
}
