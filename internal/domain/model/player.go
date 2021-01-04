package model

import "github.com/VulpesFerrilata/library/pkg/model"

type Player struct {
	model.Model
	GameID    int `gorm:"primaryKey"`
	UserID    int `gorm:"primaryKey"`
	Color     string
	TurnOrder int
	IsLeft    bool
}
