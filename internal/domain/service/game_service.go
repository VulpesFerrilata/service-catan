package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
)

type GameService interface {
	GetGameRepository() repository.ReadOnlyGameRepository
	GetById(ctx context.Context, id uint) (*model.Game, error)
	Save(ctx context.Context, game *model.Game) error
}
