package datamodel

import (
	"math/rand"

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

func (h Harbors) Shuffle() {
	rand.Shuffle(len(h), func(i, j int) { h[i].harborType, h[j].harborType = h[j].harborType, h[i].harborType })
}
