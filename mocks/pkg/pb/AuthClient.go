// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	pb "github.com/guilhermealvess/guicpay/pkg/pb"
)

// AuthClient is an autogenerated mock type for the AuthClient type
type AuthClient struct {
	mock.Mock
}

// Auth provides a mock function with given fields: ctx, in, opts
func (_m *AuthClient) Auth(ctx context.Context, in *pb.AuthRequest, opts ...grpc.CallOption) (*pb.AuthResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Auth")
	}

	var r0 *pb.AuthResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pb.AuthRequest, ...grpc.CallOption) (*pb.AuthResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pb.AuthRequest, ...grpc.CallOption) *pb.AuthResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.AuthResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pb.AuthRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAuthClient creates a new instance of AuthClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthClient {
	mock := &AuthClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
