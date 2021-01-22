package datamodel

import "fmt"

func NewHexEdge(hex hex, location hexEdgeLocation) hexEdge {
	return hexEdge{
		hex:      hex,
		location: location,
	}
}

type hexEdge struct {
	hex
	location hexEdgeLocation
}

func (h hexEdge) GetLocation() hexEdgeLocation {
	return h.location
}

func (h hexEdge) Equals(hexEdge hexEdge) bool {
	if !h.hex.Equals(hexEdge.hex) {
		return false
	}
	if h.location != hexEdge.location {
		return false
	}
	return true
}

func NewHexEdgeLocation(value string) (hexEdgeLocation, error) {
	hexEdgeLocation := hexEdgeLocation(value)
	if hexEdgeLocation == Top {
		return hexEdgeLocation, nil
	}
	if hexEdgeLocation == Bottom {
		return hexEdgeLocation, nil
	}
	return hexEdgeLocation, fmt.Errorf("hex edge location is invalid: %s", value)
}

type hexEdgeLocation string

func (h hexEdgeLocation) String() string {
	return string(h)
}

const (
	Top    hexEdgeLocation = "Top"
	Bottom hexEdgeLocation = "Bottom"
)
