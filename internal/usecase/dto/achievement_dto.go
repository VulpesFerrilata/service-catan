package dto

import "github.com/VulpesFerrilata/catan/internal/domain/model"

type AchievementDTO struct {
	ID          int
	Type        model.AchievementType
	BonusPoints int
}
