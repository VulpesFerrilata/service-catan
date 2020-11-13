package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

func NewDevelopmentCard(game *Game, developmentType datamodel.DevelopmentType) *DevelopmentCard {
	developmentCard := new(DevelopmentCard)
	developmentCard.developmentCard = new(datamodel.DevelopmentCard)
	developmentCard.developmentCard.Type = developmentType
	developmentCard.SetGame(game)
	return developmentCard
}

type DevelopmentCard struct {
	developmentCard *datamodel.DevelopmentCard
	game            *Game
	isModified      bool
}

func (dc *DevelopmentCard) GetDevelopmentCard() datamodel.DevelopmentCard {
	return *dc.developmentCard
}

func (dc *DevelopmentCard) GetId() uint {
	return dc.developmentCard.ID
}

func (dc *DevelopmentCard) GetGameId() uint {
	return dc.developmentCard.GameID
}

func (dc *DevelopmentCard) setGameId(gameId uint) {
	if dc.developmentCard.GameID != gameId {
		dc.developmentCard.GameID = gameId
		dc.isModified = true
	}
}

func (dc *DevelopmentCard) GetPlayerId() *uint {
	return dc.developmentCard.PlayerID
}

func (dc *DevelopmentCard) SetPlayerId(playerId *uint) {
	if dc.developmentCard.PlayerID != playerId {
		dc.developmentCard.PlayerID = playerId
		dc.isModified = true
	}
}

func (dc *DevelopmentCard) GetType() datamodel.DevelopmentType {
	return dc.developmentCard.Type
}

func (dc *DevelopmentCard) IsModified() bool {
	return dc.isModified
}

func (dc *DevelopmentCard) SetGame(game *Game) {
	dc.game = game
	game.developmentCards.append(dc)
}

func (dc *DevelopmentCard) GetPlayer() *Player {
	if dc.GetPlayerId() == nil {
		return nil
	}

	return dc.game.players.Filter(func(player *Player) bool {
		return player.GetId() == *dc.GetPlayerId()
	}).First()
}
