package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

func NewPlayer(game *Game, user *User) *Player {
	player := new(Player)
	player.Player = new(datamodel.Player)
	player.SetGame(game)
	player.SetUser(user)
	return player
}

type Player struct {
	*datamodel.Player
	game             *Game
	user             *User
	achievements     Achievements
	resourceCards    ResourceCards
	developmentCards DevelopmentCards
	isRemoved        bool
}

func (p *Player) SetGame(game *Game) {
	p.GameID = game.ID
	p.game = game
	game.players.append(p)
}

func (p *Player) SetUser(user *User) {
	p.UserID = user.ID
	p.user = user
}

func (p *Player) GetUser() *User {
	return p.user
}

func (p *Player) GetAchievements() Achievements {
	return p.game.achievements.Filter(func(achievement *Achievement) bool {
		return achievement.PlayerID == &p.ID
	})
}

func (p *Player) GetDevelopmentCards() DevelopmentCards {
	return p.game.developmentCards.Filter(func(developmentCard *DevelopmentCard) bool {
		return developmentCard.PlayerID == &p.ID
	})
}

func (p *Player) GetResourceCards() ResourceCards {
	return p.game.resourceCards.Filter(func(resourceCard *ResourceCard) bool {
		return resourceCard.PlayerID == &p.ID
	})
}

func (p *Player) IsRemoved() bool {
	return p.isRemoved
}

func (p *Player) Remove() {
	p.isRemoved = true
	if len(p.game.players) == 0 {
		p.game.isRemoved = true
	}
}

func (p *Player) IsHost() bool {
	minPlayerId := p.ID
	for _, player := range p.game.players {
		if player.ID < minPlayerId {
			minPlayerId = player.ID
		}
	}
	return p.ID == minPlayerId
}
