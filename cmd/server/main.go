package main

import (
	"gRPC-Todo/internal/api/entitiy/repository"
	"gRPC-Todo/internal/api/hundler"
	"gRPC-Todo/internal/api/usecase"
	"gRPC-Todo/internal/db"
	"gRPC-Todo/pkg/user"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	user.UnimplementedUserServiceServer
}

func main() {

	db := db.NewDB()

	userRepository := repository.NewUserRepository(db)
	userUsecaase := usecase.NewUserUsecase(userRepository)
	userHandler := hundler.NewUserHandler(userUsecaase)


	
	lis,err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()

	user.RegisterUserServiceServer(s, userHandler)
	log.Println("Server is running on port: 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
