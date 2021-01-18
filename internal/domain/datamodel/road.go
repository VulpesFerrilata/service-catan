package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/google/uuid"
)

func NewRoadFromRoadModel(roadModel *model.Road) *Road {
	road := new(Road)
	road.id = roadModel.ID
	road.q = roadModel.Q
	road.r = roadModel.R
	road.location = roadModel.Location
	road.playerID = road.playerID
	road.isModified = false
	road.isRemoved = false
	return road
}

type Road struct {
	base
	id       uuid.UUID
	q        int
	r        int
	location model.RoadLocation
	playerID *uuid.UUID
	game     *Game
}

func (r *Road) GetPlayer(game *Game) *Player {
	if r.playerID == nil {
		return nil
	}

	return r.game.players.Filter(func(player *Player) bool {
		return player.id == *r.playerID
	}).First()
}

func (r *Road) GetAdjacentConstructions() Constructions {
	return r.game.constructions.Filter(func(construction *Construction) bool {
		if r.location == model.TopLeft {
			return (construction.q == r.q && construction.r == r.r && construction.location == model.Top) ||
				(construction.q == r.q && construction.r == r.r-1 && construction.location == model.Bottom)
		} else if r.location == model.MiddleLeft {
			return (construction.q == r.q && construction.r == r.r-1 && construction.location == model.Bottom) ||
				(construction.q == r.q-1 && construction.r == r.r+1 && construction.location == model.Top)
		}
		return (construction.q == r.q && construction.r == r.r && construction.location == model.Bottom) ||
			(construction.q == r.q-1 && construction.r == r.r+1 && construction.location == model.Top)
	})
}

func (r *Road) GetAdjacentRoads() Roads {
	return r.game.roads.Filter(func(road *Road) bool {
		if r.location == model.TopLeft {
			return (road.q == r.q+1 && road.r == r.r-1 && road.location == model.MiddleLeft) ||
				(road.q == r.q+1 && road.r == r.r-1 && road.location == model.BottomLeft) ||
				(road.q == r.q && road.r == r.r-1 && road.location == model.BottomLeft) ||
				(road.q == r.q && road.r == r.r && road.location == model.MiddleLeft)
		} else if r.location == model.MiddleLeft {
			return (road.q == r.q && road.r == r.r && road.location == model.TopLeft) ||
				(road.q == r.q && road.r == r.r && road.location == model.BottomLeft) ||
				(road.q == r.q && road.r == r.r-1 && road.location == model.BottomLeft) ||
				(road.q == r.q-1 && road.r == r.r+1 && road.location == model.TopLeft)
		}
		return (road.q == r.q && road.r == r.r && road.location == model.MiddleLeft) ||
			(road.q == r.q-1 && road.r == r.r+1 && road.location == model.TopLeft) ||
			(road.q == r.q && road.r == r.r+1 && road.location == model.TopLeft) ||
			(road.q == r.q && road.r == r.r+1 && road.location == model.MiddleLeft)
	})
}

func (r Road) ToModel() *model.Road {
	roadModel := new(model.Road)
	roadModel.ID = r.id
	roadModel.Q = r.q
	roadModel.R = r.r
	roadModel.Location = r.location
	roadModel.PlayerID = r.playerID
	return roadModel
}
