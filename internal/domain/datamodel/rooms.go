package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/pkg/errors"
)

func NewRoomsFromGameModels(gameModels []*model.Game) (Rooms, error) {
	rooms := make(Rooms, 0)

	for _, gameModel := range gameModels {
		room, err := NewRoomFromGameModel(gameModel)
		if err != nil {
			return nil, errors.Wrap(err, "datamodel.NewRoomsFromGameModels")
		}
		rooms = append(rooms, room)
	}

	return rooms, nil
}

type Rooms []*Room
