package interactor

import (
	"context"
	"errors"
	"math/rand"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/catan/internal/domain/service"
	"github.com/VulpesFerrilata/catan/internal/usecase/request"
	"github.com/VulpesFerrilata/catan/internal/usecase/response"
	"github.com/VulpesFerrilata/grpc/protoc/user"
	"github.com/VulpesFerrilata/library/pkg/validator"
)

type CatanInteractor interface {
}

func NewCatanInteractor(validate validator.Validate,
	roomService service.RoomService,
	gameAggregateService service.GameAggregateService,
	userService user.UserService) CatanInteractor {
	return &catanInteractor{
		validate:             validate,
		roomService:          roomService,
		gameAggregateService: gameAggregateService,
		userService:          userService,
	}
}

type catanInteractor struct {
	validate             validator.Validate
	roomService          service.RoomService
	gameAggregateService service.GameAggregateService
	userService          user.UserService
}

func (ci *catanInteractor) FindRooms(ctx context.Context, roomRequest *request.RoomRequest) (response.RoomsResponse, error) {
	if err := ci.validate.Struct(ctx, roomRequest); err != nil {
		return nil, err
	}

	if roomRequest.ID != 0 {
		room, err := ci.roomService.GetRoomById(ctx, uint(roomRequest.ID))
		if err != nil {
			return nil, err
		}

		return response.NewRoomsResponse(room), nil
	}

	rooms, err := ci.roomService.FindRoomsByStatus(ctx, roomRequest.Status)
	if err != nil {
		return nil, err
	}
	return response.NewRoomsResponse(rooms...), nil
}

func (ci *catanInteractor) CreateGame(ctx context.Context) (*response.GameResponse, error) {
	game := model.NewGame()

	if err := ci.gameAggregateService.Save(ctx, game); err != nil {
		return nil, err
	}

	return response.NewGameResponse(game), nil
}

func (ci *catanInteractor) JoinGame(ctx context.Context, gameRequest *request.GameRequest) (*response.GameResponse, error) {
	userId := 0 //todo: userid from context

	game, err := ci.gameAggregateService.GetById(ctx, uint(gameRequest.ID))
	if err != nil {
		return nil, err
	}

	switch game.GetStatus() {
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
		if player.GetUserId() == nil {
			return false
		}
		return *player.GetUserId() == uint(userId)
	}).First()

	if player == nil {
		userRequestPb := new(user.UserRequest)
		userRequestPb.ID = int64(userId)
		userPb, err := ci.userService.GetUserById(ctx, userRequestPb)
		if err != nil {
			return nil, err
		}
		user := model.NewUser(userPb)
		player := model.NewPlayer()
		player.SetUser(user)
		game.AddPlayers(player)

		if err := ci.gameAggregateService.Save(ctx, game); err != nil {
			return nil, err
		}
	}

	return response.NewGameResponse(game), nil
}

func (ci *catanInteractor) LoadGame(ctx context.Context, gameRequest *request.GameRequest) (*response.GameResponse, error) {
	return nil, nil
}

func (ci *catanInteractor) UpdateGame(ctx context.Context, gameRequest *request.GameRequest) (*response.GameResponse, error) {
	return nil, nil
}

func (ci *catanInteractor) StartGame(ctx context.Context, gameRequest *request.GameRequest) (*response.GameResponse, error) {
	game, err := ci.gameAggregateService.GetById(ctx, uint(gameRequest.ID))
	if err != nil {
		return nil, err
	}

	switch game.GetStatus() {
	case datamodel.GS_STARTED:
		return nil, errors.New("game has already started")
	case datamodel.GS_FINISHED:
		return nil, errors.New("game was finished")
	}

	userId := uint(0) //todo: userid from context
	player := game.GetPlayers().Filter(func(player *model.Player) bool {
		if player.GetUserId() == nil {
			return false
		}
		return *player.GetUserId() == userId
	}).First()
	if player == nil {
		return nil, errors.New("player is not exists")
	}
	if !player.IsHost() {
		return nil, errors.New("only host player can start game")
	}

	game.SetStatus(datamodel.GS_STARTED)

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

	harbors := model.NewHarbors()
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
		player.SetTurnOrder(idx + 1)
		player.SetColor(colors[idx])
		if player.GetTurnOrder() == 1 {
			game.SetPlayerInTurn(player.GetId())
		}
	}

	if err := ci.gameAggregateService.Save(ctx, game); err != nil {
		return nil, err
	}

	return response.NewGameResponse(game), nil
}

func (ci *catanInteractor) LeaveGame(ctx context.Context, gameRequest *request.GameRequest) (*response.GameResponse, error) {
	game, err := ci.gameAggregateService.GetById(ctx, uint(gameRequest.ID))
	if err != nil {
		return nil, err
	}

	userId := uint(0) //todo: userid from context
	player := game.GetPlayers().Filter(func(player *model.Player) bool {
		if player.GetUserId() == nil {
			return false
		}
		return *player.GetUserId() == userId
	}).First()
	if player != nil {
		switch game.GetStatus() {
		case datamodel.GS_WAITING:
			player.Remove()
		case datamodel.GS_STARTED:
			player.Leave()
		}

		if err := ci.gameAggregateService.Save(ctx, game); err != nil {
			return nil, err
		}
	}

	return response.NewGameResponse(game), nil
}

