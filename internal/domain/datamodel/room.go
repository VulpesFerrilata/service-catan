package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/google/uuid"
)

func NewRoomFromGameModel(gameModel *model.Game) *Room {
	room := new(Room)
	room.id = gameModel.ID
	room.status = gameModel.Status
	return room
}

type Room struct {
	id      uuid.UUID
	status  model.GameStatus
	players Players
}

func (r *Room) AddPlayers(players ...*Player) {
	r.players = append(r.players, players...)
}
