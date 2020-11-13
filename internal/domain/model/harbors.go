package model

import (
	"math/rand"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
)

func NewHarbors(game *Game) Harbors {
	var harbors Harbors

	firstHarbor := NewHarbor(game)
	firstHarbor.harbor.Q = 1
	firstHarbor.harbor.R = 0
	firstHarbor.harbor.FieldQ = 1
	firstHarbor.harbor.FieldR = 1
	harbors.append(firstHarbor)

	secondHarbor := NewHarbor(game)
	secondHarbor.harbor.Q = 3
	secondHarbor.harbor.R = 0
	secondHarbor.harbor.FieldQ = 2
	secondHarbor.harbor.FieldR = 1
	harbors.append(secondHarbor)

	thirdHarbor := NewHarbor(game)
	thirdHarbor.harbor.Q = 4
	thirdHarbor.harbor.R = 1
	thirdHarbor.harbor.FieldQ = 3
	thirdHarbor.harbor.FieldR = 2
	harbors.append(thirdHarbor)

	fourthHarbor := NewHarbor(game)
	fourthHarbor.harbor.Q = -1
	fourthHarbor.harbor.R = 2
	fourthHarbor.harbor.FieldQ = 0
	fourthHarbor.harbor.FieldR = 2
	harbors.append(fourthHarbor)

	fifthHarbor := NewHarbor(game)
	fifthHarbor.harbor.Q = 4
	fifthHarbor.harbor.R = 3
	fifthHarbor.harbor.FieldQ = 3
	fifthHarbor.harbor.FieldR = 3
	harbors.append(fifthHarbor)

	sixthHarbor := NewHarbor(game)
	sixthHarbor.harbor.Q = -2
	sixthHarbor.harbor.R = 4
	sixthHarbor.harbor.FieldQ = -1
	sixthHarbor.harbor.FieldR = 4
	harbors.append(sixthHarbor)

	seventhHarbor := NewHarbor(game)
	seventhHarbor.harbor.Q = 2
	seventhHarbor.harbor.R = 5
	seventhHarbor.harbor.FieldQ = 2
	seventhHarbor.harbor.FieldR = 4
	harbors.append(seventhHarbor)

	eighthHarbor := NewHarbor(game)
	eighthHarbor.harbor.Q = -2
	eighthHarbor.harbor.R = 6
	eighthHarbor.harbor.FieldQ = -1
	eighthHarbor.harbor.FieldR = 5
	harbors.append(eighthHarbor)

	ninthHarbor := NewHarbor(game)
	ninthHarbor.harbor.Q = 0
	ninthHarbor.harbor.R = 6
	ninthHarbor.harbor.FieldQ = 0
	ninthHarbor.harbor.FieldR = 5
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
			harbors[harborIdx].harbor.Type = harborType
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
