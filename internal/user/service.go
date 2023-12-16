package user

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nvtrinh2001/chatapp/proto/user"
	"github.com/nvtrinh2001/chatapp/user-service/pkg/utils"
)

const (
	secretKey = "secret"
)

type service struct {
	UserServiceClient
	timeout time.Duration
}

func NewService(userService UserServiceClient) Service {
	return &service{
		userService,
		time.Duration(2) * time.Second,
	}
}

func (s *service) CreateUser(c context.Context, request *CreateUserRequest) (*CreateUserResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	// hash password
	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	user := &user.CreateUserRequest{
		Username: request.Username,
		Email:    request.Email,
		Password: hashedPassword,
	}

	u, err := s.UserServiceClient.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	response := &CreateUserResponse{
		ID:       u.Id,
		Username: u.Username,
		Email:    u.Email,
	}

	return response, nil
}

func (s *service) ChangeUsername(c context.Context, request *ChangeUsernameRequest) (*ChangeUsernameResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	req := &user.ChangeUsernameRequest{
		Id:          request.ID,
		NewUsername: request.NewUsername,
	}

	u, err := s.UserServiceClient.ChangeUsername(ctx, req)
	if err != nil {
		return nil, err
	}

	response := &ChangeUsernameResponse{
		ID:       u.Id,
		Username: u.Username,
		Email:    u.Email,
	}

	return response, nil
}

type MyJWTClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (s *service) Login(c context.Context, request *LoginUserRequest) (*LoginUserResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	req := &user.GetUserByEmailRequest{Email: request.Email}

	user, err := s.UserServiceClient.GetUserByEmail(ctx, req)
	if err != nil {
		return &LoginUserResponse{}, err
	}

	err = utils.CheckPassword(request.Password, user.Password)
	if err != nil {
		return &LoginUserResponse{}, err
	}

	// custom claims can be used for permissions and additional information of user
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		ID:       user.Id,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    user.Id,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	secretString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return &LoginUserResponse{}, err
	}

	return &LoginUserResponse{
		accessToken: secretString,
		Username:    user.Username,
		ID:          user.Id,
	}, nil
}
