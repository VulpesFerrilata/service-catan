package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
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

func NewRoads() Roads {
	var roads Roads

	minQ := 2
	maxQ := 4
	for r := 0; r <= 6; r++ {
		for q := minQ; q <= maxQ; q++ {
			if (r > 0 && q != maxQ) || r > 3 {
				topRoad := new(Road)
				topRoad.q = q
				topRoad.r = r
				topRoad.location = model.TopLeft
				roads = append(roads, topRoad)
			}

			if r > 0 && r < 6 {
				midRoad := new(Road)
				midRoad.q = q
				midRoad.r = r
				midRoad.location = datamodel.RL_MID_LEFT
				roads = append(roads, midRoad)
			}

			if (r < 6 && q != maxQ) || r < 3 {
				botRoad := new(Road)
				botRoad.q = q
				botRoad.r = r
				botRoad.location = datamodel.RL_BOT_LEFT
				roads = append(roads, botRoad)
			}
		}

		if r < 3 {
			minQ--
		} else {
			maxQ--
		}
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
