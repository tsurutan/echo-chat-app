package services

import (
	"echo-mvc-sample/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

type UserRepositoryStub struct {
	FindReturnData  *models.User
	FindReturnError error
}

func (repository *UserRepositoryStub) Find(id int) (*models.User, error) {
	return repository.FindReturnData, repository.FindReturnError
}

func TestUserServiceGetUser(t *testing.T) {
	userRepository := &UserRepositoryStub{}
	userService := &UserServiceImpl{
		userRepository,
	}

	t.Run("should return user without error", func(t *testing.T) {
		userRepository.FindReturnData = &models.User{
			ID:   1234,
			Name: "Fuga",
		}
		user, err := userService.GetUser(1234)

		assert.Equal(t, &models.User{
			ID:   1234,
			Name: "Fuga",
		}, user)
		assert.Equal(t, nil, err)
	})
}
