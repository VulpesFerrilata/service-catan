package response

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

type GameResponse struct {
	ID           int                  `json:"id"`
	Host         int                  `json:"host"`
	PlayerInTurn int                  `json:"playerInTurn"`
	Status       datamodel.GameStatus `json:"status"`
	Players      []*PlayerResponse    `json:"players"`
}
