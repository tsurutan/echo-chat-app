package services

import (
	"echo-mvc-sample/models"
	"echo-mvc-sample/repositories"
)

type UserService interface {
	GetUser(id int) (*models.User, error)
}

type UserServiceImpl struct {
	Repository repositories.UserRepository
}

func (service *UserServiceImpl) GetUser(id int) (*models.User, error) {
	return service.Repository.Find(id)
}
