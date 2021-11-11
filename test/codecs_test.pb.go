// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: codecs_test.proto

package test

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	pmongo "github.com/viamrobotics/mongo-go-driver-protobuf/pmongo"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoolValue   *wrappers.BoolValue   `protobuf:"bytes,1,opt,name=boolValue,proto3" json:"boolValue,omitempty"`
	BytesValue  *wrappers.BytesValue  `protobuf:"bytes,2,opt,name=bytesValue,proto3" json:"bytesValue,omitempty"`
	DoubleValue *wrappers.DoubleValue `protobuf:"bytes,3,opt,name=doubleValue,proto3" json:"doubleValue,omitempty"`
	FloatValue  *wrappers.FloatValue  `protobuf:"bytes,4,opt,name=floatValue,proto3" json:"floatValue,omitempty"`
	Int32Value  *wrappers.Int32Value  `protobuf:"bytes,5,opt,name=int32Value,proto3" json:"int32Value,omitempty"`
	Int64Value  *wrappers.Int64Value  `protobuf:"bytes,6,opt,name=int64Value,proto3" json:"int64Value,omitempty"`
	StringValue *wrappers.StringValue `protobuf:"bytes,7,opt,name=stringValue,proto3" json:"stringValue,omitempty"`
	Uint32Value *wrappers.UInt32Value `protobuf:"bytes,8,opt,name=uint32Value,proto3" json:"uint32Value,omitempty"`
	Uint64Value *wrappers.UInt64Value `protobuf:"bytes,9,opt,name=uint64Value,proto3" json:"uint64Value,omitempty"`
	Timestamp   *timestamp.Timestamp  `protobuf:"bytes,10,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Id          *pmongo.ObjectId      `protobuf:"bytes,11,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_codecs_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_codecs_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_codecs_test_proto_rawDescGZIP(), []int{0}
}

func (x *Data) GetBoolValue() *wrappers.BoolValue {
	if x != nil {
		return x.BoolValue
	}
	return nil
}

func (x *Data) GetBytesValue() *wrappers.BytesValue {
	if x != nil {
		return x.BytesValue
	}
	return nil
}

func (x *Data) GetDoubleValue() *wrappers.DoubleValue {
	if x != nil {
		return x.DoubleValue
	}
	return nil
}

func (x *Data) GetFloatValue() *wrappers.FloatValue {
	if x != nil {
		return x.FloatValue
	}
	return nil
}

func (x *Data) GetInt32Value() *wrappers.Int32Value {
	if x != nil {
		return x.Int32Value
	}
	return nil
}

func (x *Data) GetInt64Value() *wrappers.Int64Value {
	if x != nil {
		return x.Int64Value
	}
	return nil
}

func (x *Data) GetStringValue() *wrappers.StringValue {
	if x != nil {
		return x.StringValue
	}
	return nil
}

func (x *Data) GetUint32Value() *wrappers.UInt32Value {
	if x != nil {
		return x.Uint32Value
	}
	return nil
}

func (x *Data) GetUint64Value() *wrappers.UInt64Value {
	if x != nil {
		return x.Uint64Value
	}
	return nil
}

func (x *Data) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Data) GetId() *pmongo.ObjectId {
	if x != nil {
		return x.Id
	}
	return nil
}

var File_codecs_test_proto protoreflect.FileDescriptor

var file_codecs_test_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x73, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x70, 0x6d, 0x6f, 0x6e,
	0x67, 0x6f, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x90, 0x05, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x38, 0x0a, 0x09, 0x62, 0x6f,
	0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6c, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x73,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x0a, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3b,
	0x0a, 0x0a, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x0a, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x69, 0x6e,
	0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x75, 0x69, 0x6e, 0x74,
	0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x75, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x75, 0x69, 0x6e, 0x74,
	0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x75, 0x69, 0x6e,
	0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x20, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x70, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x52, 0x02, 0x69, 0x64, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_codecs_test_proto_rawDescOnce sync.Once
	file_codecs_test_proto_rawDescData = file_codecs_test_proto_rawDesc
)

func file_codecs_test_proto_rawDescGZIP() []byte {
	file_codecs_test_proto_rawDescOnce.Do(func() {
		file_codecs_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_codecs_test_proto_rawDescData)
	})
	return file_codecs_test_proto_rawDescData
}

var file_codecs_test_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_codecs_test_proto_goTypes = []interface{}{
	(*Data)(nil),                 // 0: test.Data
	(*wrappers.BoolValue)(nil),   // 1: google.protobuf.BoolValue
	(*wrappers.BytesValue)(nil),  // 2: google.protobuf.BytesValue
	(*wrappers.DoubleValue)(nil), // 3: google.protobuf.DoubleValue
	(*wrappers.FloatValue)(nil),  // 4: google.protobuf.FloatValue
	(*wrappers.Int32Value)(nil),  // 5: google.protobuf.Int32Value
	(*wrappers.Int64Value)(nil),  // 6: google.protobuf.Int64Value
	(*wrappers.StringValue)(nil), // 7: google.protobuf.StringValue
	(*wrappers.UInt32Value)(nil), // 8: google.protobuf.UInt32Value
	(*wrappers.UInt64Value)(nil), // 9: google.protobuf.UInt64Value
	(*timestamp.Timestamp)(nil),  // 10: google.protobuf.Timestamp
	(*pmongo.ObjectId)(nil),      // 11: pmongo.ObjectId
}
var file_codecs_test_proto_depIdxs = []int32{
	1,  // 0: test.Data.boolValue:type_name -> google.protobuf.BoolValue
	2,  // 1: test.Data.bytesValue:type_name -> google.protobuf.BytesValue
	3,  // 2: test.Data.doubleValue:type_name -> google.protobuf.DoubleValue
	4,  // 3: test.Data.floatValue:type_name -> google.protobuf.FloatValue
	5,  // 4: test.Data.int32Value:type_name -> google.protobuf.Int32Value
	6,  // 5: test.Data.int64Value:type_name -> google.protobuf.Int64Value
	7,  // 6: test.Data.stringValue:type_name -> google.protobuf.StringValue
	8,  // 7: test.Data.uint32Value:type_name -> google.protobuf.UInt32Value
	9,  // 8: test.Data.uint64Value:type_name -> google.protobuf.UInt64Value
	10, // 9: test.Data.timestamp:type_name -> google.protobuf.Timestamp
	11, // 10: test.Data.id:type_name -> pmongo.ObjectId
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_codecs_test_proto_init() }
func file_codecs_test_proto_init() {
	if File_codecs_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_codecs_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_codecs_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_codecs_test_proto_goTypes,
		DependencyIndexes: file_codecs_test_proto_depIdxs,
		MessageInfos:      file_codecs_test_proto_msgTypes,
	}.Build()
	File_codecs_test_proto = out.File
	file_codecs_test_proto_rawDesc = nil
	file_codecs_test_proto_goTypes = nil
	file_codecs_test_proto_depIdxs = nil
}
