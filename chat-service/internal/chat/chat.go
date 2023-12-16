package chat

import (
	"context"

	"github.com/gorilla/websocket"
	"github.com/nvtrinh2001/chatapp/proto/chat"
)

// client
type Client struct {
	Conn     *websocket.Conn
	Message  chan *Message
	ID       string `json:"id"`
	RoomID   string `json:"roomId"`
	Username string `json:"username"`
}

type Message struct {
	Content  string `json:"content"`
	RoomID   string `json:"roomId"`
	Username string `json:"username"`
}

// hub
type Hub struct {
	// map room id to a room object
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

type Room struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

type Repository interface {
	CreateRoom(c context.Context, room *Room) (*Room, error)
	GetRooms(c context.Context) ([]*Room, error)
	GetClients(c context.Context, roomId string) ([]*Client, error)
}

type Service interface {
	CreateRoom(c context.Context, request *chat.CreateRoomRequest) (*chat.CreateRoomResponse, error)
	GetRooms(c context.Context, request *chat.GetRoomsRequest) (*chat.GetRoomsResponse, error)
	GetClients(c context.Context, request *chat.GetClientsRequest) (*chat.GetClientsResponse, error)
}
