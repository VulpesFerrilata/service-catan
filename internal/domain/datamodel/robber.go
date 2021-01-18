package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/google/uuid"
)

func NewRobberFromRobberModel(robberModel *model.Robber) *Robber {
	robber := new(Robber)
	robber.id = robberModel.ID
	robber.status = robberModel.Status
	robber.isModified = false
	robber.isRemoved = false
	return robber
}

func NewRobber(terrains Terrains) *Robber {
	robber := new(Robber)
	robber.status = model.Idle

	desertTerrain := terrains.Filter(func(terrain *Terrain) bool {
		if terrain.terrainType == model.DesertTerrain {
			return true
		}
		return false
	}).First()
	robber.terrainID = desertTerrain.id

	return robber
}

type Robber struct {
	base
	id        uuid.UUID
	status    model.RobberStatus
	terrainID uuid.UUID
	game      *Game
}

func (r Robber) ToModel() *model.Robber {
	robberModel := new(model.Robber)
	robberModel.ID = r.id
	robberModel.Status = r.status
	robberModel.TerrainID = r.terrainID
	return robberModel
}
