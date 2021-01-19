package datamodel

import (
	"math"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func NewTerrain(q int, r int, number int, terrainType model.TerrainType) (*Terrain, error) {
	terrain := new(Terrain)
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.Wrap(err, "datamodel.NewTerrain")
	}
	terrain.id = id
	terrain.q = q
	terrain.r = r
	terrain.number = number
	terrain.terrainType = terrainType

	terrain.SetModelState(Added)

	return terrain, nil
}

func NewTerrainFromTerrainModel(terrainModel *model.Terrain) *Terrain {
	terrain := new(Terrain)
	terrain.id = terrainModel.ID
	terrain.q = terrainModel.Q
	terrain.r = terrainModel.R
	terrain.number = terrainModel.Number
	terrain.terrainType = terrainModel.TerrainType

	terrain.SetModelState(Unchanged)

	return terrain
}

type Terrain struct {
	base
	id          uuid.UUID
	q           int
	r           int
	number      int
	terrainType model.TerrainType
	game        *Game
	harbor      *Harbor
	robber      *Robber
}

func (t Terrain) GetQ() int {
	return t.q
}

func (t Terrain) GetR() int {
	return t.r
}

func (t Terrain) GetTerrainType() model.TerrainType {
	return t.terrainType
}

func (t Terrain) GetAdjacentTerrains() Terrains {
	return t.game.terrains.Filter(func(terrain *Terrain) bool {
		if terrain.q == t.q && terrain.r == t.r {
			return false
		}
		if math.Abs(float64(terrain.q - t.q)) <= 1 && math.Abs(float64(terrain.r - t.r)) <= 1 {
			return true
		}
		return return false
	})
}

func (t Terrain) GetAdjacentConstructions() Constructions {
	return t.game.constructions.Filter(func(construction *Construction) bool {
		return (construction.q == t.q+1 && construction.r == t.r-1 && construction.location == model.Bottom) ||
			(construction.q == t.q && construction.r == t.r-1 && construction.location == model.Bottom) ||
			(construction.q == t.q && construction.r == t.r && construction.location == model.Top) ||
			(construction.q == t.q && construction.r == t.r && construction.location == model.Bottom) ||
			(construction.q == t.q && construction.r == t.r+1 && construction.location == model.Top) ||
			(construction.q == t.q-1 && construction.r == t.r+1 && construction.location == model.Top)
	})
}

func (t Terrain) ToModel() *model.Terrain {
	terrainModel := new(model.Terrain)
	terrainModel.ID = t.id
	terrainModel.Q = t.q
	terrainModel.R = t.r
	terrainModel.Number = t.number
	terrainModel.TerrainType = t.terrainType
	return terrainModel
}
