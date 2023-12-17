
package stores

import (
	reflect "reflect"
	gomock "github.com/golang/mock/gomock"
	models "car_garage/models"
	gofr "gofr.dev/pkg/gofr"
)

type MockCar struct {
	ctrl     *gomock.Controller
	recorder *MockCarMockRecorder
}

type MockCarMockRecorder struct {
	mock *MockCar
}

func NewMockCar(ctrl *gomock.Controller) *MockCar {
	mock := &MockCar{ctrl: ctrl}
	mock.recorder = &MockCarMockRecorder{mock}
	return mock
}

func (m *MockCar) EXPECT() *MockCarMockRecorder {
	return m.recorder
}

func (m *MockCar) GetCars(ctx *gofr.Context, owner string) ([]models.Car, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCars", ctx, owner)
	ret0, _ := ret[0].([]models.Car)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCarMockRecorder) GetCars(ctx, owner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCars", reflect.TypeOf((*MockCar)(nil).GetCars), ctx, owner)
}

func (m *MockCar) AddCar(ctx *gofr.Context, car models.Car) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCar", ctx, car)
	ret0, _ := ret[0].(error)
	return ret0
}


func (mr *MockCarMockRecorder) AddCar(ctx, car interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCar", reflect.TypeOf((*MockCar)(nil).AddCar), ctx, car)
}

func (m *MockCar) DeleteCar(ctx *gofr.Context, id string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCar", ctx, id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCarMockRecorder) DeleteCar(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCar", reflect.TypeOf((*MockCar)(nil).DeleteCar), ctx, id)
}
