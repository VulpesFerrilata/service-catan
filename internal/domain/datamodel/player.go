package datamodel

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	GameID                  uint
	UserID                  uint
	Color                   string
	TurnOrder               int
	IsLeft                  bool
	IsRolledDices           bool
	IsPlayedDevelopmentCard bool
}
