package routes

import (
	"gRPC-Todo/pkg/user"
	"google.golang.org/grpc"
	"log"
	"net"
)

func NewServer(uh user.UserServiceServer) {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()

	user.RegisterUserServiceServer(s, uh)
	log.Println("Server is running on port: 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}

}
