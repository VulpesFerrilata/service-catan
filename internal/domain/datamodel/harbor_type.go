package datamodel

import "fmt"

func NewHarborType(value string) (HarborType, error) {
	harborType := HarborType(value)

	if _, ok := harborTypes[harborType]; ok {
		return harborType, nil
	}

	return harborType, fmt.Errorf("harbor type is invalid: %s", value)
}

type HarborType string

func (h HarborType) String() string {
	return string(h)
}

const (
	GeneralHarbor HarborType = "General"
	LumberHarbor  HarborType = "Lumber"
	BrickHarbor   HarborType = "Brick"
	WoolHarbor    HarborType = "Wool"
	GrainHarbor   HarborType = "Grain"
	OreHarbor     HarborType = "Ore"
)

var harborTypes = map[HarborType]struct{}{
	GeneralHarbor: {},
	LumberHarbor:  {},
	BrickHarbor:   {},
	WoolHarbor:    {},
	GrainHarbor:   {},
	OreHarbor:     {},
}
