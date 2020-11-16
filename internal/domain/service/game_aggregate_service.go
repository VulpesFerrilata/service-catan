package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
)

type GameAggregateService interface {
	GetById(ctx context.Context, id uint) (*model.Game, error)
	Save(ctx context.Context, game *model.Game) error
}

type gameAggregateService struct {
	gameService            GameService
	playerAggregateService PlayerAggregateService
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

func (gas *gameAggregateService) GetById(ctx context.Context, id uint) (*model.Game, error) {
	game, err := gas.gameService.GetGameRepository().GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	players, err := gas.playerAggregateService.FindByGameId(ctx, game.GetId())
	if err != nil {
		return nil, err
	}
	game.AddPlayers(players...)

	dices, err := gas.diceService.GetDiceRepository().FindByGameId(ctx, game.GetId())
	if err != nil {
		return nil, err
	}
	game.AddDices(dices...)

	achievements, err := gas.achievementService.GetAchievementRepository().FindByGameId(ctx, game.GetId())
	if err != nil {
		return nil, err
	}
	game.AddAchievements(achievements...)

	resourceCards, err := gas.resourceCardService.GetResourceCardRepository().FindByGameId(ctx, game.GetId())
	if err != nil {
		return nil, err
	}
	game.AddResourceCards(resourceCards...)

	developmentCards, err := gas.developmentCardService.GetDevelopmentCardRepository().FindByGameId(ctx, game.GetId())
	if err != nil {
		return nil, err
	}
	game.AddDevelopmentCards(developmentCards...)

	terrains, err := gas.terrainService.GetFieldRepository().FindByGameId(ctx, game.GetId())
	if err != nil {
		return nil, err
	}
	game.AddTerrains(terrains...)

	robber, err := gas.robberService.GetRobberRepository().GetByGameId(ctx, game.GetId())
	if err != nil {
		return nil, err
	}
	game.SetRobber(robber)

	constructions, err := gas.constructionService.GetConstructionRepository().FindByGameId(ctx, game.GetId())
	if err != nil {
		return nil, err
	}
	game.AddConstructions(constructions...)

	roads, err := gas.roadService.GetRoadRepository().FindByGameId(ctx, game.GetId())
	if err != nil {
		return nil, err
	}
	game.AddRoads(roads...)

	harbors, err := gas.harborService.GetHarborRepository().FindByGameId(ctx, game.GetId())
	if err != nil {
		return nil, err
	}
	game.AddHarbors(harbors...)

	return game, nil
}

func (gas *gameAggregateService) Save(ctx context.Context, game *model.Game) error {
	if err := gas.gameService.Save(ctx, game); err != nil {
		return err
	}

	players := game.GetPlayers()
	for _, player := range players {
		if err := gas.playerAggregateService.Save(ctx, player); err != nil {
			return err
		}
	}

	dices := game.GetDices()
	for _, dice := range dices {
		if err := gas.diceService.Save(ctx, dice); err != nil {
			return err
		}
	}

	achievements := game.GetAchievements()
	for _, achievement := range achievements {
		if err := gas.achievementService.Save(ctx, achievement); err != nil {
			return err
		}
	}

	resourceCards := game.GetResourceCards()
	for _, resourceCard := range resourceCards {
		if err := gas.resourceCardService.Save(ctx, resourceCard); err != nil {
			return err
		}
	}

	developmentCards := game.GetDevelopmentCards()
	for _, developmentCard := range developmentCards {
		if err := gas.developmentCardService.Save(ctx, developmentCard); err != nil {
			return err
		}
	}

	terrains := game.GetTerrains()
	for _, terrain := range terrains {
		if err := gas.terrainService.Save(ctx, terrain); err != nil {
			return err
		}
	}

	robber := game.GetRobber()
	if err := gas.robberService.Save(ctx, robber); err != nil {
		return err
	}

	constructions := game.GetConstructions()
	for _, construction := range constructions {
		if err := gas.constructionService.Save(ctx, construction); err != nil {
			return err
		}
	}

	roads := game.GetRoads()
	for _, road := range roads {
		if err := gas.roadService.Save(ctx, road); err != nil {
			return err
		}
	}

	harbors := game.GetHarbors()
	for _, harbor := range harbors {
		if err := gas.harborService.Save(ctx, harbor); err != nil {
			return err
		}
	}

	return nil
}
