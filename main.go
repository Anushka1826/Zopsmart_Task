package main

import (
	"car_garage/handlers"
	"car_garage/stores/car"

	"gofr.dev/pkg/gofr"
)

func main() {

	app := gofr.New()

	app.Server.ValidateHeaders = false

	store := car.New()

	h := handlers.New(store)

	app.GET("/car", h.GetCars)
	app.POST("/car", h.AddCar)
	app.PUT("/car/:id", h.UpdateCar)
	app.DELETE("/car/:id", h.DeleteCar)
	app.Server.HTTP.Port = 8097


	app.Start()
}
