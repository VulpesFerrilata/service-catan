package service

import (
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
)

type RoadService interface {
	GetRoadRepository() repository.RoadRepository
}

func NewRoadService(roadRepository repository.RoadRepository) RoadService {
	return &roadService{
		roadRepository: roadRepository,
	}
}

type roadService struct {
	roadRepository repository.RoadRepository
}

func (rs roadService) GetRoadRepository() repository.RoadRepository {
	return rs.roadRepository
}
