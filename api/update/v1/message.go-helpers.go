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

package update

import (
	"google.golang.org/protobuf/proto"
)

// Marshal an object of type AcceptanceInfo to the protobuf v3 wire format
func (val *AcceptanceInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type AcceptanceInfo from the protobuf v3 wire format
func (val *AcceptanceInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *AcceptanceInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two AcceptanceInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *AcceptanceInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *AcceptanceInfo
	switch t := that.(type) {
	case *AcceptanceInfo:
		that1 = t
	case AcceptanceInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type CompletionInfo to the protobuf v3 wire format
func (val *CompletionInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type CompletionInfo from the protobuf v3 wire format
func (val *CompletionInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *CompletionInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two CompletionInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *CompletionInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *CompletionInfo
	switch t := that.(type) {
	case *CompletionInfo:
		that1 = t
	case CompletionInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type UpdateInfo to the protobuf v3 wire format
func (val *UpdateInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type UpdateInfo from the protobuf v3 wire format
func (val *UpdateInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *UpdateInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two UpdateInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *UpdateInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *UpdateInfo
	switch t := that.(type) {
	case *UpdateInfo:
		that1 = t
	case UpdateInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
