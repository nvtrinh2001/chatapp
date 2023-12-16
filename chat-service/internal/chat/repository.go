package chat

import (
	"context"
	"strconv"
)

type repository struct {
	hub *Hub
}

func NewHub() *Hub {

	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message),
	}
}

func NewRespository(hub *Hub) Repository {
	return &repository{hub: hub}
}

func (r *repository) CreateRoom(c context.Context, room *Room) (*Room, error) {
	numRooms := len(r.hub.Rooms)

	room.ID = strconv.Itoa(numRooms)
	r.hub.Rooms[room.ID] = room

	return r.hub.Rooms[room.ID], nil
}

func (r *repository) GetRooms(c context.Context) ([]*Room, error) {
	rooms := make([]*Room, 0)
	for _, room := range r.hub.Rooms {
		rooms = append(rooms, room)
	}

	return rooms, nil
}

func (r *repository) GetClients(c context.Context, roomId string) ([]*Client, error) {
	if len(r.hub.Rooms) == 0 {
		return nil, nil
	}

	room := r.hub.Rooms[roomId]

	clients := make([]*Client, 0)
	for _, client := range room.Clients {
		clients = append(clients, client)
	}

	return clients, nil
}
