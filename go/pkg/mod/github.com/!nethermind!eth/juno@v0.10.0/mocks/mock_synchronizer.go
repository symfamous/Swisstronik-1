// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/NethermindEth/juno/sync (interfaces: Reader)
//
// Generated by this command:
//
//	mockgen -destination=../mocks/mock_synchronizer.go -package=mocks -mock_names Reader=MockSyncReader github.com/NethermindEth/juno/sync Reader
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	core "github.com/NethermindEth/juno/core"
	sync "github.com/NethermindEth/juno/sync"
	gomock "go.uber.org/mock/gomock"
)

// MockSyncReader is a mock of Reader interface.
type MockSyncReader struct {
	ctrl     *gomock.Controller
	recorder *MockSyncReaderMockRecorder
}

// MockSyncReaderMockRecorder is the mock recorder for MockSyncReader.
type MockSyncReaderMockRecorder struct {
	mock *MockSyncReader
}

// NewMockSyncReader creates a new mock instance.
func NewMockSyncReader(ctrl *gomock.Controller) *MockSyncReader {
	mock := &MockSyncReader{ctrl: ctrl}
	mock.recorder = &MockSyncReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSyncReader) EXPECT() *MockSyncReaderMockRecorder {
	return m.recorder
}

// HighestBlockHeader mocks base method.
func (m *MockSyncReader) HighestBlockHeader() *core.Header {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HighestBlockHeader")
	ret0, _ := ret[0].(*core.Header)
	return ret0
}

// HighestBlockHeader indicates an expected call of HighestBlockHeader.
func (mr *MockSyncReaderMockRecorder) HighestBlockHeader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HighestBlockHeader", reflect.TypeOf((*MockSyncReader)(nil).HighestBlockHeader))
}

// StartingBlockNumber mocks base method.
func (m *MockSyncReader) StartingBlockNumber() (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartingBlockNumber")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartingBlockNumber indicates an expected call of StartingBlockNumber.
func (mr *MockSyncReaderMockRecorder) StartingBlockNumber() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartingBlockNumber", reflect.TypeOf((*MockSyncReader)(nil).StartingBlockNumber))
}

// SubscribeNewHeads mocks base method.
func (m *MockSyncReader) SubscribeNewHeads() sync.HeaderSubscription {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeNewHeads")
	ret0, _ := ret[0].(sync.HeaderSubscription)
	return ret0
}

// SubscribeNewHeads indicates an expected call of SubscribeNewHeads.
func (mr *MockSyncReaderMockRecorder) SubscribeNewHeads() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeNewHeads", reflect.TypeOf((*MockSyncReader)(nil).SubscribeNewHeads))
}