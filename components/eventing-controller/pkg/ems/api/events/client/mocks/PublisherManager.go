// Code generated by mockery v2.27.1. DO NOT EDIT.

package mocks

import (
	event "github.com/cloudevents/sdk-go/v2/event"
	mock "github.com/stretchr/testify/mock"

	types "github.com/kyma-project/kyma/components/eventing-controller/pkg/ems/api/events/types"
)

// PublisherManager is an autogenerated mock type for the PublisherManager type
type PublisherManager struct {
	mock.Mock
}

// Create provides a mock function with given fields: subscription
func (_m *PublisherManager) Create(subscription *types.Subscription) (*types.CreateResponse, error) {
	ret := _m.Called(subscription)

	var r0 *types.CreateResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*types.Subscription) (*types.CreateResponse, error)); ok {
		return rf(subscription)
	}
	if rf, ok := ret.Get(0).(func(*types.Subscription) *types.CreateResponse); ok {
		r0 = rf(subscription)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.CreateResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*types.Subscription) error); ok {
		r1 = rf(subscription)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: name
func (_m *PublisherManager) Delete(name string) (*types.DeleteResponse, error) {
	ret := _m.Called(name)

	var r0 *types.DeleteResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*types.DeleteResponse, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) *types.DeleteResponse); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.DeleteResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: name
func (_m *PublisherManager) Get(name string) (*types.Subscription, *types.Response, error) {
	ret := _m.Called(name)

	var r0 *types.Subscription
	var r1 *types.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(string) (*types.Subscription, *types.Response, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) *types.Subscription); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(string) *types.Response); ok {
		r1 = rf(name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*types.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(name)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// List provides a mock function with given fields:
func (_m *PublisherManager) List() (*types.Subscriptions, *types.Response, error) {
	ret := _m.Called()

	var r0 *types.Subscriptions
	var r1 *types.Response
	var r2 error
	if rf, ok := ret.Get(0).(func() (*types.Subscriptions, *types.Response, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *types.Subscriptions); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Subscriptions)
		}
	}

	if rf, ok := ret.Get(1).(func() *types.Response); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*types.Response)
		}
	}

	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Publish provides a mock function with given fields: cloudEvent, qos
func (_m *PublisherManager) Publish(cloudEvent event.Event, qos types.Qos) (*types.PublishResponse, error) {
	ret := _m.Called(cloudEvent, qos)

	var r0 *types.PublishResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(event.Event, types.Qos) (*types.PublishResponse, error)); ok {
		return rf(cloudEvent, qos)
	}
	if rf, ok := ret.Get(0).(func(event.Event, types.Qos) *types.PublishResponse); ok {
		r0 = rf(cloudEvent, qos)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.PublishResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(event.Event, types.Qos) error); ok {
		r1 = rf(cloudEvent, qos)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TriggerHandshake provides a mock function with given fields: name
func (_m *PublisherManager) TriggerHandshake(name string) (*types.TriggerHandshake, error) {
	ret := _m.Called(name)

	var r0 *types.TriggerHandshake
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*types.TriggerHandshake, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) *types.TriggerHandshake); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.TriggerHandshake)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: name, webhookAuth
func (_m *PublisherManager) Update(name string, webhookAuth *types.WebhookAuth) (*types.UpdateResponse, error) {
	ret := _m.Called(name, webhookAuth)

	var r0 *types.UpdateResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *types.WebhookAuth) (*types.UpdateResponse, error)); ok {
		return rf(name, webhookAuth)
	}
	if rf, ok := ret.Get(0).(func(string, *types.WebhookAuth) *types.UpdateResponse); ok {
		r0 = rf(name, webhookAuth)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.UpdateResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(string, *types.WebhookAuth) error); ok {
		r1 = rf(name, webhookAuth)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateState provides a mock function with given fields: name, state
func (_m *PublisherManager) UpdateState(name string, state types.State) (*types.UpdateStateResponse, error) {
	ret := _m.Called(name, state)

	var r0 *types.UpdateStateResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(string, types.State) (*types.UpdateStateResponse, error)); ok {
		return rf(name, state)
	}
	if rf, ok := ret.Get(0).(func(string, types.State) *types.UpdateStateResponse); ok {
		r0 = rf(name, state)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.UpdateStateResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(string, types.State) error); ok {
		r1 = rf(name, state)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewPublisherManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewPublisherManager creates a new instance of PublisherManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPublisherManager(t mockConstructorTestingTNewPublisherManager) *PublisherManager {
	mock := &PublisherManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
