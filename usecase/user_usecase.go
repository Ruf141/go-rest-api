package usercase

import (
	"go-rest-api/model"
	"go-rest-api/repository"

	"golang.org/x/crypto/bcrypt"
)

type IUserUseCase interface {
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (string, error)
}

type userUsecase struct {
	ur repository.IUserRepository
}

func NewUserUsecase(ur repository.IUserRepository) IUserUseCase{
	return &userUsecase{ur}
}

func (uu *userUsecase) SignUp(user model.User)(model.UserResponse.userUsecas,error){
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password),10)
	if err != nil {
		return model.UserResponse{}, err
	}
	newUser := model.User{Email: user.Email, Password: string(hash)}
	if err := uu.ur.CreateUser(&newUser); err!=nil{
		return model.UserResponse{},err
	}
	resUser := model.UserResponse{
		ID: newUser.ID,
		Email : newUser.Email,
	}
	return resUser, nil
}