package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/middleware"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
)

type AchievementRepository interface {
	FindByGameId(ctx context.Context, gameId uuid.UUID) (datamodel.Achievements, error)
	InsertOrUpdate(ctx context.Context, achievement *datamodel.Achievement) error
}

func NewAchievementRepository(transactionMiddlewae *middleware.TransactionMiddleware,
	validate *validator.Validate) AchievementRepository {
	return &achievementRepository{
		transactionMiddlewae: transactionMiddlewae,
		validate:             validate,
	}
}

type achievementRepository struct {
	transactionMiddlewae *middleware.TransactionMiddleware
	validate             *validator.Validate
}

func (a achievementRepository) FindByGameId(ctx context.Context, gameId uuid.UUID) (datamodel.Achievements, error) {
	achievementModels := make([]*model.Achievement, 0)
	err := a.transactionMiddlewae.Get(ctx).Find(&achievementModels, "game_id = ?", gameId).Error
	if err != nil {
		return nil, errors.Wrap(err, "repository.AchievementRepository.FindByGameId")
	}

	achievements := make(datamodel.Achievements, 0)
	for _, achievementModel := range achievementModels {
		achievement, err := datamodel.NewAchievementFromModel(achievementModel)
		if err != nil {
			return nil, errors.Wrap(err, "repository.AchievementRepository.FindByGameId")
		}
		achievements = append(achievements, achievement)
	}

	return achievements, errors.Wrap(err, "repository.AchievementRepository.FindByGameId")
}

func (a achievementRepository) InsertOrUpdate(ctx context.Context, achievement *datamodel.Achievement) error {
	achievementModel := achievement.ToModel()

	if err := a.validate.StructCtx(ctx, achievementModel); err != nil {
		return errors.Wrap(err, "repository.AchievementRepository.InsertOrUpdate")
	}

	err := a.transactionMiddlewae.Get(ctx).Save(achievementModel).Error
	return errors.Wrap(err, "repository.AchievementRepository.InsertOrUpdate")
}
