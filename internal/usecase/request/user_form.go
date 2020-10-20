package request

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

type UserForm struct {
	ID          int
	Status      datamodel.GameStatus
	PlayerCount int
}
