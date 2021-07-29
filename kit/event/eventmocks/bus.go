package eventmocks

import (
	"context"
	"github.com/alexperezortuno/go-auth-with-jwt-redis-postgres/kit/event"
	"github.com/stretchr/testify/mock"
)

type Bus struct {
	mock.Mock
}

// Publish provides a mock function with given fields: _a0, _a1
func (_m *Bus) Publish(_a0 context.Context, _a1 []event.Event) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []event.Event) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Subscribe provides a mock function with given fields: _a0, _a1
func (_m *Bus) Subscribe(_a0 event.Type, _a1 event.Handler) {
	_m.Called(_a0, _a1)
}
