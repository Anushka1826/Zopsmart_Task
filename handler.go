package handlers

import (
	"car_garage/models"
	"car_garage/stores"

	"gofr.dev/pkg/gofr"
)

type handler struct {
	store stores.Car
}


func New(c stores.Car) handler {
	return handler{store: c}
}

func (h handler) GetCars(ctx *gofr.Context) (interface{}, error) {
	cars, err := h.store.GetCars(ctx)
	if err != nil {
		return nil, err
	}
	return cars, nil
}

func (h handler) AddCar(ctx *gofr.Context) (interface{}, error) {
	var c models.Car
	err := ctx.Bind(&c)
	if err != nil {
		return nil, err
	}
	err = h.store.AddCar(ctx, c)
	if err != nil {
		return nil, err
	}
	return "Car added to the garage!", nil
}

func (h handler) UpdateCar(ctx *gofr.Context) (interface{}, error) {
	id := ctx.Param("id")
	var c models.Car
	err := ctx.Bind(&c)
	if err != nil {
		return nil, err
	}
	err = h.store.UpdateCar(ctx, id, c)
	if err != nil {
		return nil, err
	}
	return "Car details updated!", nil
}

func (h handler) DeleteCar(ctx *gofr.Context) (interface{}, error) {
	id := ctx.Param("id")
	err := h.store.DeleteCar(ctx, id)
	if err != nil {
		return nil, err
	}
	return "Car deleted from the garage!", nil
}
