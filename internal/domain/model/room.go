package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

func NewRoom(game *Game) *Room {
	room := new(Room)
	room.ID = game.ID
	room.Status = game.Status
	for _, player := range game.GetPlayers() {
		room.users = append(room.users, player.GetUser())
	}
	return room
}

func NewRooms(games []*Game) []*Room {
	rooms := make([]*Room, 0)

	for _, game := range games {
		rooms = append(rooms, NewRoom(game))
	}

	return rooms
}

type Room struct {
	ID     uint
	Status datamodel.GameStatus
	users  []*User
}
