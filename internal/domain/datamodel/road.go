package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

type Road struct {
	datamodel.Road
	game       *Game
	isModified bool
}

func (r *Road) SetGame(game *Game) {
	if game != nil {
		r.GameID = &game.id
	}
	r.game = game
}

func (r *Road) GetPlayer(game *Game) *Player {
	if r.PlayerID == nil {
		return nil
	}

	return r.game.players.Filter(func(player Player) bool {
		return player.id == *r.PlayerID
	}).First()
}

func (r *Road) GetAdjacentConstructions() Constructions {
	return r.game.constructions.Filter(func(construction *Construction) bool {
		if r.Location == datamodel.RL_TOP_LEFT {
			return (construction.Q == r.Q && construction.R == r.R && construction.Location == datamodel.CL_TOP) ||
				(construction.Q == r.Q && construction.R == r.R-1 && construction.Location == datamodel.CL_BOT)
		} else if r.Location == datamodel.RL_MID_LEFT {
			return (construction.Q == r.Q && construction.R == r.R-1 && construction.Location == datamodel.CL_BOT) ||
				(construction.Q == r.Q-1 && construction.R == r.R+1 && construction.Location == datamodel.CL_TOP)
		}
		return (construction.Q == r.Q && construction.R == r.R && construction.Location == datamodel.CL_BOT) ||
			(construction.Q == r.Q-1 && construction.R == r.R+1 && construction.Location == datamodel.CL_TOP)
	})
}

func (r *Road) GetAdjacentRoads() Roads {
	return r.game.roads.Filter(func(road *Road) bool {
		if r.Location == datamodel.RL_TOP_LEFT {
			return (road.Q == r.Q+1 && road.R == r.R-1 && road.Location == datamodel.RL_MID_LEFT) ||
				(road.Q == r.Q+1 && road.R == r.R-1 && road.Location == datamodel.RL_BOT_LEFT) ||
				(road.Q == r.Q && road.R == r.R-1 && road.Location == datamodel.RL_BOT_LEFT) ||
				(road.Q == r.Q && road.R == r.R && road.Location == datamodel.RL_MID_LEFT)
		} else if r.Location == datamodel.RL_MID_LEFT {
			return (road.Q == r.Q && road.R == r.R && road.Location == datamodel.RL_TOP_LEFT) ||
				(road.Q == r.Q && road.R == r.R && road.Location == datamodel.RL_BOT_LEFT) ||
				(road.Q == r.Q && road.R == r.R-1 && road.Location == datamodel.RL_BOT_LEFT) ||
				(road.Q == r.Q-1 && road.R == r.R+1 && road.Location == datamodel.RL_TOP_LEFT)
		}
		return (road.Q == r.Q && road.R == r.R && road.Location == datamodel.RL_MID_LEFT) ||
			(road.Q == r.Q-1 && road.R == r.R+1 && road.Location == datamodel.RL_TOP_LEFT) ||
			(road.Q == r.Q && road.R == r.R+1 && road.Location == datamodel.RL_TOP_LEFT) ||
			(road.Q == r.Q && road.R == r.R+1 && road.Location == datamodel.RL_MID_LEFT)
	})
}
