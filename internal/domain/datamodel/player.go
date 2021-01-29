package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func NewPlayer(user *User) (*Player, error) {
	player := new(Player)

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.Wrap(err, "datamodel.NewPlayer")
	}
	player.id = id

	player.user = user
	return player, nil
}

func NewPlayerFromModel(playerModel *model.Player, user *User) (*Player, error) {
	player := new(Player)
	player.id = playerModel.ID
	player.color = playerModel.Color
	player.turnOrder = playerModel.TurnOrder
	player.isLeft = playerModel.IsLeft
	player.user = user
	return player, nil
}

type Player struct {
	id               uuid.UUID
	color            string
	turnOrder        int
	isLeft           bool
	game             *Game
	user             *User
	achievements     Achievements
	constructions    Constructions
	developmentCards DevelopmentCards
	resourceCards    ResourceCards
	roads            Roads
}

func (p Player) GetId() uuid.UUID {
	return p.id
}

func (p Player) GetColor() string {
	return p.color
}

func (p *Player) SetColor(color string) error {
	if p.game.status != Waiting {
		//TODO: action is unavailable in this state error
	}
	isDuplicateColor := p.game.players.Any(func(player *Player) bool {
		if player.id != p.id && player.color == color {
			return true
		}
		return false
	})
	if isDuplicateColor {
		//TODO: duplicate color error
	}
	p.color = color
	return nil
}

func (p Player) GetTurnOrder() int {
	return p.turnOrder
}

func (p *Player) SetTurnOrder(turnOrder int) error {
	isDuplicateTurnOrder := p.game.players.Any(func(player *Player) bool {
		if player.id != p.id && player.turnOrder == turnOrder {
			return true
		}
		return false
	})
	if isDuplicateTurnOrder {
		//TODO: duplicate turn order error
	}
	p.turnOrder = turnOrder
	return nil
}

func (p Player) IsLeft() bool {
	return p.isLeft
}

func (p Player) Leave() {
	p.isLeft = true
}

func (p Player) GetUser() *User {
	return p.user
}

func (p Player) GetAchievements() Achievements {
	return p.game.achievements.Filter(func(achievement *Achievement) bool {
		playerId := achievement.playerID
		if playerId == nil {
			return false
		}
		return *playerId == p.id
	})
}

func (p Player) GetDevelopmentCards() DevelopmentCards {
	return p.game.developmentCards.Filter(func(developmentCard *DevelopmentCard) bool {
		playerId := developmentCard.playerID
		if playerId == nil {
			return false
		}
		return *playerId == p.id
	})
}

func (p Player) GetResourceCards() ResourceCards {
	return p.game.resourceCards.Filter(func(resourceCard *ResourceCard) bool {
		playerId := resourceCard.playerID
		if playerId == nil {
			return false
		}
		return *playerId == p.id
	})
}

func (p Player) GetRoads() Roads {
	return p.game.roads.Filter(func(road *Road) bool {
		playerId := road.playerID
		if playerId == nil {
			return false
		}
		return *playerId == p.id
	})
}

func (p Player) GetConstructions() Constructions {
	return p.game.constructions.Filter(func(construction *Construction) bool {
		playerId := construction.playerID
		if playerId == nil {
			return false
		}
		return *playerId == p.id
	})
}

func (p *Player) IsInTurn() bool {
	if p.game.playerInTurn == nil {
		return false
	}
	return *p.game.playerInTurn == p.id
}

func (p Player) ToModel() *model.Player {
	playerModel := new(model.Player)
	playerModel.ID = p.id

	if p.game != nil {
		playerModel.GameID = p.game.id
	}

	if p.user != nil {
		playerModel.UserID = p.user.id
	}

	playerModel.Color = p.color
	playerModel.TurnOrder = p.turnOrder
	return playerModel
}
