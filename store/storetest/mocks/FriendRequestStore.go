// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost-server/v5/model"
	mock "github.com/stretchr/testify/mock"
)

// FriendRequestStore is an autogenerated mock type for the FriendRequestStore type
type FriendRequestStore struct {
	mock.Mock
}

// AcceptRequest provides a mock function with given fields: senderId, receiverId
func (_m *FriendRequestStore) AcceptRequest(senderId string, receiverId string) error {
	ret := _m.Called(senderId, receiverId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(senderId, receiverId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindFriendRequest provides a mock function with given fields: senderId, receiverId
func (_m *FriendRequestStore) FindFriendRequest(senderId string, receiverId string) (*model.FriendRequest, error) {
	ret := _m.Called(senderId, receiverId)

	var r0 *model.FriendRequest
	if rf, ok := ret.Get(0).(func(string, string) *model.FriendRequest); ok {
		r0 = rf(senderId, receiverId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.FriendRequest)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(senderId, receiverId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFriendList provides a mock function with given fields: userid
func (_m *FriendRequestStore) GetFriendList(userid string) ([]*model.FriendRequest, error) {
	ret := _m.Called(userid)

	var r0 []*model.FriendRequest
	if rf, ok := ret.Get(0).(func(string) []*model.FriendRequest); ok {
		r0 = rf(userid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.FriendRequest)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPendingList provides a mock function with given fields: senderId
func (_m *FriendRequestStore) GetPendingList(senderId string) ([]*model.FriendRequest, error) {
	ret := _m.Called(senderId)

	var r0 []*model.FriendRequest
	if rf, ok := ret.Get(0).(func(string) []*model.FriendRequest); ok {
		r0 = rf(senderId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.FriendRequest)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(senderId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReceivedList provides a mock function with given fields: receiverId
func (_m *FriendRequestStore) GetReceivedList(receiverId string) ([]*model.FriendRequest, error) {
	ret := _m.Called(receiverId)

	var r0 []*model.FriendRequest
	if rf, ok := ret.Get(0).(func(string) []*model.FriendRequest); ok {
		r0 = rf(receiverId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.FriendRequest)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(receiverId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveRequest provides a mock function with given fields: senderId, receiverId
func (_m *FriendRequestStore) RemoveRequest(senderId string, receiverId string) error {
	ret := _m.Called(senderId, receiverId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(senderId, receiverId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: request
func (_m *FriendRequestStore) Save(request *model.FriendRequest) (*model.FriendRequest, error) {
	ret := _m.Called(request)

	var r0 *model.FriendRequest
	if rf, ok := ret.Get(0).(func(*model.FriendRequest) *model.FriendRequest); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.FriendRequest)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.FriendRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}