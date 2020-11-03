package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

func NewDevelopmentCard(game *Game, developmentType datamodel.DevelopmentType) *DevelopmentCard {
	developmentCard := new(DevelopmentCard)
	developmentCard.Type = developmentType
	developmentCard.SetGame(game)
	developmentCard.SetPlayer(nil)
	return developmentCard
}

type DevelopmentCard struct {
	*datamodel.DevelopmentCard
	game   *Game
	player *Player
}

func (dc *DevelopmentCard) SetGame(game *Game) {
	dc.GameID = game.ID
	dc.game = game

	if dc.player == nil {
		dc.game.developmentCards.append(dc)
	}
}

func (dc *DevelopmentCard) SetPlayer(player *Player) {
	if dc.player != nil {
		dc.player.developmentCards.remove(dc)
	}
	if dc.game != nil {
		dc.game.developmentCards.remove(dc)
	}

	if player != nil {
		dc.PlayerID = player.ID
		dc.player = player
		player.developmentCards.append(dc)
	} else {
		dc.PlayerID = 0
		dc.player = nil
		dc.game.developmentCards.append(dc)
	}
}
