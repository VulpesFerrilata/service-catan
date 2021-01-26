package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func NewRoad(hexEdge *HexEdge) (*Road, error) {
	road := new(Road)

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.Wrap(err, "datamodel.NewRoad")
	}
	road.id = id

	road.hexEdge = hexEdge
	road.playerID = nil
	return road, nil
}

func NewRoadFromRoadModel(roadModel *model.Road) (*Road, error) {
	road := new(Road)
	road.id = roadModel.ID

	location, err := NewHexEdgeLocation(roadModel.Location)
	if err != nil {
		return nil, errors.Wrap(err, "datamodel.NewRoadFromRoadModel")
	}
	hexEdge := NewHexEdge(roadModel.Q, roadModel.R, location)
	road.hexEdge = hexEdge

	road.playerID = roadModel.PlayerID
	return road, nil
}

type Road struct {
	id       uuid.UUID
	hexEdge  *HexEdge
	playerID *uuid.UUID
	game     *Game
}

func (r Road) GetHexEdge() *HexEdge {
	return r.hexEdge
}

func (r Road) GetPlayer(game *Game) *Player {
	if r.playerID == nil {
		return nil
	}

	return r.game.players.Filter(func(player *Player) bool {
		return player.id == *r.playerID
	}).First()
}

func (r Road) GetAdjacentConstructions() Constructions {
	possibleAdjacentHexCorners := r.GetHexEdge().GetPossibleAdjacentHexCorners()
	return r.game.constructions.Filter(func(construction *Construction) bool {
		for _, possibleAdjacentHexCorner := range possibleAdjacentHexCorners {
			if construction.GetHexCorner().Equals(possibleAdjacentHexCorner) {
				return true
			}
		}
		return false
	})
}

func (r Road) GetAdjacentRoads() Roads {
	possibleAdjacentHexEdges := r.GetHexEdge().GetPossibleAdjacentHexEdges()
	return r.game.roads.Filter(func(road *Road) bool {
		for _, possibleAdjacentHexEdge := range possibleAdjacentHexEdges {
			if road.GetHexEdge().Equals(possibleAdjacentHexEdge) {
				return true
			}
		}
		return false
	})
}

func (r Road) ToModel() *model.Road {
	roadModel := new(model.Road)
	roadModel.ID = r.id
	roadModel.Q = r.GetHexEdge().GetQ()
	roadModel.R = r.GetHexEdge().GetR()
	roadModel.Location = r.GetHexEdge().GetLocation().String()
	roadModel.PlayerID = r.playerID
	return roadModel
}
