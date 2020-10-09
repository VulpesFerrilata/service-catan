package adapter

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/catan/internal/usecase/form"
)

type CatanAdapter interface {
	ParseUser(ctx context.Context, userForm *form.UserForm) (*model.Player, error)
	ParseGame(ctx context.Context, gameForm *form.GameForm) (*model.Game, error)
}

func NewCatanAdapter() CatanAdapter {
	return &catanAdapter{}
}

type catanAdapter struct {
}
