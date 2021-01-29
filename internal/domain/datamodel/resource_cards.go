package datamodel

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
