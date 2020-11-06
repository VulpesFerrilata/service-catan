package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

func NewRoads(game *Game) Roads {
	var roads Roads

	minQ := 2
	maxQ := 4
	for r := 0; r <= 6; r++ {
		for q := minQ; q <= maxQ; q++ {
			if (r > 0 && q != maxQ) || r > 3 {
				topRoad := NewRoad(game)
				topRoad.Q = q
				topRoad.R = r
				topRoad.Location = datamodel.RL_TOP_LEFT
				roads.append(topRoad)
			}

			if r > 0 && r < 6 {
				midRoad := NewRoad(game)
				midRoad.Q = q
				midRoad.R = r
				midRoad.Location = datamodel.RL_MID_LEFT
				roads.append(midRoad)
			}

			if (r < 6 && q != maxQ) || r < 3 {
				botRoad := NewRoad(game)
				botRoad.Q = q
				botRoad.R = r
				botRoad.Location = datamodel.RL_BOT_LEFT
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

func (r *Roads) SetGame(game *Game) {
	for _, road := range *r {
		road.SetGame(game)
	}
}
