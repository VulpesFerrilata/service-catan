package cache

type Game struct {
	ID               uint               `json:"id"`
	Players          []*Player          `json:"players"`
	Achievements     []*Achievement     `json:"achievements"`
	Constructions    []*Construction    `json:"constructions"`
	DevelopmentCards []*DevelopmentCard `json:"developmentCards"`
}
