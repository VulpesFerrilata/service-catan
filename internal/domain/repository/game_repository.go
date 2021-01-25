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

type GameRepository interface {
	GetById(ctx context.Context, id uint) (*datamodel.Game, error)
	InsertOrUpdate(ctx context.Context, game *datamodel.Game) error
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

func (gr gameRepository) InsertOrUpdate(ctx context.Context, game *datamodel.Game) error {
	gameModel := game.ToModel()

	if err := gr.validate.StructCtx(ctx, gameModel); err != nil {
		return errors.Wrap(err, "repository.GameRepository.InsertOrUpdate")
	}

	err := gr.transactionMiddleware.Get(ctx).Save(gameModel).Error
	if err != nil {
		return errors.Wrap(err, "repository.GameRepository.InsertOrUpdate")
	}

	players := game.GetPlayers()
	for _, player := range players {
		if err := gr.playerRepository.InsertOrUpdate(ctx, player); err != nil {
			return errors.Wrap(err, "service.GameRepository.InsertOrUpdate")
		}
	}

	dices := game.GetDices()
	for _, dice := range dices {
		if err := gr.diceRepository.InsertOrUpdate(ctx, dice); err != nil {
			return errors.Wrap(err, "service.GameRepository.InsertOrUpdate")
		}
	}

	achievements := game.GetAchievements()
	for _, achievement := range achievements {
		if err := gr.achievementRepository.InsertOrUpdate(ctx, achievement); err != nil {
			return errors.Wrap(err, "service.GameRepository.InsertOrUpdate")
		}
	}

	resourceCards := game.GetResourceCards()
	for _, resourceCard := range resourceCards {
		if err := gr.resourceCardRepository.InsertOrUpdate(ctx, resourceCard); err != nil {
			return errors.Wrap(err, "service.GameRepository.InsertOrUpdate")
		}
	}

	developmentCards := game.GetDevelopmentCards()
	for _, developmentCard := range developmentCards {
		if err := gr.developmentCardRepository.InsertOrUpdate(ctx, developmentCard); err != nil {
			return errors.Wrap(err, "service.GameRepository.InsertOrUpdate")
		}
	}

	terrains := game.GetTerrains()
	for _, terrain := range terrains {
		if err := gr.terrainRepository.InsertOrUpdate(ctx, terrain); err != nil {
			return errors.Wrap(err, "service.GameRepository.InsertOrUpdate")
		}
	}

	robber := game.GetRobber()
	if err := gr.robberRepository.InsertOrUpdate(ctx, robber); err != nil {
		return errors.Wrap(err, "service.GameRepository.InsertOrUpdate")
	}

	constructions := game.GetConstructions()
	for _, construction := range constructions {
		if err := gr.constructionRepository.InsertOrUpdate(ctx, construction); err != nil {
			return errors.Wrap(err, "service.GameRepository.InsertOrUpdate")
		}
	}

	roads := game.GetRoads()
	for _, road := range roads {
		if err := gr.roadRepository.InsertOrUpdate(ctx, road); err != nil {
			return errors.Wrap(err, "service.GameRepository.InsertOrUpdate")
		}
	}

	harbors := game.GetHarbors()
	for _, harbor := range harbors {
		if err := gr.harborRepository.InsertOrUpdate(ctx, harbor); err != nil {
			return errors.Wrap(err, "service.GameRepository.InsertOrUpdate")
		}
	}

	return nil
}
