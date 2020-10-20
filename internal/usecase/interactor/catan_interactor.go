package interactor

import (
	"context"

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
	gameService service.GameService,
	roomService service.RoomService,
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

func (ci catanInteractor) FindRooms(ctx context.Context, roomRequest *request.RoomRequest) ([]*response.RoomResponse, error) {
	if err := ci.validate.Struct(ctx, roomRequest); err != nil {
		return nil, err
	}

	if roomRequest.ID != 0 {
		room, err := ci.roomService.GetRoomById(ctx, uint(roomRequest.ID))
		if err != nil {
			return nil, err
		}

		return response.NewRoomResponses(room), nil
	}

	rooms, err := ci.roomService.FindRoomsByStatus(ctx, roomRequest.Status)
	if err != nil {
		return nil, err
	}
	return response.NewRoomResponses(rooms...), nil
}

func (ci catanInteractor) CreateGame(ctx context.Context) (*response.GameResponse, error) {
	game, err := ci.gameService.Create(ctx)
	if err != nil {
		return nil, err
	}

	userId := 0
	player := model.NewPlayer(uint(userId))
	game.AddPlayer(player)

	if err := ci.gameService.Save(ctx, game); err != nil {
		return nil, err
	}

	return nil, nil
}
