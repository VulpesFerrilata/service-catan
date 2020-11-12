package interactor

import (
	"context"
	"errors"

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

func (ci catanInteractor) FindRooms(ctx context.Context, roomRequest *request.RoomRequest) (response.RoomsResponse, error) {
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

func (ci catanInteractor) CreateGame(ctx context.Context) (*response.GameResponse, error) {
	userId := 0

	userRequestPb := new(user.UserRequest)
	userRequestPb.ID = int64(userId)
	userPb, err := ci.userService.GetUserById(ctx, userRequestPb)
	if err != nil {
		return nil, err
	}
	user := model.NewUser(userPb)

	game := model.NewGame()
	model.NewPlayer(game, user)

	if err := ci.gameAggregateService.Save(ctx, game); err != nil {
		return nil, err
	}

	return response.NewGameResponse(game), nil
}

func (ci catanInteractor) JoinGame(ctx context.Context, gameRequest *request.GameRequest) (*response.GameResponse, error) {
	userId := 0 //todo: userid from context

	game, err := ci.gameAggregateService.GetById(ctx, uint(gameRequest.ID))
	if err != nil {
		return nil, err
	}

	players := game.GetPlayers()

	player := players.Filter(func(player *model.Player) bool {
		return player.UserID == uint(userId)
	}).First()
	if player != nil {
		return response.NewGameResponse(game), nil
	}

	if game.Status == datamodel.GS_WAITING && len(players) < 4 {
		userRequestPb := new(user.UserRequest)
		userRequestPb.ID = int64(userId)
		userPb, err := ci.userService.GetUserById(ctx, userRequestPb)
		if err != nil {
			return nil, err
		}
		user := model.NewUser(userPb)
		model.NewPlayer(game, user)

		if err := ci.gameAggregateService.Save(ctx, game); err != nil {
			return nil, err
		}
	}

	return response.NewGameResponse(game), nil
}

func (ci catanInteractor) StartGame(ctx context.Context, gameRequest *request.GameRequest) (*response.GameResponse, error) {
	game, err := ci.gameAggregateService.GetById(ctx, uint(gameRequest.ID))
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
		return player.UserID == userId
	}).First()
	if player == nil {
		return nil, errors.New("player is not exists")
	}
	if !player.IsHost() {
		return nil, errors.New("only host player can start game")
	}

	game.Init()

	if err := ci.gameAggregateService.Save(ctx, game); err != nil {
		return nil, err
	}

	return response.NewGameResponse(game), nil
}

func (ci catanInteractor) LeaveGame(ctx context.Context, gameRequest *request.GameRequest) (*response.GameResponse, error) {
	game, err := ci.gameAggregateService.GetById(ctx, uint(gameRequest.ID))
	if err != nil {
		return nil, err
	}

	userId := uint(0) //todo: userid from context
	player := game.GetPlayers().Filter(func(player *model.Player) bool {
		return player.UserID == userId
	}).First()
	if player != nil {
		switch game.Status {
		case datamodel.GS_WAITING:
			player.Remove()
		case datamodel.GS_STARTED:
			player.IsLeft = true
		}

		if err := ci.gameAggregateService.Save(ctx, game); err != nil {
			return nil, err
		}
	}

	return response.NewGameResponse(game), nil
}

func (ci catanInteractor) RollDices(ctx context.Context, gameRequest *request.GameRequest) (*response.GameResponse, error) {
	userId := 0

	game, err := ci.gameAggregateService.GetById(ctx, uint(gameRequest.ID))
	if err != nil {
		return nil, err
	}

	player := game.GetPlayers().Filter(func(player *model.Player) bool {
		return player.UserID == uint(userId)
	}).First()
	if player == nil {
		return nil, errors.New("player is not exists")
	}
	if !player.IsInTurn() {
		return nil, errors.New("player is not in turn")
	}

	player.IsRolledDices = true
	dices := game.GetDices()

	dices.Roll()
	totalNumber := dices.GetTotalNumber()

	if totalNumber == 7 {

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

	if err := ci.gameAggregateService.Save(ctx, game); err != nil {
		return nil, err
	}

	return response.NewGameResponse(game), nil
}
