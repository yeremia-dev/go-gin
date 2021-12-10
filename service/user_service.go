package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/yeremia-dev/go-gin/dto"
	"github.com/yeremia-dev/go-gin/entity"
	"github.com/yeremia-dev/go-gin/repository"
)

type UserService interface {
	Update(user dto.UserUpdateDTO) entity.User
	Profile(userID string) entity.User
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Update(user dto.UserUpdateDTO) entity.User {
	userToUpdate := entity.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	updatedUser := service.userRepository.Updateuser(userToUpdate)
	return updatedUser
}

func (service *userService) Profile(userID string) entity.User {
	return service.userRepository.ProfileUser(userID)
}
