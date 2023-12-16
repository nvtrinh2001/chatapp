package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	protoUser "github.com/nvtrinh2001/chatapp/proto/user"
	"github.com/nvtrinh2001/chatapp/user-service/db"
	"github.com/nvtrinh2001/chatapp/user-service/internal/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	grpcServer := grpc.NewServer()

	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Could not initialize database connection: %s", err)
	}

	userRepository := user.NewRespository(dbConn.GetDB())

	log := hclog.Default()

	userServer := user.NewUserServer(userRepository, log, protoUser.UnimplementedUserServer{})

	protoUser.RegisterUserServer(grpcServer, userServer)

	// add grpc service to the reflection list when the client wants to achieve info of the server
	// e.g. using grpcurl list
	reflection.Register(grpcServer)

	// create a TCP socket for inbound server connections
	listeningSocket, err := net.Listen("tcp", fmt.Sprintf(":%d", 9093))
	if err != nil {
		log.Error("Unable to create a new listener", "error", err)
		os.Exit(1)
	}

	grpcServer.Serve(listeningSocket)
}
