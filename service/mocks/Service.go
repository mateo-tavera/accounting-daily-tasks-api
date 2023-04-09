// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	entity "github.com/mateo-tavera/accounting-daily-tasks/entity"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CreateTask provides a mock function with given fields: _a0
func (_m *Service) CreateTask(_a0 entity.Task) (*entity.Response, error) {
	ret := _m.Called(_a0)

	var r0 *entity.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(entity.Task) (*entity.Response, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(entity.Task) *entity.Response); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(entity.Task) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTask provides a mock function with given fields: _a0, _a1
func (_m *Service) DeleteTask(_a0 int, _a1 string) (*entity.Response, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *entity.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(int, string) (*entity.Response, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(int, string) *entity.Response); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(int, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllTasks provides a mock function with given fields: _a0, _a1
func (_m *Service) GetAllTasks(_a0 int, _a1 string) (*entity.Response, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *entity.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(int, string) (*entity.Response, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(int, string) *entity.Response); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(int, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTask provides a mock function with given fields: _a0, _a1
func (_m *Service) GetTask(_a0 int, _a1 string) (*entity.Response, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *entity.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(int, string) (*entity.Response, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(int, string) *entity.Response); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(int, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTask provides a mock function with given fields: _a0
func (_m *Service) UpdateTask(_a0 entity.Task) (*entity.Response, error) {
	ret := _m.Called(_a0)

	var r0 *entity.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(entity.Task) (*entity.Response, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(entity.Task) *entity.Response); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(entity.Task) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewService interface {
	mock.TestingT
	Cleanup(func())
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewService(t mockConstructorTestingTNewService) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
