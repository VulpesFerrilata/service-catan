package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/google/uuid"
)

func NewConstructionFromConstructionModel(constructionModel *model.Construction) *Construction {
	construction := new(Construction)
	construction.id = constructionModel.ID
	construction.q = constructionModel.Q
	construction.r = constructionModel.R
	construction.location = constructionModel.Location
	construction.constructionType = constructionModel.ConstructionType
	construction.isModified = false
	construction.isRemoved = false
	return construction
}

type Construction struct {
	base
	id               uuid.UUID
	q                int
	r                int
	location         model.ConstructionLocation
	constructionType model.ConstructionType
	game             *Game
	playerID         *uuid.UUID
}

func (c Construction) ToModel() *model.Construction {
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
