package service

import (
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
)

type HarborService interface {
	GetHarborRepository() repository.HarborRepository
}

func NewHarborService(harborRepository repository.HarborRepository) HarborService {
	return &harborService{
		harborRepository: harborRepository,
	}
}

type harborService struct {
	harborRepository repository.HarborRepository
}

func (hs harborService) GetHarborRepository() repository.HarborRepository {
	return hs.harborRepository
}
