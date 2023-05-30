package usecase

import (
	"gRPC-Todo/internal/api/entitiy/model"
	"gRPC-Todo/internal/api/entitiy/repository"
	"github.com/golang-jwt/jwt/v4"
	bycrypt "golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type IUserUsecase interface {
	LoginUser(user model.User) (string, error)
	RegisterUser(user model.User) (string, error)
}

type userUsecase struct {
	ur repository.IUserRepository
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) LoginUser(user model.User) (string, error) {
	newUser := model.User{}
	if err := uu.ur.GetUserByEmail(&newUser, user.Email); err != nil {
		return "", err
	}

	if err := bycrypt.CompareHashAndPassword(newUser.Password, user.Password); err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": newUser.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (uu *userUsecase) RegisterUser(user model.User) (string, error) {

	hash, err := bycrypt.GenerateFromPassword(user.Password, 10)
	if err != nil {
		return "", err
	}

	newUser := model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: hash,
	}

	if err := uu.ur.CreateUser(&newUser); err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": newUser.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
