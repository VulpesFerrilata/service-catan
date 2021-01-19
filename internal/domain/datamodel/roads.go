package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/model"
)

func NewRoadsFromRoadModels(roadModels []*model.Road) Roads {
	roads := make(Roads, 0)

	for _, roadModel := range roadModels {
		road := NewRoadFromRoadModel(roadModel)
		roads = append(roads, road)
	}

	return roads
}

type Roads []*Road

type RoadFilterFunc func(road *Road) bool

func (r Roads) Filter(roadFilterFunc RoadFilterFunc) Roads {
	var roads Roads

	for _, road := range r {
		if roadFilterFunc(road) {
			roads = append(roads, road)
		}
	}

	return roads
}

func (r Roads) First() *Road {
	if len(r) > 0 {
		return r[0]
	}
	return nil
}
