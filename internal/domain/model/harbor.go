package model

import (
	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/pkg/math"
)

func NewHarbor(game *Game) *Harbor {
	harbor := new(Harbor)
	harbor.harbor = new(datamodel.Harbor)
	harbor.SetGame(game)
	return harbor
}

type Harbor struct {
	harbor     *datamodel.Harbor
	game       *Game
	isModified bool
}

func (h *Harbor) GetHarbor() datamodel.Harbor {
	return *h.harbor
}

func (h *Harbor) GetId() uint {
	return h.harbor.ID
}

func (h *Harbor) GetGameId() uint {
	return h.harbor.GameID
}

func (h *Harbor) setGameId(gameId uint) {
	if h.harbor.GameID != gameId {
		h.harbor.GameID = gameId
		h.isModified = true
	}
}

func (h *Harbor) GetQ() int {
	return h.harbor.Q
}

func (h *Harbor) GetR() int {
	return h.harbor.R
}

func (h *Harbor) GetType() datamodel.HarborType {
	return h.harbor.Type
}

func (h *Harbor) GetFieldQ() int {
	return h.harbor.FieldQ
}

func (h *Harbor) GetFieldR() int {
	return h.harbor.FieldR
}

func (h *Harbor) IsModified() bool {
	return h.isModified
}

func (h *Harbor) SetGame(game *Game) {
	h.game = game
	game.harbors.append(h)
}

func (h *Harbor) GetIntersectRoad() *Road {
	return h.game.roads.Filter(func(road *Road) bool {
		if h.GetQ() == h.GetFieldQ() {
			return road.GetQ() == h.GetQ() && road.GetR() == math.Max(h.GetR(), h.GetFieldR()) && road.GetLocation() == datamodel.RL_TOP_LEFT
		} else if h.GetR() == h.GetFieldR() {
			return road.GetQ() == math.Max(h.GetQ(), h.GetFieldQ()) && road.GetR() == h.GetR() && road.GetLocation() == datamodel.RL_MID_LEFT
		}
		return road.GetQ() == math.Max(h.GetQ(), h.GetFieldQ()) && road.GetR() == math.Min(h.GetR(), h.GetFieldR()) && road.GetLocation() == datamodel.RL_BOT_LEFT
	}).First()
}
