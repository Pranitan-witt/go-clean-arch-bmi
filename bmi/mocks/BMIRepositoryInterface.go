// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// BMIRepositoryInterface is an autogenerated mock type for the BMIRepositoryInterface type
type BMIRepositoryInterface struct {
	mock.Mock
}

// SaveBMI provides a mock function with given fields: ctx, name, _a2
func (_m *BMIRepositoryInterface) SaveBMI(ctx context.Context, name string, _a2 float32) error {
	ret := _m.Called(ctx, name, _a2)

	if len(ret) == 0 {
		panic("no return value specified for SaveBMI")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, float32) error); ok {
		r0 = rf(ctx, name, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewBMIRepositoryInterface creates a new instance of BMIRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBMIRepositoryInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *BMIRepositoryInterface {
	mock := &BMIRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
