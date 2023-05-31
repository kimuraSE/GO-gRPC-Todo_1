package routes

import (
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

	s := grpc.NewServer()

	user.RegisterUserServiceServer(s, uh)
	todo.RegisterTodoServiceServer(s, th)
	log.Println("Server is running on port: 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}

}
