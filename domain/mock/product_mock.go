package mock

import (
	"context"

	"github.com/fahmilukis/go-product-svc/domain"
	"github.com/stretchr/testify/mock"
)

type ProductRepository struct {
	*mock.Mock
}

// Delete mock
func (_m *ProductRepository) Delete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)
	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update mock
func (_m *ProductRepository) Update(ctx context.Context, a *domain.Products) error {
	ret := _m.Called(ctx, a)
	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Products) error); ok {
		r0 = rf(ctx, a)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Store mock
func (_m *ProductRepository) Store(ctx context.Context, a *domain.Products) error {
	ret := _m.Called(ctx, a)
	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Products) error); ok {
		r0 = rf(ctx, a)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetch mock
func (_m *ProductRepository) Fetch(ctx context.Context, csr string, num int64) ([]domain.Products, string, error) {
	ret := _m.Called(ctx, csr, num)

	var r0 []domain.Products
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) []domain.Products); ok {
		r0 = rf(ctx, csr, num)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Products)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, string, int64) string); ok {
		r1 = rf(ctx, csr, num)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) error); ok {
		r2 = rf(ctx, csr, num)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
