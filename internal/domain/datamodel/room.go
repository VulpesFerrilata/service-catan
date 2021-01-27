package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func NewRoomFromGameModel(gameModel *model.Game) (*Room, error) {
	room := new(Room)
	room.id = gameModel.ID

	gameStatus, err := NewGameStatus(gameModel.Status)
	if err != nil {
		return nil, errors.Wrap(err, "datamodel.NewRoomFromGameModel")
	}
	room.status = gameStatus

	return room, nil
}

type Room struct {
	id      uuid.UUID
	status  gameStatus
	players Players
}

func (r *Room) AddPlayers(players ...*Player) {
	r.players = append(r.players, players...)
}
