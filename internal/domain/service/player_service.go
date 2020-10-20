package service

import (
	"context"

	model "github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
)

type PlayerService interface {
	GetPlayerRepository() repository.ReadOnlyPlayerRepository
	Create(ctx context.Context, player *model.Player) error
}
