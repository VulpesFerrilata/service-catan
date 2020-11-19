package model

import (
	"math/rand"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
)

func NewHarbors(terrains Terrains) Harbors {
	var harbors Harbors

	firstHarbor := NewHarbor()
	firstHarbor.harbor.Q = 1
	firstHarbor.harbor.R = 0
	terrain := terrains.Filter(func(terrain *Terrain) bool {
		return terrain.GetQ() == 1 && terrain.GetR() == 1
	}).First()
	firstHarbor.SetTerrain(terrain)
	harbors.append(firstHarbor)

	secondHarbor := NewHarbor()
	secondHarbor.harbor.Q = 3
	secondHarbor.harbor.R = 0
	terrain = terrains.Filter(func(terrain *Terrain) bool {
		return terrain.GetQ() == 2 && terrain.GetR() == 1
	}).First()
	firstHarbor.SetTerrain(terrain)
	harbors.append(secondHarbor)

	thirdHarbor := NewHarbor()
	thirdHarbor.harbor.Q = 4
	thirdHarbor.harbor.R = 1
	terrain = terrains.Filter(func(terrain *Terrain) bool {
		return terrain.GetQ() == 3 && terrain.GetR() == 2
	}).First()
	firstHarbor.SetTerrain(terrain)
	harbors.append(thirdHarbor)

	fourthHarbor := NewHarbor()
	fourthHarbor.harbor.Q = -1
	fourthHarbor.harbor.R = 2
	terrain = terrains.Filter(func(terrain *Terrain) bool {
		return terrain.GetQ() == 0 && terrain.GetR() == 2
	}).First()
	firstHarbor.SetTerrain(terrain)
	harbors.append(fourthHarbor)

	fifthHarbor := NewHarbor()
	fifthHarbor.harbor.Q = 4
	fifthHarbor.harbor.R = 3
	terrain = terrains.Filter(func(terrain *Terrain) bool {
		return terrain.GetQ() == 3 && terrain.GetR() == 3
	}).First()
	firstHarbor.SetTerrain(terrain)
	harbors.append(fifthHarbor)

	sixthHarbor := NewHarbor()
	sixthHarbor.harbor.Q = -2
	sixthHarbor.harbor.R = 4
	terrain = terrains.Filter(func(terrain *Terrain) bool {
		return terrain.GetQ() == -1 && terrain.GetR() == 4
	}).First()
	firstHarbor.SetTerrain(terrain)
	harbors.append(sixthHarbor)

	seventhHarbor := NewHarbor()
	seventhHarbor.harbor.Q = 2
	seventhHarbor.harbor.R = 5
	terrain = terrains.Filter(func(terrain *Terrain) bool {
		return terrain.GetQ() == 2 && terrain.GetR() == 4
	}).First()
	firstHarbor.SetTerrain(terrain)
	harbors.append(seventhHarbor)

	eighthHarbor := NewHarbor()
	eighthHarbor.harbor.Q = -2
	eighthHarbor.harbor.R = 6
	terrain = terrains.Filter(func(terrain *Terrain) bool {
		return terrain.GetQ() == -1 && terrain.GetR() == 5
	}).First()
	firstHarbor.SetTerrain(terrain)
	harbors.append(eighthHarbor)

	ninthHarbor := NewHarbor()
	ninthHarbor.harbor.Q = 0
	ninthHarbor.harbor.R = 6
	terrain = terrains.Filter(func(terrain *Terrain) bool {
		return terrain.GetQ() == 0 && terrain.GetR() == 5
	}).First()
	firstHarbor.SetTerrain(terrain)
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
