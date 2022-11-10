package controllers

import (
	"echo-mvc-sample/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type UserServiceDouble struct {
	GetUserReturnData  *models.User
	GetUserReturnError error
	GetUserCalledId    int
}

func (service *UserServiceDouble) GetUser(id int) (*models.User, error) {
	service.GetUserCalledId = id
	return service.GetUserReturnData, service.GetUserReturnError
}

func TestUserControllerGet(t *testing.T) {
	e := echo.New()
	userService := &UserServiceDouble{
		GetUserCalledId: 0,
	}
	controller := &UserController{
		UserService: userService,
	}
	req := httptest.NewRequest(http.MethodGet, "/users/1234", nil)
	var context echo.Context
	var rec *httptest.ResponseRecorder

	setup := func() {
		rec = httptest.NewRecorder()
		context = e.NewContext(req, rec)
		context.SetPath("/users/:id")
		userService.GetUserReturnData = &models.User{}
	}

	t.Run("should return status ok", func(t *testing.T) {
		setup()
		controller.getUser(context)
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("should return user json", func(t *testing.T) {
		setup()
		userService.GetUserReturnData = &models.User{
			ID:   1234,
			Name: "Yamada",
		}
		controller.getUser(context)
		assert.Equal(t, `{"id":1234,"name":"Yamada"}`+"\n", rec.Body.String())
	})

	t.Run("should call user service by param id", func(t *testing.T) {
		setup()
		context.SetParamNames("id")
		context.SetParamValues("1234")
		controller.getUser(context)
		assert.Equal(t, 1234, userService.GetUserCalledId)
	})
}
