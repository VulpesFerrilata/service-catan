package datamodel

import (
	"fmt"
)

func NewGameStatus(value string) (gameStatus, error) {
	status := gameStatus(value)

	if _, ok := gameStatuses[status]; ok {
		return status, nil
	}

	return status, fmt.Errorf("game status is invalid: %s", value)
}

type gameStatus string

func (g gameStatus) String() string {
	return string(g)
}

const (
	Waiting  gameStatus = "Waiting"
	Started  gameStatus = "Started"
	Finished gameStatus = "Finished"
)

var gameStatuses = map[gameStatus]struct{}{
	Waiting:  {},
	Started:  {},
	Finished: {},
}
