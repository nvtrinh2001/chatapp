package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nvtrinh2001/chatapp/internal/chat"
	"github.com/nvtrinh2001/chatapp/internal/user"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler, chatHandler *chat.Handler) {
	r = gin.Default()

	// http - user
	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.GET("/logout", userHandler.Logout)
	r.PATCH("/users/:userId", userHandler.ChangeUsername)

	// http - chatroom
	r.POST("/ws/rooms", chatHandler.CreateRoom)
	r.GET("/ws/rooms", chatHandler.GetRooms)
	r.GET("/ws/clients/:roomId", chatHandler.GetClients)

	// // websocket
	// r.GET("/ws/joinRoom/:roomId", chatHandler.JoinRoom)
}

func Start(addr string) error {
	return r.Run(addr)
}
