package datamodel

func NewHexEdge(hex *Hex, location HexEdgeLocation) *HexEdge {
	return &HexEdge{
		hex:      hex,
		location: location,
	}
}

type HexEdge struct {
	hex      *Hex
	location HexEdgeLocation
}

func (h HexEdge) GetHex() *Hex {
	return h.hex
}

func (h HexEdge) GetLocation() HexEdgeLocation {
	return h.location
}

func (h HexEdge) GetPossibleAdjacentHexCorners() []*HexCorner {
	hexCorners := make([]*HexCorner, 0)

	hexCornerVectorsMap := map[HexEdgeLocation][]*HexCorner{
		TopLeft: []*HexCorner{
			NewHexCorner(NewHex(0, -1), Bottom),
			NewHexCorner(NewHex(0, 0), Top),
		},
		MiddleLeft: []*HexCorner{
			NewHexCorner(NewHex(0, -1), Bottom),
			NewHexCorner(NewHex(-1, 1), Top),
		},
		BottomLeft: []*HexCorner{
			NewHexCorner(NewHex(-1, 1), Top),
			NewHexCorner(NewHex(0, 0), Bottom),
		},
	}

	hexCornerVectors := hexCornerVectorsMap[h.GetLocation()]
	for _, hexCornerVector := range hexCornerVectors {
		hexCorner := NewHexCorner(NewHex(h.GetHex().GetQ()+hexCornerVector.GetHex().GetQ(), h.GetHex().GetR()+hexCornerVector.GetHex().GetR()), hexCornerVector.GetLocation())
		hexCorners = append(hexCorners, hexCorner)
	}

	return hexCorners
}

func (h HexEdge) GetPossibleAdjacentHexEdges() []*HexEdge {
	hexEdges := make([]*HexEdge, 0)

	hexEdgeVectorsMap := map[HexEdgeLocation][]*HexEdge{
		TopLeft: []*HexEdge{
			NewHexEdge(NewHex(0, -1), BottomLeft),
			NewHexEdge(NewHex(1, -1), MiddleLeft),
			NewHexEdge(NewHex(1, -1), BottomLeft),
			NewHexEdge(NewHex(0, 0), MiddleLeft),
		},
		MiddleLeft: []*HexEdge{
			NewHexEdge(NewHex(0, -1), BottomLeft),
			NewHexEdge(NewHex(0, 0), TopLeft),
			NewHexEdge(NewHex(0, 0), BottomLeft),
			NewHexEdge(NewHex(-1, 1), TopLeft),
		},
		BottomLeft: []*HexEdge{
			NewHexEdge(NewHex(0, 0), MiddleLeft),
			NewHexEdge(NewHex(-1, 1), TopLeft),
			NewHexEdge(NewHex(0, 1), TopLeft),
			NewHexEdge(NewHex(0, 1), MiddleLeft),
		},
	}

	hexEdgeVectors := hexEdgeVectorsMap[h.GetLocation()]
	for _, hexEdgeVector := range hexEdgeVectors {
		hexEdge := NewHexEdge(NewHex(h.GetHex().GetQ()+hexEdgeVector.GetHex().GetQ(), h.GetHex().GetR()), hexEdgeVector.GetLocation())
		hexEdges = append(hexEdges, hexEdge)
	}

	return hexEdges
}

func (h HexEdge) Equals(hexEdge *HexEdge) bool {
	if !h.GetHex().Equals(hexEdge.GetHex()) {
		return false
	}
	if h.GetLocation() != hexEdge.GetLocation() {
		return false
	}
	return true
}
