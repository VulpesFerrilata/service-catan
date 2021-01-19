package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func NewConstruction(q int, r int, constructionLocation model.ConstructionLocation, constructionType model.ConstructionType) (*Construction, error) {
	construction := new(Construction)
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.Wrap(err, "datamodel.NewConstruction")
	}
	construction.id = id
	construction.q = q
	construction.r = r
	construction.location = constructionLocation
	construction.constructionType = constructionType
	construction.playerID = nil

	construction.SetModelState(Added)

	return construction, nil
}

func NewConstructionFromConstructionModel(constructionModel *model.Construction) *Construction {
	construction := new(Construction)
	construction.id = constructionModel.ID
	construction.q = constructionModel.Q
	construction.r = constructionModel.R
	construction.location = constructionModel.Location
	construction.constructionType = constructionModel.ConstructionType
	construction.playerID = constructionModel.PlayerID

	construction.SetModelState(Unchanged)

	return construction
}

type Construction struct {
	base
	id               uuid.UUID
	q                int
	r                int
	location         model.ConstructionLocation
	constructionType model.ConstructionType
	playerID         *uuid.UUID
	game             *Game
}

func (c *Construction) ToModel() *model.Construction {
	c.SetModelState(Unchanged)

	constructionModel := new(model.Construction)
	constructionModel.ID = c.id
	if c.game != nil {
		constructionModel.GameID = c.game.id
	}
	constructionModel.Q = c.q
	constructionModel.R = c.r
	constructionModel.Location = c.location
	constructionModel.PlayerID = c.playerID
	constructionModel.ConstructionType = c.constructionType
	return constructionModel
}
