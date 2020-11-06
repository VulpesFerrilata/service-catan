package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

func NewRobber(game *Game) *Robber {
	robber := new(Robber)
	robber.Robber = new(datamodel.Robber)
	robber.Status = datamodel.RS_IDLE

	desertField := game.fields.Filter(func(field *Field) bool {
		if field.Type == datamodel.FT_DESERT {
			return true
		}
		return false
	}).First()
	robber.Q = desertField.Q
	robber.R = desertField.R

	return robber
}

type Robber struct {
	*datamodel.Robber
	game *Game
}

func (r *Robber) SetGame(game *Game) {
	r.game = game
	game.robber = r
}
