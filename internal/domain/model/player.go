package model

import (
	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/pkg/errors"
)

func NewPlayer(game *Game, user *User) *Player {
	player := new(Player)

	game.AddPlayers(player)
	player.SetUser(user)

	return player
}

func EmptyPlayer() *Player {
	player := new(Player)
	player.isModified = false
	player.isRemoved = false
}

type Player struct {
	id         uint
	color      string
	turnOrder  int
	isLeft     bool
	game       *Game
	user       *User
	isModified bool
	isRemoved  bool
}

func (p Player) GetId() uint {
	return p.id
}

func (p Player) GetColor() string {
	return p.color
}

func (p *Player) SetColor(color string) error {
	if p.game.status != datamodel.Waiting {
		//TODO: action is unavailable in this state error
	}
	isDuplicateColor := p.game.players.Any(func(player Player) bool {
		if player.id != p.id && player.color == color {
			return true
		}
		return false
	})
	if isDuplicateColor {
		//TODO: duplicate color error
	}
	p.color = color
	p.isModified = true
	return nil
}

func (p Player) GetTurnOrder() int {
	return p.turnOrder
}

func (p *Player) SetTurnOrder(turnOrder int) error {
	isDuplicateTurnOrder := p.game.players.Any(func(player Player) bool {
		if player.id != p.id && player.turnOrder == turnOrder {
			return true
		}
		return false
	})
	if isDuplicateTurnOrder {
		//TODO: duplicate turn order error
	}
	p.turnOrder = turnOrder
	p.isModified = true
	return nil
}

func (p Player) IsLeft() bool {
	return p.isLeft
}

func (p Player) Leave() {
	p.isLeft = true
	p.isModified = true
}

func (p Player) GetUser() *User {
	return p.user
}

func (p *Player) SetUser(user *User) {
	p.user = user
}

func (p Player) GetAchievements() Achievements {
	return p.game.achievements.Filter(func(achievement *Achievement) bool {
		playerId := achievement.PlayerID
		if playerId == nil {
			return false
		}
		return *playerId == p.id
	})
}

func (p Player) GetDevelopmentCards() DevelopmentCards {
	return p.game.developmentCards.Filter(func(developmentCard *DevelopmentCard) bool {
		playerId := developmentCard.PlayerID
		if playerId == nil {
			return false
		}
		return *playerId == p.id
	})
}

func (p Player) GetResourceCards() ResourceCards {
	return p.game.resourceCards.Filter(func(resourceCard *ResourceCard) bool {
		playerId := resourceCard.PlayerID
		if playerId == nil {
			return false
		}
		return *playerId == p.id
	})
}

func (p Player) GetRoads() Roads {
	return p.game.roads.Filter(func(road *Road) bool {
		playerId := road.PlayerID
		if playerId == nil {
			return false
		}
		return *playerId == p.id
	})
}

func (p Player) GetConstructions() Constructions {
	return p.game.constructions.Filter(func(construction *Construction) bool {
		playerId := construction.PlayerID
		if playerId == nil {
			return false
		}
		return *playerId == p.id
	})
}

func (p Player) IsHost() bool {
	minPlayerId := p.id
	for _, player := range p.game.players {
		if player.id < minPlayerId {
			minPlayerId = player.id
		}
	}
	return p.id == minPlayerId
}

func (p *Player) IsInTurn() bool {
	return p.game.playerInTurn == p.id
}

func (p Player) IsModified() bool {
	return p.isModified
}

func (p Player) IsRemoved() bool {
	return p.isRemoved
}

func (p *Player) Remove() {
	p.isRemoved = true
	if len(p.game.players) == 0 {
		p.game.isRemoved = true
	}
}

func (p *Player) Persist(f func(player *datamodel.Player) error) error {
	player := new(datamodel.Player)
	player.ID = p.id
	if p.game != nil {
		player.GameID = p.game.id
	}
	if p.user != nil {
		player.UserID = p.user.id
	}
	player.Color = p.color
	player.TurnOrder = p.turnOrder

	if err := f(player); err != nil {
		return errors.Wrap(err, "model.Player.Persist")
	}

	p.id = player.ID
	p.color = player.Color
	p.turnOrder = player.TurnOrder
	return nil
}
