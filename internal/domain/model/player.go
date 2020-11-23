package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

type Player struct {
	datamodel.Player
	game      *Game
	user      *User
	isRemoved bool
}

func (p *Player) SetGame(game *Game) {
	if game != nil {
		p.GameID = &game.ID
	}
	p.game = game
}

func (p *Player) GetUser() *User {
	return p.user
}

func (p *Player) SetUser(user *User) {
	if user == nil {
		p.UserID = nil
	} else {
		p.UserID = &user.ID
	}
	p.user = user
}

func (p *Player) GetAchievements() Achievements {
	return p.game.achievements.Filter(func(achievement *Achievement) bool {
		playerId := achievement.PlayerID
		if playerId == nil {
			return false
		}
		return *playerId == p.ID
	})
}

func (p *Player) GetDevelopmentCards() DevelopmentCards {
	return p.game.developmentCards.Filter(func(developmentCard *DevelopmentCard) bool {
		playerId := developmentCard.PlayerID
		if playerId == nil {
			return false
		}
		return *playerId == p.ID
	})
}

func (p *Player) GetResourceCards() ResourceCards {
	return p.game.resourceCards.Filter(func(resourceCard *ResourceCard) bool {
		playerId := resourceCard.PlayerID
		if playerId == nil {
			return false
		}
		return *playerId == p.ID
	})
}

func (p *Player) GetRoads() Roads {
	return p.game.roads.Filter(func(road *Road) bool {
		playerId := road.PlayerID
		if playerId == nil {
			return false
		}
		return *playerId == p.ID
	})
}

func (p *Player) GetConstructions() Constructions {
	return p.game.constructions.Filter(func(construction *Construction) bool {
		playerId := construction.PlayerID
		if playerId == nil {
			return false
		}
		return *playerId == p.ID
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

func (p *Player) IsInTurn() bool {
	if p.game.PlayerInTurn == nil {
		return false
	}
	return *p.game.PlayerInTurn == p.ID
}
