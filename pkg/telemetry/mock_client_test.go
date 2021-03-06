// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

// Code generated by mockery v1.0.0. DO NOT EDIT.
package telemetry

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockClient is an autogenerated mock type for the client type
type MockClient struct {
	mock.Mock
	client Client
}

// Report provides a mock function with given fields: ctx
func (_m *MockClient) Report(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Run provides a mock function with given fields: ctx
func (_m *MockClient) Run(ctx context.Context) {
	client, _ := NewClient("", ClientOpts{
		Application: "qwe",
		Instance:    "",
		Interval:    2,
		PacketSize:  0,
	})

	client.Run(ctx)
}
