package datamodel

import "gorm.io/gorm"

type Harbor struct {
	gorm.Model
	GameID *uint
	Q      int
	R      int
	Type   HarborType
	FieldQ int
	FieldR int
}

type HarborType string

const (
	HT_GENERAL HarborType = "GENERAL"
	HT_LUMBER  HarborType = "LUMBER"
	HT_BRICK   HarborType = "BRICK"
	HT_WOOL    HarborType = "WOOL"
	HT_GRAIN   HarborType = "GRAIN"
	HT_ORE     HarborType = "ORE"
)
