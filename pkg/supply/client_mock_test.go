// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/SAP/cloud-authorization-buildpack/pkg/uploader (interfaces: AMSClient)

// Package supply_test is a generated GoMock package.
package supply_test

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAMSClient is a mock of AMSClient interface.
type MockAMSClient struct {
	ctrl     *gomock.Controller
	recorder *MockAMSClientMockRecorder
}

// MockAMSClientMockRecorder is the mock recorder for MockAMSClient.
type MockAMSClientMockRecorder struct {
	mock *MockAMSClient
}

// NewMockAMSClient creates a new mock instance.
func NewMockAMSClient(ctrl *gomock.Controller) *MockAMSClient {
	mock := &MockAMSClient{ctrl: ctrl}
	mock.recorder = &MockAMSClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAMSClient) EXPECT() *MockAMSClientMockRecorder {
	return m.recorder
}

// Do mocks base method.
func (m *MockAMSClient) Do(arg0 *http.Request) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", arg0)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do.
func (mr *MockAMSClientMockRecorder) Do(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockAMSClient)(nil).Do), arg0)
}
