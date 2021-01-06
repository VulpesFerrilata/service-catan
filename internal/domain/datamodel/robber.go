package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/pkg/errors"
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
	robber.terrain = desertTerrain

	return robber
}

type Robber struct {
	base
	id      int
	status  model.RobberStatus
	game    *Game
	terrain *Terrain
}

func (r *Robber) Persist(f func(robberModel *model.Robber) error) error {
	robberModel := new(model.Robber)
	robberModel.ID = r.id
	robberModel.Status = r.status
	if r.terrain != nil {
		robberModel.TerrainID = r.terrain.id
	}

	if err := f(robberModel); err != nil {
		return errors.Wrap(err, "model.Robber.Persist")
	}
	r.isModified = false
	r.isRemoved = false

	r.id = robberModel.ID
	r.status = robberModel.Status

	return nil
}
