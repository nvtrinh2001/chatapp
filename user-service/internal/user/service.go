package user

import (
	"context"
	"strconv"

	"github.com/hashicorp/go-hclog"
	"github.com/nvtrinh2001/chatapp/proto/user"
)

type UserServer struct {
	Repository
	l hclog.Logger
	user.UnimplementedUserServer
}

func NewUserServer(repo Repository, l hclog.Logger, unimplementedServer user.UnimplementedUserServer) *UserServer {
	return &UserServer{repo, l, unimplementedServer}
}

func (s *UserServer) GetUser(ctx context.Context, request *user.GetUserRequest) (*user.GetUserResponse, error) {
	s.l.Info("Handle request for GetUser", "id", request.GetId())
	user := user.GetUserResponse{
		Id:       "1",
		Username: "nvt",
		Email:    "nvt@gmail.com",
	}
	return &user, nil
}

func (s *UserServer) CreateUser(c context.Context, request *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	u := &User{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}

	res, err := s.Repository.CreateUser(c, u)
	if err != nil {
		return nil, err
	}

	response := &user.CreateUserResponse{
		Id:       strconv.Itoa(int(res.ID)),
		Email:    res.Email,
		Username: res.Username,
	}
	return response, nil
}

func (s *UserServer) GetUserByEmail(c context.Context, request *user.GetUserByEmailRequest) (*user.GetUserByEmailResponse, error) {
	res, err := s.Repository.GetUserByEmail(c, request.Email)
	if err != nil {
		return nil, err
	}

	response := &user.GetUserByEmailResponse{
		Id:       strconv.Itoa(int(res.ID)),
		Email:    res.Email,
		Username: res.Username,
		Password: res.Password,
	}

	return response, nil
}

func (s *UserServer) ChangeUsername(c context.Context, request *user.ChangeUsernameRequest) (*user.ChangeUsernameResponse, error) {
	res, err := s.Repository.ChangeUsername(c, request.Id, request.NewUsername)
	if err != nil {
		return nil, err
	}

	response := &user.ChangeUsernameResponse{
		Id:       strconv.Itoa(int(res.ID)),
		Email:    res.Email,
		Username: res.Username,
	}

	return response, nil
}
