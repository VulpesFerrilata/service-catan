package datamodel

import "fmt"

func NewHexEdgeLocation(value string) (HexEdgeLocation, error) {
	hexEdgeLocation := HexEdgeLocation(value)
	if _, ok := hexEdgeLocations[hexEdgeLocation]; ok {
		return hexEdgeLocation, nil
	}
	return hexEdgeLocation, fmt.Errorf("hex edge location is invalid: %s", value)
}

type HexEdgeLocation string

func (h HexEdgeLocation) String() string {
	return string(h)
}

const (
	TopLeft    HexEdgeLocation = "TopLeft"
	MiddleLeft HexEdgeLocation = "MiddleLeft"
	BottomLeft HexEdgeLocation = "BottomLeft"
)

var hexEdgeLocations = map[HexEdgeLocation]struct{}{
	TopLeft:    {},
	MiddleLeft: {},
	BottomLeft: {},
}
