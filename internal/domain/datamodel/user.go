package datamodel

import (
	"github.com/VulpesFerrilata/grpc/protoc/user"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func NewUserFromUserPb(userPb *user.UserResponse) (*User, error) {
	user := new(User)

	id, err := uuid.Parse(userPb.GetID())
	if err != nil {
		return nil, errors.Wrap(err, "datamodel.NewUserFromUserPb")
	}
	user.id = id
	user.username = userPb.GetUsername()
	user.displayName = userPb.GetDisplayName()
	user.email = userPb.GetEmail()

	return user, nil
}

type User struct {
	id          uuid.UUID
	username    string
	displayName string
	email       string
}

func (u User) GetId() uuid.UUID {
	return u.id
}

func (u User) GetUsername() string {
	return u.username
}

func (u User) GetDisplayName() string {
	return u.displayName
}

func (u User) GetEmail() string {
	return u.email
}
