// Code generated by MockGen. DO NOT EDIT.
// Source: go.uber.org/yarpc/v2 (interfaces: RouterMiddleware,UnaryInboundTransportMiddleware,UnaryOutboundTransportMiddleware,StreamInboundTransportMiddleware,StreamOutboundTransportMiddleware)

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package yarpctest is a generated GoMock package.
package yarpctest

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v2 "go.uber.org/yarpc/v2"
)

// MockRouterMiddleware is a mock of RouterMiddleware interface
type MockRouterMiddleware struct {
	ctrl     *gomock.Controller
	recorder *MockRouterMiddlewareMockRecorder
}

// MockRouterMiddlewareMockRecorder is the mock recorder for MockRouterMiddleware
type MockRouterMiddlewareMockRecorder struct {
	mock *MockRouterMiddleware
}

// NewMockRouterMiddleware creates a new mock instance
func NewMockRouterMiddleware(ctrl *gomock.Controller) *MockRouterMiddleware {
	mock := &MockRouterMiddleware{ctrl: ctrl}
	mock.recorder = &MockRouterMiddlewareMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRouterMiddleware) EXPECT() *MockRouterMiddlewareMockRecorder {
	return m.recorder
}

// Choose mocks base method
func (m *MockRouterMiddleware) Choose(arg0 context.Context, arg1 *v2.Request, arg2 v2.Router) (v2.TransportHandlerSpec, error) {
	ret := m.ctrl.Call(m, "Choose", arg0, arg1, arg2)
	ret0, _ := ret[0].(v2.TransportHandlerSpec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Choose indicates an expected call of Choose
func (mr *MockRouterMiddlewareMockRecorder) Choose(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Choose", reflect.TypeOf((*MockRouterMiddleware)(nil).Choose), arg0, arg1, arg2)
}

// Procedures mocks base method
func (m *MockRouterMiddleware) Procedures(arg0 v2.Router) []v2.TransportProcedure {
	ret := m.ctrl.Call(m, "Procedures", arg0)
	ret0, _ := ret[0].([]v2.TransportProcedure)
	return ret0
}

// Procedures indicates an expected call of Procedures
func (mr *MockRouterMiddlewareMockRecorder) Procedures(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Procedures", reflect.TypeOf((*MockRouterMiddleware)(nil).Procedures), arg0)
}

// MockUnaryInboundTransportMiddleware is a mock of UnaryInboundTransportMiddleware interface
type MockUnaryInboundTransportMiddleware struct {
	ctrl     *gomock.Controller
	recorder *MockUnaryInboundTransportMiddlewareMockRecorder
}

// MockUnaryInboundTransportMiddlewareMockRecorder is the mock recorder for MockUnaryInboundTransportMiddleware
type MockUnaryInboundTransportMiddlewareMockRecorder struct {
	mock *MockUnaryInboundTransportMiddleware
}

// NewMockUnaryInboundTransportMiddleware creates a new mock instance
func NewMockUnaryInboundTransportMiddleware(ctrl *gomock.Controller) *MockUnaryInboundTransportMiddleware {
	mock := &MockUnaryInboundTransportMiddleware{ctrl: ctrl}
	mock.recorder = &MockUnaryInboundTransportMiddlewareMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUnaryInboundTransportMiddleware) EXPECT() *MockUnaryInboundTransportMiddlewareMockRecorder {
	return m.recorder
}

// Handle mocks base method
func (m *MockUnaryInboundTransportMiddleware) Handle(arg0 context.Context, arg1 *v2.Request, arg2 *v2.Buffer, arg3 v2.UnaryTransportHandler) (*v2.Response, *v2.Buffer, error) {
	ret := m.ctrl.Call(m, "Handle", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v2.Response)
	ret1, _ := ret[1].(*v2.Buffer)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Handle indicates an expected call of Handle
func (mr *MockUnaryInboundTransportMiddlewareMockRecorder) Handle(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockUnaryInboundTransportMiddleware)(nil).Handle), arg0, arg1, arg2, arg3)
}

// MockUnaryOutboundTransportMiddleware is a mock of UnaryOutboundTransportMiddleware interface
type MockUnaryOutboundTransportMiddleware struct {
	ctrl     *gomock.Controller
	recorder *MockUnaryOutboundTransportMiddlewareMockRecorder
}

// MockUnaryOutboundTransportMiddlewareMockRecorder is the mock recorder for MockUnaryOutboundTransportMiddleware
type MockUnaryOutboundTransportMiddlewareMockRecorder struct {
	mock *MockUnaryOutboundTransportMiddleware
}

