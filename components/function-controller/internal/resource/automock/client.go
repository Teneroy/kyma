// Code generated by mockery v2.33.1. DO NOT EDIT.

package automock

import (
	context "context"

	client "sigs.k8s.io/controller-runtime/pkg/client"

	labels "k8s.io/apimachinery/pkg/labels"

	mock "github.com/stretchr/testify/mock"

	resource "github.com/kyma-project/kyma/components/function-controller/internal/resource"

	types "k8s.io/apimachinery/pkg/types"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, object
func (_m *Client) Create(ctx context.Context, object resource.Object) error {
	ret := _m.Called(ctx, object)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, resource.Object) error); ok {
		r0 = rf(ctx, object)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateWithReference provides a mock function with given fields: ctx, parent, object
func (_m *Client) CreateWithReference(ctx context.Context, parent resource.Object, object resource.Object) error {
	ret := _m.Called(ctx, parent, object)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, resource.Object, resource.Object) error); ok {
		r0 = rf(ctx, parent, object)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, resourceType
func (_m *Client) Delete(ctx context.Context, resourceType resource.Object) error {
	ret := _m.Called(ctx, resourceType)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, resource.Object) error); ok {
		r0 = rf(ctx, resourceType)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAllBySelector provides a mock function with given fields: ctx, resourceType, namespace, selector
func (_m *Client) DeleteAllBySelector(ctx context.Context, resourceType resource.Object, namespace string, selector labels.Selector) error {
	ret := _m.Called(ctx, resourceType, namespace, selector)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, resource.Object, string, labels.Selector) error); ok {
		r0 = rf(ctx, resourceType, namespace, selector)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, key, object
func (_m *Client) Get(ctx context.Context, key types.NamespacedName, object resource.Object) error {
	ret := _m.Called(ctx, key, object)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.NamespacedName, resource.Object) error); ok {
		r0 = rf(ctx, key, object)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListByLabel provides a mock function with given fields: ctx, namespace, _a2, object
func (_m *Client) ListByLabel(ctx context.Context, namespace string, _a2 map[string]string, object client.ObjectList) error {
	ret := _m.Called(ctx, namespace, _a2, object)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, map[string]string, client.ObjectList) error); ok {
		r0 = rf(ctx, namespace, _a2, object)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Status provides a mock function with given fields:
func (_m *Client) Status() client.StatusWriter {
	ret := _m.Called()

	var r0 client.StatusWriter
	if rf, ok := ret.Get(0).(func() client.StatusWriter); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.StatusWriter)
		}
	}

	return r0
}

// Update provides a mock function with given fields: ctx, object
func (_m *Client) Update(ctx context.Context, object resource.Object) error {
	ret := _m.Called(ctx, object)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, resource.Object) error); ok {
		r0 = rf(ctx, object)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
