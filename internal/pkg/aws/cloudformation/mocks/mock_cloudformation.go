// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/aws/cloudformation/interfaces.go

// Package mocks is a generated GoMock package.
package mocks

import (
	aws "github.com/aws/aws-sdk-go/aws"
	request "github.com/aws/aws-sdk-go/aws/request"
	cloudformation "github.com/aws/aws-sdk-go/service/cloudformation"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockchangeSetAPI is a mock of changeSetAPI interface
type MockchangeSetAPI struct {
	ctrl     *gomock.Controller
	recorder *MockchangeSetAPIMockRecorder
}

// MockchangeSetAPIMockRecorder is the mock recorder for MockchangeSetAPI
type MockchangeSetAPIMockRecorder struct {
	mock *MockchangeSetAPI
}

// NewMockchangeSetAPI creates a new mock instance
func NewMockchangeSetAPI(ctrl *gomock.Controller) *MockchangeSetAPI {
	mock := &MockchangeSetAPI{ctrl: ctrl}
	mock.recorder = &MockchangeSetAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockchangeSetAPI) EXPECT() *MockchangeSetAPIMockRecorder {
	return m.recorder
}

// CreateChangeSet mocks base method
func (m *MockchangeSetAPI) CreateChangeSet(arg0 *cloudformation.CreateChangeSetInput) (*cloudformation.CreateChangeSetOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateChangeSet", arg0)
	ret0, _ := ret[0].(*cloudformation.CreateChangeSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateChangeSet indicates an expected call of CreateChangeSet
func (mr *MockchangeSetAPIMockRecorder) CreateChangeSet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChangeSet", reflect.TypeOf((*MockchangeSetAPI)(nil).CreateChangeSet), arg0)
}

// WaitUntilChangeSetCreateCompleteWithContext mocks base method
func (m *MockchangeSetAPI) WaitUntilChangeSetCreateCompleteWithContext(arg0 aws.Context, arg1 *cloudformation.DescribeChangeSetInput, arg2 ...request.WaiterOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WaitUntilChangeSetCreateCompleteWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitUntilChangeSetCreateCompleteWithContext indicates an expected call of WaitUntilChangeSetCreateCompleteWithContext
func (mr *MockchangeSetAPIMockRecorder) WaitUntilChangeSetCreateCompleteWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitUntilChangeSetCreateCompleteWithContext", reflect.TypeOf((*MockchangeSetAPI)(nil).WaitUntilChangeSetCreateCompleteWithContext), varargs...)
}

// DescribeChangeSet mocks base method
func (m *MockchangeSetAPI) DescribeChangeSet(arg0 *cloudformation.DescribeChangeSetInput) (*cloudformation.DescribeChangeSetOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeChangeSet", arg0)
	ret0, _ := ret[0].(*cloudformation.DescribeChangeSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeChangeSet indicates an expected call of DescribeChangeSet
func (mr *MockchangeSetAPIMockRecorder) DescribeChangeSet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeChangeSet", reflect.TypeOf((*MockchangeSetAPI)(nil).DescribeChangeSet), arg0)
}

// ExecuteChangeSet mocks base method
func (m *MockchangeSetAPI) ExecuteChangeSet(arg0 *cloudformation.ExecuteChangeSetInput) (*cloudformation.ExecuteChangeSetOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteChangeSet", arg0)
	ret0, _ := ret[0].(*cloudformation.ExecuteChangeSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecuteChangeSet indicates an expected call of ExecuteChangeSet
func (mr *MockchangeSetAPIMockRecorder) ExecuteChangeSet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteChangeSet", reflect.TypeOf((*MockchangeSetAPI)(nil).ExecuteChangeSet), arg0)
}

// DeleteChangeSet mocks base method
func (m *MockchangeSetAPI) DeleteChangeSet(arg0 *cloudformation.DeleteChangeSetInput) (*cloudformation.DeleteChangeSetOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteChangeSet", arg0)
	ret0, _ := ret[0].(*cloudformation.DeleteChangeSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteChangeSet indicates an expected call of DeleteChangeSet
func (mr *MockchangeSetAPIMockRecorder) DeleteChangeSet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteChangeSet", reflect.TypeOf((*MockchangeSetAPI)(nil).DeleteChangeSet), arg0)
}

// Mockapi is a mock of api interface
type Mockapi struct {
	ctrl     *gomock.Controller
	recorder *MockapiMockRecorder
}

// MockapiMockRecorder is the mock recorder for Mockapi
type MockapiMockRecorder struct {
	mock *Mockapi
}

// NewMockapi creates a new mock instance
func NewMockapi(ctrl *gomock.Controller) *Mockapi {
	mock := &Mockapi{ctrl: ctrl}
	mock.recorder = &MockapiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Mockapi) EXPECT() *MockapiMockRecorder {
	return m.recorder
}

// CreateChangeSet mocks base method
func (m *Mockapi) CreateChangeSet(arg0 *cloudformation.CreateChangeSetInput) (*cloudformation.CreateChangeSetOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateChangeSet", arg0)
	ret0, _ := ret[0].(*cloudformation.CreateChangeSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateChangeSet indicates an expected call of CreateChangeSet
func (mr *MockapiMockRecorder) CreateChangeSet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChangeSet", reflect.TypeOf((*Mockapi)(nil).CreateChangeSet), arg0)
}

// WaitUntilChangeSetCreateCompleteWithContext mocks base method
func (m *Mockapi) WaitUntilChangeSetCreateCompleteWithContext(arg0 aws.Context, arg1 *cloudformation.DescribeChangeSetInput, arg2 ...request.WaiterOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WaitUntilChangeSetCreateCompleteWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitUntilChangeSetCreateCompleteWithContext indicates an expected call of WaitUntilChangeSetCreateCompleteWithContext
func (mr *MockapiMockRecorder) WaitUntilChangeSetCreateCompleteWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitUntilChangeSetCreateCompleteWithContext", reflect.TypeOf((*Mockapi)(nil).WaitUntilChangeSetCreateCompleteWithContext), varargs...)
}

// DescribeChangeSet mocks base method
func (m *Mockapi) DescribeChangeSet(arg0 *cloudformation.DescribeChangeSetInput) (*cloudformation.DescribeChangeSetOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeChangeSet", arg0)
	ret0, _ := ret[0].(*cloudformation.DescribeChangeSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeChangeSet indicates an expected call of DescribeChangeSet
func (mr *MockapiMockRecorder) DescribeChangeSet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeChangeSet", reflect.TypeOf((*Mockapi)(nil).DescribeChangeSet), arg0)
}

// ExecuteChangeSet mocks base method
func (m *Mockapi) ExecuteChangeSet(arg0 *cloudformation.ExecuteChangeSetInput) (*cloudformation.ExecuteChangeSetOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteChangeSet", arg0)
	ret0, _ := ret[0].(*cloudformation.ExecuteChangeSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecuteChangeSet indicates an expected call of ExecuteChangeSet
func (mr *MockapiMockRecorder) ExecuteChangeSet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteChangeSet", reflect.TypeOf((*Mockapi)(nil).ExecuteChangeSet), arg0)
}

// DeleteChangeSet mocks base method
func (m *Mockapi) DeleteChangeSet(arg0 *cloudformation.DeleteChangeSetInput) (*cloudformation.DeleteChangeSetOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteChangeSet", arg0)
	ret0, _ := ret[0].(*cloudformation.DeleteChangeSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteChangeSet indicates an expected call of DeleteChangeSet
func (mr *MockapiMockRecorder) DeleteChangeSet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteChangeSet", reflect.TypeOf((*Mockapi)(nil).DeleteChangeSet), arg0)
}

// DescribeStacks mocks base method
func (m *Mockapi) DescribeStacks(arg0 *cloudformation.DescribeStacksInput) (*cloudformation.DescribeStacksOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeStacks", arg0)
	ret0, _ := ret[0].(*cloudformation.DescribeStacksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeStacks indicates an expected call of DescribeStacks
func (mr *MockapiMockRecorder) DescribeStacks(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeStacks", reflect.TypeOf((*Mockapi)(nil).DescribeStacks), arg0)
}

// DescribeStackEvents mocks base method
func (m *Mockapi) DescribeStackEvents(arg0 *cloudformation.DescribeStackEventsInput) (*cloudformation.DescribeStackEventsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeStackEvents", arg0)
	ret0, _ := ret[0].(*cloudformation.DescribeStackEventsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeStackEvents indicates an expected call of DescribeStackEvents
func (mr *MockapiMockRecorder) DescribeStackEvents(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeStackEvents", reflect.TypeOf((*Mockapi)(nil).DescribeStackEvents), arg0)
}

// GetTemplate mocks base method
func (m *Mockapi) GetTemplate(input *cloudformation.GetTemplateInput) (*cloudformation.GetTemplateOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTemplate", input)
	ret0, _ := ret[0].(*cloudformation.GetTemplateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTemplate indicates an expected call of GetTemplate
func (mr *MockapiMockRecorder) GetTemplate(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTemplate", reflect.TypeOf((*Mockapi)(nil).GetTemplate), input)
}

// DeleteStack mocks base method
func (m *Mockapi) DeleteStack(arg0 *cloudformation.DeleteStackInput) (*cloudformation.DeleteStackOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStack", arg0)
	ret0, _ := ret[0].(*cloudformation.DeleteStackOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteStack indicates an expected call of DeleteStack
func (mr *MockapiMockRecorder) DeleteStack(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStack", reflect.TypeOf((*Mockapi)(nil).DeleteStack), arg0)
}

// WaitUntilStackCreateCompleteWithContext mocks base method
func (m *Mockapi) WaitUntilStackCreateCompleteWithContext(arg0 aws.Context, arg1 *cloudformation.DescribeStacksInput, arg2 ...request.WaiterOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WaitUntilStackCreateCompleteWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitUntilStackCreateCompleteWithContext indicates an expected call of WaitUntilStackCreateCompleteWithContext
func (mr *MockapiMockRecorder) WaitUntilStackCreateCompleteWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitUntilStackCreateCompleteWithContext", reflect.TypeOf((*Mockapi)(nil).WaitUntilStackCreateCompleteWithContext), varargs...)
}

// WaitUntilStackUpdateCompleteWithContext mocks base method
func (m *Mockapi) WaitUntilStackUpdateCompleteWithContext(arg0 aws.Context, arg1 *cloudformation.DescribeStacksInput, arg2 ...request.WaiterOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WaitUntilStackUpdateCompleteWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitUntilStackUpdateCompleteWithContext indicates an expected call of WaitUntilStackUpdateCompleteWithContext
func (mr *MockapiMockRecorder) WaitUntilStackUpdateCompleteWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitUntilStackUpdateCompleteWithContext", reflect.TypeOf((*Mockapi)(nil).WaitUntilStackUpdateCompleteWithContext), varargs...)
}

// WaitUntilStackDeleteCompleteWithContext mocks base method
func (m *Mockapi) WaitUntilStackDeleteCompleteWithContext(arg0 aws.Context, arg1 *cloudformation.DescribeStacksInput, arg2 ...request.WaiterOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WaitUntilStackDeleteCompleteWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitUntilStackDeleteCompleteWithContext indicates an expected call of WaitUntilStackDeleteCompleteWithContext
func (mr *MockapiMockRecorder) WaitUntilStackDeleteCompleteWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitUntilStackDeleteCompleteWithContext", reflect.TypeOf((*Mockapi)(nil).WaitUntilStackDeleteCompleteWithContext), varargs...)
}
