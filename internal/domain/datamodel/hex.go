package datamodel

func NewHex(q int, r int) *Hex {
	return &Hex{
		q: q,
		r: r,
	}
}

type Hex struct {
	q int
	r int
}

func (h Hex) GetQ() int {
	return h.q
}

func (h Hex) GetR() int {
	return h.r
}

func (h Hex) GetPossibleAdjacentHexes() []*Hex {
	hexes := make([]*Hex, 0)

	hexVectors := []*Hex{
		NewHex(0, -1), //top left
		NewHex(1, -1), //top right
		NewHex(1, 0),  //middle right
		NewHex(0, 1),  //bottom right
		NewHex(-1, 1), //bottom left
		NewHex(-1, 0), //middle left
	}

	for _, hexVector := range hexVectors {
		hex := NewHex(h.GetQ()+hexVector.GetQ(), h.GetR()+hexVector.GetR())
		hexes = append(hexes, hex)
	}

	return hexes
}

func (h Hex) GetPossibleAdjacentHexCorners() []*HexCorner {
	hexCorners := make([]*HexCorner, 0)

	hexCornerVectors := []*HexCorner{
		NewHexCorner(NewHex(0, -1), Bottom), //top left
		NewHexCorner(NewHex(0, 0), Top),     //top
		NewHexCorner(NewHex(1, -1), Bottom), //top right
		NewHexCorner(NewHex(0, 1), Top),     //bottom right
		NewHexCorner(NewHex(0, 0), Bottom),  //bottom
		NewHexCorner(NewHex(-1, 1), Top),    //bottom left
	}

	for _, hexCornerVector := range hexCornerVectors {
		hexCorner := NewHexCorner(NewHex(h.GetQ()+hexCornerVector.GetHex().GetQ(), h.GetR()+hexCornerVector.GetHex().GetR()), hexCornerVector.location)
		hexCorners = append(hexCorners, hexCorner)
	}

	return hexCorners
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
