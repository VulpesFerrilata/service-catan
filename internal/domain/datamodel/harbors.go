package datamodel

import (
	"math/rand"
)

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
