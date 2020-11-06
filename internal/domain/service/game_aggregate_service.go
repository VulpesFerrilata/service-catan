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
	fieldService           FieldService
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

	players, err := gas.playerAggregateService.FindByGameId(ctx, game.ID)
	if err != nil {
		return nil, err
	}
	players.SetGame(game)

	dices, err := gas.diceService.GetDiceRepository().FindByGameId(ctx, game.ID)
	if err != nil {
		return nil, err
	}
	dices.SetGame(game)

	achievements, err := gas.achievementService.GetAchievementRepository().FindByGameId(ctx, game.ID)
	if err != nil {
		return nil, err
	}
	achievements.SetGame(game)

	resourceCards, err := gas.resourceCardService.GetResourceCardRepository().FindByGameId(ctx, game.ID)
	if err != nil {
		return nil, err
	}
	resourceCards.SetGame(game)

	developmentCards, err := gas.developmentCardService.GetDevelopmentCardRepository().FindByGameId(ctx, game.ID)
	if err != nil {
		return nil, err
	}
	developmentCards.SetGame(game)

	fields, err := gas.fieldService.GetFieldRepository().FindByGameId(ctx, game.ID)
	if err != nil {
		return nil, err
	}
	fields.SetGame(game)

	robber, err := gas.robberService.GetRobberRepository().GetByGameId(ctx, game.ID)
	if err != nil {
		return nil, err
	}
	robber.SetGame(game)

	constructions, err := gas.constructionService.GetConstructionRepository().FindByGameId(ctx, game.ID)
	if err != nil {
		return nil, err
	}
	constructions.SetGame(game)

	roads, err := gas.roadService.GetRoadRepository().FindByGameId(ctx, game.ID)
	if err != nil {
		return nil, err
	}
	roads.SetGame(game)

	harbors, err := gas.harborService.GetHarborRepository().FindByGameId(ctx, game.ID)
	if err != nil {
		return nil, err
	}
	harbors.SetGame(game)

	return game, nil
}

func (gas *gameAggregateService) Save(ctx context.Context, game *model.Game) error {
	if err := gas.gameService.Save(ctx, game); err != nil {
		return err
	}

	players := game.GetPlayers()
	for _, player := range players {
		player.GameID = game.ID
		if err := gas.playerAggregateService.Save(ctx, player); err != nil {
			return err
		}
	}

	dices := game.GetDices()
	for _, dice := range dices {
		dice.GameID = game.ID
		if err := gas.diceService.Save(ctx, dice); err != nil {
			return err
		}
	}

	achievements := game.GetAchievements()
	for _, achievement := range achievements {
		achievement.GameID = game.ID
		if err := gas.achievementService.Save(ctx, achievement); err != nil {
			return err
		}
	}

	resourceCards := game.GetResourceCards()
	for _, resourceCard := range resourceCards {
		resourceCard.GameID = game.ID
		if err := gas.resourceCardService.Save(ctx, resourceCard); err != nil {
			return err
		}
	}

	developmentCards := game.GetDevelopmentCards()
	for _, developmentCard := range developmentCards {
		developmentCard.GameID = game.ID
		if err := gas.developmentCardService.Save(ctx, developmentCard); err != nil {
			return err
		}
	}

	fields := game.GetFields()
	for _, field := range fields {
		field.GameID = game.ID
		if err := gas.fieldService.Save(ctx, field); err != nil {
			return err
		}
	}

	robber := game.GetRobber()
	robber.GameID = game.ID
	if err := gas.robberService.Save(ctx, robber); err != nil {
		return err
	}

	constructions := game.GetConstructions()
	for _, construction := range constructions {
		construction.GameID = game.ID
		if err := gas.constructionService.Save(ctx, construction); err != nil {
			return err
		}
	}

	roads := game.GetRoads()
	for _, road := range roads {
		road.GameID = game.ID
		if err := gas.roadService.Save(ctx, road); err != nil {
			return err
		}
	}

	harbors := game.GetHarbors()
	for _, harbor := range harbors {
		harbor.GameID = game.ID
		if err := gas.harborService.Save(ctx, harbor); err != nil {
			return err
		}
	}

	return nil
}
