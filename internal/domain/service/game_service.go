package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
)

type GameService interface {
	GetById(ctx context.Context, id uint) (*model.Game, error)
	Save(ctx context.Context, game *model.Game) error
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

func (gs *gameService) GetById(ctx context.Context, id uint) (*model.Game, error) {
	game, err := gs.gameRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	players, err := gs.playerAggregateService.FindByGameId(ctx, id)
	if err != nil {
		return nil, err
	}
	game.AddPlayers(players...)

	dices, err := gs.diceService.GetDiceRepository().FindByGameId(ctx, id)
	if err != nil {
		return nil, err
	}
	game.AddDices(dices...)

	achievements, err := gs.achievementService.GetAchievementRepository().FindByGameId(ctx, id)
	if err != nil {
		return nil, err
	}
	game.AddAchievements(achievements...)

	resourceCards, err := gs.resourceCardService.GetResourceCardRepository().FindByGameId(ctx, id)
	if err != nil {
		return nil, err
	}
	game.AddResourceCards(resourceCards...)

	developmentCards, err := gs.developmentCardService.GetDevelopmentCardRepository().FindByGameId(ctx, id)
	if err != nil {
		return nil, err
	}
	game.AddDevelopmentCards(developmentCards...)

	terrains, err := gs.terrainService.GetFieldRepository().FindByGameId(ctx, id)
	if err != nil {
		return nil, err
	}
	game.AddTerrains(terrains...)

	robber, err := gs.robberService.GetRobberRepository().GetByGameId(ctx, id)
	if err != nil {
		return nil, err
	}
	game.SetRobber(robber)

	constructions, err := gs.constructionService.GetConstructionRepository().FindByGameId(ctx, id)
	if err != nil {
		return nil, err
	}
	game.AddConstructions(constructions...)

	roads, err := gs.roadService.GetRoadRepository().FindByGameId(ctx, id)
	if err != nil {
		return nil, err
	}
	game.AddRoads(roads...)

	harbors, err := gs.harborService.GetHarborRepository().FindByGameId(ctx, id)
	if err != nil {
		return nil, err
	}
	game.AddHarbors(harbors...)

	return game, nil
}

func (gs *gameService) Save(ctx context.Context, game *model.Game) error {
	if err := gs.gameService.Save(ctx, game); err != nil {
		return err
	}

	players := game.GetPlayers()
	for _, player := range players {
		if err := gs.playerAggregateService.Save(ctx, player); err != nil {
			return err
		}
	}

	dices := game.GetDices()
	for _, dice := range dices {
		if err := gs.diceService.Save(ctx, dice); err != nil {
			return err
		}
	}

	achievements := game.GetAchievements()
	for _, achievement := range achievements {
		if err := gs.achievementService.Save(ctx, achievement); err != nil {
			return err
		}
	}

	resourceCards := game.GetResourceCards()
	for _, resourceCard := range resourceCards {
		if err := gs.resourceCardService.Save(ctx, resourceCard); err != nil {
			return err
		}
	}

	developmentCards := game.GetDevelopmentCards()
	for _, developmentCard := range developmentCards {
		if err := gs.developmentCardService.Save(ctx, developmentCard); err != nil {
			return err
		}
	}

	terrains := game.GetTerrains()
	for _, terrain := range terrains {
		if err := gs.terrainService.Save(ctx, terrain); err != nil {
			return err
		}
	}

	robber := game.GetRobber()
	if err := gs.robberService.Save(ctx, robber); err != nil {
		return err
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
