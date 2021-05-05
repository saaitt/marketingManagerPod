package service

import (
	"errors"
	"github.com/saaitt/marketingManagerPod/model"
	"github.com/saaitt/marketingManagerPod/request"
	"github.com/saaitt/marketingManagerPod/response"
)

type UserRepo interface {
	Create(user *model.User) error
	FindByUsername(username string) (*model.User, error)
}

type UserService struct {
	Repo UserRepo
}

func (i UserService) Create(request request.CreateUserRequest) (*response.UserResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	user := model.User{
		Username:     request.Username,
		UserType:     request.UserType,
	}
	if err := user.SetPassword(request.Password); err != nil {
		return nil, err
	}
	if err := i.Repo.Create(&user); err != nil {
		return nil, err
	}
	return userToResponse(user), nil
}

func (i UserService) Authenticate(request request.AuthenticationRequest) (*response.UserResponse, error) {
	user := model.User{
		Username:     request.Username,
		PasswordHash: request.Password,
	}
	existingUser, err := i.Repo.FindByUsername(user.Username)
	if err != nil {
		return nil, err
	}
	if existingUser.ValidatePassword(user.PasswordHash) {
		return userToResponse(user), nil
	}
	return nil, errors.New("wrong username or password")
}

func userToResponse(user model.User) *response.UserResponse {
	return &response.UserResponse{
		ID:       user.ID,
		Username: user.Username,
	}
}


