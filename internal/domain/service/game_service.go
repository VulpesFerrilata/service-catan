package service

import (
	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
	"github.com/pkg/errors"
)

type GameService interface {
	GetGameRepository() repository.GameRepository
	InitGame(game *datamodel.Game) error
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

func (g gameService) GetGameRepository() repository.GameRepository {
	return g.gameRepository
}

func (g gameService) InitGame(game *datamodel.Game) error {
	dices, err := g.diceService.InitDices()
	if err != nil {
		return errors.Wrap(err, "service.GameService.InitGame")
	}
	game.AddDices(dices...)

	achievements, err := g.achievementService.InitAchievements()
	if err != nil {
		return errors.Wrap(err, "service.GameService.InitGame")
	}
	game.AddAchievements(achievements...)

	resourceCards, err := g.resourceCardService.InitResourceCards()
	if err != nil {
		return errors.Wrap(err, "service.GameService.InitGame")
	}
	game.AddResourceCards(resourceCards...)

	developmentCards, err := g.developmentCardService.InitDevelopmentCards()
	if err != nil {
		return errors.Wrap(err, "service.GameService.InitGame")
	}
	game.AddDevelopmentCards(developmentCards...)

	terrains, err := g.terrainService.InitTerrains()
	if err != nil {
		return errors.Wrap(err, "service.GameService.InitGame")
	}
	game.AddTerrains(terrains...)

	robber, err := g.robberService.InitRobber(terrains)
	if err != nil {
		return errors.Wrap(err, "service.GameService.InitGame")
	}
	game.SetRobber(robber)

	constructions, err := g.constructionService.InitConstructions(terrains)
	if err != nil {
		return errors.Wrap(err, "service.GameService.InitGame")
	}
	game.AddConstructions(constructions...)

	roads, err := g.roadService.InitRoads(terrains)
	if err != nil {
		return errors.Wrap(err, "service.GameService.InitGame")
	}
	game.AddRoads(roads...)

	harbors, err := g.harborService.InitHarbors(terrains)
	if err != nil {
		return errors.Wrap(err, "service.GameService.InitGame")
	}
	game.AddHarbors(harbors...)

	return nil
}
