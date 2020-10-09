package dto

import "github.com/VulpesFerrilata/catan/internal/domain/model"

type ConstructionDTO struct {
	ID               int
	Q                int
	R                int
	Location         model.ConstructionLocation
	IsUpgradedCastle bool
}
