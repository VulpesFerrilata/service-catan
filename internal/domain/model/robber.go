package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

func NewRobber(terrains Terrains) *Robber {
	robber := new(Robber)
	robber.robber = new(datamodel.Robber)
	robber.robber.Status = datamodel.RS_IDLE

	desertField := terrains.Filter(func(terrain *Terrain) bool {
		if terrain.GetType() == datamodel.TT_DESERT {
			return true
		}
		return false
	}).First()
	robber.robber.Q = desertField.GetQ()
	robber.robber.R = desertField.GetR()

	return robber
}

type Robber struct {
	robber     *datamodel.Robber
	game       *Game
	isModified bool
}

func (r *Robber) GetRobber() datamodel.Robber {
	return *r.robber
}

func (r *Robber) GetId() uint {
	return r.robber.ID
}

func (r *Robber) GetGameId() *uint {
	return r.robber.GameID
}

func (r *Robber) setGame(game *Game) {
	if game != nil {
		r.robber.GameID = &game.game.ID
		r.game = game
	}
}

func (r *Robber) GetQ() int {
	return r.robber.Q
}

func (r *Robber) GetR() int {
	return r.robber.R
}

func (r *Robber) GetStatus() datamodel.RobberStatus {
	return r.robber.Status
}

func (r *Robber) SetStatus(robberStatus datamodel.RobberStatus) {
	if r.robber.Status != robberStatus {
		r.robber.Status = robberStatus
		r.isModified = true
	}
}

func (r *Robber) IsModified() bool {
	return r.isModified
}
