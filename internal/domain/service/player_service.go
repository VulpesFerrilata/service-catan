package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
)

type PlayerService interface {
	GetPlayerRepository() repository.SafePlayerRepository
	Save(ctx context.Context, player *datamodel.Player) error
}

func NewPlayerService(playerRepository repository.PlayerRepository) PlayerService {
	return &playerService{
		playerRepository: playerRepository,
	}
}

type playerService struct {
	playerRepository repository.PlayerRepository
}

func (ps playerService) GetPlayerRepository() repository.SafePlayerRepository {
	return ps.playerRepository
}
