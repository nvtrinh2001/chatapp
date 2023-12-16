proto:
	protoc proto/user.proto --go_out=. --go-grpc_out=.
	protoc proto/chat.proto --go_out=. --go-grpc_out=.

list-services:
	grpcurl --plaintext localhost:9092 list

list-user-methods:
	grpcurl --plaintext localhost:9092 list User        

get-user-request-details:
	grpcurl --plaintext localhost:9092 describe .GetUserRequest

example-execute-get-user-request:
	grpcurl --plaintext localhost:9092 User/GetUser

.PHONY: proto
