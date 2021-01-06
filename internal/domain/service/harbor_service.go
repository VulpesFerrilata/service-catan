package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
	"github.com/pkg/errors"
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

func (hs harborService) Save(ctx context.Context, harbor *datamodel.Harbor) error {
	if harbor.IsRemoved() {
		err := hs.harborRepository.Delete(ctx, harbor)
		return errors.Wrap(err, "service.HarborService.Save")
	}
	if harbor.IsModified() {
		err := hs.harborRepository.InsertOrUpdate(ctx, harbor)
		return errors.Wrap(err, "service.HarborService.Save")
	}
	return nil
}
