package model

type Players []*Player

func (p *Players) append(player *Player) {
	*p = append(*p, player)
}

type PlayerFilterFunc func(player *Player) bool

func (p Players) Filter(playerFilterFunc PlayerFilterFunc) Players {
	var players Players
	for _, player := range p {
		if playerFilterFunc(player) {
			players.append(player)
		}
	}
	return players
}

func (p Players) First() *Player {
	if len(p) > 0 {
		return p[0]
	}
	return nil
}
