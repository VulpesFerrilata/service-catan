package model

import (
	"math"
	"math/rand"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
)

func NewTerrains() Terrains {
	var terrains Terrains

	minQ := 1
	maxQ := 3
	for r := 1; r <= 5; r++ {
		for q := minQ; q <= maxQ; q++ {
			terrain := NewTerrain()
			terrain.terrain.Q = q
			terrain.terrain.R = r
			terrains.append(terrain)
		}

		if r < 3 {
			minQ--
		} else {
			maxQ--
		}
	}

	normalTerrains, specialTerrains, desertTerrain := terrains.splitRandomly()

	desertTerrain.terrain.Type = datamodel.TT_DESERT
	desertTerrain.terrain.Number = 7

	specialNumbers := map[int]int{
		6: 2,
		8: 2,
	}
	rand.Shuffle(len(specialTerrains), func(i, j int) { specialTerrains[i], specialTerrains[j] = specialTerrains[j], specialTerrains[i] })
	specialTerrainIdx := 0
	for specialNumber, quantity := range specialNumbers {
		for i := 1; i <= quantity; i++ {
			specialTerrains[specialTerrainIdx].terrain.Number = specialNumber
			specialTerrainIdx++
		}
	}

	numbers := map[int]int{
		2:  1,
		3:  2,
		4:  2,
		5:  2,
		9:  2,
		10: 2,
		11: 2,
		12: 1,
	}
	rand.Shuffle(len(normalTerrains), func(i, j int) { normalTerrains[i], normalTerrains[j] = normalTerrains[j], normalTerrains[i] })
	normalTerrainIdx := 0
	for numbers, quantity := range numbers {
		for i := 1; i <= quantity; i++ {
			normalTerrains[normalTerrainIdx].terrain.Number = numbers
			normalTerrainIdx++
		}
	}

	terrainTypes := map[datamodel.TerrainType]int{
		datamodel.TT_FOREST:   4,
		datamodel.TT_HILL:     3,
		datamodel.TT_PASTURE:  4,
		datamodel.TT_FIELD:    4,
		datamodel.TT_MOUNTAIN: 3,
	}

	rand.Shuffle(len(terrains), func(i, j int) { terrains[i], terrains[j] = terrains[j], terrains[i] })
	terrainIdx := 0
	for terrainType, quantity := range terrainTypes {
		for i := 1; i <= quantity; i++ {
			if terrains[terrainIdx].terrain.Type == datamodel.TT_DESERT {
				terrainIdx++
			}
			terrains[terrainIdx].terrain.Type = terrainType
			terrainIdx++
		}
	}

	return terrains
}

type Terrains []*Terrain

func (t *Terrains) append(terrain *Terrain) {
	*t = append(*t, terrain)
}

func (t *Terrains) splitRandomly() (Terrains, Terrains, *Terrain) {
	var specialFields Terrains
	var normalFields Terrains
	var desertField *Terrain

	whitelistFields := *t

	desertFieldIdx := rand.Intn(len(whitelistFields))
	desertField = whitelistFields[desertFieldIdx]
	whitelistFields = whitelistFields.Filter(func(terrain *Terrain) bool {
		if terrain == desertField {
			return false
		}
		return true
	})

	for i := 1; i <= 4; i++ {
		idx := rand.Intn(len(whitelistFields))
		specialField := whitelistFields[idx]
		specialFields.append(specialField)
		whitelistFields = whitelistFields.Filter(func(terrain *Terrain) bool {
			if math.Abs(float64(terrain.GetQ()-specialField.GetQ())) <= 1 && math.Abs(float64(terrain.GetR()-specialField.GetR())) <= 1 {
				return false
			}
			return true
		})
	}

	normalFields = t.Filter(func(terrain *Terrain) bool {
		for _, specialField := range specialFields {
			if terrain == specialField {
				return false
			}
		}
		if terrain == desertField {
			return false
		}
		return true
	})

	return normalFields, specialFields, desertField
}

type TerrainFilterFunc func(terrain *Terrain) bool

func (t Terrains) Filter(terrainFilterFunc TerrainFilterFunc) Terrains {
	var terrains Terrains
	for _, terrain := range t {
		if terrainFilterFunc(terrain) {
			terrains.append(terrain)
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
