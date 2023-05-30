package hundler

import (
	"context"
	"gRPC-Todo/internal/api/usecase"
	"gRPC-Todo/pkg/user"
)

type IUserHandler interface {
	RegisterUser(ctx context.Context, in *user.RegisterRequest) (*user.RegisterResponse, error)
}

type userHandler struct {
	user.UnimplementedUserServiceServer
	uu usecase.IUserUsecase
}

