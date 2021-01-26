package service

import (
	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
	"github.com/pkg/errors"
)

type RobberService interface {
	GetRobberRepository() repository.RobberRepository
	InitRobber(terrains datamodel.Terrains) (*datamodel.Robber, error)
}

func NewRobberService(robberRepository repository.RobberRepository) RobberService {
	return &robberService{
		robberRepository: robberRepository,
	}
}

type robberService struct {
	robberRepository repository.RobberRepository
}

func (rs robberService) GetRobberRepository() repository.RobberRepository {
	return rs.robberRepository
}

func (rs robberService) InitRobber(terrains datamodel.Terrains) (*datamodel.Robber, error) {
	terrain := terrains.Filter(func(terrain *datamodel.Terrain) bool {
		return terrain.GetTerrainType() == datamodel.DesertTerrain
	}).First()
	robber, err := datamodel.NewRobber(terrain)
	return robber, errors.Wrap(err, "service.RobberService.InitRobber")
}
