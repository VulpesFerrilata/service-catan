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
	isRemoved bool
}

func (p *Player) setGame(game *Game) {
	p.GameID = game.ID
	p.game = game
}

func (p *Player) SetUser(user *User) {
	p.UserID = user.ID
	p.user = user
}

func (p Player) GetUser() *User {
	return p.user
}

func (p Player) IsRemoved() bool {
	return p.isRemoved
}

func (p *Player) Remove() {
	p.isRemoved = true
}

func (p Player) IsHost() bool {
	minPlayerId := p.ID
	for _, player := range p.game.players {
		if player.ID < minPlayerId {
			minPlayerId = player.ID
		}
	}
	return p.ID == minPlayerId
}
