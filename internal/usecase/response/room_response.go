package response

import (
	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
)

func NewRoomResponse(room *datamodel.Room) *RoomResponse {
	roomResponse := new(RoomResponse)
	roomResponse.ID = room.GetId().String()
	roomResponse.Status = room.GetStatus().String()
	roomResponse.PlayerCount = len(room.GetPlayers())
	return roomResponse
}

type RoomResponse struct {
	ID          string `json:"id"`
	Status      string `json:"status"`
	PlayerCount int    `json:"playerCount"`
}
