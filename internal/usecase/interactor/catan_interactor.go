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
	gameService service.GameService,
	playerService service.PlayerService,
	userService user.UserService) CatanInteractor {
	return &catanInteractor{
		validate:      validate,
		roomService:   roomService,
		gameService:   gameService,
		playerService: playerService,
		userService:   userService,
	}
}

type catanInteractor struct {
	validate      validator.Validate
	roomService   service.RoomService
	gameService   service.GameService
	playerService service.PlayerService
	userService   user.UserService
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

	game := model.NewGame()

	player := model.NewPlayer()
	game.AddPlayer(player)

	user := model.NewUser(userPb)
	player.SetUser(user)

	if err := ci.gameService.Save(ctx, game); err != nil {
		return nil, err
	}

	return response.NewGameResponse(game), nil
}

func (ci catanInteractor) JoinGame(ctx context.Context, gameRequest *request.GameRequest) (*response.GameResponse, error) {
	game, err := ci.gameService.GetById(ctx, uint(gameRequest.ID))
	if err != nil {
		return nil, err
	}

	switch game.Status {
	case datamodel.GS_STARTED:
		return nil, errors.New("game has already started")
	case datamodel.GS_FINISHED:
		return nil, errors.New("game was finished")
	}

	player := model.NewPlayer()
	game.AddPlayer(player)

	userId := 0 //todo: userid from context

	userRequestPb := new(user.UserRequest)
	userRequestPb.ID = int64(userId)
	userPb, err := ci.userService.GetUserById(ctx, userRequestPb)
	if err != nil {
		return nil, err
	}
	user := model.NewUser(userPb)
	player.SetUser(user)

	if err := ci.gameService.Save(ctx, game); err != nil {
		return nil, err
	}

	return response.NewGameResponse(game), nil
}

func (ci catanInteractor) StartGame(ctx context.Context, gameRequest *request.GameRequest) (*response.GameResponse, error) {
	game, err := ci.gameService.GetById(ctx, uint(gameRequest.ID))
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
	player, err := game.GetPlayerByUserId(userId)
	if err != nil {
		return nil, err
	}
	if !player.IsHost() {
		return nil, errors.New("only host player can start game")
	}

	game.Init()

	if err := ci.gameService.Save(ctx, game); err != nil {
		return nil, err
	}

	return response.NewGameResponse(game), nil
}

func (ci catanInteractor) LeaveGame(ctx context.Context, gameRequest *request.GameRequest) (*response.GameResponse, error) {
	game, err := ci.gameService.GetById(ctx, uint(gameRequest.ID))
	if err != nil {
		return nil, err
	}

	userId := uint(0) //todo: userid from context
	player, err := game.GetPlayerByUserId(userId)
	if err != nil {
		return nil, errors.New("")
	}

	switch game.Status {
	case datamodel.GS_WAITING:
		player.IsDeleted = true

		if game.GetPlayerQuantity() != 0 {
			game.SwitchHost()
		} else {
			game.IsDeleted = true
		}
	case datamodel.GS_STARTED:
		player.IsLeft = true
	}

	if err := ci.gameService.Save(ctx, game); err != nil {
		return nil, err
	}

	return response.NewGameResponse(game), nil
}
