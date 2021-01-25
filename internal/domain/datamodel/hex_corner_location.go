package datamodel

import "fmt"

func NewHexCornerLocation(value string) (HexCornerLocation, error) {
	hexCornerLocation := HexCornerLocation(value)
	if _, ok := hexCornerLocations[hexCornerLocation]; ok {
		return hexCornerLocation, nil
	}
	return hexCornerLocation, fmt.Errorf("hex corner location is invalid: %s", value)
}

type HexCornerLocation string

func (h HexCornerLocation) String() string {
	return string(h)
}

const (
	Top    HexCornerLocation = "Top"
	Bottom HexCornerLocation = "Bottom"
)

var hexCornerLocations = map[HexCornerLocation]struct{}{
	Top:    struct{}{},
	Bottom: struct{}{},
}
