package datamodel

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

type Room struct {
	id      int
	status  datamodel.GameStatus
	players Players
}

func (r *Room) AddPlayers(players ...*Player) {
	r.players = append(r.players, players...)
}
