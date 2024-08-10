package main

import (
	"go-rest-template/internal/repository"
	"go-rest-template/internal/rest"
	"go-rest-template/usecase/todo"
	"go-rest-template/usecase/user"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const defaultAddress = ":8080"

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	userRepo := repository.NewUserRepository()
	todoRepo := repository.NewToDoRepository()
	userSvc := user.NewService(userRepo)
	todoSvc := todo.NewService(todoRepo, userRepo)

	rest.NewUserHandler(e, userSvc)
	rest.NewToDoHandler(e, todoSvc)

	e.GET("/health", healthCheck)

	address := os.Getenv("SERVER_ADDRESS")
	if address == "" {
		address = defaultAddress
	}
	log.Fatal(e.Start(address))
}

func healthCheck(c echo.Context) error {
	return c.String(200, "OK")
}
