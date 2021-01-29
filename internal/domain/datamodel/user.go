package datamodel

import (
	"fmt"

	"github.com/VulpesFerrilata/grpc/protoc/user"
	"github.com/google/uuid"
)

func NewUserFromUserPb(userPb *user.UserResponse) (*User, error) {
	user := new(User)

	id, err := uuid.Parse(userPb.GetID())
	if err != nil {
		return nil, fmt.Errorf("user's id is invalid: %s", userPb.GetID())
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
