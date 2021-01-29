package datamodel

import (
	"fmt"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func NewConstruction(hexCorner *HexCorner) (*Construction, error) {
	construction := new(Construction)

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.Wrap(err, "datamodel.NewConstruction")
	}
	construction.id = id

	construction.hexCorner = hexCorner
	construction.constructionType = Land
	construction.playerID = nil
	return construction, nil
}

func NewConstructionFromModel(constructionModel *model.Construction) (*Construction, error) {
	construction := new(Construction)
	construction.id = constructionModel.ID

	location, err := NewHexCornerLocation(constructionModel.Location)
	if err != nil {
		return nil, fmt.Errorf("hex corner location is invalid: %s", constructionModel.Location)
	}
	hexCorner := NewHexCorner(constructionModel.Q, constructionModel.R, location)
	construction.hexCorner = hexCorner

	constructionType, err := NewConstructionType(constructionModel.ConstructionType)
	if err != nil {
		return nil, fmt.Errorf("construction type is invalid: %s", constructionModel.ConstructionType)
	}
	construction.constructionType = constructionType

	construction.playerID = constructionModel.PlayerID
	return construction, nil
}

type Construction struct {
	id        uuid.UUID
	hexCorner *HexCorner
	constructionType
	playerID *uuid.UUID
	game     *Game
}

func (c Construction) GetHexCorner() *HexCorner {
	return c.hexCorner
}

func (c Construction) ToModel() *model.Construction {
	constructionModel := new(model.Construction)
	constructionModel.ID = c.id

	if c.game != nil {
		constructionModel.GameID = c.game.id
	}

	constructionModel.Q = c.GetHexCorner().GetQ()
	constructionModel.R = c.GetHexCorner().GetR()
	constructionModel.Location = c.GetHexCorner().GetLocation().String()
	constructionModel.PlayerID = c.playerID
	constructionModel.ConstructionType = c.constructionType.String()
	return constructionModel
}
