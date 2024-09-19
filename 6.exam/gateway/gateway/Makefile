gen-protoc:
	protoc --go_out=. --go-grpc_out=. ./proto/users/user.proto
	
swag-gen:
  ~/go/bin/swag init -g ./internal/api/router/router.go -o ./swagger/docs force 1