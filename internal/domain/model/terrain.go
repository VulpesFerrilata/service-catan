package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

func NewTerrain(game *Game) *Terrain {
	terrain := new(Terrain)
	terrain.Terrain = new(datamodel.Terrain)
	terrain.SetGame(game)
	return terrain
}

type Terrain struct {
	*datamodel.Terrain
	game *Game
}

func (t *Terrain) SetGame(game *Game) {
	t.game = game
	game.terrains.append(t)
}

func (t *Terrain) HasRobber() bool {
	return t.game.robber.Q == t.Q && t.game.robber.R == t.R
}

func (t *Terrain) GetAdjacentConstructions() Constructions {
	return t.game.constructions.Filter(func(construction *Construction) bool {
		return (construction.Q == t.Q+1 && construction.R == t.R-1 && construction.Location == datamodel.CL_BOT) ||
			(construction.Q == t.Q && construction.R == t.R-1 && construction.Location == datamodel.CL_BOT) ||
			(construction.Q == t.Q && construction.R == t.R && construction.Location == datamodel.CL_TOP) ||
			(construction.Q == t.Q && construction.R == t.R && construction.Location == datamodel.CL_BOT) ||
			(construction.Q == t.Q && construction.R == t.R+1 && construction.Location == datamodel.CL_TOP) ||
			(construction.Q == t.Q-1 && construction.R == t.R+1 && construction.Location == datamodel.CL_TOP)
	})
}
