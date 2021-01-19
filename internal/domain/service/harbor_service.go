package service

import (
	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
	"github.com/pkg/errors"
)

type HarborService interface {
	GetHarborRepository() repository.HarborRepository
}

func NewHarborService(harborRepository repository.HarborRepository) HarborService {
	return &harborService{
		harborRepository: harborRepository,
	}
}

type harborService struct {
	harborRepository repository.HarborRepository
}

func (hs harborService) GetHarborRepository() repository.HarborRepository {
	return hs.harborRepository
}

func (hs harborService) InitHarbors(terrains datamodel.Terrains) (datamodel.Harbors, error) {
	harbors := make(datamodel.Harbors, 0)

	harborTypes := []model.HarborType{
		model.GeneralHarbor,
		model.GeneralHarbor,
		model.GeneralHarbor,
		model.GeneralHarbor,
		model.GeneralHarbor,
		model.LumberHarbor,
		model.BrickHarbor,
		model.WoolHarbor,
		model.GrainHarbor,
		model.OreHarbor,
	}

	terrain := terrains.Filter(func(terrain *datamodel.Terrain) bool {
		return terrain.GetQ() == 1 && terrain.GetR() == 1
	}).First()
	harborType, harborTypes := harborTypes[0], harborTypes[1:]
	firstHarbor, err := datamodel.NewHarbor(1, 0, harborType, terrain)
	if err != nil {
		return nil, errors.Wrap(err, "service.HarborService.InitHarbors")
	}
	harbors = append(harbors, firstHarbor)

	terrain = terrains.Filter(func(terrain *datamodel.Terrain) bool {
		return terrain.GetQ() == 2 && terrain.GetR() == 1
	}).First()
	harborType, harborTypes = harborTypes[0], harborTypes[1:]
	secondHarbor, err := datamodel.NewHarbor(3, 0, harborType, terrain)
	if err != nil {
		return nil, errors.Wrap(err, "service.HarborService.InitHarbors")
	}
	harbors = append(harbors, secondHarbor)

	terrain = terrains.Filter(func(terrain *datamodel.Terrain) bool {
		return terrain.GetQ() == 3 && terrain.GetR() == 2
	}).First()
	harborType, harborTypes = harborTypes[0], harborTypes[1:]
	thirdHarbor, err := datamodel.NewHarbor(4, 1, harborType, terrain)
	if err != nil {
		return nil, errors.Wrap(err, "service.HarborService.InitHarbors")
	}
	harbors = append(harbors, thirdHarbor)

	terrain = terrains.Filter(func(terrain *datamodel.Terrain) bool {
		return terrain.GetQ() == 0 && terrain.GetR() == 2
	}).First()
	harborType, harborTypes = harborTypes[0], harborTypes[1:]
	fourthHarbor, err := datamodel.NewHarbor(-1, 2, harborType, terrain)
	if err != nil {
		return nil, errors.Wrap(err, "service.HarborService.InitHarbors")
	}
	harbors = append(harbors, fourthHarbor)

	terrain = terrains.Filter(func(terrain *datamodel.Terrain) bool {
		return terrain.GetQ() == 3 && terrain.GetR() == 3
	}).First()
	harborType, harborTypes = harborTypes[0], harborTypes[1:]
	fifthHarbor, err := datamodel.NewHarbor(4, 3, harborType, terrain)
	if err != nil {
		return nil, errors.Wrap(err, "service.HarborService.InitHarbors")
	}
	harbors = append(harbors, fifthHarbor)

	terrain = terrains.Filter(func(terrain *datamodel.Terrain) bool {
		return terrain.GetQ() == -1 && terrain.GetR() == 4
	}).First()
	harborType, harborTypes = harborTypes[0], harborTypes[1:]
	sixthHarbor, err := datamodel.NewHarbor(-2, 4, harborType, terrain)
	if err != nil {
		return nil, errors.Wrap(err, "service.HarborService.InitHarbors")
	}
	harbors = append(harbors, sixthHarbor)

	terrain = terrains.Filter(func(terrain *datamodel.Terrain) bool {
		return terrain.GetQ() == 2 && terrain.GetR() == 4
	}).First()
	harborType, harborTypes = harborTypes[0], harborTypes[1:]
	seventhHarbor, err := datamodel.NewHarbor(2, 5, harborType, terrain)
	if err != nil {
		return nil, errors.Wrap(err, "service.HarborService.InitHarbors")
	}
	harbors = append(harbors, seventhHarbor)

	terrain = terrains.Filter(func(terrain *datamodel.Terrain) bool {
		return terrain.GetQ() == -1 && terrain.GetR() == 5
	}).First()
	harborType, harborTypes = harborTypes[0], harborTypes[1:]
	eighthHarbor, err := datamodel.NewHarbor(-2, 6, harborType, terrain)
	if err != nil {
		return nil, errors.Wrap(err, "service.HarborService.InitHarbors")
	}
	harbors = append(harbors, eighthHarbor)

	terrain = terrains.Filter(func(terrain *datamodel.Terrain) bool {
		return terrain.GetQ() == 0 && terrain.GetR() == 5
	}).First()
	harborType, harborTypes = harborTypes[0], harborTypes[1:]
	ninthHarbor, err := datamodel.NewHarbor(0, 6, harborType, terrain)
	if err != nil {
		return nil, errors.Wrap(err, "service.HarborService.InitHarbors")
	}
	harbors = append(harbors, ninthHarbor)

	harbors.Shuffle()

	return harbors, nil
}
