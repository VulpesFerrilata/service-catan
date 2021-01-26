package datamodel

import (
	"fmt"
)

func NewGameStatus(value string) (GameStatus, error) {
	gameStatus := GameStatus(value)

	if _, ok := gameStatuses[gameStatus]; ok {
		return gameStatus, nil
	}

	return gameStatus, fmt.Errorf("game status is invalid: %s", value)
}

type GameStatus string

func (g GameStatus) String() string {
	return string(g)
}

const (
	Waiting  GameStatus = "Waiting"
	Started  GameStatus = "Started"
	Finished GameStatus = "Finished"
)

var gameStatuses = map[GameStatus]struct{}{
	Waiting:  {},
	Started:  {},
	Finished: {},
}
