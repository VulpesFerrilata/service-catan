package model

import (
	"math/rand"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
)

func NewHarbors(game *Game) Harbors {
	var harbors Harbors

	firstHarbor := NewHarbor(game)
	firstHarbor.Q = 1
	firstHarbor.R = 0
	firstHarbor.FieldQ = 1
	firstHarbor.FieldR = 1
	harbors.append(firstHarbor)

	secondHarbor := NewHarbor(game)
	secondHarbor.Q = 3
	secondHarbor.R = 0
	secondHarbor.FieldQ = 2
	secondHarbor.FieldR = 1
	harbors.append(secondHarbor)

	thirdHarbor := NewHarbor(game)
	thirdHarbor.Q = 4
	thirdHarbor.R = 1
	thirdHarbor.FieldQ = 3
	thirdHarbor.FieldR = 2
	harbors.append(thirdHarbor)

	fourthHarbor := NewHarbor(game)
	fourthHarbor.Q = -1
	fourthHarbor.R = 2
	fourthHarbor.FieldQ = 0
	fourthHarbor.FieldR = 2
	harbors.append(fourthHarbor)

	fifthHarbor := NewHarbor(game)
	fifthHarbor.Q = 4
	fifthHarbor.R = 3
	fifthHarbor.FieldQ = 3
	fifthHarbor.FieldR = 3
	harbors.append(fifthHarbor)

	sixthHarbor := NewHarbor(game)
	sixthHarbor.Q = -2
	sixthHarbor.R = 4
	sixthHarbor.FieldQ = -1
	sixthHarbor.FieldR = 4
	harbors.append(sixthHarbor)

	seventhHarbor := NewHarbor(game)
	seventhHarbor.Q = 2
	seventhHarbor.R = 5
	seventhHarbor.FieldQ = 2
	seventhHarbor.FieldR = 4
	harbors.append(seventhHarbor)

	eighthHarbor := NewHarbor(game)
	eighthHarbor.Q = -2
	eighthHarbor.R = 6
	eighthHarbor.FieldQ = -1
	eighthHarbor.FieldR = 5
	harbors.append(eighthHarbor)

	ninthHarbor := NewHarbor(game)
	ninthHarbor.Q = 0
	ninthHarbor.R = 6
	ninthHarbor.FieldQ = 0
	ninthHarbor.FieldR = 5
	harbors.append(ninthHarbor)

	harborTypes := map[datamodel.HarborType]int{
		datamodel.HT_GENERAL: 4,
		datamodel.HT_LUMBER:  1,
		datamodel.HT_BRICK:   1,
		datamodel.HT_WOOL:    1,
		datamodel.HT_GRAIN:   1,
		datamodel.HT_ORE:     1,
	}
	rand.Shuffle(len(harbors), func(i, j int) { harbors[i], harbors[j] = harbors[j], harbors[i] })
	harborIdx := 0
	for harborType, quantity := range harborTypes {
		for i := 1; i <= quantity; i++ {
			harbors[harborIdx].Type = harborType
			harborIdx++
		}
	}

	return harbors
}

type Harbors []*Harbor

func (h *Harbors) append(harbor *Harbor) {
	*h = append(*h, harbor)
}

func (h *Harbors) SetGame(game *Game) {
	for _, harbor := range *h {
		harbor.SetGame(game)
	}
}

type HarborFilterFunc func(harbor *Harbor) bool

func (h Harbors) Filter(harborFilterFunc HarborFilterFunc) Harbors {
	var harbors Harbors

	for _, harbor := range h {
		if harborFilterFunc(harbor) {
			harbors.append(harbor)
		}
	}

	return harbors
}
