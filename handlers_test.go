package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"car_garage/models"
	"car_garage/stores"
	"car_garage/stores/car"


	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/request"
)

func initializeHandlersTest(t *testing.T) (*stores.MockCar, handler, *gofr.Gofr) {
	ctrl := gomock.NewController(t)

	store := stores.NewMockCar(ctrl)
	storee := car.New()
	h := New(storee)
	app := gofr.New()

	return store, h, app
}

func TestHandler_Get(t *testing.T) {
	tests := []struct {
		desc        string
		queryParams string
		name        string
		resp        interface{}
		err         error
	}{
		{"get without params", "", "", []models.Car{{Make: "Toyota", Model: "Camry", Year: 2022, Owner: "John Doe", InGarage: true}}, nil},
		{"get with name", "name=John", "John", []models.Car{{Make: "Toyota", Model: "Corolla", Year: 2019, Owner: "John", InGarage: true}}, nil},
		{"get with invalid name", "name=1", "1", nil, errors.InvalidParam{Param: []string{"name"}}},
	}

	store, h, app := initializeHandlersTest(t)

	for i, tc := range tests {
		req := httptest.NewRequest(http.MethodGet, "/car?"+tc.queryParams, nil)
		r := request.NewHTTPRequest(req)
		ctx := gofr.NewContext(nil, r, app)

		store.EXPECT().GetCars(ctx, tc.name).Return(tc.resp, tc.err)

		resp, err := h.GetCars(ctx)

		assert.Equal(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)

		assert.Equal(t, tc.resp, resp, "TEST[%d], failed.\n%s", i, tc.desc)
	}
}

func TestHandler_Create_Invalid_Input_Error(t *testing.T) {
	expErr := errors.Error("test error")

	_, h, app := initializeHandlersTest(t)
	req := httptest.NewRequest(http.MethodGet, "/dummy", nil)
	r := request.NewHTTPRequest(req)
	ctx := gofr.NewContext(nil, r, app)

	_, err := h.AddCar(ctx)

	assert.Equal(t, expErr, err)
}

func TestHandler_Create_Invalid_JSON(t *testing.T) {
	input := `{"make":"Ford","model":"Fusion","year":"2020","owner":"Alice","in_garage":"true"}`
	expErr := &json.UnmarshalTypeError{
		Value:  "string",
		Type:   reflect.TypeOf(2020),
		Offset: 26,
		Struct: "Car",
		Field:  "year",
	}

	_, h, app := initializeHandlersTest(t)

	inputReader := strings.NewReader(input)
	req := httptest.NewRequest(http.MethodGet, "/dummy", inputReader)
	r := request.NewHTTPRequest(req)
	ctx := gofr.NewContext(nil, r, app)

	_, err := h.AddCar(ctx)

	assert.Equal(t, expErr, err)
}

func TestHandler_Create(t *testing.T) {
	carJSON := `{"make":"Ford","model":"Fusion","year":2020,"owner":"Alice","in_garage":true}`
	c := models.Car{Make: "Ford", Model: "Fusion", Year: 2020, Owner: "Alice", InGarage: true}
	tests := []struct {
		desc string
		resp string
		err  error
	}{
		{"create success", "New Car Added!!", nil},
		{"create fail", "", errors.Error("test error")},
	}

	store, h, app := initializeHandlersTest(t)

	for i, tc := range tests {
		input := strings.NewReader(carJSON)

		req := httptest.NewRequest(http.MethodPost, "/car", input)
		r := request.NewHTTPRequest(req)
		ctx := gofr.NewContext(nil, r, app)

		store.EXPECT().AddCar(ctx, c).Return(tc.err)

		_, err := h.AddCar(ctx)

		assert.Equal(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)
	}
}

func TestHandler_Delete(t *testing.T) {
	tests := []struct {
		desc  string
		id    string
		count int
		resp  interface{}
		err   error
	}{
		{"delete invalid entity", "1", 0, nil, errors.InvalidParam{Param: []string{"id"}}},
		{"delete multiple entities", "5", 2, "2 Cars Deleted!", nil},
		{"delete single entity", "3", 1, "1 Car Deleted!", nil},
	}

	store, h, app := initializeHandlersTest(t)

	for i, tc := range tests {
		req := httptest.NewRequest(http.MethodDelete, "/car/"+tc.id, nil)
		r := request.NewHTTPRequest(req)
		ctx := gofr.NewContext(nil, r, app)

		store.EXPECT().DeleteCar(ctx, tc.id).Return(tc.count, tc.err).Times(1)

		resp, err := h.DeleteCar(ctx)

		assert.Equal(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)

		assert.Equal(t, tc.resp, resp, "TEST[%d], failed.\n%s", i, tc.desc)
	}
}
