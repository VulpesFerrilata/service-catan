package repository

import (
	"context"

	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/app_error"
	"github.com/VulpesFerrilata/library/pkg/middleware"
	"gorm.io/gorm"
)

type SafeGameRepository interface {
	GetById(ctx context.Context, id uint) (*datamodel.Game, error)
}

type GameRepository interface {
	SafeGameRepository
	Save(ctx context.Context, game *datamodel.Game) error
}

func NewGameRepository(transactionMiddleware *middleware.TransactionMiddleware,
	validate *validator.Validate,
	playerRepository PlayerRepository,
	diceRepository DiceRepository,
	achievementRepository AchievementRepository,
	resourceCardRepository ResourceCardRepository,
	developmentCardRepository DevelopmentCardRepository,
	terrainRepository TerrainRepository,
	robberRepository RobberRepository,
	constructionRepository ConstructionRepository,
	roadRepository RoadRepository,
	harborRepository HarborRepository) GameRepository {
	return &gameRepository{
		transactionMiddleware:     transactionMiddleware,
		validate:                  validate,
		playerRepository:          playerRepository,
		diceRepository:            diceRepository,
		achievementRepository:     achievementRepository,
		resourceCardRepository:    resourceCardRepository,
		developmentCardRepository: developmentCardRepository,
		terrainRepository:         terrainRepository,
		robberRepository:          robberRepository,
		constructionRepository:    constructionRepository,
		roadRepository:            roadRepository,
		harborRepository:          harborRepository,
	}
}

type gameRepository struct {
	transactionMiddleware     *middleware.TransactionMiddleware
	validate                  *validator.Validate
	playerRepository          PlayerRepository
	diceRepository            DiceRepository
	achievementRepository     AchievementRepository
	resourceCardRepository    ResourceCardRepository
	developmentCardRepository DevelopmentCardRepository
	terrainRepository         TerrainRepository
	robberRepository          RobberRepository
	constructionRepository    ConstructionRepository
	roadRepository            RoadRepository
	harborRepository          HarborRepository
}

func (gr gameRepository) GetById(ctx context.Context, gameId uint) (*datamodel.Game, error) {
	gameModel := new(model.Game)
	err := gr.transactionMiddleware.Get(ctx).First(&gameModel, gameId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, app_error.NewNotFoundError("game")
	}
	if err != nil {
		return nil, errors.Wrap(err, "repository.GameRepository.GetById")
	}

	game := datamodel.NewGameFromGameModel(gameModel)

	players, err := gr.playerRepository.FindByGameId(ctx, gameModel.ID)
	if err != nil {
		return nil, errors.Wrap(err, "repository.GameRepository.GetById")
	}
	game.AddPlayers(players...)

	return game, nil
}

func (gr gameRepository) insertOrUpdate(ctx context.Context, game *datamodel.Game) error {
	return game.Persist(func(gameModel *model.Game) error {
		if err := gr.validate.StructCtx(ctx, gameModel); err != nil {
			if fieldErrors, ok := errors.Cause(err).(validator.ValidationErrors); ok {
				err = app_error.NewEntityValidationError(gameModel, fieldErrors)
			}
			return errors.Wrap(err, "repository.GameRepository.InsertOrUpdate")
		}

		err := gr.transactionMiddleware.Get(ctx).Save(gameModel).Error
		return errors.Wrap(err, "repository.GameRepository.InsertOrUpdate")
	})
}

func (gr gameRepository) delete(ctx context.Context, game *datamodel.Game) error {
	return game.Persist(func(gameModel *model.Game) error {
		err := gr.transactionMiddleware.Get(ctx).Delete(gameModel).Error
		return errors.Wrap(err, "repository.GameRepository.Delete")
	})
}

func (gr gameRepository) save(ctx context.Context, game *datamodel.Game) error {
	if game.IsRemoved() {
		err := gr.delete(ctx, game)
		return errors.Wrap(err, "service.GameRepository.save")
	}
	if game.IsModified() {
		err := gr.insertOrUpdate(ctx, game)
		return errors.Wrap(err, "service.GameRepository.save")
	}
	return nil
}

func (gr gameRepository) Save(ctx context.Context, game *datamodel.Game) error {
	if err := gr.save(ctx, game); err != nil {
		return errors.Wrap(err, "service.GameRepository.Save")
	}

	players := game.GetPlayers()
	for _, player := range players {
		if err := gr.playerRepository.Save(ctx, player); err != nil {
			return errors.Wrap(err, "service.GameRepository.Save")
		}
	}

	dices := game.GetDices()
	for _, dice := range dices {
		if err := gr.diceRepository.Save(ctx, dice); err != nil {
			return errors.Wrap(err, "service.GameRepository.Save")
		}
	}

	achievements := game.GetAchievements()
	for _, achievement := range achievements {
		if err := gr.achievementRepository.Save(ctx, achievement); err != nil {
			return errors.Wrap(err, "service.GameRepository.Save")
		}
	}

	resourceCards := game.GetResourceCards()
	for _, resourceCard := range resourceCards {
		if err := gr.resourceCardRepository.Save(ctx, resourceCard); err != nil {
			return errors.Wrap(err, "service.GameRepository.Save")
		}
	}

	developmentCards := game.GetDevelopmentCards()
	for _, developmentCard := range developmentCards {
		if err := gr.developmentCardRepository.Save(ctx, developmentCard); err != nil {
			return errors.Wrap(err, "service.GameRepository.Save")
		}
	}

	terrains := game.GetTerrains()
	for _, terrain := range terrains {
		if err := gr.terrainRepository.Save(ctx, terrain); err != nil {
			return errors.Wrap(err, "service.GameRepository.Save")
		}
	}

	robber := game.GetRobber()
	if err := gr.robberRepository.Save(ctx, robber); err != nil {
		return errors.Wrap(err, "service.GameRepository.Save")
	}

	constructions := game.GetConstructions()
	for _, construction := range constructions {
		if err := gr.constructionRepository.Save(ctx, construction); err != nil {
			return err
		}
	}

	roads := game.GetRoads()
	for _, road := range roads {
		if err := gr.roadRepository.Save(ctx, road); err != nil {
			return err
		}
	}

	harbors := game.GetHarbors()
	for _, harbor := range harbors {
		if err := gr.harborRepository.Save(ctx, harbor); err != nil {
			return err
		}
	}

	return nil
}
