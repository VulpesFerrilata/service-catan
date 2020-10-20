package response

type PlayerResponse struct {
	ID        int           `json:"id"`
	GameID    int           `json:"gameId"`
	User      *UserResponse `json:"user"`
	Color     string        `json:"color"`
	TurnOrder int           `json:"turnOrder"`
}
