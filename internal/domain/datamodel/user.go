package model

import "github.com/VulpesFerrilata/grpc/protoc/user"

func NewUser(userPb *user.UserResponse) *User {
	user := new(User)
	user.id = uint(userPb.GetID())
	return user
}

type User struct {
	id       uint
	username string
}

func (u User) GetUsername() string {
	return u.username
}
