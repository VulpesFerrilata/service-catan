package model

import (
	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
)

func NewGame() *Game {
	game := new(Game)
	game.game = new(datamodel.Game)
	game.game.Status = datamodel.GS_WAITING
	game.game.Turn = 1
	return game
}

type Game struct {
	game             *datamodel.Game
	players          Players
	dices            Dices
	achievements     Achievements
	resourceCards    ResourceCards
	developmentCards DevelopmentCards
	terrains         Terrains
	robber           *Robber
	constructions    Constructions
	roads            Roads
	harbors          Harbors
	isRemoved        bool
	isModified       bool
}

func (g *Game) GetGame() datamodel.Game {
	return *g.game
}

func (g *Game) GetId() uint {
	return g.game.ID
}

func (g *Game) GetTurn() int {
	return g.game.Turn
}

func (g *Game) NextTurn() {
	g.game.Turn++
	g.isModified = true
}

func (g *Game) GetPlayerInTurn() uint {
	return g.game.PlayerInTurn
}

func (g *Game) SetPlayerInTurn(playerId uint) {
	g.game.PlayerInTurn = playerId
	g.isModified = true
}

func (g *Game) GetStatus() datamodel.GameStatus {
	return g.game.Status
}

func (g *Game) SetStatus(status datamodel.GameStatus) {
	g.game.Status = status
	g.isModified = true
}

func (g *Game) IsModified() bool {
	return g.isModified
}

func (g *Game) GetPlayers() Players {
	return g.players
}

func (g *Game) AddPlayers(players ...*Player) {
	for _, player := range players {
		g.players.append(player)
		player.setGame(g)
	}
}

func (g *Game) GetDices() Dices {
	return g.dices
}

func (g *Game) AddDices(dices ...*Dice) {
	for _, dice := range dices {
		g.dices.append(dice)
		dice.setGame(g)
	}
}

func (g *Game) GetAchievements() Achievements {
	return g.achievements
}

func (g *Game) AddAchievements(achievements ...*Achievement) {
	for _, achievement := range achievements {
		g.achievements.append(achievement)
		achievement.setGame(g)
	}
}

func (g *Game) GetResourceCards() ResourceCards {
	return g.resourceCards
}

func (g *Game) AddResourceCards(resourceCards ...*ResourceCard) {
	for _, resourceCard := range resourceCards {
		g.resourceCards.append(resourceCard)
		resourceCard.setGame(g)
	}
}

func (g *Game) GetDevelopmentCards() DevelopmentCards {
	return g.developmentCards
}

func (g *Game) AddDevelopmentCards(developmentCards ...*DevelopmentCard) {
	for _, developmentCard := range developmentCards {
		g.developmentCards.append(developmentCard)
		developmentCard.setGame(g)
	}
}

func (g *Game) GetTerrains() Terrains {
	return g.terrains
}

func (g *Game) AddTerrains(terrains ...*Terrain) {
	for _, terrain := range terrains {
		g.terrains.append(terrain)
		terrain.setGame(g)
	}
}

func (g *Game) GetRobber() *Robber {
	return g.robber
}

func (g *Game) SetRobber(robber *Robber) {
	g.robber = robber
	robber.setGame(g)
}

func (g *Game) GetConstructions() Constructions {
	return g.constructions
}

func (g *Game) AddConstructions(constructions ...*Construction) {
	for _, construction := range constructions {
		g.constructions.append(construction)
		construction.setGame(g)
	}
}

func (g *Game) GetRoads() Roads {
	return g.roads
}

func (g *Game) AddRoads(roads ...*Road) {
	for _, road := range roads {
		g.roads.append(road)
		road.setGame(g)
	}
}

func (g *Game) GetHarbors() Harbors {
	return g.harbors
}

func (g *Game) AddHarbors(harbors ...*Harbor) {
	for _, harbor := range harbors {
		g.harbors.append(harbor)
		harbor.setGame(g)
	}
}

func (g *Game) IsRemoved() bool {
	return g.isRemoved
}
