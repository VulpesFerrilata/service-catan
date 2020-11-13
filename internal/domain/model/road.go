package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

func NewRoad(game *Game) *Road {
	road := new(Road)
	road.road = new(datamodel.Road)
	road.SetGame(game)
	return road
}

type Road struct {
	road       *datamodel.Road
	game       *Game
	isModified bool
}

func (r *Road) GetRoad() datamodel.Road {
	return *r.road
}

func (r *Road) GetId() uint {
	return r.road.ID
}

func (r *Road) GetGameId() uint {
	return r.road.GameID
}

func (r *Road) setGameId(gameId uint) {
	if r.road.GameID != gameId {
		r.road.GameID = gameId
		r.isModified = true
	}
}

func (r *Road) GetPlayerId() *uint {
	return r.road.PlayerID
}

func (r *Road) SetPlayerId(playerId *uint) {
	if r.road.PlayerID != playerId {
		r.road.PlayerID = playerId
		r.isModified = true
	}
}

func (r *Road) GetQ() int {
	return r.road.Q
}

func (r *Road) GetR() int {
	return r.road.R
}

func (r *Road) GetLocation() datamodel.RoadLocation {
	return r.road.Location
}

func (r *Road) IsModified() bool {
	return r.isModified
}

func (r *Road) SetGame(game *Game) {
	r.game = game
	game.roads.append(r)
}

func (r *Road) GetPlayer(game *Game) *Player {
	if r.GetPlayerId() == nil {
		return nil
	}

	return r.game.players.Filter(func(player *Player) bool {
		return player.GetId() == *r.GetPlayerId()
	}).First()
}

func (r *Road) GetAdjacentConstructions() Constructions {
	return r.game.constructions.Filter(func(construction *Construction) bool {
		if r.GetLocation() == datamodel.RL_TOP_LEFT {
			return (construction.GetQ() == r.GetQ() && construction.GetR() == r.GetR() && construction.GetLocation() == datamodel.CL_TOP) ||
				(construction.GetQ() == r.GetQ() && construction.GetR() == r.GetR()-1 && construction.GetLocation() == datamodel.CL_BOT)
		} else if r.GetLocation() == datamodel.RL_MID_LEFT {
			return (construction.GetQ() == r.GetQ() && construction.GetR() == r.GetR()-1 && construction.GetLocation() == datamodel.CL_BOT) ||
				(construction.GetQ() == r.GetQ()-1 && construction.GetR() == r.GetR()+1 && construction.GetLocation() == datamodel.CL_TOP)
		}
		return (construction.GetQ() == r.GetQ() && construction.GetR() == r.GetR() && construction.GetLocation() == datamodel.CL_BOT) ||
			(construction.GetQ() == r.GetQ()-1 && construction.GetR() == r.GetR()+1 && construction.GetLocation() == datamodel.CL_TOP)
	})
}

func (r *Road) GetAdjacentRoads() Roads {
	return r.game.roads.Filter(func(road *Road) bool {
		if r.GetLocation() == datamodel.RL_TOP_LEFT {
			return (road.GetQ() == r.GetQ()+1 && road.GetR() == r.GetR()-1 && road.GetLocation() == datamodel.RL_MID_LEFT) ||
				(road.GetQ() == r.GetQ()+1 && road.GetR() == r.GetR()-1 && road.GetLocation() == datamodel.RL_BOT_LEFT) ||
				(road.GetQ() == r.GetQ() && road.GetR() == r.GetR()-1 && road.GetLocation() == datamodel.RL_BOT_LEFT) ||
				(road.GetQ() == r.GetQ() && road.GetR() == r.GetR() && road.GetLocation() == datamodel.RL_MID_LEFT)
		} else if r.GetLocation() == datamodel.RL_MID_LEFT {
			return (road.GetQ() == r.GetQ() && road.GetR() == r.GetR() && road.GetLocation() == datamodel.RL_TOP_LEFT) ||
				(road.GetQ() == r.GetQ() && road.GetR() == r.GetR() && road.GetLocation() == datamodel.RL_BOT_LEFT) ||
				(road.GetQ() == r.GetQ() && road.GetR() == r.GetR()-1 && road.GetLocation() == datamodel.RL_BOT_LEFT) ||
				(road.GetQ() == r.GetQ()-1 && road.GetR() == r.GetR()+1 && road.GetLocation() == datamodel.RL_TOP_LEFT)
		}
		return (road.GetQ() == r.GetQ() && road.GetR() == r.GetR() && road.GetLocation() == datamodel.RL_MID_LEFT) ||
			(road.GetQ() == r.GetQ()-1 && road.GetR() == r.GetR()+1 && road.GetLocation() == datamodel.RL_TOP_LEFT) ||
			(road.GetQ() == r.GetQ() && road.GetR() == r.GetR()+1 && road.GetLocation() == datamodel.RL_TOP_LEFT) ||
			(road.GetQ() == r.GetQ() && road.GetR() == r.GetR()+1 && road.GetLocation() == datamodel.RL_MID_LEFT)
	})
}
