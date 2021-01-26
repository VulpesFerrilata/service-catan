package datamodel

import "fmt"

func NewConstructionType(value string) (constructionType, error) {
	constructionType := constructionType(value)
	if _, ok := constructionTypes[constructionType]; ok {
		return constructionType, nil
	}
	return constructionType, fmt.Errorf("construction type is invalid: %s", value)
}

type constructionType string

func (c constructionType) String() string {
	return c.String()
}

const (
	Land       constructionType = "Land"
	Settlement constructionType = "Settlement"
	City       constructionType = "City"
)

var constructionTypes = map[constructionType]struct{}{
	Land:       {},
	Settlement: {},
	City:       {},
}
