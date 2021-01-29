package datamodel

type DevelopmentCards []*DevelopmentCard

func (dc *DevelopmentCards) append(developmentCard *DevelopmentCard) {
	*dc = append(*dc, developmentCard)
}

type DevelopmentCardFilterFunc func(developmentCard *DevelopmentCard) bool

func (dc DevelopmentCards) Filter(developmentCardFilterFunc DevelopmentCardFilterFunc) DevelopmentCards {
	var developmentCards DevelopmentCards
	for _, developmentCard := range dc {
		if developmentCardFilterFunc(developmentCard) {
			developmentCards.append(developmentCard)
		}
	}
	return developmentCards
}
