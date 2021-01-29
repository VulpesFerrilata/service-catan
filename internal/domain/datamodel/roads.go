package datamodel

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