// NewMockUnaryOutboundTransportMiddleware creates a new mock instance
func NewMockUnaryOutboundTransportMiddleware(ctrl *gomock.Controller) *MockUnaryOutboundTransportMiddleware {
	mock := &MockUnaryOutboundTransportMiddleware{ctrl: ctrl}
	mock.recorder = &MockUnaryOutboundTransportMiddlewareMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUnaryOutboundTransportMiddleware) EXPECT() *MockUnaryOutboundTransportMiddlewareMockRecorder {
	return m.recorder
}

// Call mocks base method
func (m *MockUnaryOutboundTransportMiddleware) Call(arg0 context.Context, arg1 *v2.Request, arg2 *v2.Buffer, arg3 v2.UnaryOutbound) (*v2.Response, *v2.Buffer, error) {
	ret := m.ctrl.Call(m, "Call", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v2.Response)
	ret1, _ := ret[1].(*v2.Buffer)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Call indicates an expected call of Call
func (mr *MockUnaryOutboundTransportMiddlewareMockRecorder) Call(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockUnaryOutboundTransportMiddleware)(nil).Call), arg0, arg1, arg2, arg3)
}

// MockStreamInboundTransportMiddleware is a mock of StreamInboundTransportMiddleware interface
type MockStreamInboundTransportMiddleware struct {
	ctrl     *gomock.Controller
	recorder *MockStreamInboundTransportMiddlewareMockRecorder
}

// MockStreamInboundTransportMiddlewareMockRecorder is the mock recorder for MockStreamInboundTransportMiddleware
type MockStreamInboundTransportMiddlewareMockRecorder struct {
	mock *MockStreamInboundTransportMiddleware
}

// NewMockStreamInboundTransportMiddleware creates a new mock instance
func NewMockStreamInboundTransportMiddleware(ctrl *gomock.Controller) *MockStreamInboundTransportMiddleware {
	mock := &MockStreamInboundTransportMiddleware{ctrl: ctrl}
	mock.recorder = &MockStreamInboundTransportMiddlewareMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStreamInboundTransportMiddleware) EXPECT() *MockStreamInboundTransportMiddlewareMockRecorder {
	return m.recorder
}

// HandleStream mocks base method
func (m *MockStreamInboundTransportMiddleware) HandleStream(arg0 *v2.ServerStream, arg1 v2.StreamTransportHandler) error {
	ret := m.ctrl.Call(m, "HandleStream", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleStream indicates an expected call of HandleStream
func (mr *MockStreamInboundTransportMiddlewareMockRecorder) HandleStream(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleStream", reflect.TypeOf((*MockStreamInboundTransportMiddleware)(nil).HandleStream), arg0, arg1)
}

// MockStreamOutboundTransportMiddleware is a mock of StreamOutboundTransportMiddleware interface
type MockStreamOutboundTransportMiddleware struct {
	ctrl     *gomock.Controller
	recorder *MockStreamOutboundTransportMiddlewareMockRecorder
}

// MockStreamOutboundTransportMiddlewareMockRecorder is the mock recorder for MockStreamOutboundTransportMiddleware
type MockStreamOutboundTransportMiddlewareMockRecorder struct {
	mock *MockStreamOutboundTransportMiddleware
}

// NewMockStreamOutboundTransportMiddleware creates a new mock instance
func NewMockStreamOutboundTransportMiddleware(ctrl *gomock.Controller) *MockStreamOutboundTransportMiddleware {
	mock := &MockStreamOutboundTransportMiddleware{ctrl: ctrl}
	mock.recorder = &MockStreamOutboundTransportMiddlewareMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStreamOutboundTransportMiddleware) EXPECT() *MockStreamOutboundTransportMiddlewareMockRecorder {
	return m.recorder
}

// CallStream mocks base method
func (m *MockStreamOutboundTransportMiddleware) CallStream(arg0 context.Context, arg1 *v2.Request, arg2 v2.StreamOutbound) (*v2.ClientStream, error) {
	ret := m.ctrl.Call(m, "CallStream", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v2.ClientStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CallStream indicates an expected call of CallStream
func (mr *MockStreamOutboundTransportMiddlewareMockRecorder) CallStream(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallStream", reflect.TypeOf((*MockStreamOutboundTransportMiddleware)(nil).CallStream), arg0, arg1, arg2)
}
