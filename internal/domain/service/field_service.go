package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
)

type FieldService interface {
	GetFieldRepository() repository.SafeFieldRepository
	Save(ctx context.Context, field *model.Field) error
}

type fieldService struct {
	fieldRepository repository.FieldRepository
}

func (fs *fieldService) GetFieldRepository() repository.SafeFieldRepository {
	return fs.fieldRepository
}

func (fs *fieldService) validate(ctx context.Context, field *model.Field) error {
	//TODO: validate field
	return nil
}

func (fs *fieldService) Save(ctx context.Context, field *model.Field) error {
	if err := fs.validate(ctx, field); err != nil {
		return err
	}

	return fs.fieldRepository.InsertOrUpdate(ctx, field)
}
