package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
	"github.com/pkg/errors"
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

func (ps playerService) Save(ctx context.Context, player *datamodel.Player) error {
	if player.IsRemoved() {
		err := ps.playerRepository.Delete(ctx, player)
		return errors.Wrap(err, "service.PlayerService.Save")
	}
	if player.IsModified() {
		err := ps.playerRepository.InsertOrUpdate(ctx, player)
		return errors.Wrap(err, "service.PlayerService.Save")
	}
	return nil
}
