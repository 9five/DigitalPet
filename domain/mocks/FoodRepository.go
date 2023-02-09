// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "goDDD/domain"

	mock "github.com/stretchr/testify/mock"
)

// FoodRepository is an autogenerated mock type for the FoodRepository type
type FoodRepository struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, food, optsWhere
func (_m *FoodRepository) Get(ctx context.Context, food *domain.Food, optsWhere ...map[string]interface{}) (*domain.Food, error) {
	_va := make([]interface{}, len(optsWhere))
	for _i := range optsWhere {
		_va[_i] = optsWhere[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, food)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *domain.Food
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Food, ...map[string]interface{}) *domain.Food); ok {
		r0 = rf(ctx, food, optsWhere...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Food)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Food, ...map[string]interface{}) error); ok {
		r1 = rf(ctx, food, optsWhere...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Gets provides a mock function with given fields: ctx, food, optsWhere
func (_m *FoodRepository) Gets(ctx context.Context, food *domain.Food, optsWhere ...map[string]interface{}) ([]*domain.Food, error) {
	_va := make([]interface{}, len(optsWhere))
	for _i := range optsWhere {
		_va[_i] = optsWhere[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, food)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*domain.Food
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Food, ...map[string]interface{}) []*domain.Food); ok {
		r0 = rf(ctx, food, optsWhere...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Food)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Food, ...map[string]interface{}) error); ok {
		r1 = rf(ctx, food, optsWhere...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// New provides a mock function with given fields: ctx, food
func (_m *FoodRepository) New(ctx context.Context, food *domain.Food) (*domain.Food, error) {
	ret := _m.Called(ctx, food)

	var r0 *domain.Food
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Food) *domain.Food); ok {
		r0 = rf(ctx, food)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Food)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Food) error); ok {
		r1 = rf(ctx, food)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewFoodRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewFoodRepository creates a new instance of FoodRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFoodRepository(t mockConstructorTestingTNewFoodRepository) *FoodRepository {
	mock := &FoodRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}