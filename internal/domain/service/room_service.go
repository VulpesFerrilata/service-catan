package service

import (
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
)

type RoomService interface {
	GetRoomRepository() repository.RoomRepository
}

func NewRoomService(roomRepository repository.RoomRepository) RoomService {
	return &roomService{
		roomRepository: roomRepository,
	}
}

type roomService struct {
	roomRepository repository.RoomRepository
}

func (rs roomService) GetRoomRepository() repository.RoomRepository {
	return rs.roomRepository
}
