package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/pkg/errors"
)

func NewGame() *Game {
	game := new(Game)
	game.status = model.Waiting
	return game
}

func NewGameFromGameModel(gameModel *model.Game) *Game {
	game := new(Game)
	game.id = gameModel.ID
	game.playerInTurn = gameModel.PlayerInTurn
	game.turn = gameModel.Turn
	game.status = gameModel.Status
	game.isModified = false
	game.isRemoved = false
	return game
}

type Game struct {
	base
	id               int
	playerInTurn     int
	turn             int
	status           model.GameStatus
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
}

func (g Game) GetId() int {
	return g.id
}

func (g Game) GetPlayerInTurn() *Player {
	return g.players.Filter(func(player *Player) bool {
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
			for {
				idx++
				if idx >= len(players) {
					idx = 0
				}
				if !players[idx].isLeft {
					g.playerInTurn = players[idx].id
					g.isModified = true
					return
				}
			}
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
		g.dices = append(g.dices, dice)
		dice.game = g
	}
}

func (g Game) GetAchievements() Achievements {
	return g.achievements
}

func (g *Game) AddAchievements(achievements ...*Achievement) {
	for _, achievement := range achievements {
		g.achievements = append(g.achievements, achievement)
		achievement.game = g
	}
}

func (g Game) GetResourceCards() ResourceCards {
	return g.resourceCards
}

func (g *Game) AddResourceCards(resourceCards ...*ResourceCard) {
	for _, resourceCard := range resourceCards {
		g.resourceCards = append(g.resourceCards, resourceCard)
		resourceCard.game = g
	}
}

func (g Game) GetDevelopmentCards() DevelopmentCards {
	return g.developmentCards
}

func (g *Game) AddDevelopmentCards(developmentCards ...*DevelopmentCard) {
	for _, developmentCard := range developmentCards {
		g.developmentCards = append(g.developmentCards, developmentCard)
		developmentCard.game = g
	}
}

func (g Game) GetTerrains() Terrains {
	return g.terrains
}

func (g *Game) AddTerrains(terrains ...*Terrain) {
	for _, terrain := range terrains {
		g.terrains = append(g.terrains, terrain)
		terrain.game = g
	}
}

func (g Game) GetRobber() *Robber {
	return g.robber
}

func (g *Game) SetRobber(robber *Robber) {
	g.robber = robber
	robber.game = g
}

func (g Game) GetConstructions() Constructions {
	return g.constructions
}

func (g *Game) AddConstructions(constructions ...*Construction) {
	for _, construction := range constructions {
		g.constructions = append(g.constructions, construction)
		construction.game = g
	}
}

func (g Game) GetRoads() Roads {
	return g.roads
}

func (g *Game) AddRoads(roads ...*Road) {
	for _, road := range roads {
		g.roads = append(g.roads, road)
		road.game = g
	}
}

func (g Game) GetHarbors() Harbors {
	return g.harbors
}

func (g *Game) AddHarbors(harbors ...*Harbor) {
	for _, harbor := range harbors {
		g.harbors = append(g.harbors, harbor)
		harbor.game = g
	}
}

func (g Game) IsModified() bool {
	return g.isModified
}

func (g Game) IsRemoved() bool {
	return g.isRemoved
}

func (g *Game) Persist(f func(gameModel *model.Game) error) error {
	gameModel := new(model.Game)
	gameModel.ID = g.id
	gameModel.PlayerInTurn = g.playerInTurn
	gameModel.Turn = g.turn
	gameModel.Status = g.status

	if err := f(gameModel); err != nil {
		return errors.Wrap(err, "datamodel.Game.Persist")
	}
	g.isModified = false
	g.isRemoved = false

	g.id = gameModel.ID
	g.playerInTurn = gameModel.PlayerInTurn
	g.turn = gameModel.Turn
	g.status = gameModel.Status

	return nil
}
