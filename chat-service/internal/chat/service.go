package chat

import (
	"context"

	"github.com/hashicorp/go-hclog"
	"github.com/nvtrinh2001/chatapp/proto/chat"
)

type ChatServer struct {
	Repository
	l hclog.Logger
	chat.UnimplementedChatServer
}

func NewChatServer(repo Repository, l hclog.Logger, unimplementedServer chat.UnimplementedChatServer) *ChatServer {
	return &ChatServer{repo, l, unimplementedServer}
}

func (s *ChatServer) CreateRoom(ctx context.Context, request *chat.CreateRoomRequest) (*chat.CreateRoomResponse, error) {
	room := &Room{
		Name: request.Name,
	}

	res, err := s.Repository.CreateRoom(ctx, room)
	if err != nil {
		return nil, err
	}

	response := &chat.CreateRoomResponse{
		Id:   res.ID,
		Name: res.Name,
	}

	return response, nil
}

func (s *ChatServer) GetRooms(ctx context.Context, request *chat.GetRoomsRequest) (*chat.GetRoomsResponse, error) {
	res, err := s.Repository.GetRooms(ctx)
	if err != nil {
		return nil, err
	}

	rooms := make([]*chat.Room, len(res))

	for i, room := range res {
		// Create a new Room object for each iteration
		rooms[i] = &chat.Room{
			Id:   room.ID,
			Name: room.Name,
		}
	}

	response := &chat.GetRoomsResponse{Rooms: rooms}

	return response, nil
}

func (s *ChatServer) GetClients(ctx context.Context, request *chat.GetClientsRequest) (*chat.GetClientsResponse, error) {
	res, err := s.Repository.GetClients(ctx, request.RoomId)
	if err != nil {
		return nil, err
	}

	clients := make([]*chat.Client, len(res))

	for i, client := range res {
		// Create a new Room object for each iteration
		clients[i] = &chat.Client{
			Id:       client.ID,
			Username: client.Username,
		}
	}

	response := &chat.GetClientsResponse{Clients: clients}

	return response, nil
}
