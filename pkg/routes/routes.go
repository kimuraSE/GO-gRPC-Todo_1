package routes

import (
	"context"
	"errors"
	"gRPC-Todo/pkg/todo"
	"gRPC-Todo/pkg/user"
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
)

func NewServer(uh user.UserServiceServer, th todo.TodoServiceServer) {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_auth.UnaryServerInterceptor(authorize),
		)),
	)

	user.RegisterUserServiceServer(s, uh)
	todo.RegisterTodoServiceServer(s, th)
	log.Println("Server is running on port: 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}

}

func authorize(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "Bearer")
	if err != nil {
		return nil, err
	}
	if token != "test" {
		return nil, errors.New("bad token")
	}
	return ctx, nil
}
