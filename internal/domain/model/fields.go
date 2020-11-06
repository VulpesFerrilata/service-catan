package model

import (
	"math"
	"math/rand"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
)

func NewFields(game *Game) Fields {
	var fields Fields

	minQ := 1
	maxQ := 3
	for r := 1; r <= 5; r++ {
		for q := minQ; q <= maxQ; q++ {
			field := NewField(game)
			field.Q = q
			field.R = r
			fields.append(field)
		}

		if r < 3 {
			minQ--
		} else {
			maxQ--
		}
	}

	normalFields, specialFields, desertField := fields.split()

	desertField.Type = datamodel.FT_DESERT
	desertField.Number = 7

	specialNumbers := map[int]int{
		6: 2,
		8: 2,
	}
	rand.Shuffle(len(specialFields), func(i, j int) { specialFields[i], specialFields[j] = specialFields[j], specialFields[i] })
	specialFieldIdx := 0
	for specialNumber, quantity := range specialNumbers {
		for i := 1; i <= quantity; i++ {
			specialFields[specialFieldIdx].Number = specialNumber
			specialFieldIdx++
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
	rand.Shuffle(len(normalFields), func(i, j int) { normalFields[i], normalFields[j] = normalFields[j], normalFields[i] })
	normalFieldIdx := 0
	for numbers, quantity := range numbers {
		for i := 1; i <= quantity; i++ {
			normalFields[normalFieldIdx].Number = numbers
			normalFieldIdx++
		}
	}

	fieldTypes := map[datamodel.FieldType]int{
		datamodel.FT_WOOD:  4,
		datamodel.FT_BRICK: 3,
		datamodel.FT_SHEEP: 4,
		datamodel.FT_GRAIN: 4,
		datamodel.FT_STONE: 3,
	}

	rand.Shuffle(len(fields), func(i, j int) { fields[i], fields[j] = fields[j], fields[i] })
	fieldIdx := 0
	for fieldType, quantity := range fieldTypes {
		for i := 1; i <= quantity; i++ {
			if fields[fieldIdx].Type == datamodel.FT_DESERT {
				fieldIdx++
			}
			fields[fieldIdx].Type = fieldType
			fieldIdx++
		}
	}

	return fields
}

type Fields []*Field

func (f *Fields) append(field *Field) {
	*f = append(*f, field)
}

func (f *Fields) split() (Fields, Fields, *Field) {
	var specialFields Fields
	var normalFields Fields
	var desertField *Field

	whitelistFields := *f

	desertFieldIdx := rand.Intn(len(whitelistFields))
	desertField = whitelistFields[desertFieldIdx]
	whitelistFields = whitelistFields.Filter(func(field *Field) bool {
		if field == desertField {
			return false
		}
		return true
	})

	for i := 1; i <= 4; i++ {
		idx := rand.Intn(len(whitelistFields))
		specialField := whitelistFields[idx]
		specialFields.append(specialField)
		whitelistFields = whitelistFields.Filter(func(field *Field) bool {
			if math.Abs(float64(field.Q-specialField.Q)) <= 1 && math.Abs(float64(field.R-specialField.R)) <= 1 {
				return false
			}
			return true
		})
	}

	normalFields = f.Filter(func(field *Field) bool {
		for _, specialField := range specialFields {
			if field == specialField {
				return false
			}
		}
		if field == desertField {
			return false
		}
		return true
	})

	return normalFields, specialFields, desertField
}

func (f *Fields) SetGame(game *Game) {
	for _, field := range *f {
		field.SetGame(game)
	}
}

type FieldFilterFunc func(field *Field) bool

func (f *Fields) Filter(fieldFilterFunc FieldFilterFunc) Fields {
	var fields Fields
	for _, field := range *f {
		if fieldFilterFunc(field) {
			fields.append(field)
		}
	}
	return fields
}

func (f Fields) First() *Field {
	if len(f) > 0 {
		return (f)[0]
	}
	return nil
}
