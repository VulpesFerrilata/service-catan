package request

type RoomRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
