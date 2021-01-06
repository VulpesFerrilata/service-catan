package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/pkg/errors"
)

func NewTerrainFromTerrainModel(terrainModel *model.Terrain) *Terrain {
	terrain := new(Terrain)
	terrain.id = terrainModel.ID
	terrain.q = terrainModel.Q
	terrain.r = terrainModel.R
	terrain.number = terrainModel.Number
	terrain.terrainType = terrainModel.TerrainType
	terrain.isModified = false
	terrain.isRemoved = false
	return terrain
}

type Terrain struct {
	base
	id          int
	q           int
	r           int
	number      int
	terrainType model.TerrainType
	game        *Game
	harbor      *Harbor
	robber      *Robber
}

func (t Terrain) GetAdjacentConstructions() Constructions {
	return t.game.constructions.Filter(func(construction *Construction) bool {
		return (construction.q == t.q+1 && construction.r == t.r-1 && construction.location == datamodel.CL_BOT) ||
			(construction.q == t.q && construction.r == t.r-1 && construction.location == datamodel.CL_BOT) ||
			(construction.q == t.q && construction.r == t.r && construction.location == datamodel.CL_TOP) ||
			(construction.q == t.q && construction.r == t.r && construction.location == datamodel.CL_BOT) ||
			(construction.q == t.q && construction.r == t.r+1 && construction.location == datamodel.CL_TOP) ||
			(construction.q == t.q-1 && construction.r == t.r+1 && construction.location == datamodel.CL_TOP)
	})
}

func (t *Terrain) Persist(f func(terrainModel *model.Terrain) error) error {
	terrainModel := new(model.Terrain)
	terrainModel.ID = t.id
	terrainModel.Q = t.q
	terrainModel.R = t.r
	terrainModel.Number = t.number
	terrainModel.TerrainType = t.terrainType

	if err := f(terrainModel); err != nil {
		return errors.Wrap(err, "model.Terrain.Persist")
	}
	t.isModified = false
	t.isRemoved = false

	t.id = terrainModel.ID
	t.q = terrainModel.Q
	t.r = terrainModel.R
	t.number = terrainModel.Number
	t.terrainType = terrainModel.TerrainType

	return nil
}
