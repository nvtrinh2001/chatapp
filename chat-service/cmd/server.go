package main

import (
	"fmt"
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/nvtrinh2001/chatapp/chat-service/internal/chat"
	protoChat "github.com/nvtrinh2001/chatapp/proto/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	grpcServer := grpc.NewServer()

	log := hclog.Default()

	hub := chat.NewHub()
	repo := chat.NewRespository(hub)

	chatServer := chat.NewChatServer(repo, log, protoChat.UnimplementedChatServer{})

	protoChat.RegisterChatServer(grpcServer, chatServer)

	// add grpc service to the reflection list when the client wants to achieve info of the server
	// e.g. using grpcurl list
	reflection.Register(grpcServer)

	// create a TCP socket for inbound server connections
	listeningSocket, err := net.Listen("tcp", fmt.Sprintf(":%d", 9092))
	if err != nil {
		log.Error("Unable to create a new listener", "error", err)
		os.Exit(1)
	}

	grpcServer.Serve(listeningSocket)

	// 	hub := ws.NewHub()
	// 	wsHandler := ws.NewHandler(hub)
	// 	go hub.Run()
	//
	// 	router.InitRouter(userHandler, wsHandler)
	// 	router.Start("0.0.0.0:8080")

}
