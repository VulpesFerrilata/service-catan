package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
	"github.com/pkg/errors"
)

type GameService interface {
	GetGameRepository() repository.SafeGameRepository
	Save(ctx context.Context, game *datamodel.Game) error
}

func NewGameService(gameRepository repository.GameRepository,
	playerService PlayerService,
	diceService DiceService,
	achievementService AchievementService,
	resourceCardService ResourceCardService,
	developmentCardService DevelopmentCardService,
	terrainService TerrainService,
	robberService RobberService,
	constructionService ConstructionService,
	roadService RoadService,
	harborService HarborService) GameService {
	return &gameService{
		gameRepository:         gameRepository,
		playerService:          playerService,
		diceService:            diceService,
		achievementService:     achievementService,
		resourceCardService:    resourceCardService,
		developmentCardService: developmentCardService,
		terrainService:         terrainService,
		robberService:          robberService,
		constructionService:    constructionService,
		roadService:            roadService,
		harborService:          harborService,
	}
}

type gameService struct {
	gameRepository         repository.GameRepository
	playerService          PlayerService
	diceService            DiceService
	achievementService     AchievementService
	resourceCardService    ResourceCardService
	developmentCardService DevelopmentCardService
	terrainService         TerrainService
	robberService          RobberService
	constructionService    ConstructionService
	roadService            RoadService
	harborService          HarborService
}

func (gs gameService) GetGameRepository() repository.SafeGameRepository {
	return gs.gameRepository
}

func (gs gameService) save(ctx context.Context, game *datamodel.Game) error {
	if game.IsRemoved() {
		err := gs.gameRepository.Delete(ctx, game)
		return errors.Wrap(err, "service.GameService.save")
	}
	if game.IsModified() {
		err := gs.gameRepository.InsertOrUpdate(ctx, game)
		return errors.Wrap(err, "service.GameService.save")
	}
	return nil
}

func (gs gameService) Save(ctx context.Context, game *datamodel.Game) error {
	if err := gs.save(ctx, game); err != nil {
		return errors.Wrap(err, "service.GameService.Save")
	}

	players := game.GetPlayers()
	for _, player := range players {
		if err := gs.playerService.Save(ctx, player); err != nil {
			return errors.Wrap(err, "service.GameService.Save")
		}
	}

	dices := game.GetDices()
	for _, dice := range dices {
		if err := gs.diceService.Save(ctx, dice); err != nil {
			return errors.Wrap(err, "service.GameService.Save")
		}
	}

	achievements := game.GetAchievements()
	for _, achievement := range achievements {
		if err := gs.achievementService.Save(ctx, achievement); err != nil {
			return errors.Wrap(err, "service.GameService.Save")
		}
	}

	resourceCards := game.GetResourceCards()
	for _, resourceCard := range resourceCards {
		if err := gs.resourceCardService.Save(ctx, resourceCard); err != nil {
			return errors.Wrap(err, "service.GameService.Save")
		}
	}

	developmentCards := game.GetDevelopmentCards()
	for _, developmentCard := range developmentCards {
		if err := gs.developmentCardService.Save(ctx, developmentCard); err != nil {
			return errors.Wrap(err, "service.GameService.Save")
		}
	}

	terrains := game.GetTerrains()
	for _, terrain := range terrains {
		if err := gs.terrainService.Save(ctx, terrain); err != nil {
			return errors.Wrap(err, "service.GameService.Save")
		}
	}

	robber := game.GetRobber()
	if err := gs.robberService.Save(ctx, robber); err != nil {
		return errors.Wrap(err, "service.GameService.Save")
	}

	constructions := game.GetConstructions()
	for _, construction := range constructions {
		if err := gs.constructionService.Save(ctx, construction); err != nil {
			return err
		}
	}

	roads := game.GetRoads()
	for _, road := range roads {
		if err := gs.roadService.Save(ctx, road); err != nil {
			return err
		}
	}

	harbors := game.GetHarbors()
	for _, harbor := range harbors {
		if err := gs.harborService.Save(ctx, harbor); err != nil {
			return err
		}
	}

	return nil
}
