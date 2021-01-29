package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func NewTerrain(hex *Hex, number int, terrainType TerrainType) (*Terrain, error) {
	terrain := new(Terrain)

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.Wrap(err, "datamodel.NewTerrain")
	}
	terrain.id = id

	terrain.hex = hex
	terrain.number = number
	terrain.terrainType = terrainType

	return terrain, nil
}

func NewTerrainFromModel(terrainModel *model.Terrain, hex *Hex) (*Terrain, error) {
	terrain := new(Terrain)
	terrain.id = terrainModel.ID
	terrain.number = terrainModel.Number

	terrainType, err := NewTerrainType(terrainModel.TerrainType)
	if err != nil {
		return nil, errors.Wrap(err, "datamodel.NewTerrainFromTerrainModel")
	}
	terrain.terrainType = terrainType

	terrain.hex = hex
	return terrain, nil
}

type Terrain struct {
	id          uuid.UUID
	hex         *Hex
	number      int
	terrainType TerrainType
	game        *Game
	harbor      *Harbor
	robber      *Robber
}

func (t Terrain) GetHex() *Hex {
	return t.hex
}

func (t Terrain) GetTerrainType() TerrainType {
	return t.terrainType
}

func (t Terrain) GetAdjacentTerrains() (Terrains, error) {
	possibleAdjacentHexes, err := t.GetHex().GetPossibleAdjacentHexes()
	if err != nil {
		return nil, errors.Wrap(err, "datamodel.Terrain.GetAdjacentTerrains")
	}
	return t.game.terrains.Filter(func(terrain *Terrain) bool {
		for _, possibleAdjacentHex := range possibleAdjacentHexes {
			if terrain.GetHex().Equals(possibleAdjacentHex) {
				return true
			}
		}
		return false
	}), nil
}

func (t Terrain) GetAdjacentConstructions() Constructions {
	possibleAdjacentHexCorners := t.GetHex().GetPossibleAdjacentHexCorners()
	return t.game.constructions.Filter(func(construction *Construction) bool {
		for _, possibleAdjacentHexCorner := range possibleAdjacentHexCorners {
			if construction.GetHexCorner().Equals(possibleAdjacentHexCorner) {
				return true
			}
		}
		return false
	})
}

func (t Terrain) ToModel() *model.Terrain {
	terrainModel := new(model.Terrain)
	terrainModel.ID = t.id
	terrainModel.GameID = t.game.id
	terrainModel.HexID = t.hex.id
	terrainModel.Number = t.number
	terrainModel.TerrainType = t.terrainType.String()
	return terrainModel
}
