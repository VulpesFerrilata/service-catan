package service

import (
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
)

type GameService interface {
	GetGameRepository() repository.GameRepository
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

func (gs gameService) GetGameRepository() repository.GameRepository {
	return gs.gameRepository
}
