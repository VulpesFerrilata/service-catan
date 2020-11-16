package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

func NewLongestRoadAchievement() *Achievement {
	achievement := new(Achievement)
	achievement.achievement = new(datamodel.Achievement)
	achievement.achievement.Type = datamodel.AT_LONGEST_ROAD
	achievement.achievement.BonusPoints = 2
	return achievement
}

func NewLargestArmyAchievement() *Achievement {
	achievement := new(Achievement)
	achievement.achievement = new(datamodel.Achievement)
	achievement.achievement.Type = datamodel.AT_LARGEST_ARMY
	achievement.achievement.BonusPoints = 2
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

func (a *Achievement) GetGameId() *uint {
	return a.achievement.GameID
}

func (a *Achievement) setGame(game *Game) {
	if game != nil {
		a.achievement.GameID = &game.game.ID
		a.game = game
	}
}

func (a *Achievement) GetPlayerId() *uint {
	return a.achievement.PlayerID
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

func (a *Achievement) GetPlayer() *Player {
	if a.GetPlayerId() == nil {
		return nil
	}

	return a.game.players.Filter(func(player *Player) bool {
		return player.GetId() == *a.GetPlayerId()
	}).First()
}

func (a *Achievement) SetPlayer(player *Player) {
	if player == nil {
		a.achievement.PlayerID = nil
	} else {
		a.achievement.PlayerID = &player.player.ID
	}
	a.isModified = true
}
