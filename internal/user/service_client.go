package user

import (
	"context"

	"github.com/nvtrinh2001/chatapp/proto/user"
)

type rpcClient struct {
	user.UserClient
}

func NewRPCClient(uc user.UserClient) *rpcClient {
	return &rpcClient{uc}
}

func (r *rpcClient) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	return r.UserClient.CreateUser(ctx, req)
}

func (r *rpcClient) GetUserByEmail(ctx context.Context, req *user.GetUserByEmailRequest) (*user.GetUserByEmailResponse, error) {
	return r.UserClient.GetUserByEmail(ctx, req)
}

func (r *rpcClient) ChangeUsername(ctx context.Context, req *user.ChangeUsernameRequest) (*user.ChangeUsernameResponse, error) {
	return r.UserClient.ChangeUsername(ctx, req)
}
