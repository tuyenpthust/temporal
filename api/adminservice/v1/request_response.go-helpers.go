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

package adminservice

import (
	"google.golang.org/protobuf/proto"
)

// Marshal an object of type RebuildMutableStateRequest to the protobuf v3 wire format
func (val *RebuildMutableStateRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RebuildMutableStateRequest from the protobuf v3 wire format
func (val *RebuildMutableStateRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RebuildMutableStateRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RebuildMutableStateRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RebuildMutableStateRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RebuildMutableStateRequest
	switch t := that.(type) {
	case *RebuildMutableStateRequest:
		that1 = t
	case RebuildMutableStateRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RebuildMutableStateResponse to the protobuf v3 wire format
func (val *RebuildMutableStateResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RebuildMutableStateResponse from the protobuf v3 wire format
func (val *RebuildMutableStateResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RebuildMutableStateResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RebuildMutableStateResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RebuildMutableStateResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RebuildMutableStateResponse
	switch t := that.(type) {
	case *RebuildMutableStateResponse:
		that1 = t
	case RebuildMutableStateResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ImportWorkflowExecutionRequest to the protobuf v3 wire format
func (val *ImportWorkflowExecutionRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ImportWorkflowExecutionRequest from the protobuf v3 wire format
func (val *ImportWorkflowExecutionRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ImportWorkflowExecutionRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ImportWorkflowExecutionRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ImportWorkflowExecutionRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ImportWorkflowExecutionRequest
	switch t := that.(type) {
	case *ImportWorkflowExecutionRequest:
		that1 = t
	case ImportWorkflowExecutionRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ImportWorkflowExecutionResponse to the protobuf v3 wire format
func (val *ImportWorkflowExecutionResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ImportWorkflowExecutionResponse from the protobuf v3 wire format
func (val *ImportWorkflowExecutionResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ImportWorkflowExecutionResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ImportWorkflowExecutionResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ImportWorkflowExecutionResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ImportWorkflowExecutionResponse
	switch t := that.(type) {
	case *ImportWorkflowExecutionResponse:
		that1 = t
	case ImportWorkflowExecutionResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type DescribeMutableStateRequest to the protobuf v3 wire format
func (val *DescribeMutableStateRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type DescribeMutableStateRequest from the protobuf v3 wire format
func (val *DescribeMutableStateRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *DescribeMutableStateRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two DescribeMutableStateRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DescribeMutableStateRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DescribeMutableStateRequest
	switch t := that.(type) {
	case *DescribeMutableStateRequest:
		that1 = t
	case DescribeMutableStateRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type DescribeMutableStateResponse to the protobuf v3 wire format
func (val *DescribeMutableStateResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type DescribeMutableStateResponse from the protobuf v3 wire format
func (val *DescribeMutableStateResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *DescribeMutableStateResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two DescribeMutableStateResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DescribeMutableStateResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DescribeMutableStateResponse
	switch t := that.(type) {
	case *DescribeMutableStateResponse:
		that1 = t
	case DescribeMutableStateResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type DescribeHistoryHostRequest to the protobuf v3 wire format
func (val *DescribeHistoryHostRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type DescribeHistoryHostRequest from the protobuf v3 wire format
func (val *DescribeHistoryHostRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *DescribeHistoryHostRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two DescribeHistoryHostRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DescribeHistoryHostRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DescribeHistoryHostRequest
	switch t := that.(type) {
	case *DescribeHistoryHostRequest:
		that1 = t
	case DescribeHistoryHostRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type DescribeHistoryHostResponse to the protobuf v3 wire format
func (val *DescribeHistoryHostResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type DescribeHistoryHostResponse from the protobuf v3 wire format
func (val *DescribeHistoryHostResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *DescribeHistoryHostResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two DescribeHistoryHostResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DescribeHistoryHostResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DescribeHistoryHostResponse
	switch t := that.(type) {
	case *DescribeHistoryHostResponse:
		that1 = t
	case DescribeHistoryHostResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type CloseShardRequest to the protobuf v3 wire format
func (val *CloseShardRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type CloseShardRequest from the protobuf v3 wire format
func (val *CloseShardRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *CloseShardRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two CloseShardRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *CloseShardRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *CloseShardRequest
	switch t := that.(type) {
	case *CloseShardRequest:
		that1 = t
	case CloseShardRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type CloseShardResponse to the protobuf v3 wire format
func (val *CloseShardResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type CloseShardResponse from the protobuf v3 wire format
func (val *CloseShardResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *CloseShardResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two CloseShardResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *CloseShardResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *CloseShardResponse
	switch t := that.(type) {
	case *CloseShardResponse:
		that1 = t
	case CloseShardResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetShardRequest to the protobuf v3 wire format
func (val *GetShardRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetShardRequest from the protobuf v3 wire format
func (val *GetShardRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetShardRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetShardRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetShardRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetShardRequest
	switch t := that.(type) {
	case *GetShardRequest:
		that1 = t
	case GetShardRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetShardResponse to the protobuf v3 wire format
func (val *GetShardResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetShardResponse from the protobuf v3 wire format
func (val *GetShardResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetShardResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetShardResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetShardResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetShardResponse
	switch t := that.(type) {
	case *GetShardResponse:
		that1 = t
	case GetShardResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ListHistoryTasksRequest to the protobuf v3 wire format
func (val *ListHistoryTasksRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ListHistoryTasksRequest from the protobuf v3 wire format
func (val *ListHistoryTasksRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ListHistoryTasksRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ListHistoryTasksRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ListHistoryTasksRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ListHistoryTasksRequest
	switch t := that.(type) {
	case *ListHistoryTasksRequest:
		that1 = t
	case ListHistoryTasksRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ListHistoryTasksResponse to the protobuf v3 wire format
func (val *ListHistoryTasksResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ListHistoryTasksResponse from the protobuf v3 wire format
func (val *ListHistoryTasksResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ListHistoryTasksResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ListHistoryTasksResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ListHistoryTasksResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ListHistoryTasksResponse
	switch t := that.(type) {
	case *ListHistoryTasksResponse:
		that1 = t
	case ListHistoryTasksResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type Task to the protobuf v3 wire format
func (val *Task) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type Task from the protobuf v3 wire format
func (val *Task) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *Task) Size() int {
	return proto.Size(val)
}

// Equal returns whether two Task values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *Task) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *Task
	switch t := that.(type) {
	case *Task:
		that1 = t
	case Task:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RemoveTaskRequest to the protobuf v3 wire format
func (val *RemoveTaskRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RemoveTaskRequest from the protobuf v3 wire format
func (val *RemoveTaskRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RemoveTaskRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RemoveTaskRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RemoveTaskRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RemoveTaskRequest
	switch t := that.(type) {
	case *RemoveTaskRequest:
		that1 = t
	case RemoveTaskRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RemoveTaskResponse to the protobuf v3 wire format
func (val *RemoveTaskResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RemoveTaskResponse from the protobuf v3 wire format
func (val *RemoveTaskResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RemoveTaskResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RemoveTaskResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RemoveTaskResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RemoveTaskResponse
	switch t := that.(type) {
	case *RemoveTaskResponse:
		that1 = t
	case RemoveTaskResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetWorkflowExecutionRawHistoryV2Request to the protobuf v3 wire format
func (val *GetWorkflowExecutionRawHistoryV2Request) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetWorkflowExecutionRawHistoryV2Request from the protobuf v3 wire format
func (val *GetWorkflowExecutionRawHistoryV2Request) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetWorkflowExecutionRawHistoryV2Request) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetWorkflowExecutionRawHistoryV2Request values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetWorkflowExecutionRawHistoryV2Request) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetWorkflowExecutionRawHistoryV2Request
	switch t := that.(type) {
	case *GetWorkflowExecutionRawHistoryV2Request:
		that1 = t
	case GetWorkflowExecutionRawHistoryV2Request:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetWorkflowExecutionRawHistoryV2Response to the protobuf v3 wire format
func (val *GetWorkflowExecutionRawHistoryV2Response) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetWorkflowExecutionRawHistoryV2Response from the protobuf v3 wire format
func (val *GetWorkflowExecutionRawHistoryV2Response) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetWorkflowExecutionRawHistoryV2Response) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetWorkflowExecutionRawHistoryV2Response values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetWorkflowExecutionRawHistoryV2Response) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetWorkflowExecutionRawHistoryV2Response
	switch t := that.(type) {
	case *GetWorkflowExecutionRawHistoryV2Response:
		that1 = t
	case GetWorkflowExecutionRawHistoryV2Response:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetReplicationMessagesRequest to the protobuf v3 wire format
func (val *GetReplicationMessagesRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetReplicationMessagesRequest from the protobuf v3 wire format
func (val *GetReplicationMessagesRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetReplicationMessagesRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetReplicationMessagesRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetReplicationMessagesRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetReplicationMessagesRequest
	switch t := that.(type) {
	case *GetReplicationMessagesRequest:
		that1 = t
	case GetReplicationMessagesRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetReplicationMessagesResponse to the protobuf v3 wire format
func (val *GetReplicationMessagesResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetReplicationMessagesResponse from the protobuf v3 wire format
func (val *GetReplicationMessagesResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetReplicationMessagesResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetReplicationMessagesResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetReplicationMessagesResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetReplicationMessagesResponse
	switch t := that.(type) {
	case *GetReplicationMessagesResponse:
		that1 = t
	case GetReplicationMessagesResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetNamespaceReplicationMessagesRequest to the protobuf v3 wire format
func (val *GetNamespaceReplicationMessagesRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetNamespaceReplicationMessagesRequest from the protobuf v3 wire format
func (val *GetNamespaceReplicationMessagesRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetNamespaceReplicationMessagesRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetNamespaceReplicationMessagesRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetNamespaceReplicationMessagesRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetNamespaceReplicationMessagesRequest
	switch t := that.(type) {
	case *GetNamespaceReplicationMessagesRequest:
		that1 = t
	case GetNamespaceReplicationMessagesRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetNamespaceReplicationMessagesResponse to the protobuf v3 wire format
func (val *GetNamespaceReplicationMessagesResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetNamespaceReplicationMessagesResponse from the protobuf v3 wire format
func (val *GetNamespaceReplicationMessagesResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetNamespaceReplicationMessagesResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetNamespaceReplicationMessagesResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetNamespaceReplicationMessagesResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetNamespaceReplicationMessagesResponse
	switch t := that.(type) {
	case *GetNamespaceReplicationMessagesResponse:
		that1 = t
	case GetNamespaceReplicationMessagesResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetDLQReplicationMessagesRequest to the protobuf v3 wire format
func (val *GetDLQReplicationMessagesRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetDLQReplicationMessagesRequest from the protobuf v3 wire format
func (val *GetDLQReplicationMessagesRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetDLQReplicationMessagesRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetDLQReplicationMessagesRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetDLQReplicationMessagesRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetDLQReplicationMessagesRequest
	switch t := that.(type) {
	case *GetDLQReplicationMessagesRequest:
		that1 = t
	case GetDLQReplicationMessagesRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetDLQReplicationMessagesResponse to the protobuf v3 wire format
func (val *GetDLQReplicationMessagesResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetDLQReplicationMessagesResponse from the protobuf v3 wire format
func (val *GetDLQReplicationMessagesResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetDLQReplicationMessagesResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetDLQReplicationMessagesResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetDLQReplicationMessagesResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetDLQReplicationMessagesResponse
	switch t := that.(type) {
	case *GetDLQReplicationMessagesResponse:
		that1 = t
	case GetDLQReplicationMessagesResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ReapplyEventsRequest to the protobuf v3 wire format
func (val *ReapplyEventsRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ReapplyEventsRequest from the protobuf v3 wire format
func (val *ReapplyEventsRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ReapplyEventsRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ReapplyEventsRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ReapplyEventsRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ReapplyEventsRequest
	switch t := that.(type) {
	case *ReapplyEventsRequest:
		that1 = t
	case ReapplyEventsRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ReapplyEventsResponse to the protobuf v3 wire format
func (val *ReapplyEventsResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ReapplyEventsResponse from the protobuf v3 wire format
func (val *ReapplyEventsResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ReapplyEventsResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ReapplyEventsResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ReapplyEventsResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ReapplyEventsResponse
	switch t := that.(type) {
	case *ReapplyEventsResponse:
		that1 = t
	case ReapplyEventsResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type AddSearchAttributesRequest to the protobuf v3 wire format
func (val *AddSearchAttributesRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type AddSearchAttributesRequest from the protobuf v3 wire format
func (val *AddSearchAttributesRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *AddSearchAttributesRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two AddSearchAttributesRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *AddSearchAttributesRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *AddSearchAttributesRequest
	switch t := that.(type) {
	case *AddSearchAttributesRequest:
		that1 = t
	case AddSearchAttributesRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type AddSearchAttributesResponse to the protobuf v3 wire format
func (val *AddSearchAttributesResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type AddSearchAttributesResponse from the protobuf v3 wire format
func (val *AddSearchAttributesResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *AddSearchAttributesResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two AddSearchAttributesResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *AddSearchAttributesResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *AddSearchAttributesResponse
	switch t := that.(type) {
	case *AddSearchAttributesResponse:
		that1 = t
	case AddSearchAttributesResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RemoveSearchAttributesRequest to the protobuf v3 wire format
func (val *RemoveSearchAttributesRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RemoveSearchAttributesRequest from the protobuf v3 wire format
func (val *RemoveSearchAttributesRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RemoveSearchAttributesRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RemoveSearchAttributesRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RemoveSearchAttributesRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RemoveSearchAttributesRequest
	switch t := that.(type) {
	case *RemoveSearchAttributesRequest:
		that1 = t
	case RemoveSearchAttributesRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RemoveSearchAttributesResponse to the protobuf v3 wire format
func (val *RemoveSearchAttributesResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RemoveSearchAttributesResponse from the protobuf v3 wire format
func (val *RemoveSearchAttributesResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RemoveSearchAttributesResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RemoveSearchAttributesResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RemoveSearchAttributesResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RemoveSearchAttributesResponse
	switch t := that.(type) {
	case *RemoveSearchAttributesResponse:
		that1 = t
	case RemoveSearchAttributesResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetSearchAttributesRequest to the protobuf v3 wire format
func (val *GetSearchAttributesRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetSearchAttributesRequest from the protobuf v3 wire format
func (val *GetSearchAttributesRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetSearchAttributesRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetSearchAttributesRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetSearchAttributesRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetSearchAttributesRequest
	switch t := that.(type) {
	case *GetSearchAttributesRequest:
		that1 = t
	case GetSearchAttributesRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetSearchAttributesResponse to the protobuf v3 wire format
func (val *GetSearchAttributesResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetSearchAttributesResponse from the protobuf v3 wire format
func (val *GetSearchAttributesResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetSearchAttributesResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetSearchAttributesResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetSearchAttributesResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetSearchAttributesResponse
	switch t := that.(type) {
	case *GetSearchAttributesResponse:
		that1 = t
	case GetSearchAttributesResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type DescribeClusterRequest to the protobuf v3 wire format
func (val *DescribeClusterRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type DescribeClusterRequest from the protobuf v3 wire format
func (val *DescribeClusterRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *DescribeClusterRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two DescribeClusterRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DescribeClusterRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DescribeClusterRequest
	switch t := that.(type) {
	case *DescribeClusterRequest:
		that1 = t
	case DescribeClusterRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type DescribeClusterResponse to the protobuf v3 wire format
func (val *DescribeClusterResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type DescribeClusterResponse from the protobuf v3 wire format
func (val *DescribeClusterResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *DescribeClusterResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two DescribeClusterResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DescribeClusterResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DescribeClusterResponse
	switch t := that.(type) {
	case *DescribeClusterResponse:
		that1 = t
	case DescribeClusterResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ListClustersRequest to the protobuf v3 wire format
func (val *ListClustersRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ListClustersRequest from the protobuf v3 wire format
func (val *ListClustersRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ListClustersRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ListClustersRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ListClustersRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ListClustersRequest
	switch t := that.(type) {
	case *ListClustersRequest:
		that1 = t
	case ListClustersRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ListClustersResponse to the protobuf v3 wire format
func (val *ListClustersResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ListClustersResponse from the protobuf v3 wire format
func (val *ListClustersResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ListClustersResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ListClustersResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ListClustersResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ListClustersResponse
	switch t := that.(type) {
	case *ListClustersResponse:
		that1 = t
	case ListClustersResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type AddOrUpdateRemoteClusterRequest to the protobuf v3 wire format
func (val *AddOrUpdateRemoteClusterRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type AddOrUpdateRemoteClusterRequest from the protobuf v3 wire format
func (val *AddOrUpdateRemoteClusterRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *AddOrUpdateRemoteClusterRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two AddOrUpdateRemoteClusterRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *AddOrUpdateRemoteClusterRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *AddOrUpdateRemoteClusterRequest
	switch t := that.(type) {
	case *AddOrUpdateRemoteClusterRequest:
		that1 = t
	case AddOrUpdateRemoteClusterRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type AddOrUpdateRemoteClusterResponse to the protobuf v3 wire format
func (val *AddOrUpdateRemoteClusterResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type AddOrUpdateRemoteClusterResponse from the protobuf v3 wire format
func (val *AddOrUpdateRemoteClusterResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *AddOrUpdateRemoteClusterResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two AddOrUpdateRemoteClusterResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *AddOrUpdateRemoteClusterResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *AddOrUpdateRemoteClusterResponse
	switch t := that.(type) {
	case *AddOrUpdateRemoteClusterResponse:
		that1 = t
	case AddOrUpdateRemoteClusterResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RemoveRemoteClusterRequest to the protobuf v3 wire format
func (val *RemoveRemoteClusterRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RemoveRemoteClusterRequest from the protobuf v3 wire format
func (val *RemoveRemoteClusterRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RemoveRemoteClusterRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RemoveRemoteClusterRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RemoveRemoteClusterRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RemoveRemoteClusterRequest
	switch t := that.(type) {
	case *RemoveRemoteClusterRequest:
		that1 = t
	case RemoveRemoteClusterRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RemoveRemoteClusterResponse to the protobuf v3 wire format
func (val *RemoveRemoteClusterResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RemoveRemoteClusterResponse from the protobuf v3 wire format
func (val *RemoveRemoteClusterResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RemoveRemoteClusterResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RemoveRemoteClusterResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RemoveRemoteClusterResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RemoveRemoteClusterResponse
	switch t := that.(type) {
	case *RemoveRemoteClusterResponse:
		that1 = t
	case RemoveRemoteClusterResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ListClusterMembersRequest to the protobuf v3 wire format
func (val *ListClusterMembersRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ListClusterMembersRequest from the protobuf v3 wire format
func (val *ListClusterMembersRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ListClusterMembersRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ListClusterMembersRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ListClusterMembersRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ListClusterMembersRequest
	switch t := that.(type) {
	case *ListClusterMembersRequest:
		that1 = t
	case ListClusterMembersRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ListClusterMembersResponse to the protobuf v3 wire format
func (val *ListClusterMembersResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ListClusterMembersResponse from the protobuf v3 wire format
func (val *ListClusterMembersResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ListClusterMembersResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ListClusterMembersResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ListClusterMembersResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ListClusterMembersResponse
	switch t := that.(type) {
	case *ListClusterMembersResponse:
		that1 = t
	case ListClusterMembersResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetDLQMessagesRequest to the protobuf v3 wire format
func (val *GetDLQMessagesRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetDLQMessagesRequest from the protobuf v3 wire format
func (val *GetDLQMessagesRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetDLQMessagesRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetDLQMessagesRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetDLQMessagesRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetDLQMessagesRequest
	switch t := that.(type) {
	case *GetDLQMessagesRequest:
		that1 = t
	case GetDLQMessagesRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetDLQMessagesResponse to the protobuf v3 wire format
func (val *GetDLQMessagesResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetDLQMessagesResponse from the protobuf v3 wire format
func (val *GetDLQMessagesResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetDLQMessagesResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetDLQMessagesResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetDLQMessagesResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetDLQMessagesResponse
	switch t := that.(type) {
	case *GetDLQMessagesResponse:
		that1 = t
	case GetDLQMessagesResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type PurgeDLQMessagesRequest to the protobuf v3 wire format
func (val *PurgeDLQMessagesRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type PurgeDLQMessagesRequest from the protobuf v3 wire format
func (val *PurgeDLQMessagesRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *PurgeDLQMessagesRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two PurgeDLQMessagesRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *PurgeDLQMessagesRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *PurgeDLQMessagesRequest
	switch t := that.(type) {
	case *PurgeDLQMessagesRequest:
		that1 = t
	case PurgeDLQMessagesRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type PurgeDLQMessagesResponse to the protobuf v3 wire format
func (val *PurgeDLQMessagesResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type PurgeDLQMessagesResponse from the protobuf v3 wire format
func (val *PurgeDLQMessagesResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *PurgeDLQMessagesResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two PurgeDLQMessagesResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *PurgeDLQMessagesResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *PurgeDLQMessagesResponse
	switch t := that.(type) {
	case *PurgeDLQMessagesResponse:
		that1 = t
	case PurgeDLQMessagesResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type MergeDLQMessagesRequest to the protobuf v3 wire format
func (val *MergeDLQMessagesRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type MergeDLQMessagesRequest from the protobuf v3 wire format
func (val *MergeDLQMessagesRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *MergeDLQMessagesRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two MergeDLQMessagesRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *MergeDLQMessagesRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *MergeDLQMessagesRequest
	switch t := that.(type) {
	case *MergeDLQMessagesRequest:
		that1 = t
	case MergeDLQMessagesRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type MergeDLQMessagesResponse to the protobuf v3 wire format
func (val *MergeDLQMessagesResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type MergeDLQMessagesResponse from the protobuf v3 wire format
func (val *MergeDLQMessagesResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *MergeDLQMessagesResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two MergeDLQMessagesResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *MergeDLQMessagesResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *MergeDLQMessagesResponse
	switch t := that.(type) {
	case *MergeDLQMessagesResponse:
		that1 = t
	case MergeDLQMessagesResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RefreshWorkflowTasksRequest to the protobuf v3 wire format
func (val *RefreshWorkflowTasksRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RefreshWorkflowTasksRequest from the protobuf v3 wire format
func (val *RefreshWorkflowTasksRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RefreshWorkflowTasksRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RefreshWorkflowTasksRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RefreshWorkflowTasksRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RefreshWorkflowTasksRequest
	switch t := that.(type) {
	case *RefreshWorkflowTasksRequest:
		that1 = t
	case RefreshWorkflowTasksRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RefreshWorkflowTasksResponse to the protobuf v3 wire format
func (val *RefreshWorkflowTasksResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RefreshWorkflowTasksResponse from the protobuf v3 wire format
func (val *RefreshWorkflowTasksResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RefreshWorkflowTasksResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RefreshWorkflowTasksResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RefreshWorkflowTasksResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RefreshWorkflowTasksResponse
	switch t := that.(type) {
	case *RefreshWorkflowTasksResponse:
		that1 = t
	case RefreshWorkflowTasksResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ResendReplicationTasksRequest to the protobuf v3 wire format
func (val *ResendReplicationTasksRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ResendReplicationTasksRequest from the protobuf v3 wire format
func (val *ResendReplicationTasksRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ResendReplicationTasksRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ResendReplicationTasksRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ResendReplicationTasksRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ResendReplicationTasksRequest
	switch t := that.(type) {
	case *ResendReplicationTasksRequest:
		that1 = t
	case ResendReplicationTasksRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ResendReplicationTasksResponse to the protobuf v3 wire format
func (val *ResendReplicationTasksResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ResendReplicationTasksResponse from the protobuf v3 wire format
func (val *ResendReplicationTasksResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ResendReplicationTasksResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ResendReplicationTasksResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ResendReplicationTasksResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ResendReplicationTasksResponse
	switch t := that.(type) {
	case *ResendReplicationTasksResponse:
		that1 = t
	case ResendReplicationTasksResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetTaskQueueTasksRequest to the protobuf v3 wire format
func (val *GetTaskQueueTasksRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetTaskQueueTasksRequest from the protobuf v3 wire format
func (val *GetTaskQueueTasksRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetTaskQueueTasksRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetTaskQueueTasksRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetTaskQueueTasksRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetTaskQueueTasksRequest
	switch t := that.(type) {
	case *GetTaskQueueTasksRequest:
		that1 = t
	case GetTaskQueueTasksRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetTaskQueueTasksResponse to the protobuf v3 wire format
func (val *GetTaskQueueTasksResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetTaskQueueTasksResponse from the protobuf v3 wire format
func (val *GetTaskQueueTasksResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetTaskQueueTasksResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetTaskQueueTasksResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetTaskQueueTasksResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetTaskQueueTasksResponse
	switch t := that.(type) {
	case *GetTaskQueueTasksResponse:
		that1 = t
	case GetTaskQueueTasksResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type DeleteWorkflowExecutionRequest to the protobuf v3 wire format
func (val *DeleteWorkflowExecutionRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type DeleteWorkflowExecutionRequest from the protobuf v3 wire format
func (val *DeleteWorkflowExecutionRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *DeleteWorkflowExecutionRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two DeleteWorkflowExecutionRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DeleteWorkflowExecutionRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DeleteWorkflowExecutionRequest
	switch t := that.(type) {
	case *DeleteWorkflowExecutionRequest:
		that1 = t
	case DeleteWorkflowExecutionRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type DeleteWorkflowExecutionResponse to the protobuf v3 wire format
func (val *DeleteWorkflowExecutionResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type DeleteWorkflowExecutionResponse from the protobuf v3 wire format
func (val *DeleteWorkflowExecutionResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *DeleteWorkflowExecutionResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two DeleteWorkflowExecutionResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DeleteWorkflowExecutionResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DeleteWorkflowExecutionResponse
	switch t := that.(type) {
	case *DeleteWorkflowExecutionResponse:
		that1 = t
	case DeleteWorkflowExecutionResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type StreamWorkflowReplicationMessagesRequest to the protobuf v3 wire format
func (val *StreamWorkflowReplicationMessagesRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type StreamWorkflowReplicationMessagesRequest from the protobuf v3 wire format
func (val *StreamWorkflowReplicationMessagesRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *StreamWorkflowReplicationMessagesRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two StreamWorkflowReplicationMessagesRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *StreamWorkflowReplicationMessagesRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *StreamWorkflowReplicationMessagesRequest
	switch t := that.(type) {
	case *StreamWorkflowReplicationMessagesRequest:
		that1 = t
	case StreamWorkflowReplicationMessagesRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type StreamWorkflowReplicationMessagesResponse to the protobuf v3 wire format
func (val *StreamWorkflowReplicationMessagesResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type StreamWorkflowReplicationMessagesResponse from the protobuf v3 wire format
func (val *StreamWorkflowReplicationMessagesResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *StreamWorkflowReplicationMessagesResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two StreamWorkflowReplicationMessagesResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *StreamWorkflowReplicationMessagesResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *StreamWorkflowReplicationMessagesResponse
	switch t := that.(type) {
	case *StreamWorkflowReplicationMessagesResponse:
		that1 = t
	case StreamWorkflowReplicationMessagesResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetNamespaceRequest to the protobuf v3 wire format
func (val *GetNamespaceRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetNamespaceRequest from the protobuf v3 wire format
func (val *GetNamespaceRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetNamespaceRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetNamespaceRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetNamespaceRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetNamespaceRequest
	switch t := that.(type) {
	case *GetNamespaceRequest:
		that1 = t
	case GetNamespaceRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetNamespaceResponse to the protobuf v3 wire format
func (val *GetNamespaceResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetNamespaceResponse from the protobuf v3 wire format
func (val *GetNamespaceResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetNamespaceResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetNamespaceResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetNamespaceResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetNamespaceResponse
	switch t := that.(type) {
	case *GetNamespaceResponse:
		that1 = t
	case GetNamespaceResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetDLQTasksRequest to the protobuf v3 wire format
func (val *GetDLQTasksRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetDLQTasksRequest from the protobuf v3 wire format
func (val *GetDLQTasksRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetDLQTasksRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetDLQTasksRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetDLQTasksRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetDLQTasksRequest
	switch t := that.(type) {
	case *GetDLQTasksRequest:
		that1 = t
	case GetDLQTasksRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetDLQTasksResponse to the protobuf v3 wire format
func (val *GetDLQTasksResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetDLQTasksResponse from the protobuf v3 wire format
func (val *GetDLQTasksResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetDLQTasksResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetDLQTasksResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetDLQTasksResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetDLQTasksResponse
	switch t := that.(type) {
	case *GetDLQTasksResponse:
		that1 = t
	case GetDLQTasksResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type PurgeDLQTasksRequest to the protobuf v3 wire format
func (val *PurgeDLQTasksRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type PurgeDLQTasksRequest from the protobuf v3 wire format
func (val *PurgeDLQTasksRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *PurgeDLQTasksRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two PurgeDLQTasksRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *PurgeDLQTasksRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *PurgeDLQTasksRequest
	switch t := that.(type) {
	case *PurgeDLQTasksRequest:
		that1 = t
	case PurgeDLQTasksRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type PurgeDLQTasksResponse to the protobuf v3 wire format
func (val *PurgeDLQTasksResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type PurgeDLQTasksResponse from the protobuf v3 wire format
func (val *PurgeDLQTasksResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *PurgeDLQTasksResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two PurgeDLQTasksResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *PurgeDLQTasksResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *PurgeDLQTasksResponse
	switch t := that.(type) {
	case *PurgeDLQTasksResponse:
		that1 = t
	case PurgeDLQTasksResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type DLQJobToken to the protobuf v3 wire format
func (val *DLQJobToken) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type DLQJobToken from the protobuf v3 wire format
func (val *DLQJobToken) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *DLQJobToken) Size() int {
	return proto.Size(val)
}

// Equal returns whether two DLQJobToken values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DLQJobToken) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DLQJobToken
	switch t := that.(type) {
	case *DLQJobToken:
		that1 = t
	case DLQJobToken:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type MergeDLQTasksRequest to the protobuf v3 wire format
func (val *MergeDLQTasksRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type MergeDLQTasksRequest from the protobuf v3 wire format
func (val *MergeDLQTasksRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *MergeDLQTasksRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two MergeDLQTasksRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *MergeDLQTasksRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *MergeDLQTasksRequest
	switch t := that.(type) {
	case *MergeDLQTasksRequest:
		that1 = t
	case MergeDLQTasksRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type MergeDLQTasksResponse to the protobuf v3 wire format
func (val *MergeDLQTasksResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type MergeDLQTasksResponse from the protobuf v3 wire format
func (val *MergeDLQTasksResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *MergeDLQTasksResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two MergeDLQTasksResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *MergeDLQTasksResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *MergeDLQTasksResponse
	switch t := that.(type) {
	case *MergeDLQTasksResponse:
		that1 = t
	case MergeDLQTasksResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type DescribeDLQJobRequest to the protobuf v3 wire format
func (val *DescribeDLQJobRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type DescribeDLQJobRequest from the protobuf v3 wire format
func (val *DescribeDLQJobRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *DescribeDLQJobRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two DescribeDLQJobRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DescribeDLQJobRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DescribeDLQJobRequest
	switch t := that.(type) {
	case *DescribeDLQJobRequest:
		that1 = t
	case DescribeDLQJobRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type DescribeDLQJobResponse to the protobuf v3 wire format
func (val *DescribeDLQJobResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type DescribeDLQJobResponse from the protobuf v3 wire format
func (val *DescribeDLQJobResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *DescribeDLQJobResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two DescribeDLQJobResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DescribeDLQJobResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DescribeDLQJobResponse
	switch t := that.(type) {
	case *DescribeDLQJobResponse:
		that1 = t
	case DescribeDLQJobResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type CancelDLQJobRequest to the protobuf v3 wire format
func (val *CancelDLQJobRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type CancelDLQJobRequest from the protobuf v3 wire format
func (val *CancelDLQJobRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *CancelDLQJobRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two CancelDLQJobRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *CancelDLQJobRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *CancelDLQJobRequest
	switch t := that.(type) {
	case *CancelDLQJobRequest:
		that1 = t
	case CancelDLQJobRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type CancelDLQJobResponse to the protobuf v3 wire format
func (val *CancelDLQJobResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type CancelDLQJobResponse from the protobuf v3 wire format
func (val *CancelDLQJobResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *CancelDLQJobResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two CancelDLQJobResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *CancelDLQJobResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *CancelDLQJobResponse
	switch t := that.(type) {
	case *CancelDLQJobResponse:
		that1 = t
	case CancelDLQJobResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type AddTasksRequest to the protobuf v3 wire format
func (val *AddTasksRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type AddTasksRequest from the protobuf v3 wire format
func (val *AddTasksRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *AddTasksRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two AddTasksRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *AddTasksRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *AddTasksRequest
	switch t := that.(type) {
	case *AddTasksRequest:
		that1 = t
	case AddTasksRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type AddTasksResponse to the protobuf v3 wire format
func (val *AddTasksResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type AddTasksResponse from the protobuf v3 wire format
func (val *AddTasksResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *AddTasksResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two AddTasksResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *AddTasksResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *AddTasksResponse
	switch t := that.(type) {
	case *AddTasksResponse:
		that1 = t
	case AddTasksResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ListQueuesRequest to the protobuf v3 wire format
func (val *ListQueuesRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ListQueuesRequest from the protobuf v3 wire format
func (val *ListQueuesRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ListQueuesRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ListQueuesRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ListQueuesRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ListQueuesRequest
	switch t := that.(type) {
	case *ListQueuesRequest:
		that1 = t
	case ListQueuesRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ListQueuesResponse to the protobuf v3 wire format
func (val *ListQueuesResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ListQueuesResponse from the protobuf v3 wire format
func (val *ListQueuesResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ListQueuesResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ListQueuesResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ListQueuesResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ListQueuesResponse
	switch t := that.(type) {
	case *ListQueuesResponse:
		that1 = t
	case ListQueuesResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
