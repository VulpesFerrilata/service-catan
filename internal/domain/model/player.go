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

func (p *Player) GetGameId() *uint {
	return p.player.GameID
}

func (p *Player) setGame(game *Game) {
	if game != nil {
		p.player.GameID = &game.game.ID
		p.game = game
	}
}

func (p *Player) GetUserId() *uint {
	return p.player.UserID
}

func (p *Player) GetColor() string {
	return p.player.Color
}

func (p *Player) SetColor(color string) {
	p.player.Color = color
	p.isModified = true
}

func (p *Player) GetTurnOrder() int {
	return p.player.TurnOrder
}

func (p *Player) SetTurnOrder(turnOrder int) {
	p.player.TurnOrder = turnOrder
	p.isModified = true
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

func (p *Player) GetUser() *User {
	return p.user
}

func (p *Player) SetUser(user *User) {
	if user == nil {
		p.player.UserID = nil
	} else {
		p.player.UserID = &user.ID
	}
	p.user = user
}

func (p *Player) GetAchievements() Achievements {
	return p.game.achievements.Filter(func(achievement *Achievement) bool {
		playerId := achievement.GetPlayerId()
		if playerId == nil {
			return false
		}
		return *playerId == p.GetId()
	})
}

func (p *Player) GetDevelopmentCards() DevelopmentCards {
	return p.game.developmentCards.Filter(func(developmentCard *DevelopmentCard) bool {
		playerId := developmentCard.GetPlayerId()
		if playerId == nil {
			return false
		}
		return *playerId == p.GetId()
	})
}

func (p *Player) GetResourceCards() ResourceCards {
	return p.game.resourceCards.Filter(func(resourceCard *ResourceCard) bool {
		playerId := resourceCard.GetPlayerId()
		if playerId == nil {
			return false
		}
		return *playerId == p.GetId()
	})
}

func (p *Player) GetRoads() Roads {
	return p.game.roads.Filter(func(road *Road) bool {
		playerId := road.GetPlayerId()
		if playerId == nil {
			return false
		}
		return *playerId == p.GetId()
	})
}

func (p *Player) GetConstructions() Constructions {
	return p.game.constructions.Filter(func(construction *Construction) bool {
		playerId := construction.GetPlayerId()
		if playerId == nil {
			return false
		}
		return *playerId == p.GetId()
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
	return p.game.GetPlayerInTurn() == p.GetId()
}
