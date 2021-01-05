package model

import "github.com/VulpesFerrilata/library/pkg/model"

type Player struct {
	model.Model
	ID        int `gorm:"primaryKey"`
	GameID    int
	UserID    int
	Color     string
	TurnOrder int
	IsLeft    bool
}
