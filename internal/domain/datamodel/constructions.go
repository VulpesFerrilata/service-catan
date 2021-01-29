package datamodel

type Constructions []*Construction

type ConstructionFilterFunc func(construction *Construction) bool

func (c Constructions) Filter(constructionFilterFunc ConstructionFilterFunc) Constructions {
	var constructions Constructions

	for _, construction := range c {
		if constructionFilterFunc(construction) {
			constructions = append(constructions, construction)
		}
	}

	return constructions
}
