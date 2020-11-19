package cache

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

type Achievement struct {
	ID         int                       `json:"id"`
	GameID     int                       `json:"gameId"`
	PlayerID   int                       `json:"playerId"`
	Type       datamodel.AchievementType `json:"type"`
	BonusPoint int                       `json:"bonusPoint"`
}
