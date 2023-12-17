package car

import (
	"car_garage/models"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"

	"go.mongodb.org/mongo-driver/bson"
)

type store struct{}

func New() store {
	return store{}
}

func (s store) GetCars(ctx *gofr.Context) ([]models.Car, error) {
	collection := ctx.MongoDB.Collection("cars")

	filter := bson.D{}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, errors.DB{Err: err}
	}
	defer cur.Close(ctx)

	var cars []models.Car
	for cur.Next(ctx) {
		var c models.Car
		if err := cur.Decode(&c); err != nil {
			return nil, errors.DB{Err: err}
		}
		cars = append(cars, c)
	}
	return cars, nil
}

func (s store) AddCar(ctx *gofr.Context, c models.Car) error {
	collection := ctx.MongoDB.Collection("cars")
	_, err := collection.InsertOne(ctx, c)
	return err
}

func (s store) UpdateCar(ctx *gofr.Context, id string, c models.Car) error {
	collection := ctx.MongoDB.Collection("cars")
	_, err := collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": c})
	return err
}

func (s store) DeleteCar(ctx *gofr.Context, id string) error {
	collection := ctx.MongoDB.Collection("cars")
	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
