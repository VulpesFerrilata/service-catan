package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/grpc/protoc/user"
)

type PlayerAggregateService interface {
	FindByGameId(ctx context.Context, gameId uint) (model.Players, error)
	Save(ctx context.Context, player *model.Player) error
}

type playerAggregateService struct {
	playerService          PlayerService
	userService            user.UserService
	achievementService     AchievementService
	resourceCardService    ResourceCardService
	developmentCardService DevelopmentCardService
}

func (pas *playerAggregateService) FindByGameId(ctx context.Context, gameId uint) (model.Players, error) {
	players, err := pas.playerService.GetPlayerRepository().FindByGameId(ctx, gameId)
	if err != nil {
		return nil, err
	}

	for _, player := range players {
		userId := player.GetUserId()
		if userId != nil {
			userRequest := new(user.UserRequest)
			userRequest.ID = int64(*userId)
			userPb, err := pas.userService.GetUserById(ctx, userRequest)
			if err != nil {
				return nil, err
			}
			user := model.NewUser(userPb)
			player.SetUser(user)
		}
	}

	return players, nil
}

func (pas *playerAggregateService) Save(ctx context.Context, player *model.Player) error {
	return pas.playerService.Save(ctx, player)
}
