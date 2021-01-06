package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
)

type SafeRoomRepository interface {
	GetRoomByGameId(ctx context.Context, gameId int) (*datamodel.Room, error)
	FindRoomByGameStatus(ctx context.Context, status model.GameStatus) ([]*datamodel.Room, error)
}

type RoomRepository interface {
	SafeRoomRepository
}
