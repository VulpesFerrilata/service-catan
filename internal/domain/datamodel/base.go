package datamodel

type base struct {
	modelState ModelState
}

func (b *base) SetModelState(state ModelState) {
	if b.modelState == Added {
		return
	}
	if b.modelState == Deleted {
		return
	}
	b.modelState = state
}

func (b base) GetModelState() ModelState {
	return b.modelState
}

type ModelState int

const (
	Unchanged ModelState = iota
	Added
	Modified
	Deleted
)
