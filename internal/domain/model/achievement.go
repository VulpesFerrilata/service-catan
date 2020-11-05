package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

func NewLongestRoadAchievement(game *Game) *Achievement {
	achievement := new(Achievement)
	achievement.Type = datamodel.AT_LONGEST_ROAD
	achievement.BonusPoints = 2
	achievement.SetGame(game)
	return achievement
}

func NewLargestArmyAchievement(game *Game) *Achievement {
	achievement := new(Achievement)
	achievement.Type = datamodel.AT_LARGEST_ARMY
	achievement.BonusPoints = 2
	achievement.SetGame(game)
	return achievement
}

type Achievement struct {
	*datamodel.Achievement
	game *Game
}

func (a *Achievement) SetGame(game *Game) {
	a.GameID = game.ID
	a.game = game
	a.game.achievements.append(a)
}

func (a *Achievement) GetPlayer() *Player {
	return a.game.players.Filter(func(player *Player) bool {
		return player.ID == *a.PlayerID
	}).First()
}
