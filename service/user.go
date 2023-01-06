package service

import (
	"log"

	"github.com/linothomas14/hadir-in-api/helper/response"
	"github.com/linothomas14/hadir-in-api/model"
	"github.com/linothomas14/hadir-in-api/repository"
)

type UserService interface {
	Update(user model.User) (response.UserResponse, error)
	GetProfile(userId int) (response.UserResponse, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRep repository.UserRepository) UserService {
	return &userService{
		userRepository: userRep,
	}
}

func (service *userService) Update(userParam model.User) (response.UserResponse, error) {
	var userRes response.UserResponse

	user, err := service.userRepository.GetUser(int(userParam.ID))
	userParam.CreatedAt = user.CreatedAt
	userParam.UpdatedAt = user.UpdatedAt
	if err != nil {
		return response.UserResponse{}, err
	}

	if userParam.Name == "" {
		log.Println("Masuk sini")
		userParam.Name = user.Name
	}

	if userParam.Email == "" {
		userParam.Email = user.Email
	}

	user, err = service.userRepository.UpdateUser(userParam)

	if err != nil {
		return response.UserResponse{}, err
	}

	userRes.ID = user.ID
	userRes.Name = user.Name
	userRes.Email = user.Email
	return userRes, nil
}

func (service *userService) GetProfile(userId int) (response.UserResponse, error) {

	var userRes response.UserResponse

	user, err := service.userRepository.GetUser(userId)
	if err != nil {
		return response.UserResponse{}, err
	}

	userRes.ID = user.ID
	userRes.Name = user.Name
	userRes.Email = user.Email

	return userRes, err
}
