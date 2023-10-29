package main

import (
	"log"

	"github.com/nvtrinh2001/chatapp/db"
	"github.com/nvtrinh2001/chatapp/internal/user"
	"github.com/nvtrinh2001/chatapp/router"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Could not initialize database connection: %s", err)
	}

	userRepository := user.NewRespository(dbConn.GetDB())
	userService := user.NewService(userRepository)
	userHandler := user.NewHandler(userService)

	router.InitRouter(userHandler)
	router.Start("0.0.0.0:8080")
}
