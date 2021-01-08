package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
)

type HarborService interface {
	GetHarborRepository() repository.SafeHarborRepository
	Save(ctx context.Context, harbor *datamodel.Harbor) error
}

func NewHarborService(harborRepository repository.HarborRepository) HarborService {
	return &harborService{
		harborRepository: harborRepository,
	}
}

type harborService struct {
	harborRepository repository.HarborRepository
}

func (hs harborService) GetHarborRepository() repository.SafeHarborRepository {
	return hs.harborRepository
}
