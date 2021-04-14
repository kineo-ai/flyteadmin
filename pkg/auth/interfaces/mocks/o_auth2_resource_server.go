// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	interfaces "github.com/flyteorg/flyteadmin/pkg/auth/interfaces"
	mock "github.com/stretchr/testify/mock"
)

// OAuth2ResourceServer is an autogenerated mock type for the OAuth2ResourceServer type
type OAuth2ResourceServer struct {
	mock.Mock
}

type OAuth2ResourceServer_ValidateAccessToken struct {
	*mock.Call
}

func (_m OAuth2ResourceServer_ValidateAccessToken) Return(_a0 interfaces.IdentityContext, _a1 error) *OAuth2ResourceServer_ValidateAccessToken {
	return &OAuth2ResourceServer_ValidateAccessToken{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *OAuth2ResourceServer) OnValidateAccessToken(ctx context.Context, tokenStr string) *OAuth2ResourceServer_ValidateAccessToken {
	c := _m.On("ValidateAccessToken", ctx, tokenStr)
	return &OAuth2ResourceServer_ValidateAccessToken{Call: c}
}

func (_m *OAuth2ResourceServer) OnValidateAccessTokenMatch(matchers ...interface{}) *OAuth2ResourceServer_ValidateAccessToken {
	c := _m.On("ValidateAccessToken", matchers...)
	return &OAuth2ResourceServer_ValidateAccessToken{Call: c}
}

// ValidateAccessToken provides a mock function with given fields: ctx, tokenStr
func (_m *OAuth2ResourceServer) ValidateAccessToken(ctx context.Context, tokenStr string) (interfaces.IdentityContext, error) {
	ret := _m.Called(ctx, tokenStr)

	var r0 interfaces.IdentityContext
	if rf, ok := ret.Get(0).(func(context.Context, string) interfaces.IdentityContext); ok {
		r0 = rf(ctx, tokenStr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.IdentityContext)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, tokenStr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
