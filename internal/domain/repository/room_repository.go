package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/middleware"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type RoomRepository interface {
	GetRoomByGameId(ctx context.Context, gameId uuid.UUID) (*datamodel.Room, error)
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

func (r roomRepository) GetRoomByGameId(ctx context.Context, gameId uuid.UUID) (*datamodel.Room, error) {
	gameModel := new(model.Game)

	err := r.transactionMiddleware.Get(ctx).First(gameModel, gameId).Error
	if err != nil {
		return nil, errors.Wrap(err, "repository.RoomRepository.GetRoomByGameId")
	}

	room, err := datamodel.NewRoomFromGameModel(gameModel)
	return room, errors.Wrap(err, "repository.RoomRepository.GetRoomByGameId")
}
