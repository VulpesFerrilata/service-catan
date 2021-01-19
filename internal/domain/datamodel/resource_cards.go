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
