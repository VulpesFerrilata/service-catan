package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
	"github.com/VulpesFerrilata/library/pkg/app_error"
	"github.com/pkg/errors"
)

type PlayerService interface {
	GetPlayerRepository() repository.PlayerRepository
	NewPlayer(ctx context.Context, user *datamodel.User) (*datamodel.Player, error)
}

func NewPlayerService(playerRepository repository.PlayerRepository) PlayerService {
	return &playerService{
		playerRepository: playerRepository,
	}
}

type playerService struct {
	playerRepository repository.PlayerRepository
}

func (p playerService) GetPlayerRepository() repository.PlayerRepository {
	return p.playerRepository
}

func (p playerService) isExists(ctx context.Context, user *datamodel.User) (bool, error) {
	_, err := p.playerRepository.GetByUserId(ctx, user.GetId())
	if err != nil {
		if _, ok := errors.Cause(err).(*app_error.NotFoundError); ok {
			return false, nil
		}
		return false, errors.Wrap(err, "service.PlayerService.isExists")
	}
	return true, nil
}

func (p playerService) NewPlayer(ctx context.Context, user *datamodel.User) (*datamodel.Player, error) {
	isExists, err := p.isExists(ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, "service.PlayerService.NewPlayer")
	}
	if isExists {
		return nil, app_error.NewAlreadyExistsError("player")
	}

	player, err := datamodel.NewPlayer()
	if err != nil {
		return nil, errors.Wrap(err, "service.PlayerService.NewPlayer")
	}
	player.SetUser(user)

	return player, nil
}
