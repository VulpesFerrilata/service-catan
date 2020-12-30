package model

import (
	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
)

func NewGame() *Game {
	game := new(Game)
	game.status = datamodel.Waiting
	return game
}

type Game struct {
	id               uint
	playerInTurn     uint
	turn             int
	status           datamodel.GameStatus
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
	isModified       bool
	isRemoved        bool
}

func (g Game) GetId() uint {
	return g.id
}

func (g Game) GetPlayerInTurn() *Player {
	return g.players.Filter(func(player Player) bool {
		return player.id == g.playerInTurn
	}).First()
}

func (g *Game) NextPlayerInTurn() {
	players := g.players

	players.Sort(func(a, b Player) bool {
		return a.turnOrder < b.turnOrder
	})

	for idx, player := range players {
		if player.id == g.playerInTurn {
			if idx+1 == len(players) {
				g.playerInTurn = players[0].id
			} else {
				g.playerInTurn = players[idx+1].id
			}
			g.isModified = true
			return
		}
	}
}

func (g Game) GetTurn() int {
	return g.turn
}

func (g *Game) NextTurn() {
	g.turn++
	g.isModified = true
}

func (g Game) GetPlayers() Players {
	return g.players
}

func (g *Game) AddPlayers(players ...*Player) {
	for _, player := range players {
		g.players = append(g.players, player)
		player.game = g
	}
}

func (g Game) GetDices() Dices {
	return g.dices
}

func (g *Game) AddDices(dices ...*Dice) {
	for _, dice := range dices {
		g.dices.append(dice)
		dice.SetGame(g)
	}
}

func (g Game) GetAchievements() Achievements {
	return g.achievements
}

func (g *Game) AddAchievements(achievements ...*Achievement) {
	for _, achievement := range achievements {
		g.achievements.append(achievement)
		achievement.SetGame(g)
	}
}

func (g Game) GetResourceCards() ResourceCards {
	return g.resourceCards
}

func (g *Game) AddResourceCards(resourceCards ...*ResourceCard) {
	for _, resourceCard := range resourceCards {
		g.resourceCards.append(resourceCard)
		resourceCard.SetGame(g)
	}
}

func (g Game) GetDevelopmentCards() DevelopmentCards {
	return g.developmentCards
}

func (g *Game) AddDevelopmentCards(developmentCards ...*DevelopmentCard) {
	for _, developmentCard := range developmentCards {
		g.developmentCards.append(developmentCard)
		developmentCard.SetGame(g)
	}
}

func (g Game) GetTerrains() Terrains {
	return g.terrains
}

func (g *Game) AddTerrains(terrains ...*Terrain) {
	for _, terrain := range terrains {
		g.terrains.append(terrain)
		terrain.SetGame(g)
	}
}

func (g Game) GetRobber() *Robber {
	return g.robber
}

func (g *Game) SetRobber(robber *Robber) {
	g.robber = robber
	robber.SetGame(g)
}

func (g Game) GetConstructions() Constructions {
	return g.constructions
}

func (g *Game) AddConstructions(constructions ...*Construction) {
	for _, construction := range constructions {
		g.constructions.append(construction)
		construction.SetGame(g)
	}
}

func (g Game) GetRoads() Roads {
	return g.roads
}

func (g *Game) AddRoads(roads ...*Road) {
	for _, road := range roads {
		g.roads.append(road)
		road.SetGame(g)
	}
}

func (g Game) GetHarbors() Harbors {
	return g.harbors
}

func (g *Game) AddHarbors(harbors ...*Harbor) {
	for _, harbor := range harbors {
		g.harbors.append(harbor)
		harbor.SetGame(g)
	}
}

func (g Game) IsModified() bool {
	return g.isModified
}

func (g Game) IsRemoved() bool {
	return g.isRemoved
}
