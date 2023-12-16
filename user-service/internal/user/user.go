package user

import (
	"context"

	"github.com/nvtrinh2001/chatapp/proto/user"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Repository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	ChangeUsername(ctx context.Context, id string, newUsername string) (*User, error)
}

type Service interface {
	CreateUser(c context.Context, request *user.CreateUserRequest) (*user.CreateUserResponse, error)
	GetUserByEmail(c context.Context, request *user.GetUserByEmailRequest) (*user.GetUserByEmailResponse, error)
	ChangeUsername(c context.Context, request *user.ChangeUsernameRequest) (*user.ChangeUsernameResponse, error)
}
