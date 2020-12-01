package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
)

type RoomService interface {
	GetRoomById(ctx context.Context, id uint) (*model.Room, error)
	FindRoomsByStatus(ctx context.Context, status datamodel.GameStatus) ([]*model.Room, error)
}

func NewRoomService() RoomService {
	return &roomService{}
}

type roomService struct {
	gameService            GameService
	playerAggregateService PlayerAggregateService
}

func (rs *roomService) GetRoomById(ctx context.Context, id uint) (*model.Room, error) {
	game, err := rs.gameService.GetGameRepository().GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	players, err := rs.playerAggregateService.FindByGameId(ctx, id)
	if err != nil {
		return nil, err
	}
	game.AddPlayers(players...)

	return model.NewRoom(game), nil
}

func (rs *roomService) FindRoomsByStatus(ctx context.Context, status datamodel.GameStatus) ([]*model.Room, error) {
	games, err := rs.gameService.GetGameRepository().FindByGameStatus(ctx, status)
	if err != nil {
		return nil, err
	}

	for _, game := range games {
		players, err := rs.playerAggregateService.FindByGameId(ctx, game.ID)
		if err != nil {
			return nil, err
		}
		game.AddPlayers(players...)
	}

	return model.NewRooms(games), nil
}