func (ci *catanInteractor) RollDices(ctx context.Context, gameRequest *request.GameRequest) (*response.GameResponse, error) {
	userId := 0

	game, err := ci.gameAggregateService.GetById(ctx, uint(gameRequest.ID))
	if err != nil {
		return nil, err
	}

	player := game.GetPlayers().Filter(func(player *model.Player) bool {
		if player.GetUserId() == nil {
			return false
		}
		return *player.GetUserId() == uint(userId)
	}).First()
	if player == nil {
		return nil, errors.New("player is not exists")
	}
	if !player.IsInTurn() {
		return nil, errors.New("player is not in turn")
	}

	if game.GetTurn() == 1 || game.GetTurn() == 2 {
		return nil, errors.New("action is restricted in setup phase")
	}

	dices := game.GetDices()
	if dices.IsRolled() {
		return nil, errors.New("player already rolled dices")
	}

	totalNumber := dices.Roll()
	if totalNumber == 7 {
		robber := game.GetRobber()
		robber.SetStatus(datamodel.RS_MOVE)
	} else {
		terrains := game.GetTerrains().Filter(func(terrain *model.Terrain) bool {
			return terrain.GetNumber() == totalNumber && !terrain.HasRobber()
		})
		for _, terrain := range terrains {
			constructions := terrain.GetAdjacentConstructions().Filter(func(construction *model.Construction) bool {
				return construction.GetPlayerId() != nil
			})
			resourceCards := game.GetResourceCards().Filter(func(resourceCard *model.ResourceCard) bool {
				var resourceType datamodel.ResourceType
				switch terrain.GetType() {
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
				return resourceCard.GetPlayerId() == nil && resourceCard.GetType() == resourceType
			})

			resourceCardsDemand := 0
			for _, construction := range constructions {
				resourceCardsDemand++
				if construction.IsUpgradedCastle() {
					resourceCardsDemand++
				}
			}
			if resourceCardsDemand > len(resourceCards) {
				continue
			}

			resourceCardIdx := 0
			for _, construction := range constructions {
				resourceDispatchQuantity := 1
				if construction.IsUpgradedCastle() {
					resourceDispatchQuantity++
				}

				for i := 1; i <= resourceDispatchQuantity; i++ {
					resourceCards[resourceCardIdx].SetPlayer(construction.GetPlayer())
					resourceCardIdx++
				}
			}
		}
	}

	if err := ci.gameAggregateService.Save(ctx, game); err != nil {
		return nil, err
	}

	return response.NewGameResponse(game), nil
}

func (ci *catanInteractor) BuildRoad(ctx context.Context, roadRequest *request.RoadRequest) (*response.GameResponse, error) {
	userId := 0

	game, err := ci.gameAggregateService.GetById(ctx, uint(roadRequest.GameID))
	if err != nil {
		return nil, err
	}

	player := game.GetPlayers().Filter(func(player *model.Player) bool {
		if player.GetUserId() == nil {
			return false
		}
		return *player.GetUserId() == uint(userId)
	}).First()
	if player == nil {
		return nil, errors.New("player is not exists")
	}
	if !player.IsInTurn() {
		return nil, errors.New("player is not in turn")
	}

	road := game.GetRoads().Filter(func(road *model.Road) bool {
		return road.GetId() == uint(roadRequest.ID)
	}).First()
	if road == nil {
		return nil, errors.New("road is not exists")
	}
	if road.GetPlayerId() != nil {
		return nil, errors.New("road is owned by another player")
	}

	gameTurn := game.GetTurn()
	if gameTurn == 1 || gameTurn == 2 {
		roads := player.GetRoads()

		if len(roads) >= gameTurn {
			return nil, errors.New("player has already build road")
		}
	} else {
		lumberResourceCards := player.GetResourceCards().Filter(func(resourceCard *model.ResourceCard) bool {
			return resourceCard.GetType() == datamodel.RT_LUMBER
		})

		brickResourceCards := player.GetResourceCards().Filter(func(resourceCard *model.ResourceCard) bool {
			return resourceCard.GetType() == datamodel.RT_BRICK
		})

		if len(lumberResourceCards) < 1 || len(brickResourceCards) < 1 {
			return nil, errors.New("insufficient resources")
		}

		lumberResourceCards[0].SetPlayer(nil)
		brickResourceCards[0].SetPlayer(nil)
	}

	road.SetPlayer(player)

	if err := ci.gameAggregateService.Save(ctx, game); err != nil {
		return nil, err
	}

	return response.NewGameResponse(game), nil
}
