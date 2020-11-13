package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

func NewPlayer() *Player {
	player := new(Player)
	player.player = new(datamodel.Player)
	return player
}

type Player struct {
	player     *datamodel.Player
	game       *Game
	user       *User
	isModified bool
	isRemoved  bool
}

func (p *Player) GetPlayer() datamodel.Player {
	return *p.player
}

func (p *Player) GetId() uint {
	return p.player.ID
}

func (p *Player) GetGameId() uint {
	return p.player.GameID
}

func (p *Player) setGameId(gameId uint) {
	if p.player.GameID != gameId {
		p.player.GameID = gameId
		p.isModified = true
	}
}

func (p *Player) SetGame(game *Game) {
	p.game = game
	game.players.append(p)
}

func (p *Player) GetUserId() uint {
	return p.player.UserID
}

func (p *Player) GetColor() string {
	return p.player.Color
}

func (p *Player) GetTurnOrder() int {
	return p.player.TurnOrder
}

func (p *Player) IsLeft() bool {
	return p.player.IsLeft
}

func (p *Player) Leave() {
	if !p.player.IsLeft {
		p.player.IsLeft = true
		p.isModified = true
	}
}

func (p *Player) SetUser(user *User) {
	if user != nil && p.player.UserID != user.ID {
		p.player.UserID = user.ID
		p.isModified = true
	}
	p.user = user
}

func (p *Player) GetUser() *User {
	return p.user
}

func (p *Player) GetAchievements() Achievements {
	return p.game.achievements.Filter(func(achievement *Achievement) bool {
		return achievement.GetPlayerId() == &p.player.ID
	})
}

func (p *Player) GetDevelopmentCards() DevelopmentCards {
	return p.game.developmentCards.Filter(func(developmentCard *DevelopmentCard) bool {
		return developmentCard.GetPlayerId() == &p.player.ID
	})
}

func (p *Player) GetResourceCards() ResourceCards {
	return p.game.resourceCards.Filter(func(resourceCard *ResourceCard) bool {
		return resourceCard.resourceCard.PlayerID == &p.player.ID
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
	minPlayerId := p.GetId()
	for _, player := range p.game.players {
		if player.GetId() < minPlayerId {
			minPlayerId = player.GetId()
		}
	}
	return p.player.ID == minPlayerId
}

func (p *Player) IsInTurn() bool {
	return p.game.PlayerInTurn == p.GetId()
}
