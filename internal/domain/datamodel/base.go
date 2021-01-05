package datamodel

type base struct {
	isModified bool
	isRemoved  bool
}

func (b base) IsModified() bool {
	return b.isModified
}

func (b base) IsRemoved() bool {
	return b.isRemoved
}
