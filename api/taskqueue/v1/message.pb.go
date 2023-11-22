// The MIT License
//
// Copyright (c) 2020 Temporal Technologies, Inc.
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// source: temporal/server/api/taskqueue/v1/message.proto

package taskqueue

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// TaskVersionDirective controls how matching should direct a task.
type TaskVersionDirective struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Default (if value is not present) is "unversioned":
	// Use the unversioned task queue, even if the task queue has versioning data.
	//
	// Types that are assignable to Value:
	//
	//	*TaskVersionDirective_UseDefault
	//	*TaskVersionDirective_BuildId
	Value isTaskVersionDirective_Value `protobuf_oneof:"value"`
}

func (x *TaskVersionDirective) Reset() {
	*x = TaskVersionDirective{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_server_api_taskqueue_v1_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskVersionDirective) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskVersionDirective) ProtoMessage() {}

func (x *TaskVersionDirective) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_taskqueue_v1_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskVersionDirective.ProtoReflect.Descriptor instead.
func (*TaskVersionDirective) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_taskqueue_v1_message_proto_rawDescGZIP(), []int{0}
}

func (m *TaskVersionDirective) GetValue() isTaskVersionDirective_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *TaskVersionDirective) GetUseDefault() *emptypb.Empty {
	if x, ok := x.GetValue().(*TaskVersionDirective_UseDefault); ok {
		return x.UseDefault
	}
	return nil
}

func (x *TaskVersionDirective) GetBuildId() string {
	if x, ok := x.GetValue().(*TaskVersionDirective_BuildId); ok {
		return x.BuildId
	}
	return ""
}

type isTaskVersionDirective_Value interface {
	isTaskVersionDirective_Value()
}

type TaskVersionDirective_UseDefault struct {
	// If use_default is present, the task should be assigned the default
	// version for the task queue. This will typically be set for the first
	// workflow task in a workflow.
	UseDefault *emptypb.Empty `protobuf:"bytes,1,opt,name=use_default,json=useDefault,proto3,oneof"`
}

type TaskVersionDirective_BuildId struct {
	// If build_id is present, use the default version in the compatible set
	// containing this build id.
	BuildId string `protobuf:"bytes,2,opt,name=build_id,json=buildId,proto3,oneof"`
}

func (*TaskVersionDirective_UseDefault) isTaskVersionDirective_Value() {}

func (*TaskVersionDirective_BuildId) isTaskVersionDirective_Value() {}

var File_temporal_server_api_taskqueue_v1_message_proto protoreflect.FileDescriptor

var file_temporal_server_api_taskqueue_v1_message_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x20, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e,
	0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x77, 0x0a, 0x14, 0x54, 0x61, 0x73, 0x6b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x39, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x5f, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x48, 0x00, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x44, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x12, 0x1b, 0x0a, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x42,
	0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x6f, 0x2e, 0x74,
	0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2f,
	0x76, 0x31, 0x3b, 0x74, 0x61, 0x73, 0x6b, 0x71, 0x75, 0x65, 0x75, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temporal_server_api_taskqueue_v1_message_proto_rawDescOnce sync.Once
	file_temporal_server_api_taskqueue_v1_message_proto_rawDescData = file_temporal_server_api_taskqueue_v1_message_proto_rawDesc
)

func file_temporal_server_api_taskqueue_v1_message_proto_rawDescGZIP() []byte {
	file_temporal_server_api_taskqueue_v1_message_proto_rawDescOnce.Do(func() {
		file_temporal_server_api_taskqueue_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_server_api_taskqueue_v1_message_proto_rawDescData)
	})
	return file_temporal_server_api_taskqueue_v1_message_proto_rawDescData
}

var file_temporal_server_api_taskqueue_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_temporal_server_api_taskqueue_v1_message_proto_goTypes = []interface{}{
	(*TaskVersionDirective)(nil), // 0: temporal.server.api.taskqueue.v1.TaskVersionDirective
	(*emptypb.Empty)(nil),        // 1: google.protobuf.Empty
}
var file_temporal_server_api_taskqueue_v1_message_proto_depIdxs = []int32{
	1, // 0: temporal.server.api.taskqueue.v1.TaskVersionDirective.use_default:type_name -> google.protobuf.Empty
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_temporal_server_api_taskqueue_v1_message_proto_init() }
func file_temporal_server_api_taskqueue_v1_message_proto_init() {
	if File_temporal_server_api_taskqueue_v1_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_temporal_server_api_taskqueue_v1_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskVersionDirective); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_temporal_server_api_taskqueue_v1_message_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*TaskVersionDirective_UseDefault)(nil),
		(*TaskVersionDirective_BuildId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_temporal_server_api_taskqueue_v1_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_server_api_taskqueue_v1_message_proto_goTypes,
		DependencyIndexes: file_temporal_server_api_taskqueue_v1_message_proto_depIdxs,
		MessageInfos:      file_temporal_server_api_taskqueue_v1_message_proto_msgTypes,
	}.Build()
	File_temporal_server_api_taskqueue_v1_message_proto = out.File
	file_temporal_server_api_taskqueue_v1_message_proto_rawDesc = nil
	file_temporal_server_api_taskqueue_v1_message_proto_goTypes = nil
	file_temporal_server_api_taskqueue_v1_message_proto_depIdxs = nil
}
