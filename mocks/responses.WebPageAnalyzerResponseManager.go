package mocks

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockAnalyzerInfo is a mock implementation of the AnalyzerInfo interface.
type MockAnalyzerInfo struct {
	ctrl     *gomock.Controller
	recorder *MockAnalyzerInfoMockRecorder
}

// NewMockAnalyzerInfo creates a new mock instance.
func NewMockAnalyzerInfo(ctrl *gomock.Controller) *MockAnalyzerInfo {
	mock := &MockAnalyzerInfo{ctrl: ctrl}
	mock.recorder = &MockAnalyzerInfoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAnalyzerInfo) EXPECT() *MockAnalyzerInfoMockRecorder {
	return m.recorder
}

// GetBody mocks the GetBody method.
func (m *MockAnalyzerInfo) GetBody() string {
	ret := m.ctrl.Call(m, "GetBody")
	ret0, _ := ret[0].(string)
	return ret0
}

// MockAnalyzerInfoMockRecorder is the mock recorder for MockAnalyzerInfo.
type MockAnalyzerInfoMockRecorder struct {
	mock *MockAnalyzerInfo
}

// GetBody expects a call to GetBody.
func (mr *MockAnalyzerInfoMockRecorder) GetBody() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBody", reflect.TypeOf((*MockAnalyzerInfo)(nil).GetBody))
}

// MockWebPageAnalyzerResponseManager is a mock implementation of the WebPageAnalyzerResponseManager interface.
type MockWebPageAnalyzerResponseManager struct {
	ctrl     *gomock.Controller
	recorder *MockWebPageAnalyzerResponseManagerMockRecorder
}

// NewMockWebPageAnalyzerResponseManager creates a new mock instance.
func NewMockWebPageAnalyzerResponseManager(ctrl *gomock.Controller) *MockWebPageAnalyzerResponseManager {
	mock := &MockWebPageAnalyzerResponseManager{ctrl: ctrl}
	mock.recorder = &MockWebPageAnalyzerResponseManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWebPageAnalyzerResponseManager) EXPECT() *MockWebPageAnalyzerResponseManagerMockRecorder {
	return m.recorder
}

// AddHeadingLevel mocks the AddHeadingLevel method.
func (m *MockWebPageAnalyzerResponseManager) AddHeadingLevel(tag, level string) {
	m.ctrl.Call(m, "AddHeadingLevel", tag, level)
}

// MockWebPageAnalyzerResponseManagerMockRecorder is the mock recorder for MockWebPageAnalyzerResponseManager.
type MockWebPageAnalyzerResponseManagerMockRecorder struct {
	mock *MockWebPageAnalyzerResponseManager
}

// AddHeadingLevel expects a call to AddHeadingLevel with the specified parameters.
func (mr *MockWebPageAnalyzerResponseManagerMockRecorder) AddHeadingLevel(tag, level interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddHeadingLevel", reflect.TypeOf((*MockWebPageAnalyzerResponseManager)(nil).AddHeadingLevel), tag, level)
}
