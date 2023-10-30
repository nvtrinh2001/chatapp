package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nvtrinh2001/chatapp/internal/user"
	"github.com/nvtrinh2001/chatapp/internal/ws"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler, wsHandler *ws.Handler) {
	r = gin.Default()

	// http - user
	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.GET("/logout", userHandler.Logout)

	// http - chatroom
	r.POST("/ws/rooms", wsHandler.CreateRoom)
	r.GET("/ws/rooms", wsHandler.GetRooms)
	r.GET("/ws/clients/:roomId", wsHandler.GetClients)

	// websocket
	r.GET("/ws/joinRoom/:roomId", wsHandler.JoinRoom)
}

func Start(addr string) error {
	return r.Run(addr)
}
