// Code generated by mockery v2.53.4. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost/server/public/model"
	mock "github.com/stretchr/testify/mock"
)

// StatusStore is an autogenerated mock type for the StatusStore type
type StatusStore struct {
	mock.Mock
}

// Get provides a mock function with given fields: userID
func (_m *StatusStore) Get(userID string) (*model.Status, error) {
	ret := _m.Called(userID)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *model.Status
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.Status, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(string) *model.Status); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Status)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByIds provides a mock function with given fields: userIds
func (_m *StatusStore) GetByIds(userIds []string) ([]*model.Status, error) {
	ret := _m.Called(userIds)

	if len(ret) == 0 {
		panic("no return value specified for GetByIds")
	}

	var r0 []*model.Status
	var r1 error
	if rf, ok := ret.Get(0).(func([]string) ([]*model.Status, error)); ok {
		return rf(userIds)
	}
	if rf, ok := ret.Get(0).(func([]string) []*model.Status); ok {
		r0 = rf(userIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Status)
		}
	}

	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(userIds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalActiveUsersCount provides a mock function with no fields
func (_m *StatusStore) GetTotalActiveUsersCount() (int64, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetTotalActiveUsersCount")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func() (int64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResetAll provides a mock function with no fields
func (_m *StatusStore) ResetAll() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ResetAll")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveOrUpdate provides a mock function with given fields: status
func (_m *StatusStore) SaveOrUpdate(status *model.Status) error {
	ret := _m.Called(status)

	if len(ret) == 0 {
		panic("no return value specified for SaveOrUpdate")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Status) error); ok {
		r0 = rf(status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveOrUpdateMany provides a mock function with given fields: statuses
func (_m *StatusStore) SaveOrUpdateMany(statuses map[string]*model.Status) error {
	ret := _m.Called(statuses)

	if len(ret) == 0 {
		panic("no return value specified for SaveOrUpdateMany")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(map[string]*model.Status) error); ok {
		r0 = rf(statuses)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateExpiredDNDStatuses provides a mock function with no fields
func (_m *StatusStore) UpdateExpiredDNDStatuses() ([]*model.Status, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for UpdateExpiredDNDStatuses")
	}

	var r0 []*model.Status
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*model.Status, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*model.Status); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Status)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateLastActivityAt provides a mock function with given fields: userID, lastActivityAt
func (_m *StatusStore) UpdateLastActivityAt(userID string, lastActivityAt int64) error {
	ret := _m.Called(userID, lastActivityAt)

	if len(ret) == 0 {
		panic("no return value specified for UpdateLastActivityAt")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int64) error); ok {
		r0 = rf(userID, lastActivityAt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewStatusStore creates a new instance of StatusStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStatusStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *StatusStore {
	mock := &StatusStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
