package cache

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

type DevelopmentCard struct {
	ID   int                       `json:"id"`
	Type datamodel.DevelopmentType `json:"type"`
}
