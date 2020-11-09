package model

import (
	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/pkg/math"
)

func NewHarbor(game *Game) *Harbor {
	harbor := new(Harbor)
	harbor.Harbor = new(datamodel.Harbor)
	harbor.SetGame(game)
	return harbor
}

type Harbor struct {
	*datamodel.Harbor
	game *Game
}

func (h *Harbor) SetGame(game *Game) {
	h.game = game
	game.harbors.append(h)
}

func (h *Harbor) GetIntersectRoad() *Road {
	return h.game.roads.Filter(func(road *Road) bool {
		if h.Q == h.FieldQ {
			return road.Q == h.Q && road.R == math.Max(h.R, h.FieldR) && road.Location == datamodel.RL_TOP_LEFT
		} else if h.R == h.FieldR {
			return road.Q == math.Max(h.Q, h.FieldQ) && road.R == h.R && road.Location == datamodel.RL_MID_LEFT
		}
		return road.Q == math.Max(h.Q, h.FieldQ) && road.R == math.Min(h.R, h.FieldR) && road.Location == datamodel.RL_BOT_LEFT
	}).First()
}
