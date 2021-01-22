package datamodel

func NewHex(q int, r int) hex {
	return hex{
		q: q,
		r: r,
	}
}

type hex struct {
	q int
	r int
}

func (h hex) GetQ() int {
	return h.q
}

func (h hex) GetR() int {
	return h.r
}

func (h hex) Equals(hex hex) bool {
	if h.q != hex.q {
		return false
	}
	if h.r != hex.r {
		return false
	}
	return true
}
