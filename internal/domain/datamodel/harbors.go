package datamodel

import (
	"math/rand"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
)

func NewHarborsFromHarborModels(harborModels []*model.Harbor) Harbors {
	harbors := make(Harbors, 0)

	for _, harborModel := range harborModels {
		harbor := NewHarborFromHarborModel(harborModel)
		harbors = append(harbors, harbor)
	}

	return harbors
}

func NewHarbors(terrains Terrains) Harbors {
	var harbors Harbors

	firstHarbor := new(Harbor)
	firstHarbor.Q = 1
	firstHarbor.R = 0
	terrain := terrains.Filter(func(terrain *Terrain) bool {
		return terrain.Q == 1 && terrain.R == 1
	}).First()
	firstHarbor.TerrainID = &terrain.ID
	harbors.append(firstHarbor)

	secondHarbor := new(Harbor)
	secondHarbor.Q = 3
	secondHarbor.R = 0
	terrain = terrains.Filter(func(terrain *Terrain) bool {
		return terrain.Q == 2 && terrain.R == 1
	}).First()
	secondHarbor.TerrainID = &terrain.ID
	harbors.append(secondHarbor)

	thirdHarbor := new(Harbor)
	thirdHarbor.Q = 4
	thirdHarbor.R = 1
	terrain = terrains.Filter(func(terrain *Terrain) bool {
		return terrain.Q == 3 && terrain.R == 2
	}).First()
	thirdHarbor.TerrainID = &terrain.ID
	harbors.append(thirdHarbor)

	fourthHarbor := new(Harbor)
	fourthHarbor.Q = -1
	fourthHarbor.R = 2
	terrain = terrains.Filter(func(terrain *Terrain) bool {
		return terrain.Q == 0 && terrain.R == 2
	}).First()
	fourthHarbor.TerrainID = &terrain.ID
	harbors.append(fourthHarbor)

	fifthHarbor := new(Harbor)
	fifthHarbor.Q = 4
	fifthHarbor.R = 3
	terrain = terrains.Filter(func(terrain *Terrain) bool {
		return terrain.Q == 3 && terrain.R == 3
	}).First()
	fifthHarbor.TerrainID = &terrain.ID
	harbors.append(fifthHarbor)

	sixthHarbor := new(Harbor)
	sixthHarbor.Q = -2
	sixthHarbor.R = 4
	terrain = terrains.Filter(func(terrain *Terrain) bool {
		return terrain.Q == -1 && terrain.R == 4
	}).First()
	sixthHarbor.TerrainID = &terrain.ID
	harbors.append(sixthHarbor)

	seventhHarbor := new(Harbor)
	seventhHarbor.Q = 2
	seventhHarbor.R = 5
	terrain = terrains.Filter(func(terrain *Terrain) bool {
		return terrain.Q == 2 && terrain.R == 4
	}).First()
	seventhHarbor.TerrainID = &terrain.ID
	harbors.append(seventhHarbor)

	eighthHarbor := new(Harbor)
	eighthHarbor.Q = -2
	eighthHarbor.R = 6
	terrain = terrains.Filter(func(terrain *Terrain) bool {
		return terrain.Q == -1 && terrain.R == 5
	}).First()
	eighthHarbor.TerrainID = &terrain.ID
	harbors.append(eighthHarbor)

	ninthHarbor := new(Harbor)
	ninthHarbor.Q = 0
	ninthHarbor.R = 6
	terrain = terrains.Filter(func(terrain *Terrain) bool {
		return terrain.Q == 0 && terrain.R == 5
	}).First()
	ninthHarbor.TerrainID = &terrain.ID
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

type HarborFilterFunc func(harbor *Harbor) bool

func (h Harbors) Filter(harborFilterFunc HarborFilterFunc) Harbors {
	var harbors Harbors

	for _, harbor := range h {
		if harborFilterFunc(harbor) {
			harbors = append(harbors, harbor)
		}
	}

	return harbors
}
