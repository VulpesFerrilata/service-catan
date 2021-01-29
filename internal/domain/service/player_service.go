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

func (p playerService) GetPlayerRepository() repository.PlayerRepository {
	return p.playerRepository
}
