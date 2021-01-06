package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/pkg/errors"
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
	id               int
	q                int
	r                int
	location         model.ConstructionLocation
	constructionType model.ConstructionType
	game             *Game
	player           *Player
}

func (c Construction) GetPlayer() *Player {
	return c.player
}

func (c *Construction) Persist(f func(constructionModel *model.Construction) error) error {
	constructionModel := new(model.Construction)
	constructionModel.ID = c.id
	if c.game != nil {
		constructionModel.GameID = c.game.id
	}
	constructionModel.Q = c.q
	constructionModel.R = c.r
	constructionModel.Location = c.location
	if c.player != nil {
		constructionModel.PlayerID = &c.player.id
	}
	constructionModel.ConstructionType = c.constructionType

	if err := f(constructionModel); err != nil {
		return errors.Wrap(err, "datamodel.Construction.Persist")
	}
	c.isModified = false
	c.isRemoved = false

	c.id = constructionModel.ID
	c.q = constructionModel.Q
	c.r = constructionModel.R
	c.location = constructionModel.Location
	c.constructionType = constructionModel.ConstructionType

	return nil
}
