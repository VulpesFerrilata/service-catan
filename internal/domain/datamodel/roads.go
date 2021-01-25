package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/pkg/errors"
)

func NewRoadsFromRoadModels(roadModels []*model.Road) (Roads, error) {
	roads := make(Roads, 0)

	for _, roadModel := range roadModels {
		road, err := NewRoadFromRoadModel(roadModel)
		if err != nil {
			return nil, errors.Wrap(err, "datamodel.NewRoadsFromRoadModels")
		}
		roads = append(roads, road)
	}

	return roads, nil
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
