package datamodel

import "github.com/VulpesFerrilata/library/pkg/model"

type Dice struct {
	model.Model
	ID     int `gorm:"primaryKey"`
	GameID int
	Number int
}
