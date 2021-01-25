package datamodel

func NewHexCorner(hex *Hex, location HexCornerLocation) *HexCorner {
	return &HexCorner{
		hex:      hex,
		location: location,
	}
}

type HexCorner struct {
	hex      *Hex
	location HexCornerLocation
}

func (h HexCorner) GetHex() *Hex {
	return h.hex
}

func (h HexCorner) GetLocation() HexCornerLocation {
	return h.location
}

func (h HexCorner) GetPossibleAdjacentHexes() []*Hex {
	hexes := make([]*Hex, 0)

	hexVectorsMap := map[HexCornerLocation][]*Hex{
		Top: []*Hex{
			NewHex(0, -1),
			NewHex(0, 0),
			NewHex(1, -1),
		},
		Bottom: []*Hex{
			NewHex(-1, 1),
			NewHex(0, 0),
			NewHex(0, 1),
		},
	}

	hexVectors := hexVectorsMap[h.GetLocation()]
	for _, hexVector := range hexVectors {
		hex := NewHex(h.GetHex().GetQ()+hexVector.GetQ(), h.GetHex().GetR()+hexVector.GetR())
		hexes = append(hexes, hex)
	}

	return hexes
}

func (h HexCorner) GetPossibleAdjacentHexCorners() []*HexCorner {
	hexCorners := make([]*HexCorner, 0)

	hexCornerVectorsMap := map[HexCornerLocation][]*HexCorner{
		Top: []*HexCorner{
			NewHexCorner(NewHex(0, -1), Bottom),
			NewHexCorner(NewHex(1, -2), Bottom),
			NewHexCorner(NewHex(1, -1), Bottom),
		},
		Bottom: []*HexCorner{
			NewHexCorner(NewHex(-1, 1), Top),
			NewHexCorner(NewHex(-1, 2), Top),
			NewHexCorner(NewHex(0, 1), Top),
		},
	}

	hexCornerVectors := hexCornerVectorsMap[h.GetLocation()]
	for _, hexCornerVector := range hexCornerVectors {
		hexCorner := NewHexCorner(NewHex(h.GetHex().GetQ()+hexCornerVector.GetHex().GetQ(), h.GetHex().GetR()+hexCornerVector.GetHex().GetR()), hexCornerVector.GetLocation())
		hexCorners = append(hexCorners, hexCorner)
	}

	return hexCorners
}

func (h HexCorner) GetPossibleAdjacentHexEdges() []*HexEdge {
	hexEdges := make([]*HexEdge, 0)

	hexEdgeVectorsMap := map[HexCornerLocation][]*HexEdge{
		Top: []*HexEdge{
			NewHexEdge(NewHex(1, -1), MiddleLeft),
			NewHexEdge(NewHex(1, -1), BottomLeft),
			NewHexEdge(NewHex(0, 0), TopLeft),
		},
		Bottom: []*HexEdge{
			NewHexEdge(NewHex(0, 0), BottomLeft),
			NewHexEdge(NewHex(0, 1), TopLeft),
			NewHexEdge(NewHex(0, 1), MiddleLeft),
		},
	}

	hexEdgeVectors := hexEdgeVectorsMap[h.GetLocation()]
	for _, hexEdgeVector := range hexEdgeVectors {
		hexEdge := NewHexEdge(NewHex(h.GetHex().GetQ()+hexEdgeVector.GetHex().GetQ(), h.GetHex().GetR()+hexEdgeVector.GetHex().GetR()), hexEdgeVector.GetLocation())
		hexEdges = append(hexEdges, hexEdge)
	}

	return hexEdges
}

func (h HexCorner) Equals(hexCorner *HexCorner) bool {
	if !h.GetHex().Equals(hexCorner.GetHex()) {
		return false
	}
	if h.GetLocation() != hexCorner.GetLocation() {
		return false
	}
	return true
}
