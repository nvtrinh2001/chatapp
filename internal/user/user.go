package user

import (
	"context"

	"github.com/nvtrinh2001/chatapp/proto/user"
)

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type ChangeUsernameRequest struct {
	ID          string `json:"id"`
	NewUsername string `json:"new-username"`
}

type ChangeUsernameResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	accessToken string
	ID          string `json:"id"`
	Username    string `json:"username"`
}

type UserServiceClient interface {
	CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error)
	GetUserByEmail(ctx context.Context, req *user.GetUserByEmailRequest) (*user.GetUserByEmailResponse, error)
	ChangeUsername(ctx context.Context, req *user.ChangeUsernameRequest) (*user.ChangeUsernameResponse, error)
}

type Service interface {
	CreateUser(c context.Context, request *CreateUserRequest) (*CreateUserResponse, error)
	Login(c context.Context, request *LoginUserRequest) (*LoginUserResponse, error)
	ChangeUsername(c context.Context, request *ChangeUsernameRequest) (*ChangeUsernameResponse, error)
}
