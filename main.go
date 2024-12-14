package main

import (
	"github.com/DevisArya/reservasi_lapangan/config"
	"github.com/DevisArya/reservasi_lapangan/handler"
	"github.com/DevisArya/reservasi_lapangan/repository"
	"github.com/DevisArya/reservasi_lapangan/routes"
	"github.com/DevisArya/reservasi_lapangan/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	db := config.NewDB()
	config.InitialMigration(db)
	validate := validator.New()
	e := echo.New()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userHandler := handler.NewUserHandler(userService)

	route := routes.NewRouter(e, userHandler)

	route.Logger.Fatal(route.Start(":8080"))
}
