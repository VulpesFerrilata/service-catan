package model

type Players []*Player

func (p Players) Append(player *Player) {
	p = append(p, player)
}

func (p Players) Remove() {
	for _, player := range p {
		player.Remove()
	}
}

func (p Players) First() *Player {
	if len(p) > 0 {
		return p[0]
	}
	return nil
}
