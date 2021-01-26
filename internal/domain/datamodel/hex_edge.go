package datamodel

func NewHexEdge(q int, r int, location HexEdgeLocation) *HexEdge {
	return &HexEdge{
		q:        q,
		r:        r,
		location: location,
	}
}

func NewHexEdgeFromVector(hex hexType, hexEdgeVector *HexEdgeVector) *HexEdge {
	q := hex.GetQ() + hexEdgeVector.Q
	r := hex.GetR() + hexEdgeVector.R
	location := hexEdgeVector.Location
	return NewHexEdge(q, r, location)
}

type HexEdge struct {
	q        int
	r        int
	location HexEdgeLocation
}

func (h HexEdge) GetQ() int {
	return h.q
}

func (h HexEdge) GetR() int {
	return h.r
}

func (h HexEdge) GetLocation() HexEdgeLocation {
	return h.location
}

func (h HexEdge) GetPossibleAdjacentHexCorners() []*HexCorner {
	hexCorners := make([]*HexCorner, 0)

	hexCornerVectorsMap := map[HexEdgeLocation][]*HexCornerVector{
		TopLeft: {
			{Q: 0, R: -1, Location: Bottom},
			{Q: 0, R: 0, Location: Top},
		},
		MiddleLeft: {
			{Q: 0, R: -1, Location: Bottom},
			{Q: -1, R: 1, Location: Top},
		},
		BottomLeft: {
			{Q: -1, R: 1, Location: Top},
			{Q: 0, R: 0, Location: Bottom},
		},
	}

	hexCornerVectors := hexCornerVectorsMap[h.GetLocation()]
	for _, hexCornerVector := range hexCornerVectors {
		hexCorner := NewHexCornerFromVector(&h, hexCornerVector)
		hexCorners = append(hexCorners, hexCorner)
	}

	return hexCorners
}

func (h HexEdge) GetPossibleAdjacentHexEdges() []*HexEdge {
	hexEdges := make([]*HexEdge, 0)

	hexEdgeVectorsMap := map[HexEdgeLocation][]*HexEdgeVector{
		TopLeft: {
			{Q: 0, R: -1, Location: BottomLeft},
			{Q: 1, R: -1, Location: MiddleLeft},
			{Q: 1, R: -1, Location: BottomLeft},
			{Q: 0, R: 0, Location: MiddleLeft},
		},
		MiddleLeft: {
			{Q: 0, R: -1, Location: BottomLeft},
			{Q: 0, R: 0, Location: TopLeft},
			{Q: 0, R: 0, Location: BottomLeft},
			{Q: -1, R: 1, Location: TopLeft},
		},
		BottomLeft: {
			{Q: 0, R: 0, Location: MiddleLeft},
			{Q: -1, R: 1, Location: TopLeft},
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

func (h HexEdge) Equals(hexEdge *HexEdge) bool {
	if h.GetQ() != hexEdge.GetQ() {
		return false
	}
	if h.GetR() != hexEdge.GetR() {
		return false
	}
	if h.GetLocation() != hexEdge.GetLocation() {
		return false
	}
	return true
}
