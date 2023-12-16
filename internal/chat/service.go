package chat

import (
	"context"
	"time"

	"github.com/nvtrinh2001/chatapp/proto/chat"
)

const (
	secretKey = "secret"
)

type service struct {
	ChatServiceClient
	timeout time.Duration
}

func NewService(chatService ChatServiceClient) Service {
	return &service{
		chatService,
		time.Duration(2) * time.Second,
	}
}

func (s *service) CreateRoom(c context.Context, request *chat.CreateRoomRequest) (*chat.CreateRoomResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	r, err := s.ChatServiceClient.CreateRoom(ctx, request)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *service) GetRooms(c context.Context, request *chat.GetRoomsRequest) (*chat.GetRoomsResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	r, err := s.ChatServiceClient.GetRooms(ctx, request)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *service) GetClients(c context.Context, request *chat.GetClientsRequest) (*chat.GetClientsResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	r, err := s.ChatServiceClient.GetClients(ctx, request)
	if err != nil {
		return nil, err
	}

	return r, nil
}
