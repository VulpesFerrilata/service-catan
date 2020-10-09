package dto

import "github.com/VulpesFerrilata/catan/internal/domain/model"

type ResourceCardDTO struct {
	ID           int
	ResourceType model.ResourceType
}
