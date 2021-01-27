package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/middleware"
	"github.com/pkg/errors"
)

type RoomRepository interface {
	Count(ctx context.Context) (int, error)
	Find(ctx context.Context, limit int, offset int) (datamodel.Rooms, error)
}

func NewRoomRepository(transactionMiddleware *middleware.TransactionMiddleware,
	playerRepository PlayerRepository) RoomRepository {
	return &roomRepository{
		transactionMiddleware: transactionMiddleware,
		playerRepository:      playerRepository,
	}
}

type roomRepository struct {
	transactionMiddleware *middleware.TransactionMiddleware
	playerRepository      PlayerRepository
}

func (r roomRepository) Count(ctx context.Context) (int, error) {
	var count int64
	err := r.transactionMiddleware.Get(ctx).Model(&model.Game{}).Count(&count).Error
	return int(count), errors.Wrap(err, "repository.RoomRepository.GetRoomByGameId")
}

func (r roomRepository) Find(ctx context.Context, limit int, offset int) (datamodel.Rooms, error) {
	gameModels := make([]*model.Game, 0)

	err := r.transactionMiddleware.Get(ctx).Limit(limit).Offset(offset).Find(&gameModels).Error
	if err != nil {
		return nil, errors.Wrap(err, "repository.RoomRepository.Find")
	}

	rooms := make(datamodel.Rooms, 0)
	for _, gameModel := range gameModels {
		room, err := datamodel.NewRoomFromGameModel(gameModel)
		if err != nil {
			return nil, errors.Wrap(err, "repository.RoomRepository.Find")
		}

		players, err := r.playerRepository.FindByGameId(ctx, gameModel.ID)
		if err != nil {
			return nil, errors.Wrap(err, "repository.RoomRepository.Find")
		}
		room.AddPlayers(players...)

		rooms = append(rooms, room)
	}

	return rooms, nil
}
