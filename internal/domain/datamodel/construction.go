package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func NewConstruction(q int, r int, hexEdge hexEdge) (*Construction, error) {
	construction := new(Construction)

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.Wrap(err, "datamodel.NewConstruction")
	}
	construction.id = id

	construction.hexEdge = hexEdge
	construction.constructionType = Land
	construction.playerID = nil

	construction.SetModelState(Added)

	return construction, nil
}

func NewConstructionFromConstructionModel(constructionModel *model.Construction) (*Construction, error) {
	construction := new(Construction)
	construction.id = constructionModel.ID

	hex := NewHex(constructionModel.Q, constructionModel.R)
	location, err := NewHexEdgeLocation(constructionModel.Location)
	if err != nil {
		return nil, errors.Wrap(err, "datamodel.NewConstructionFromConstructionModel")
	}
	hexEdge := NewHexEdge(hex, location)
	construction.hexEdge = hexEdge

	constructionType, err := NewConstructionType(constructionModel.ConstructionType)
	if err != nil {
		return nil, errors.Wrap(err, "datamodel.NewConstructionFromConstructionModel")
	}
	construction.constructionType = constructionType
	construction.playerID = constructionModel.PlayerID

	construction.SetModelState(Unchanged)

	return construction, nil
}

type Construction struct {
	base
	id uuid.UUID
	hexEdge
	constructionType
	playerID *uuid.UUID
	game     *Game
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
	constructionModel.Location = c.location.String()
	constructionModel.PlayerID = c.playerID
	constructionModel.ConstructionType = c.constructionType.String()

	return constructionModel
}
