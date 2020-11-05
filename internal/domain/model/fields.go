package model

import (
	"math/rand"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
)

func NewFields(game *Game) Fields {
	var fields Fields

	for r := 0; r <= 4; r++ {
		minQ := 1
		maxQ := 3
		for q := minQ; q <= maxQ; q++ {
			field := NewField(game)
			field.Q = q
			field.R = r
			fields.append(field)
		}
		if r < 2 {
			minQ--
		} else {
			maxQ--
		}
	}

	fieldTypes := map[datamodel.FieldType]int{
		datamodel.FT_WOOD:   4,
		datamodel.FT_BRICK:  3,
		datamodel.FT_SHEEP:  4,
		datamodel.FT_GRAIN:  4,
		datamodel.FT_STONE:  3,
		datamodel.FT_DESERT: 1,
	}

	rand.Shuffle(len(fields), func(i, j int) { fields[i], fields[j] = fields[j], fields[i] })
	fieldIdx := 0
	for fieldType, quantity := range fieldTypes {
		for i := 1; i <= quantity; i++ {
			fields[fieldIdx].Type = fieldType
			fieldIdx++
		}
	}

	numbers := map[int]int{
		2:  1,
		3:  2,
		4:  2,
		5:  2,
		6:  2,
		8:  2,
		9:  2,
		10: 2,
		11: 2,
		12: 1,
	}

	rand.Shuffle(len(fields), func(i, j int) { fields[i], fields[j] = fields[j], fields[i] })
	fieldIdx = 0
	for number, quantity := range numbers {
		for i := 1; i <= quantity; i++ {
			if fields[fieldIdx].Type == datamodel.FT_DESERT {
				fieldIdx++
			}
			fields[fieldIdx].Number = number
			fieldIdx++
		}
	}

	return fields
}

type Fields []*Field

func (f Fields) append(field *Field) {
	f = append(f, field)
}

func (f Fields) SetGame(game *Game) {
	for _, field := range f {
		field.SetGame(game)
	}
}

type FieldFilterFunc func(field *Field) bool

func (f Fields) Filter(fieldFilterFunc FieldFilterFunc) Fields {
	var fields Fields
	for _, field := range f {
		if fieldFilterFunc(field) {
			fields.append(field)
		}
	}
	return fields
}
