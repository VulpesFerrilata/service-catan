package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func NewRobber(terrain *Terrain) (*Robber, error) {
	robber := new(Robber)
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.Wrap(err, "datamodel.NewRobber")
	}
	robber.id = id
	robber.status = model.Idle
	robber.terrainID = terrain.id

	robber.SetModelState(Added)

	return robber, nil
}

func NewRobberFromRobberModel(robberModel *model.Robber) *Robber {
	robber := new(Robber)
	robber.id = robberModel.ID
	robber.status = robberModel.Status

	robber.SetModelState(Unchanged)

	return robber
}

type Robber struct {
	base
	id        uuid.UUID
	status    model.RobberStatus
	terrainID uuid.UUID
	game      *Game
}

func (r *Robber) ToModel() *model.Robber {
	r.SetModelState(Unchanged)

	robberModel := new(model.Robber)
	robberModel.ID = r.id
	robberModel.Status = r.status
	robberModel.TerrainID = r.terrainID
	return robberModel
}
