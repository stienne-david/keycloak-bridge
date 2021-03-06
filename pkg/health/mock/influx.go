// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudtrust/common-healthcheck (interfaces: InfluxHealthChecker)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	common_healthcheck "github.com/cloudtrust/common-healthcheck"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// InfluxHealthChecker is a mock of InfluxHealthChecker interface
type InfluxHealthChecker struct {
	ctrl     *gomock.Controller
	recorder *InfluxHealthCheckerMockRecorder
}

// InfluxHealthCheckerMockRecorder is the mock recorder for InfluxHealthChecker
type InfluxHealthCheckerMockRecorder struct {
	mock *InfluxHealthChecker
}

// NewInfluxHealthChecker creates a new mock instance
func NewInfluxHealthChecker(ctrl *gomock.Controller) *InfluxHealthChecker {
	mock := &InfluxHealthChecker{ctrl: ctrl}
	mock.recorder = &InfluxHealthCheckerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *InfluxHealthChecker) EXPECT() *InfluxHealthCheckerMockRecorder {
	return m.recorder
}

// HealthChecks mocks base method
func (m *InfluxHealthChecker) HealthChecks(arg0 context.Context) []common_healthcheck.InfluxReport {
	ret := m.ctrl.Call(m, "HealthChecks", arg0)
	ret0, _ := ret[0].([]common_healthcheck.InfluxReport)
	return ret0
}

// HealthChecks indicates an expected call of HealthChecks
func (mr *InfluxHealthCheckerMockRecorder) HealthChecks(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthChecks", reflect.TypeOf((*InfluxHealthChecker)(nil).HealthChecks), arg0)
}
