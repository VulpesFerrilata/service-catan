package interactor

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/service"
	"github.com/VulpesFerrilata/catan/internal/usecase/request"
	"github.com/VulpesFerrilata/catan/internal/usecase/response"
	"github.com/VulpesFerrilata/library/pkg/app_error"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
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

func (r *roomInteractor) FindRooms(ctx context.Context, roomRequest *request.RoomRequest) (*response.RoomsResponse, error) {
	if err := r.validate.StructCtx(ctx, roomRequest); err != nil {
		if fieldErrors, ok := errors.Cause(err).(validator.ValidationErrors); ok {
			err = app_error.NewValidationError(fieldErrors)
		}
		return nil, errors.Wrap(err, "interactor.RoomInteractor.GetUserById")
		return nil, err
	}

	count, err := r.roomService.GetRoomRepository().Count(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "interactor.RoomInteractor.FindRooms")
	}

	rooms, err := r.roomService.GetRoomRepository().Find(ctx, roomRequest.Limit, roomRequest.Offset)
	if err != nil {
		return nil, errors.Wrap(err, "interactor.RoomInteractor.FindRooms")
	}

	return response.NewRoomsResponse(count, rooms), err
}
