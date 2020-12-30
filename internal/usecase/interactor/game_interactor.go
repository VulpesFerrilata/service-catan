package interactor

import (
	"context"
	"math/rand"

	"github.com/pkg/errors"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/catan/internal/domain/service"
	"github.com/VulpesFerrilata/catan/internal/usecase/request"
	"github.com/VulpesFerrilata/catan/internal/usecase/response"
	"github.com/VulpesFerrilata/grpc/protoc/user"
	"github.com/VulpesFerrilata/library/pkg/validator"
)

type GameInteractor interface {
}

func NewGameInteractor(validate validator.Validate,
	gameAggregateService service.GameAggregateService,
	userService user.UserService) GameInteractor {
	return &gameInteractor{
		validate:             validate,
		gameAggregateService: gameAggregateService,
		userService:          userService,
	}
}

type gameInteractor struct {
	validate             validator.Validate
	gameAggregateService service.GameAggregateService
	userService          user.UserService
}

func (gi *gameInteractor) CreateGame(ctx context.Context) (*response.GameResponse, error) {
	game := model.NewGame()

	if err := gi.gameAggregateService.Save(ctx, game); err != nil {
		return nil, err
	}

	return nil, nil
}

func (gi *gameInteractor) JoinGame(ctx context.Context, gameRequest *request.GameRequest) (*response.GameResponse, error) {
	userId := 0 //todo: userid from context

	game, err := gi.gameAggregateService.GetById(ctx, uint(gameRequest.ID))
	if err != nil {
		return nil, err
	}

	switch game.Status {
	case datamodel.GS_STARTED:
		return nil, errors.New("game has already started")
	case datamodel.GS_FINISHED:
		return nil, errors.New("game was finished")
	}

	players := game.GetPlayers()
	if len(players) >= 4 {
		return nil, errors.New("room is full")
	}

	player := players.Filter(func(player *model.Player) bool {
		if player.UserID == nil {
			return false
		}
		return *player.UserID == uint(userId)
	}).First()

	if player == nil {
		userRequestPb := new(user.UserRequest)
		userRequestPb.ID = int64(userId)
		userPb, err := gi.userService.GetUserById(ctx, userRequestPb)
		if err != nil {
			return nil, err
		}
		user := model.NewUser(userPb)
		player := new(model.Player)
		player.SetUser(user)
		game.AddPlayers(player)

		if err := gi.gameAggregateService.Save(ctx, game); err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func (gi *gameInteractor) LoadGame(ctx context.Context, gameRequest *request.GameRequest) (*response.GameResponse, error) {
	return nil, nil
}

func (gi *gameInteractor) UpdateGame(ctx context.Context, gameRequest *request.GameRequest) (*response.GameResponse, error) {
	return nil, nil
}

func (gi *gameInteractor) StartGame(ctx context.Context, gameRequest *request.GameRequest) (*response.GameResponse, error) {
	game, err := gi.gameAggregateService.GetById(ctx, uint(gameRequest.ID))
	if err != nil {
		return nil, err
	}

	switch game.Status {
	case datamodel.GS_STARTED:
		return nil, errors.New("game has already started")
	case datamodel.GS_FINISHED:
		return nil, errors.New("game was finished")
	}

	userId := uint(0) //todo: userid from context
	player := game.GetPlayers().Filter(func(player *model.Player) bool {
		if player.UserID == nil {
			return false
		}
		return *player.UserID == userId
	}).First()
	if player == nil {
		return nil, errors.New("player is not exists")
	}
	if !player.IsHost() {
		return nil, errors.New("only host player can start game")
	}

	game.Status = datamodel.GS_STARTED

	dices := model.NewDices()
	game.AddDices(dices...)

	achievements := model.NewAchievements()
	game.AddAchievements(achievements...)

	resourceCards := model.NewResourceCards()
	game.AddResourceCards(resourceCards...)

	developmentCards := model.NewDevelopmentCards()
	game.AddDevelopmentCards(developmentCards...)

	terrains := model.NewTerrains()
	game.AddTerrains(terrains...)

	robber := model.NewRobber(terrains)
	game.SetRobber(robber)

	constructions := model.NewConstructions()
	game.AddConstructions(constructions...)

	roads := model.NewRoads()
	game.AddRoads(roads...)

	harbors := model.NewHarbors(terrains)
	game.AddHarbors(harbors...)

	players := game.GetPlayers()
	colors := []string{
		"#ff0000",
		"#0000ff",
		"#ffffff",
		"#000000",
	}
	rand.Shuffle(len(colors), func(i, j int) { colors[i], colors[j] = colors[j], colors[i] })

	rand.Shuffle(len(players), func(i, j int) { players[i], players[j] = players[j], players[i] })
	for idx, player := range players {
		player.TurnOrder = idx + 1
		player.Color = colors[idx]
		if player.TurnOrder == 1 {
			game.PlayerInTurn = &player.ID
		}
	}

	if err := gi.gameAggregateService.Save(ctx, game); err != nil {
		return nil, err
	}

	return nil, nil
}

func (gi *gameInteractor) LeaveGame(ctx context.Context, gameRequest *request.GameRequest) (*response.GameResponse, error) {
	game, err := gi.gameAggregateService.GetById(ctx, uint(gameRequest.ID))
	if err != nil {
		return nil, err
	}

	userId := uint(0) //todo: userid from context
	player := game.GetPlayers().Filter(func(player *model.Player) bool {
		if player.UserID == nil {
			return false
		}
		return *player.UserID == userId
	}).First()
	if player != nil {
		switch game.Status {
		case datamodel.GS_WAITING:
			player.Remove()
		case datamodel.GS_STARTED:
			player.IsLeft = true
		}

		if err := gi.gameAggregateService.Save(ctx, game); err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func (gi *gameInteractor) RollDices(ctx context.Context, gameRequest *request.GameRequest) (*response.GameResponse, error) {
	userId := 0

	game, err := gi.gameAggregateService.GetById(ctx, uint(gameRequest.ID))
	if err != nil {
		return nil, err
	}

	player := game.GetPlayers().Filter(func(player *model.Player) bool {
		if player.UserID == nil {
			return false
		}
		return *player.UserID == uint(userId)
	}).First()
	if player == nil {
		return nil, errors.New("player is not exists")
	}
	if !player.IsInTurn() {
		return nil, errors.New("player is not in turn")
	}

	if game.Turn == 1 || game.Turn == 2 {
		return nil, errors.New("action is restricted in setup phase")
	}

	dices := game.GetDices()
	if dices.IsRolled() {
		return nil, errors.New("player already rolled dices")
	}

	totalNumber := dices.Roll()
	if totalNumber == 7 {
		robber := game.GetRobber()
		robber.Status = datamodel.RS_MOVE
	} else {
		terrains := game.GetTerrains().Filter(func(terrain *model.Terrain) bool {
			return terrain.Number == totalNumber && !terrain.HasRobber()
		})
		for _, terrain := range terrains {
			constructions := terrain.GetAdjacentConstructions().Filter(func(construction *model.Construction) bool {
				return construction.PlayerID != nil
			})
			resourceCards := game.GetResourceCards().Filter(func(resourceCard *model.ResourceCard) bool {
				var resourceType datamodel.ResourceType
				switch terrain.Type {
				case datamodel.TT_FOREST:
					resourceType = datamodel.RT_LUMBER
				case datamodel.TT_HILL:
					resourceType = datamodel.RT_BRICK
				case datamodel.TT_PASTURE:
					resourceType = datamodel.RT_WOOL
				case datamodel.TT_FIELD:
					resourceType = datamodel.RT_GRAIN
				case datamodel.TT_MOUNTAIN:
					resourceType = datamodel.RT_ORE
				}
				return resourceCard.PlayerID == nil && resourceCard.Type == resourceType
			})

			resourceCardsDemand := 0
			for _, construction := range constructions {
				resourceCardsDemand++
				if construction.IsUpgradedCastle {
					resourceCardsDemand++
				}
			}
			if resourceCardsDemand > len(resourceCards) {
				continue
			}

			resourceCardIdx := 0
			for _, construction := range constructions {
				resourceDispatchQuantity := 1
				if construction.IsUpgradedCastle {
					resourceDispatchQuantity++
				}

				for i := 1; i <= resourceDispatchQuantity; i++ {
					resourceCards[resourceCardIdx].PlayerID = construction.PlayerID
					resourceCardIdx++
				}
			}
		}
	}

	if err := gi.gameAggregateService.Save(ctx, game); err != nil {
		return nil, err
	}

	return nil, nil
}

func (gi *gameInteractor) BuildRoad(ctx context.Context, roadRequest *request.RoadRequest) (*response.GameResponse, error) {
	userId := 0

	game, err := gi.gameAggregateService.GetById(ctx, uint(roadRequest.GameID))
	if err != nil {
		return nil, err
	}

	player := game.GetPlayers().Filter(func(player *model.Player) bool {
		if player.UserID == nil {
			return false
		}
		return *player.UserID == uint(userId)
	}).First()
	if player == nil {
		return nil, errors.New("player is not exists")
	}
	if !player.IsInTurn() {
		return nil, errors.New("player is not in turn")
	}

	road := game.GetRoads().Filter(func(road *model.Road) bool {
		return road.ID == uint(roadRequest.ID)
	}).First()
	if road == nil {
		return nil, errors.New("road is not exists")
	}
	if road.PlayerID != nil {
		return nil, errors.New("road is owned by another player")
	}

	gameTurn := game.Turn
	if gameTurn == 1 || gameTurn == 2 {
		roads := player.GetRoads()

		if len(roads) >= gameTurn {
			return nil, errors.New("player has already build road")
		}
	} else {
		nearbyConstructions := road.GetAdjacentConstructions().Filter(func(construction *model.Construction) bool {
			if construction.PlayerID == nil {
				return false
			}
			return *construction.PlayerID == player.ID
		})

		nearbyRoads := road.GetAdjacentRoads().Filter(func(road *model.Road) bool {
			if road.PlayerID == nil {
				return false
			}
			return *road.PlayerID == player.ID
		})

		if len(nearbyConstructions) == 0 && len(nearbyRoads) == 0 {
			return nil, errors.New("road should be build next to another road or construction")
		}

		lumberResourceCards := player.GetResourceCards().Filter(func(resourceCard *model.ResourceCard) bool {
			return resourceCard.Type == datamodel.RT_LUMBER
		})

		brickResourceCards := player.GetResourceCards().Filter(func(resourceCard *model.ResourceCard) bool {
			return resourceCard.Type == datamodel.RT_BRICK
		})

		if len(lumberResourceCards) < 1 || len(brickResourceCards) < 1 {
			return nil, errors.New("insuffigient resources")
		}

		lumberResourceCards[0].PlayerID = nil
		brickResourceCards[0].PlayerID = nil
	}
	road.PlayerID = &player.ID

	if err := gi.gameAggregateService.Save(ctx, game); err != nil {
		return nil, err
	}

	return nil, nil
}
