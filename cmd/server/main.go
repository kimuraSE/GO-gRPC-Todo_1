package main

import (
	"gRPC-Todo/internal/api/entitiy/repository"
	"gRPC-Todo/internal/api/hundler"
	"gRPC-Todo/internal/api/usecase"
	"gRPC-Todo/internal/db"
	"gRPC-Todo/pkg/routes"
	"gRPC-Todo/pkg/todo"
	"gRPC-Todo/pkg/user"
)

type server struct {
	user.UnimplementedUserServiceServer
	todo.UnimplementedTodoServiceServer
}

func main() {

	db := db.NewDB()

	userRepository := repository.NewUserRepository(db)
	userUsecaase := usecase.NewUserUsecase(userRepository)
	userHandler := hundler.NewUserHandler(userUsecaase)

	todoRepository := repository.NewTodoRepository(db)
	todoUsecaase := usecase.NewTodoUsecase(todoRepository)
	todoHandler := hundler.NewTodoHundler(todoUsecaase)

	routes.NewServer(userHandler, todoHandler)
}
