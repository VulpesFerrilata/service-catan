package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
)

type HarborService interface {
	GetHarborRepository() repository.SafeHarborRepository
	Save(ctx context.Context, harbor *model.Harbor) error
}

type harborService struct {
	harborRepository repository.HarborRepository
}

func (hs *harborService) GetHarborRepository() repository.SafeHarborRepository {
	return hs.harborRepository
}

func (hs *harborService) validate(ctx context.Context, harbor *model.Harbor) error {
	//TODO: validate harbor
	return nil
}

func (hs *harborService) Save(ctx context.Context, harbor *model.Harbor) error {
	if err := hs.validate(ctx, harbor); err != nil {
		return err
	}

	return hs.harborRepository.InsertOrUpdate(ctx, harbor)
}
