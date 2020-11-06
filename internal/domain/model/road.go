package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

func NewRoad(game *Game) *Road {
	road := new(Road)
	road.Road = new(datamodel.Road)
	road.SetGame(game)
	return road
}

type Road struct {
	*datamodel.Road
	game *Game
}

func (r *Road) SetGame(game *Game) {
	r.game = game
	game.roads.append(r)
}

func (r *Road) GetPlayer(game *Game) *Player {
	if r.PlayerID == nil {
		return nil
	}

	return r.game.players.Filter(func(player *Player) bool {
		return player.ID == *r.PlayerID
	}).First()
}
