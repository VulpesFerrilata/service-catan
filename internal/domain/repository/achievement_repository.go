package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/middleware"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
)

type AchievementRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (datamodel.Achievements, error)
	Save(ctx context.Context, achievement *datamodel.Achievement) error
}

func NewAchievementRepository(transactionMiddleware *middleware.TransactionMiddleware,
	validate *validator.Validate) AchievementRepository {
	return &achievementRepository{
		transactionMiddleware: transactionMiddleware,
		validate:              validate,
	}
}

type achievementRepository struct {
	transactionMiddleware *middleware.TransactionMiddleware
	validate              *validator.Validate
}

func (ar achievementRepository) FindByGameId(ctx context.Context, gameId uint) (datamodel.Achievements, error) {
	achievementModels := make([]*model.Achievement, 0)
	err := ar.transactionMiddleware.Get(ctx).Find(&achievementModels, "game_id = ?", gameId).Error
	return datamodel.NewAchievementsFromAchievementModels(achievementModels), errors.Wrap(err, "repository.AchievementRepository.FindByGameId")
}

func (ar achievementRepository) insertOrUpdate(ctx context.Context, achievement *datamodel.Achievement) error {
	return achievement.Persist(func(achievementModel *model.Achievement) error {
		if err := ar.validate.StructCtx(ctx, achievementModel); err != nil {
			return errors.Wrap(err, "repository.AchievementRepository.InsertOrUpdate")
		}

		err := ar.transactionMiddleware.Get(ctx).Save(achievementModel).Error
		return errors.Wrap(err, "repository.AchievementRepository.InsertOrUpdate")
	})
}

func (ar achievementRepository) delete(ctx context.Context, achievement *datamodel.Achievement) error {
	return achievement.Persist(func(achievementModel *model.Achievement) error {
		err := ar.transactionMiddleware.Get(ctx).Delete(achievementModel).Error
		return errors.Wrap(err, "repository.AchievementRepository.Delete")
	})
}

func (ar achievementRepository) Save(ctx context.Context, achievement *datamodel.Achievement) error {
	if achievement.IsRemoved() {
		err := ar.delete(ctx, achievement)
		return errors.Wrap(err, "service.AchievementService.Save")
	}
	if achievement.IsModified() {
		err := ar.insertOrUpdate(ctx, achievement)
		return errors.Wrap(err, "service.AchievementService.Save")
	}
	return nil
}
