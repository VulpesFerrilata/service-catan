package datamodel

import "gorm.io/gorm"

type Field struct {
	gorm.Model
	GameID uint
	Q      int
	R      int
	Number int
	Type   FieldType
}

type FieldType string

const (
	FT_BRICK  FieldType = "BRICK"
	FT_GRAIN  FieldType = "GRAIN"
	FT_SHEEP  FieldType = "SHEEP"
	FT_STONE  FieldType = "STONE"
	FT_WOOD   FieldType = "WOOD"
	FT_DESERT FieldType = "DESERT"
)
