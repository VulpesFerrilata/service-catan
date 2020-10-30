package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

func NewPlayer() *Player {
	player := new(Player)
	player.Player = new(datamodel.Player)
	return player
}

type Player struct {
	*datamodel.Player
	game      *Game
	user      *User
	IsDeleted bool
}

func (p *Player) setGame(game *Game) {
	p.GameID = game.ID
	p.game = game
}

func (p *Player) SetUser(user *User) {
	p.UserID = user.ID
	p.user = user
}

func (p *Player) GetUser() *User {
	return p.user
}

func (p Player) IsHost() bool {
	return p.game.Host == p.ID
}
