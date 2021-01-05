package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"sort"
)

func NewPlayersFromPlayerModels(playerModels []*model.Player) Players {
	players := make(Players, 0)

	for _, playerModel := range playerModels {
		player := NewPlayerFromPlayerModel(playerModel)
		players = append(players, player)
	}

	return players
}

type Players []*Player

type PlayerFilterFunc func(player *Player) bool

func (p Players) Filter(f PlayerFilterFunc) Players {
	var players Players
	for _, player := range p {
		if f(player) {
			players = append(players, player)
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

func (p Players) Any(f PlayerFilterFunc) bool {
	for _, player := range p {
		if f(player) {
			return true
		}
	}
	return false
}

type playerSorter struct {
	players Players
	by      PlayerCompareFunc
}

func (ps playerSorter) Len() int {
	return len(ps.players)
}

func (ps playerSorter) Swap(i, j int) {
	ps.players[i], ps.players[j] = ps.players[j], ps.players[i]
}

func (ps playerSorter) Less(i, j int) bool {
	return ps.by(*ps.players[i], *ps.players[j])
}

type PlayerCompareFunc func(a Player, b Player) bool

func (p Players) Sort(f PlayerCompareFunc) {
	playerSorter := &playerSorter{
		players: p,
		by:      f,
	}
	sort.Sort(playerSorter)
}
