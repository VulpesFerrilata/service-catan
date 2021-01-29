package datamodel

type Terrains []*Terrain

type TerrainFilterFunc func(terrain *Terrain) bool

func (t Terrains) Filter(f TerrainFilterFunc) Terrains {
	var terrains Terrains
	for _, terrain := range t {
		if f(terrain) {
			terrains = append(terrains, terrain)
		}
	}
	return terrains
}

func (t Terrains) First() *Terrain {
	if len(t) > 0 {
		return (t)[0]
	}
	return nil
}

func (t Terrains) Any(f TerrainFilterFunc) bool {
	for _, terrain := range t {
		if f(terrain) {
			return true
		}
	}
	return false
}
