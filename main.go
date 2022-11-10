package main

import (
	"echo-mvc-sample/controllers"
	"echo-mvc-sample/models"
	"echo-mvc-sample/repositories"
	"echo-mvc-sample/services"
	"fmt"
	"github.com/labstack/echo/v4"
)

func logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("This is the sample middleware")
		return next(c)
	}
}

func main() {
	e := echo.New()
	e.Use(logger)
	var handlers []controllers.Controller
	handlers = append(handlers, &controllers.UserController{
		UserService: &services.UserServiceImpl{
			Repository: &repositories.UserRepositoryImpl{
				DB: make(map[string]*models.User),
			},
		},
	})

	for _, handler := range handlers {
		handler.Register(e)
	}
	e.Logger.Fatal(e.Start("localhost:3000"))
}
