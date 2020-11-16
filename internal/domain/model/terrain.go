package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

func NewTerrain() *Terrain {
	terrain := new(Terrain)
	terrain.terrain = new(datamodel.Terrain)
	return terrain
}

type Terrain struct {
	terrain    *datamodel.Terrain
	game       *Game
	isModified bool
}

func (t *Terrain) GetTerrain() datamodel.Terrain {
	return *t.terrain
}

func (t *Terrain) GetId() uint {
	return t.terrain.ID
}

func (t *Terrain) GetGameId() *uint {
	return t.terrain.GameID
}

func (t *Terrain) setGame(game *Game) {
	if game != nil {
		t.terrain.GameID = &game.game.ID
		t.game = game
	}
}

func (t *Terrain) GetQ() int {
	return t.terrain.Q
}

func (t *Terrain) GetR() int {
	return t.terrain.R
}

func (t *Terrain) GetNumber() int {
	return t.terrain.Number
}

func (t *Terrain) GetType() datamodel.TerrainType {
	return t.terrain.Type
}

func (t *Terrain) IsModified() bool {
	return t.isModified
}

func (t *Terrain) HasRobber() bool {
	return t.game.robber.GetQ() == t.GetQ() && t.game.robber.GetR() == t.GetR()
}

func (t *Terrain) GetAdjacentConstructions() Constructions {
	return t.game.constructions.Filter(func(construction *Construction) bool {
		return (construction.GetQ() == t.GetQ()+1 && construction.GetR() == t.GetR()-1 && construction.GetLocation() == datamodel.CL_BOT) ||
			(construction.GetQ() == t.GetQ() && construction.GetR() == t.GetR()-1 && construction.GetLocation() == datamodel.CL_BOT) ||
			(construction.GetQ() == t.GetQ() && construction.GetR() == t.GetR() && construction.GetLocation() == datamodel.CL_TOP) ||
			(construction.GetQ() == t.GetQ() && construction.GetR() == t.GetR() && construction.GetLocation() == datamodel.CL_BOT) ||
			(construction.GetQ() == t.GetQ() && construction.GetR() == t.GetR()+1 && construction.GetLocation() == datamodel.CL_TOP) ||
			(construction.GetQ() == t.GetQ()-1 && construction.GetR() == t.GetR()+1 && construction.GetLocation() == datamodel.CL_TOP)
	})
}
