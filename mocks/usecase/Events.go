// Code generated by mockery v2.4.0. DO NOT EDIT.

package mocks

import (
	model "github.com/javiertlopez/golang-bootcamp-2020/model"
	mock "github.com/stretchr/testify/mock"
)

// Events is an autogenerated mock type for the Events type
type Events struct {
	mock.Mock
}

// AddReservations provides a mock function with given fields: id, reservations
func (_m *Events) AddReservations(id string, reservations []model.Reservation) ([]model.Reservation, error) {
	ret := _m.Called(id, reservations)

	var r0 []model.Reservation
	if rf, ok := ret.Get(0).(func(string, []model.Reservation) []model.Reservation); ok {
		r0 = rf(id, reservations)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Reservation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []model.Reservation) error); ok {
		r1 = rf(id, reservations)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: event
func (_m *Events) Create(event model.Event) (model.Event, error) {
	ret := _m.Called(event)

	var r0 model.Event
	if rf, ok := ret.Get(0).(func(model.Event) model.Event); ok {
		r0 = rf(event)
	} else {
		r0 = ret.Get(0).(model.Event)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.Event) error); ok {
		r1 = rf(event)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *Events) GetAll() ([]model.Event, error) {
	ret := _m.Called()

	var r0 []model.Event
	if rf, ok := ret.Get(0).(func() []model.Event); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *Events) GetByID(id string) (model.Event, error) {
	ret := _m.Called(id)

	var r0 model.Event
	if rf, ok := ret.Get(0).(func(string) model.Event); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(model.Event)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReservations provides a mock function with given fields: id
func (_m *Events) GetReservations(id string) ([]model.Reservation, error) {
	ret := _m.Called(id)

	var r0 []model.Reservation
	if rf, ok := ret.Get(0).(func(string) []model.Reservation); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Reservation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}