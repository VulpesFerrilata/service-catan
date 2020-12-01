package interactor

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/service"
	"github.com/VulpesFerrilata/catan/internal/usecase/request"
	"github.com/VulpesFerrilata/catan/internal/usecase/response"
	"github.com/VulpesFerrilata/library/pkg/validator"
)

type RoomInteractor interface {
}

func NewRoomInteractor(validate validator.Validate,
	roomService service.RoomService) RoomInteractor {
	return &roomInteractor{
		validate:    validate,
		roomService: roomService,
	}
}

type roomInteractor struct {
	validate    validator.Validate
	roomService service.RoomService
}

func (ri *roomInteractor) GetRoomById(ctx context.Context, roomRequest *request.RoomRequest) (*response.RoomResponse, error) {
	if err := ri.validate.StructPartial(ctx, roomRequest, "id"); err != nil {
		return nil, err
	}

	room, err := ri.roomService.GetRoomById(ctx, uint(roomRequest.ID))
	return response.NewRoomResponse(room), err
}

func (ri *roomInteractor) FindRoomsByStatus(ctx context.Context, roomRequest *request.RoomRequest) (response.RoomsResponse, error) {
	if err := ri.validate.StructPartial(ctx, roomRequest, "status"); err != nil {
		return nil, err
	}

	rooms, err := ri.roomService.FindRoomsByStatus(ctx, roomRequest.Status)
	return response.NewRoomsResponse(rooms), err
}
