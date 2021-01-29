package response

type GameResponse struct {
	ID           int               `json:"id"`
	Host         int               `json:"host"`
	PlayerInTurn int               `json:"playerInTurn"`
	Status       string            `json:"status"`
	Players      []*PlayerResponse `json:"players"`
}
