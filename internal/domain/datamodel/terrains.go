package datamodel

import (
	"math"
	"math/rand"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
)

func NewTerrainsFromTerrainModels(terrainModels []*model.Terrain) Terrains {
	terrains := make(Terrains, 0)

	for _, terrainModel := range terrainModels {
		terrain := NewTerrainFromTerrainModel(terrainModel)
		terrains = append(terrains, terrain)
	}

	return terrains
}

func NewTerrains() Terrains {
	var terrains Terrains

	minQ := 1
	maxQ := 3
	for r := 1; r <= 5; r++ {
		for q := minQ; q <= maxQ; q++ {
			terrain := new(Terrain)
			terrain.Q = q
			terrain.R = r
			terrains.append(terrain)
		}

		if r < 3 {
			minQ--
		} else {
			maxQ--
		}
	}

	normalTerrains, specialTerrains, desertTerrain := terrains.splitRandomly()

	desertTerrain.Type = datamodel.TT_DESERT
	desertTerrain.Number = 7

	specialNumbers := map[int]int{
		6: 2,
		8: 2,
	}
	rand.Shuffle(len(specialTerrains), func(i, j int) { specialTerrains[i], specialTerrains[j] = specialTerrains[j], specialTerrains[i] })
	specialTerrainIdx := 0
	for specialNumber, quantity := range specialNumbers {
		for i := 1; i <= quantity; i++ {
			specialTerrains[specialTerrainIdx].Number = specialNumber
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
			normalTerrains[normalTerrainIdx].Number = numbers
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
			if terrains[terrainIdx].Type == datamodel.TT_DESERT {
				terrainIdx++
			}
			terrains[terrainIdx].Type = terrainType
			terrainIdx++
		}
	}

	return terrains
}

type Terrains []*Terrain

func (t *Terrains) splitRandomly() (Terrains, Terrains, *Terrain) {
	var specialTerrains Terrains
	var normalTerrains Terrains
	var desertTerrain *Terrain

	whitelistTerrains := *t

	desertTerrainIdx := rand.Intn(len(whitelistTerrains))
	desertTerrain = whitelistTerrains[desertTerrainIdx]
	whitelistTerrains = whitelistTerrains.Filter(func(terrain *Terrain) bool {
		if terrain == desertTerrain {
			return false
		}
		return true
	})

	for i := 1; i <= 4; i++ {
		idx := rand.Intn(len(whitelistTerrains))
		specialTerrain := whitelistTerrains[idx]
		specialTerrains = append(specialTerrains, specialTerrain)
		whitelistTerrains = whitelistTerrains.Filter(func(terrain *Terrain) bool {
			if math.Abs(float64(terrain.q-specialTerrain.q)) <= 1 && math.Abs(float64(terrain.r-specialTerrain.r)) <= 1 {
				return false
			}
			return true
		})
	}

	normalTerrains = t.Filter(func(terrain *Terrain) bool {
		for _, specialTerrain := range specialTerrains {
			if terrain == specialTerrain {
				return false
			}
		}
		if terrain == desertTerrain {
			return false
		}
		return true
	})

	return normalTerrains, specialTerrains, desertTerrain
}

type TerrainFilterFunc func(terrain *Terrain) bool

func (t Terrains) Filter(terrainFilterFunc TerrainFilterFunc) Terrains {
	var terrains Terrains
	for _, terrain := range t {
		if terrainFilterFunc(terrain) {
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
