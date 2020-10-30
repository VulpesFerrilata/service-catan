package model

import "github.com/VulpesFerrilata/grpc/protoc/user"

func NewUser(userPb *user.UserResponse) *User {
	user := new(User)
	user.ID = uint(userPb.GetID())
	return user
}

type User struct {
	ID       uint
	Username string
}
