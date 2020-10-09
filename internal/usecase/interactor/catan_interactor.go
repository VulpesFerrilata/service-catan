package interactor

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/catan/internal/domain/service"
	"github.com/VulpesFerrilata/catan/internal/usecase/adapter"
	"github.com/VulpesFerrilata/catan/internal/usecase/form"
)

type CatanInteractor interface {
	NewGame(ctx context.Context, userForm *form.UserForm) error
	JoinGame(ctx context.Context, userForm form.UserForm) error
	StartGame(ctx context.Context, userForm form.UserForm) error
	LeaveGame(ctx context.Context, userForm form.UserForm) error
}

func NewCatanInteractor(catanAdapter adapter.CatanAdapter,
	gameService service.GameService,
	playerService service.PlayerService) CatanInteractor {
	return &catanInteractor{
		catanAdapter:  catanAdapter,
		gameService:   gameService,
		playerService: playerService,
	}
}

type catanInteractor struct {
	catanAdapter  adapter.CatanAdapter
	gameService   service.GameService
	playerService service.PlayerService
}

func (ci catanInteractor) NewGame(ctx context.Context, userForm *form.UserForm) error {
	player, err := ci.catanAdapter.ParseUser(ctx, userForm)
	if err != nil {
		return err
	}

	game, err := ci.gameService.New(ctx)
	if err != nil {
		return err
	}
	player.GameID = game.ID

	if err := ci.playerService.Create(ctx, player); err != nil {
		return err
	}

	players, err := ci.playerService.GetPlayerRepository().FindByGameId(ctx, game.ID)
	if err != nil {
		return err
	}

	return nil
}

func (ci catanInteractor) JoinGame(ctx context.Context, userForm *form.UserForm, gameForm *form.GameForm) error {
	player, err := ci.catanAdapter.ParseUser(ctx, userForm)
	if err != nil {
		return err
	}

	game, err := ci.catanAdapter.ParseGame(ctx, gameForm)
	if err != nil {
		return err
	}

	game, err = ci.gameService.GetGameRepository().GetById(ctx, game.ID)
	if err != nil {
		return err
	}
	player.GameID = game.ID

	isExists, err := ci.playerService.IsExists(ctx, player)
	if err != nil {
		return err
	}
	if !isExists {
		switch game.Status {
		case model.GS_WAITED:
			if err := ci.playerService.Create(ctx, player); err != nil {
				return err
			}
		case model.GS_STARTED:
		case model.GS_FINISHED:
		}
	}

	players, err = ci.playerService.GetPlayerRepository().FindByGameId(ctx, game.ID)
	if err != nil {
		return err
	}

	return nil
}
