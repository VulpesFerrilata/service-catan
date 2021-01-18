package datamodel

import "github.com/VulpesFerrilata/catan/internal/domain/model"

func NewRoomsFromGameModels(gameModels []*model.Game) Rooms {
	rooms := make(Rooms, 0)

	for _, gameModel := range gameModels {
		room := NewRoomFromGameModel(gameModel)
		rooms = append(rooms, room)
	}

	return rooms
}

type Rooms []*Room
