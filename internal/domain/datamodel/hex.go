package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func NewHex(q int, r int) (*Hex, error) {
	hex := new(Hex)

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.Wrap(err, "datamodel.NewHex")
	}
	hex.id = id

	hex.q = q
	hex.r = r
	return hex, nil
}

func NewHexFromVector(hex hexType, hexVector *HexVector) (*Hex, error) {
	q := hex.GetQ() + hexVector.Q
	r := hex.GetR() + hexVector.R
	return NewHex(q, r)
}

func NewHexFromModel(hexModel *model.Hex) *Hex {
	hex := new(Hex)
	hex.id = hexModel.ID
	hex.q = hexModel.Q
	hex.r = hexModel.R
	return hex
}

type hexType interface {
	GetQ() int
	GetR() int
}

type Hex struct {
	id uuid.UUID
	q  int
	r  int
}

func (h Hex) GetQ() int {
	return h.q
}

func (h Hex) GetR() int {
	return h.r
}

func (h Hex) GetPossibleAdjacentHexes() ([]*Hex, error) {
	hexes := make([]*Hex, 0)

	hexVectors := []*HexVector{
		{Q: 0, R: -1}, //top left
		{Q: 1, R: -1}, //top right
		{Q: 1, R: 0},  //middle right
		{Q: 0, R: 1},  //bottom right
		{Q: -1, R: 1}, //bottom left
		{Q: -1, R: 0}, //middle left
	}

	for _, hexVector := range hexVectors {
		// hex := NewHex(h.GetQ()+hexVector.GetQ(), h.GetR()+hexVector.GetR())
		hex, err := NewHexFromVector(&h, hexVector)
		if err != nil {
			return nil, errors.Wrap(err, "datamodel.Hex.GetPossibleAdjacentHexes")
		}
		hexes = append(hexes, hex)
	}

	return hexes, nil
}

func (h Hex) GetPossibleAdjacentHexCorners() []*HexCorner {
	hexCorners := make([]*HexCorner, 0)

	hexCornerVectors := []*HexCornerVector{
		{Q: 0, R: -1, Location: Bottom}, //top left
		{Q: 0, R: 0, Location: Top},     //top
		{Q: 1, R: -1, Location: Bottom}, //top right
		{Q: 0, R: 1, Location: Top},     //bottom right
		{Q: 0, R: 0, Location: Bottom},  //bottom
		{Q: -1, R: 1, Location: Top},    //bottom left
	}

	for _, hexCornerVector := range hexCornerVectors {
		hexCorner := NewHexCornerFromVector(&h, hexCornerVector)
		hexCorners = append(hexCorners, hexCorner)
	}

	return hexCorners
}

func (h Hex) GetPossibleAdjacentHexEdges() []*HexEdge {
	hexEdges := make([]*HexEdge, 0)

	hexEdgeVectors := []*HexEdgeVector{
		{Q: 0, R: 0, Location: TopLeft},     //top left
		{Q: 1, R: -1, Location: BottomLeft}, //top right
		{Q: 1, R: 0, Location: MiddleLeft},  //middle right
		{Q: 0, R: 1, Location: TopLeft},     //bottom right
		{Q: 0, R: 0, Location: BottomLeft},  //bottom left
		{Q: 0, R: 0, Location: MiddleLeft},  //middle left
	}

	for _, hexEdgeVector := range hexEdgeVectors {
		hexEdge := NewHexEdgeFromVector(&h, hexEdgeVector)
		hexEdges = append(hexEdges, hexEdge)
	}

	return hexEdges
}

func (h Hex) GetPossibleIntersectHexEdge(hex *Hex) *HexEdge {
	hexVector := &HexVector{
		Q: hex.GetQ() - h.GetQ(),
		R: hex.GetR() - h.GetR(),
	}

	if hexVector.Q == 0 && hexVector.R == -1 {
		return NewHexEdge(0, 0, TopLeft)
	}

	if hexVector.Q == 1 && hexVector.R == -1 {
		return NewHexEdge(1, -1, BottomLeft)
	}

	if hexVector.Q == 1 && hexVector.R == 0 {
		return NewHexEdge(1, 0, MiddleLeft)
	}

	if hexVector.Q == 0 && hexVector.R == 1 {
		return NewHexEdge(0, 1, TopLeft)
	}

	if hexVector.Q == -1 && hexVector.R == 1 {
		return NewHexEdge(0, 0, BottomLeft)
	}

	if hexVector.Q == -1 && hexVector.R == 0 {
		return NewHexEdge(0, 0, MiddleLeft)
	}

	return nil
}

func (h Hex) Equals(hex *Hex) bool {
	if h.GetQ() != hex.GetQ() {
		return false
	}
	if h.GetR() != hex.GetR() {
		return false
	}
	return true
}

func (h Hex) ToModel() *model.Hex {
	hexModel := new(model.Hex)
	hexModel.ID = h.id
	hexModel.Q = h.q
	hexModel.R = h.r
	return hexModel
}
