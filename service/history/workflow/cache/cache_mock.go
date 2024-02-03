// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
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

// Code generated by MockGen. DO NOT EDIT.
// Source: cache.go

// Package cache is a generated GoMock package.
package cache

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "go.temporal.io/api/common/v1"
	namespace "go.temporal.io/server/common/namespace"
	shard "go.temporal.io/server/service/history/shard"
	workflow "go.temporal.io/server/service/history/workflow"
)

// MockCache is a mock of Cache interface.
type MockCache struct {
	ctrl     *gomock.Controller
	recorder *MockCacheMockRecorder
}

// MockCacheMockRecorder is the mock recorder for MockCache.
type MockCacheMockRecorder struct {
	mock *MockCache
}

// NewMockCache creates a new mock instance.
func NewMockCache(ctrl *gomock.Controller) *MockCache {
	mock := &MockCache{ctrl: ctrl}
	mock.recorder = &MockCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCache) EXPECT() *MockCacheMockRecorder {
	return m.recorder
}

// GetOrCreateCurrentWorkflowExecution mocks base method.
func (m *MockCache) GetOrCreateCurrentWorkflowExecution(ctx context.Context, shardContext shard.Context, namespaceID namespace.ID, workflowID string, lockPriority workflow.LockPriority) (ReleaseCacheFunc, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrCreateCurrentWorkflowExecution", ctx, shardContext, namespaceID, workflowID, lockPriority)
	ret0, _ := ret[0].(ReleaseCacheFunc)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrCreateCurrentWorkflowExecution indicates an expected call of GetOrCreateCurrentWorkflowExecution.
func (mr *MockCacheMockRecorder) GetOrCreateCurrentWorkflowExecution(ctx, shardContext, namespaceID, workflowID, lockPriority interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrCreateCurrentWorkflowExecution", reflect.TypeOf((*MockCache)(nil).GetOrCreateCurrentWorkflowExecution), ctx, shardContext, namespaceID, workflowID, lockPriority)
}

// GetOrCreateWorkflowExecution mocks base method.
func (m *MockCache) GetOrCreateWorkflowExecution(ctx context.Context, shardContext shard.Context, namespaceID namespace.ID, execution *v1.WorkflowExecution, lockPriority workflow.LockPriority) (workflow.Context, ReleaseCacheFunc, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrCreateWorkflowExecution", ctx, shardContext, namespaceID, execution, lockPriority)
	ret0, _ := ret[0].(workflow.Context)
	ret1, _ := ret[1].(ReleaseCacheFunc)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetOrCreateWorkflowExecution indicates an expected call of GetOrCreateWorkflowExecution.
func (mr *MockCacheMockRecorder) GetOrCreateWorkflowExecution(ctx, shardContext, namespaceID, execution, lockPriority interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrCreateWorkflowExecution", reflect.TypeOf((*MockCache)(nil).GetOrCreateWorkflowExecution), ctx, shardContext, namespaceID, execution, lockPriority)
}
