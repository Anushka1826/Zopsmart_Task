package stores

import (
	"car_garage/models"

	"gofr.dev/pkg/gofr"
)


type Car interface {
	GetCars(ctx *gofr.Context) ([]models.Car, error)
	AddCar(ctx *gofr.Context, c models.Car) error
	UpdateCar(ctx *gofr.Context, id string, c models.Car) error
	DeleteCar(ctx *gofr.Context, id string) error
}