package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/grpc/protoc/user"
)

type GameAggregateService interface {
	GetById(ctx context.Context, id uint) (*model.Game, error)
	Save(ctx context.Context, game *model.Game) error
}

type gameAggregateService struct {
	gameService   GameService
	playerService PlayerService
	userService   user.UserService
	diceService   DiceService
}

func (gas gameAggregateService) GetById(ctx context.Context, id uint) (*model.Game, error) {
	game, err := gas.gameService.GetGameRepository().GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	players, err := gas.playerService.GetPlayerRepository().FindByGameId(ctx, game.ID)
	if err != nil {
		return nil, err
	}
	for _, player := range players {
		game.AddPlayer(player)

		userRequest := new(user.UserRequest)
		userRequest.ID = int64(player.UserID)
		userPb, err := gas.userService.GetUserById(ctx, userRequest)
		if err != nil {
			return nil, err
		}
		user := model.NewUser(userPb)
		player.SetUser(user)
	}

	dices, err := gas.diceService.GetDiceRepository().FindByGameId(ctx, game.ID)
	if err != nil {
		return nil, err
	}
	for _, dice := range dices {
		game.AddDice(dice)
	}

	return game, nil
}

func (gas gameAggregateService) Save(ctx context.Context, game *model.Game) error {
	if err := gas.gameService.Save(ctx, game); err != nil {
		return err
	}

	players := game.FilterPlayers(func(p *model.Player) bool {
		return true
	})
	for _, player := range players {
		if err := gas.playerService.Save(ctx, player); err != nil {
			return err
		}
	}

	dices := game.GetDices()
	for _, dice := range dices {
		if err := gas.diceService.Save(ctx, dice); err != nil {
			return err
		}
	}

	return nil
}
