package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

type Terrain struct {
	datamodel.Terrain
	game *Game
}

func (t *Terrain) SetGame(game *Game) {
	if game != nil {
		t.GameID = &game.id
	}
	t.game = game
}

func (t *Terrain) HasRobber() bool {
	if t.game.robber.TerrainID == nil {
		return false
	}
	return *t.game.robber.TerrainID == t.ID
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
