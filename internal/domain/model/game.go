package model

import (
	"math/rand"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
)

func NewGame() *Game {
	game := new(Game)
	game.Game = new(datamodel.Game)
	game.Status = datamodel.GS_WAITING
	return game
}

type Game struct {
	*datamodel.Game
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
}

func (g *Game) GetPlayers() Players {
	return g.players
}

func (g *Game) GetDices() Dices {
	return g.dices
}

func (g *Game) GetAchievements() Achievements {
	return g.achievements
}

func (g *Game) GetResourceCards() ResourceCards {
	return g.resourceCards
}

func (g *Game) GetDevelopmentCards() DevelopmentCards {
	return g.developmentCards
}

func (g *Game) GetTerrains() Terrains {
	return g.terrains
}

func (g *Game) GetRobber() *Robber {
	return g.robber
}

func (g *Game) GetConstructions() Constructions {
	return g.constructions
}

func (g *Game) GetRoads() Roads {
	return g.roads
}

func (g *Game) GetHarbors() Harbors {
	return g.harbors
}

func (g *Game) IsRemoved() bool {
	return g.isRemoved
}

func (g *Game) Init() {
	g.Status = datamodel.GS_STARTED

	NewDices(g)
	NewAchievements(g)
	NewResourceCards(g)
	NewDevelopmentCards(g)
	NewFields(g)
	NewRobber(g)
	NewConstructions(g)
	NewRoads(g)
	NewHarbors(g)

	colors := []string{
		"#ff0000",
		"#0000ff",
		"#ffffff",
		"#000000",
	}
	rand.Shuffle(len(colors), func(i, j int) { colors[i], colors[j] = colors[j], colors[i] })

	orders := []int{
		1,
		2,
		3,
		4,
	}
	rand.Shuffle(len(orders), func(i, j int) { orders[i], orders[j] = orders[j], orders[i] })

	for idx, player := range g.players {
		player.TurnOrder = orders[idx]
		player.Color = colors[idx]
	}

	g.Turn = 1
	g.PlayerInTurn = g.players[0].ID
}
