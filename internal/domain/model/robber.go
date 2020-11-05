package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

type Robber struct {
	*datamodel.Robber
	game *Game
}

func (r *Robber) SetGame(game *Game) {
	r.game.ID = game.ID
	r.game = game
	r.game.robber = r
}
