package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/grpc/protoc/user"
	"github.com/VulpesFerrilata/library/pkg/app_error"
	"github.com/VulpesFerrilata/library/pkg/middleware"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"
)

type PlayerRepository interface {
	FindByGameId(ctx context.Context, gameId uuid.UUID) (datamodel.Players, error)
	GetByUserId(ctx context.Context, userId uuid.UUID) (*datamodel.Player, error)
	InsertOrUpdate(ctx context.Context, player *datamodel.Player) error
}

func NewPlayerRepository(transactionMiddleware *middleware.TransactionMiddleware,
	validate *validator.Validate,
	userService user.UserService) PlayerRepository {
	return &playerRepository{
		transactionMiddleware: transactionMiddleware,
		validate:              validate,
		userService:           userService,
	}
}

type playerRepository struct {
	transactionMiddleware *middleware.TransactionMiddleware
	validate              *validator.Validate
	userService           user.UserService
}

func (p playerRepository) FindByGameId(ctx context.Context, gameId uuid.UUID) (datamodel.Players, error) {
	playerModels := make([]*model.Player, 0)
	err := p.transactionMiddleware.Get(ctx).Find(&playerModels, "game_id = ?", gameId).Error
	if err != nil {
		return nil, errors.Wrap(err, "repository.PlayerRepository.FindByGameId")
	}

	players := make(datamodel.Players, 0)
	for _, playerModel := range playerModels {
		userRequestPb := new(user.UserRequest)
		userRequestPb.ID = playerModel.UserID.String()
		userResponsePb, err := p.userService.GetUserById(ctx, userRequestPb)
		if err != nil {
			return nil, errors.Wrap(err, "repository.PlayerRepository.FindByGameId")
		}

		user, err := datamodel.NewUserFromUserPb(userResponsePb)
		if err != nil {
			return nil, errors.Wrap(err, "repository.PlayerRepository.FindByGameId")
		}

		player, err := datamodel.NewPlayerFromModel(playerModel, user)
		if err != nil {
			return nil, errors.Wrap(err, "repository.PlayerRepository.FindByGameId")
		}

		players = append(players, player)
	}

	return players, errors.Wrap(err, "repository.PlayerRepository.FindByGameId")
}

func (p playerRepository) GetByUserId(ctx context.Context, userId uuid.UUID) (*datamodel.Player, error) {
	playerModel := new(model.Player)

	err := p.transactionMiddleware.Get(ctx).First(playerModel, "user_id = ?", userId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, app_error.NewNotFoundError("player")
	}
	if err != nil {
		return nil, errors.Wrap(err, "repository.PlayerRepository.GetByUserId")
	}

	userRequestPb := new(user.UserRequest)
	userRequestPb.ID = playerModel.UserID.String()
	userResponsePb, err := p.userService.GetUserById(ctx, userRequestPb)
	if err != nil {
		return nil, errors.Wrap(err, "repository.PlayerRepository.GetByUserId")
	}

	user, err := datamodel.NewUserFromUserPb(userResponsePb)
	if err != nil {
		return nil, errors.Wrap(err, "repository.PlayerRepository.FindByGameId")
	}

	player, err := datamodel.NewPlayerFromModel(playerModel, user)
	if err != nil {
		return nil, errors.Wrap(err, "repository.PlayerRepository.GetByUserId")
	}

	return player, errors.Wrap(err, "repository.PlayerRepository.GetByUserId")
}

func (p playerRepository) InsertOrUpdate(ctx context.Context, player *datamodel.Player) error {
	playerModel := player.ToModel()

	if err := p.validate.StructCtx(ctx, playerModel); err != nil {
		return errors.Wrap(err, "repository.PlayerRepository.InsertOrUpdate")
	}

	err := p.transactionMiddleware.Get(ctx).Save(playerModel).Error
	return errors.Wrap(err, "repository.PlayerRepository.InsertOrUpdate")
}
