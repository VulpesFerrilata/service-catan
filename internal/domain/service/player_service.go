package service

import (
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
)

type PlayerService interface {
	GetPlayerRepository() repository.PlayerRepository
}

func NewPlayerService(playerRepository repository.PlayerRepository) PlayerService {
	return &playerService{
		playerRepository: playerRepository,
	}
}

type playerService struct {
	playerRepository repository.PlayerRepository
}

func (ps playerService) GetPlayerRepository() repository.PlayerRepository {
	return ps.playerRepository
}
