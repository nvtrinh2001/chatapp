package user

import (
	"context"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nvtrinh2001/chatapp/pkg/utils"
)

const (
	secretKey = "secret"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		repository,
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

	user := &User{
		Username: request.Username,
		Email:    request.Email,
		Password: hashedPassword,
	}

	u, err := s.Repository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	response := &CreateUserResponse{
		ID:       strconv.Itoa(int(u.ID)),
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

	user, err := s.Repository.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return &LoginUserResponse{}, err
	}

	err = utils.CheckPassword(request.Password, user.Password)
	if err != nil {
		return &LoginUserResponse{}, err
	}

	// custom claims can be used for permissions and additional information of user
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		ID:       strconv.Itoa(int(user.ID)),
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(user.ID)),
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
		ID:          strconv.Itoa(int(user.ID)),
	}, nil
}
