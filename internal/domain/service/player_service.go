package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
	"github.com/pkg/errors"
)

type PlayerService interface {
	GetPlayerRepository() repository.SafePlayerRepository
}

func NewPlayerService(playerRepository repository.PlayerRepository) PlayerService {
	return &playerService{
		playerRepository: playerRepository,
	}
}

type playerService struct {
	playerRepository repository.PlayerRepository
}

func (ps *playerService) GetPlayerRepository() repository.SafePlayerRepository {
	return ps.playerRepository
}

func (ps playerService) FindByGameId(ctx context.Context, gameId int) (datamodel.Players, error) {
	players, err := ps.playerRepository.FindByGameId(ctx, gameId)
	if err != nil {
		return nil, errors.Wrap(err, "service.PlayerService.FindByGameId")
	}
}
