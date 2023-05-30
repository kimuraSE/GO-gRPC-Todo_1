package hundler

import (
	"context"
	"gRPC-Todo/internal/api/entitiy/model"
	"gRPC-Todo/internal/api/usecase"
	"gRPC-Todo/pkg/user"
)

type IUserHandler interface {
	RegisterUser(ctx context.Context, in *user.RegisterRequest) (*user.RegisterResponse, error)
	LoginUser(ctx context.Context, in *user.LoginRequest) (*user.LoginResponse, error)
}

type userHandler struct {
	user.UnimplementedUserServiceServer
	uu usecase.IUserUsecase
}

func NewUserHandler(uu usecase.IUserUsecase) user.UserServiceServer {
	return &userHandler{uu: uu}
}

func (h *userHandler) RegisterUser(ctx context.Context, in *user.RegisterRequest) (*user.RegisterResponse, error) {

	newUser := model.User{
		Name:     in.Name,
		Email:    in.Email,
		Password: []byte(in.Password),
	}

	token, err := h.uu.RegisterUser(newUser)
	if err != nil {
		return &user.RegisterResponse{}, err
	}

	return &user.RegisterResponse{Token: token}, nil

}

func (h *userHandler) LoginUser(ctx context.Context, in *user.LoginRequest) (*user.LoginResponse, error) {

	newUser := model.User{
		Email:    in.Email,
		Password: []byte(in.Password),
	}

	token, err := h.uu.LoginUser(newUser)
	if err != nil {
		return &user.LoginResponse{}, err
	}

	return &user.LoginResponse{Token: token}, nil

}
