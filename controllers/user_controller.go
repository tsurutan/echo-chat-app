package controllers

import (
	"echo-mvc-sample/services"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type UserController struct {
	UserService services.UserService
}

type UserResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (controller *UserController) getUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// TODO
	}
	user, err := controller.UserService.GetUser(id)
	if err != nil {
		// TODO
	}
	return c.JSON(http.StatusOK, &UserResponse{Id: user.ID, Name: user.Name})
}

type UserRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (controller *UserController) saveUser(c echo.Context) error {
	// TODO: create user
	return c.String(http.StatusCreated, "Created User")
}

func (controller *UserController) Register(echo *echo.Echo) {
	echo.GET("/users/:id", controller.getUser)
	echo.POST("/users", controller.saveUser)
}
