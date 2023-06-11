package routes

import (
	"context"
	"gRPC-Todo/pkg/todo"
	"gRPC-Todo/pkg/user"
	"log"
	"net"
	"google.golang.org/grpc"
)

func NewServer(uh user.UserServiceServer, th todo.TodoServiceServer) {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(myLogging()))

	user.RegisterUserServiceServer(s, uh)
	todo.RegisterTodoServiceServer(s, th)
	log.Println("Server is running on port: 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}

}

func myLogging() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		log.Println("Request:"+req.(string))
		resp, err := handler(ctx, req)
		if err != nil {
			return nil, err
		}
		log.Println("Response:"+resp.(string))
		return resp, err
	}
}
