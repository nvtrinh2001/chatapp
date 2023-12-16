package main

import (
	"github.com/nvtrinh2001/chatapp/internal/chat"
	"github.com/nvtrinh2001/chatapp/internal/user"
	protoChat "github.com/nvtrinh2001/chatapp/proto/chat"
	protoUser "github.com/nvtrinh2001/chatapp/proto/user"
	"github.com/nvtrinh2001/chatapp/router"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	userConn, err := grpc.Dial("localhost:9093", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	defer userConn.Close()

	chatConn, err := grpc.Dial("localhost:9092", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	defer chatConn.Close()

	// create client
	userClient := protoUser.NewUserClient(userConn)
	rpcUserClient := user.NewRPCClient(userClient)

	chatClient := protoChat.NewChatClient(chatConn)
	rpcChatClient := chat.NewRPCClient(chatClient)

	// business
	userService := user.NewService(rpcUserClient)
	chatService := chat.NewService(rpcChatClient)

	// handler
	userHandler := user.NewHandler(userService)
	chatHandler := chat.NewHandler(chatService)

	router.InitRouter(userHandler, chatHandler)
	router.Start("0.0.0.0:8080")
}
