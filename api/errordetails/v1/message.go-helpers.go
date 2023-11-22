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

package errordetails

import (
	"google.golang.org/protobuf/proto"
)

// Marshal an object of type TaskAlreadyStartedFailure to the protobuf v3 wire format
func (val *TaskAlreadyStartedFailure) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type TaskAlreadyStartedFailure from the protobuf v3 wire format
func (val *TaskAlreadyStartedFailure) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *TaskAlreadyStartedFailure) Size() int {
	return proto.Size(val)
}

// Equal returns whether two TaskAlreadyStartedFailure values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *TaskAlreadyStartedFailure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *TaskAlreadyStartedFailure
	switch t := that.(type) {
	case *TaskAlreadyStartedFailure:
		that1 = t
	case TaskAlreadyStartedFailure:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type CurrentBranchChangedFailure to the protobuf v3 wire format
func (val *CurrentBranchChangedFailure) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type CurrentBranchChangedFailure from the protobuf v3 wire format
func (val *CurrentBranchChangedFailure) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *CurrentBranchChangedFailure) Size() int {
	return proto.Size(val)
}

// Equal returns whether two CurrentBranchChangedFailure values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *CurrentBranchChangedFailure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *CurrentBranchChangedFailure
	switch t := that.(type) {
	case *CurrentBranchChangedFailure:
		that1 = t
	case CurrentBranchChangedFailure:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ShardOwnershipLostFailure to the protobuf v3 wire format
func (val *ShardOwnershipLostFailure) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ShardOwnershipLostFailure from the protobuf v3 wire format
func (val *ShardOwnershipLostFailure) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ShardOwnershipLostFailure) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ShardOwnershipLostFailure values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ShardOwnershipLostFailure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ShardOwnershipLostFailure
	switch t := that.(type) {
	case *ShardOwnershipLostFailure:
		that1 = t
	case ShardOwnershipLostFailure:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RetryReplicationFailure to the protobuf v3 wire format
func (val *RetryReplicationFailure) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RetryReplicationFailure from the protobuf v3 wire format
func (val *RetryReplicationFailure) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RetryReplicationFailure) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RetryReplicationFailure values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RetryReplicationFailure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RetryReplicationFailure
	switch t := that.(type) {
	case *RetryReplicationFailure:
		that1 = t
	case RetryReplicationFailure:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type StickyWorkerUnavailableFailure to the protobuf v3 wire format
func (val *StickyWorkerUnavailableFailure) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type StickyWorkerUnavailableFailure from the protobuf v3 wire format
func (val *StickyWorkerUnavailableFailure) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *StickyWorkerUnavailableFailure) Size() int {
	return proto.Size(val)
}

// Equal returns whether two StickyWorkerUnavailableFailure values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *StickyWorkerUnavailableFailure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *StickyWorkerUnavailableFailure
	switch t := that.(type) {
	case *StickyWorkerUnavailableFailure:
		that1 = t
	case StickyWorkerUnavailableFailure:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
