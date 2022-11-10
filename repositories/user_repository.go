package repositories

import "echo-mvc-sample/models"

type UserRepository interface {
	Find(id int) (*models.User, error)
}

type UserRepositoryImpl struct {
	DB map[string]*models.User
}

func (repository *UserRepositoryImpl) Find(id int) (*models.User, error) {
	return &models.User{
		ID:   1234,
		Name: "Yes man",
	}, nil
}
