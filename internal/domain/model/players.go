package model

type Players []*Player

func (p Players) append(player *Player) {
	p = append(p, player)
}

func (p Players) SetGame(game *Game) {
	for _, player := range p {
		player.SetGame(game)
	}
}

func (p Players) GetByUserId(userId uint) *Player {
	for _, player := range p {
		if player.UserID == userId {
			return player
		}
	}
	return nil
}
