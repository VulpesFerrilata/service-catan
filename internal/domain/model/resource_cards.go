package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

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
			resourceCard.Type = resourceType
			resourceCards.append(resourceCard)
		}
	}

	return resourceCards
}

type ResourceCards []*ResourceCard

func (rc *ResourceCards) append(resourceCard *ResourceCard) {
	*rc = append(*rc, resourceCard)
}

type ResourceCardFilterFunc func(resourceCard *ResourceCard) bool

func (rc ResourceCards) Filter(resourceCardFilterFunc ResourceCardFilterFunc) ResourceCards {
	var resourceCards ResourceCards
	for _, resourceCard := range rc {
		if resourceCardFilterFunc(resourceCard) {
			resourceCards.append(resourceCard)
		}
	}
	return resourceCards
}
