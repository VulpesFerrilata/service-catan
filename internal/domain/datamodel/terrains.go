package datamodel

import (
	"math"
	"math/rand"

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

func (t Terrains) Shuffle() {
	rand.Shuffle(len(t), func(i, j int) {
		t[i].terrainType, t[j].terrainType = t[j].terrainType, t[i].terrainType
		t[i].number, t[j].number = t[j].number, t[i].number
	})

	rand.Shuffle(len(t), func(i, j int) {
		if t[i].number == 7 || t[j].number == 7 {
			return
		}

		if t[i].number == 6 || t[i].number == 8 {
			if t[j].number == 6 || t[j].number == 8 {
				return
			}

			isExist := t[j].GetAdjacentTerrains().Any(func(terrain *Terrain) bool {
				return terrain.number == 6 || terrain.number == 8
			})
			if isExist {
				return
			}
		}

		if t[j].number == 6 || t[j].number == 8 {
			if t[i].number == 6 || t[i].number == 8 {
				return
			}

			isExist := t[i].GetAdjacentTerrains().Any(func(terrain *Terrain) bool {
				return terrain.number == 6 || terrain.number == 8
			})
			if isExist {
				return
			}
		}

		t[i].number, t[j].number = t[j].number, t[i].number
	})
}
