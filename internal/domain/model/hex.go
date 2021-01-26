package model

import "github.com/google/uuid"

type Hex struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`
	Q  int       `gorm:"unique:uc_hex_location"`
	R  int       `gorm:"unique:uc_hex_location"`
}
