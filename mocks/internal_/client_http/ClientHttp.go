// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"
	http "net/http"

	json "encoding/json"

	mock "github.com/stretchr/testify/mock"

	url "net/url"
)

// ClientHttp is an autogenerated mock type for the ClientHttp type
type ClientHttp struct {
	mock.Mock
}

// Bind provides a mock function with given fields: res, v
func (_m *ClientHttp) Bind(res *http.Response, v interface{}) error {
	ret := _m.Called(res, v)

	if len(ret) == 0 {
		panic("no return value specified for Bind")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*http.Response, interface{}) error); ok {
		r0 = rf(res, v)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Send provides a mock function with given fields: ctx, method, _a2, query, bodyJSON, pathParams
func (_m *ClientHttp) Send(ctx context.Context, method string, _a2 string, query url.Values, bodyJSON json.RawMessage, pathParams ...string) (*http.Response, error) {
	_va := make([]interface{}, len(pathParams))
	for _i := range pathParams {
		_va[_i] = pathParams[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, method, _a2, query, bodyJSON)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Send")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, url.Values, json.RawMessage, ...string) (*http.Response, error)); ok {
		return rf(ctx, method, _a2, query, bodyJSON, pathParams...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, url.Values, json.RawMessage, ...string) *http.Response); ok {
		r0 = rf(ctx, method, _a2, query, bodyJSON, pathParams...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, url.Values, json.RawMessage, ...string) error); ok {
		r1 = rf(ctx, method, _a2, query, bodyJSON, pathParams...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewClientHttp creates a new instance of ClientHttp. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClientHttp(t interface {
	mock.TestingT
	Cleanup(func())
}) *ClientHttp {
	mock := &ClientHttp{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
