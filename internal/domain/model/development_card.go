package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

func NewDevelopmentCard(game *Game, developmentType datamodel.DevelopmentType) *DevelopmentCard {
	developmentCard := new(DevelopmentCard)
	developmentCard.DevelopmentCard = new(datamodel.DevelopmentCard)
	developmentCard.Type = developmentType
	developmentCard.SetGame(game)
	return developmentCard
}

type DevelopmentCard struct {
	*datamodel.DevelopmentCard
	game *Game
}

func (dc *DevelopmentCard) SetGame(game *Game) {
	dc.game = game
	game.developmentCards.append(dc)
}

func (dc *DevelopmentCard) GetPlayer() *Player {
	if dc.PlayerID == nil {
		return nil
	}

	return dc.game.players.Filter(func(player *Player) bool {
		return player.ID == *dc.PlayerID
	}).First()
}
