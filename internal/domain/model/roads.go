package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

func NewRoads() Roads {
	var roads Roads

	minQ := 2
	maxQ := 4
	for r := 0; r <= 6; r++ {
		for q := minQ; q <= maxQ; q++ {
			if (r > 0 && q != maxQ) || r > 3 {
				topRoad := NewRoad()
				topRoad.road.Q = q
				topRoad.road.R = r
				topRoad.road.Location = datamodel.RL_TOP_LEFT
				roads.append(topRoad)
			}

			if r > 0 && r < 6 {
				midRoad := NewRoad()
				midRoad.road.Q = q
				midRoad.road.R = r
				midRoad.road.Location = datamodel.RL_MID_LEFT
				roads.append(midRoad)
			}

			if (r < 6 && q != maxQ) || r < 3 {
				botRoad := NewRoad()
				botRoad.road.Q = q
				botRoad.road.R = r
				botRoad.road.Location = datamodel.RL_BOT_LEFT
				roads.append(botRoad)
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

func (r *Roads) append(road *Road) {
	*r = append(*r, road)
}

type RoadFilterFunc func(road *Road) bool

func (r Roads) Filter(roadFilterFunc RoadFilterFunc) Roads {
	var roads Roads

	for _, road := range r {
		if roadFilterFunc(road) {
			roads.append(road)
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
