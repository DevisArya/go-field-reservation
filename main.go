package main

import (
	"log"

	"github.com/DevisArya/reservasi_lapangan/app"
	"github.com/DevisArya/reservasi_lapangan/config"
	"github.com/DevisArya/reservasi_lapangan/routes"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	db := config.NewDB()
	config.InitialMigration(db)
	validate := validator.New()
	e := echo.New()

	appContainer := app.NewAppContainer(db, validate)

	routes.NewRouter(e, appContainer)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
