// package main
//
// import (
// 	"log"
//
// 	"github.com/nvtrinh2001/chatapp/router"
// 	"github.com/nvtrinh2001/chatapp/user-service/db"
// 	"github.com/nvtrinh2001/chatapp/user-service/internal/user"
// 	"github.com/nvtrinh2001/chatapp/user-service/internal/ws"
// )
//
// func main() {
// 	dbConn, err := db.NewDatabase()
// 	if err != nil {
// 		log.Fatalf("Could not initialize database connection: %s", err)
// 	}
//
// 	userRepository := user.NewRespository(dbConn.GetDB())
// 	userService := user.NewService(userRepository)
// 	userHandler := user.NewHandler(userService)
//
// 	hub := ws.NewHub()
// 	wsHandler := ws.NewHandler(hub)
// 	go hub.Run()
//
// 	router.InitRouter(userHandler, wsHandler)
// 	router.Start("0.0.0.0:8080")
// }
