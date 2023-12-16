package chat

import (
	"context"

	"github.com/nvtrinh2001/chatapp/proto/chat"
)

type rpcClient struct {
	chat.ChatClient
}

func NewRPCClient(uc chat.ChatClient) *rpcClient {
	return &rpcClient{uc}
}

func (r *rpcClient) CreateRoom(ctx context.Context, req *chat.CreateRoomRequest) (*chat.CreateRoomResponse, error) {
	return r.ChatClient.CreateRoom(ctx, req)
}

func (r *rpcClient) GetRooms(ctx context.Context, req *chat.GetRoomsRequest) (*chat.GetRoomsResponse, error) {
	return r.ChatClient.GetRooms(ctx, req)
}

func (r *rpcClient) GetClients(ctx context.Context, req *chat.GetClientsRequest) (*chat.GetClientsResponse, error) {
	return r.ChatClient.GetClients(ctx, req)
}
