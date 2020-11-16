package request

type RoadRequest struct {
	ID     int `json:"id" validate:"required"`
	GameID int `json:"gameId" validate:"required"`
}
