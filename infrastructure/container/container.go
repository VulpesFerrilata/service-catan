package container

import (
	"github.com/VulpesFerrilata/catan/infrastructure/iris/router"
	"github.com/VulpesFerrilata/catan/infrastructure/iris/server"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
	"github.com/VulpesFerrilata/catan/internal/domain/service"
	"github.com/VulpesFerrilata/library/config"
	"github.com/VulpesFerrilata/library/pkg/database"
	"github.com/VulpesFerrilata/library/pkg/db"
	"github.com/VulpesFerrilata/library/pkg/middleware"
	"github.com/VulpesFerrilata/library/pkg/translator"
	"github.com/VulpesFerrilata/library/pkg/validator"
	"go.uber.org/dig"
)

func NewContainer() *dig.Container {
	container := dig.New()

	//--Config
	container.Provide(config.NewConfig)

	//--Domain
	container.Provide(repository.NewAchievementRepository)
	container.Provide(repository.NewConstructionRepository)
	container.Provide(repository.NewDevelopmentCardRepository)
	container.Provide(repository.NewDiceRepository)
	container.Provide(repository.NewGameRepository)
	container.Provide(repository.NewHarborRepository)
	container.Provide(repository.NewPlayerRepository)
	container.Provide(repository.NewResourceCardRepository)
	container.Provide(repository.NewRoadRepository)
	container.Provide(repository.NewRobberRepository)
	container.Provide(repository.NewTerrainRepository)

	container.Provide(service.NewAchievementService)
	container.Provide(service.NewConstructionService)
	container.Provide(service.NewDevelopmentCardService)
	container.Provide(service.NewDiceService)
	container.Provide(service.NewGameService)
	container.Provide(service.NewGameAggregateService)
	container.Provide(service.NewHarborService)
	container.Provide(service.NewPlayerService)
	container.Provide(service.NewPlayerAggregateService)
	container.Provide(service.NewResourceCardService)
	container.Provide(service.NewRoadService)
	container.Provide(service.NewRobberService)
	container.Provide(service.NewTerrainService)
	//--Usecase
	container.Provide(interactor.NewAuthInteractor)
	//--Gateways
	container.Provide(gateway.NewUserService)

	//--Utility
	container.Provide(database.NewGorm)
	container.Provide(db.NewDbContext)
	container.Provide(translator.NewTranslator)
	container.Provide(validator.NewValidate)

	//--Middleware
	container.Provide(middleware.NewTransactionMiddleware)
	container.Provide(middleware.NewTranslatorMiddleware)
	container.Provide(middleware.NewErrorMiddleware)

	//--Controller
	container.Provide(controller.NewAuthController)
	//--Router
	container.Provide(router.NewRouter)
	//--Server
	container.Provide(server.NewServer)

	//--Grpc
	container.Provide(handler.NewAuthHandler)

	return container
}
