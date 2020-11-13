package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

func NewLongestRoadAchievement(game *Game) *Achievement {
	achievement := new(Achievement)
	achievement.achievement = new(datamodel.Achievement)
	achievement.achievement.Type = datamodel.AT_LONGEST_ROAD
	achievement.achievement.BonusPoints = 2
	achievement.SetGame(game)
	return achievement
}

func NewLargestArmyAchievement(game *Game) *Achievement {
	achievement := new(Achievement)
	achievement.achievement = new(datamodel.Achievement)
	achievement.achievement.Type = datamodel.AT_LARGEST_ARMY
	achievement.achievement.BonusPoints = 2
	achievement.SetGame(game)
	return achievement
}

type Achievement struct {
	achievement *datamodel.Achievement
	game        *Game
	isModified  bool
}

func (a *Achievement) GetAchievement() datamodel.Achievement {
	return *a.achievement
}

func (a *Achievement) GetId() uint {
	return a.achievement.ID
}

func (a *Achievement) GetGameId() uint {
	return a.achievement.GameID
}

func (a *Achievement) setGameId(gameId uint) {
	if a.achievement.GameID != gameId {
		a.achievement.GameID = gameId
		a.isModified = true
	}
}

func (a *Achievement) GetPlayerId() *uint {
	return a.achievement.PlayerID
}

func (a *Achievement) SetPlayerId(playerId *uint) {
	if a.achievement.PlayerID != playerId {
		a.achievement.PlayerID = playerId
		a.isModified = true
	}
}

func (a *Achievement) GetAchievementType() datamodel.AchievementType {
	return a.achievement.Type
}

func (a *Achievement) GetBonusPoints() int {
	return a.achievement.BonusPoints
}

func (a *Achievement) IsModified() bool {
	return a.isModified
}

func (a *Achievement) SetGame(game *Game) {
	a.game = game
	game.achievements.append(a)
}

func (a *Achievement) GetPlayer() *Player {
	if a.GetPlayerId() == nil {
		return nil
	}

	return a.game.players.Filter(func(player *Player) bool {
		return player.GetId() == *a.GetPlayerId()
	}).First()
}
