package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/model"
)

func NewResourceCardsFromResourceCardModels(resourceCardModels []*model.ResourceCard) ResourceCards {
	resourceCards := make(ResourceCards, 0)

	for _, resourceCardModel := range resourceCardModels {
		resourceCard := NewResourceCardFromResourceCardModel(resourceCardModel)
		resourceCards = append(resourceCards, resourceCard)
	}

	return resourceCards
}

func NewResourceCards() ResourceCards {
	var resourceCards ResourceCards

	resourceTypes := map[datamodel.ResourceType]int{
		datamodel.RT_LUMBER: 19,
		datamodel.RT_BRICK:  19,
		datamodel.RT_WOOL:   19,
		datamodel.RT_GRAIN:  19,
		datamodel.RT_ORE:    19,
	}
	for resourceType, quantity := range resourceTypes {
		for i := 1; i <= quantity; i++ {
			resourceCard := new(ResourceCard)
			resourceCard.resourceType = resourceType
			resourceCards = append(resourceCards, resourceCard)
		}
	}

	return resourceCards
}

type ResourceCards []*ResourceCard

type ResourceCardFilterFunc func(resourceCard *ResourceCard) bool

func (rc ResourceCards) Filter(resourceCardFilterFunc ResourceCardFilterFunc) ResourceCards {
	var resourceCards ResourceCards
	for _, resourceCard := range rc {
		if resourceCardFilterFunc(resourceCard) {
			resourceCards = append(resourceCards, resourceCard)
		}
	}
	return resourceCards
}
