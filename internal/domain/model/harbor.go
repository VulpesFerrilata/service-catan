package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

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
