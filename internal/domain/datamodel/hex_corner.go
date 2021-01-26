package datamodel

import "github.com/pkg/errors"

func NewHexCorner(q int, r int, location HexCornerLocation) *HexCorner {
	return &HexCorner{
		q:        q,
		r:        r,
		location: location,
	}
}

func NewHexCornerFromVector(hex hexType, hexCornerVector *HexCornerVector) *HexCorner {
	q := hex.GetQ() + hexCornerVector.Q
	r := hex.GetR() + hexCornerVector.R
	location := hexCornerVector.Location
	return NewHexCorner(q, r, location)
}

type HexCorner struct {
	q        int
	r        int
	location HexCornerLocation
}

func (h HexCorner) GetQ() int {
	return h.q
}

func (h HexCorner) GetR() int {
	return h.r
}

func (h HexCorner) GetLocation() HexCornerLocation {
	return h.location
}

func (h HexCorner) GetPossibleAdjacentHexes() ([]*Hex, error) {
	hexes := make([]*Hex, 0)

	hexVectorsMap := map[HexCornerLocation][]*HexVector{
		Top: {
			{Q: 0, R: -1},
			{Q: 0, R: 0},
			{Q: 1, R: -1},
		},
		Bottom: {
			{Q: -1, R: 1},
			{Q: 0, R: 0},
			{Q: 0, R: 1},
		},
	}

	hexVectors := hexVectorsMap[h.GetLocation()]
	for _, hexVector := range hexVectors {
		hex, err := NewHexFromVector(&h, hexVector)
		if err != nil {
			return nil, errors.Wrap(err, "datamodel.HexCorner.GetPossibleAdjacentHexes")
		}
		hexes = append(hexes, hex)
	}

	return hexes, nil
}

func (h HexCorner) GetPossibleAdjacentHexCorners() []*HexCorner {
	hexCorners := make([]*HexCorner, 0)

	hexCornerVectorsMap := map[HexCornerLocation][]*HexCornerVector{
		Top: {
			{Q: 0, R: -1, Location: Bottom},
			{Q: 1, R: -2, Location: Bottom},
			{Q: 1, R: -1, Location: Bottom},
		},
		Bottom: {
			{Q: -1, R: 1, Location: Top},
			{Q: -1, R: 2, Location: Top},
			{Q: 0, R: 1, Location: Top},
		},
	}

	hexCornerVectors := hexCornerVectorsMap[h.GetLocation()]
	for _, hexCornerVector := range hexCornerVectors {
		hexCorner := NewHexCornerFromVector(&h, hexCornerVector)
		hexCorners = append(hexCorners, hexCorner)
	}

	return hexCorners
}

func (h HexCorner) GetPossibleAdjacentHexEdges() []*HexEdge {
	hexEdges := make([]*HexEdge, 0)

	hexEdgeVectorsMap := map[HexCornerLocation][]*HexEdgeVector{
		Top: {
			{Q: 1, R: -1, Location: MiddleLeft},
			{Q: 1, R: -1, Location: BottomLeft},
			{Q: 0, R: 0, Location: TopLeft},
		},
		Bottom: {
			{Q: 0, R: 0, Location: BottomLeft},
			{Q: 0, R: 1, Location: TopLeft},
			{Q: 0, R: 1, Location: MiddleLeft},
		},
	}

	hexEdgeVectors := hexEdgeVectorsMap[h.GetLocation()]
	for _, hexEdgeVector := range hexEdgeVectors {
		hexEdge := NewHexEdgeFromVector(&h, hexEdgeVector)
		hexEdges = append(hexEdges, hexEdge)
	}

	return hexEdges
}

func (h HexCorner) Equals(hexCorner *HexCorner) bool {
	if h.GetQ() != hexCorner.GetQ() {
		return false
	}
	if h.GetR() != hexCorner.GetR() {
		return false
	}
	if h.GetLocation() != hexCorner.GetLocation() {
		return false
	}
	return true
}
