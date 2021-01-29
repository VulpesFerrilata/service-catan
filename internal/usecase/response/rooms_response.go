package response

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

func NewRoomsResponse(count int, rooms datamodel.Rooms) *RoomsResponse {
	roomsResponse := new(RoomsResponse)
	roomsResponse.Count = count

	for _, room := range rooms {
		roomResponse := NewRoomResponse(room)
		roomsResponse.Data = append(roomsResponse.Data, roomResponse)
	}

	return roomsResponse
}

type RoomsResponse struct {
	Count int             `json:"count"`
	Rooms []*RoomResponse `json:"rooms"`
}
