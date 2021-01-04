package datamodel

import "github.com/VulpesFerrilata/grpc/protoc/user"

func NewUserFromUserPb(userPb *user.UserResponse) *User {
	user := new(User)
	user.id = int(userPb.GetID())
	return user
}

type User struct {
	id       int
	username string
}

func (u User) GetUsername() string {
	return u.username
}
