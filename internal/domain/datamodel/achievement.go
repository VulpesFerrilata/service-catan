package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/pkg/errors"
)

func NewAchievementFromAchievementModel(achievementModel *model.Achievement) *Achievement {
	achievement := new(Achievement)
	achievement.id = achievementModel.ID
	achievement.achievementType = achievementModel.AchievementType
	achievement.playerID = achievementModel.PlayerID
	achievement.bonusPoints = achievementModel.BonusPoints
	achievement.isModified = false
	achievement.isRemoved = false
	return achievement
}

type Achievement struct {
	base
	id              int
	achievementType model.AchievementType
	playerID        *int
	bonusPoints     int
	game            *Game
}

func (a Achievement) GetPlayer() *Player {
	if a.playerID == nil {
		return nil
	}

	return a.game.players.Filter(func(player *Player) bool {
		return player.id == *a.playerID
	}).First()
}

func (a *Achievement) Persist(f func(achievementModel *model.Achievement) error) error {
	achievementModel := new(model.Achievement)
	achievementModel.ID = a.id
	if a.game != nil {
		achievementModel.GameID = a.game.id
	}
	achievementModel.AchievementType = a.achievementType
	achievementModel.PlayerID = a.playerID
	achievementModel.BonusPoints = a.bonusPoints

	if err := f(achievementModel); err != nil {
		return errors.Wrap(err, "datamodel.Achievement.Persist")
	}
	a.isModified = false
	a.isRemoved = false

	a.id = achievementModel.ID
	a.achievementType = achievementModel.AchievementType
	a.playerID = achievementModel.PlayerID
	a.bonusPoints = achievementModel.BonusPoints

	return nil
}
