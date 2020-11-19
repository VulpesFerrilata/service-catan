package cache

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

type Construction struct {
	ID               int                            `json:"id"`
	Q                int                            `json:"q"`
	R                int                            `json:"r"`
	Location         datamodel.ConstructionLocation `json:"location"`
	IsUpgradedCastle bool                           `json:"isUpgradedCastle"`
}
