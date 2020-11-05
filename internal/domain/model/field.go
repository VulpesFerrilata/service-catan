package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

func NewField(game *Game) *Field {
	field := new(Field)
	field.Field = new(datamodel.Field)
	field.SetGame(game)
	return field
}

type Field struct {
	*datamodel.Field
	game *Game
}

func (f *Field) SetGame(game *Game) {
	f.GameID = game.ID
	f.game = game
	f.game.fields.append(f)
}

func (f *Field) HasRobber() bool {
	return f.game.robber.Q == f.Q && f.game.robber.R == f.R
}
