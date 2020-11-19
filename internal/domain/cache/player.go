package cache

type Player struct {
	ID               int                `json:"id"`
	User             *User              `json:"user"`
	Achievements     []*Achievement     `json:"achievements"`
	Constructions    []*Construction    `json:"constructions"`
	DevelopmentCards []*DevelopmentCard `json:"developmentCards"`
	Color            string             `json:"color"`
	TurnOrder        int                `json:"turnOrder"`
	IsLeft           bool               `json:"isLeft"`
}
