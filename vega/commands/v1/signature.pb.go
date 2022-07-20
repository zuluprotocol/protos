// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: vega/commands/v1/signature.proto

package v1

import (
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

// A signature to authenticate a transaction and to be verified by the Vega
// network.
type Signature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The bytes of the signature (hex-encoded).
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	// The algorithm used to create the signature.
	Algo string `protobuf:"bytes,2,opt,name=algo,proto3" json:"algo,omitempty"`
	// The version of the signature used to create the signature.
	Version uint32 `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Signature) Reset() {
	*x = Signature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vega_commands_v1_signature_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature) ProtoMessage() {}

func (x *Signature) ProtoReflect() protoreflect.Message {
	mi := &file_vega_commands_v1_signature_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signature.ProtoReflect.Descriptor instead.
func (*Signature) Descriptor() ([]byte, []int) {
	return file_vega_commands_v1_signature_proto_rawDescGZIP(), []int{0}
}

func (x *Signature) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Signature) GetAlgo() string {
	if x != nil {
		return x.Algo
	}
	return ""
}

func (x *Signature) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

var File_vega_commands_v1_signature_proto protoreflect.FileDescriptor

var file_vega_commands_v1_signature_proto_rawDesc = []byte{
	0x0a, 0x20, 0x76, 0x65, 0x67, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x10, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x22, 0x4f, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x6c, 0x67, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x6c, 0x67, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x2e, 0x5a, 0x2c, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x65,
	0x67, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x65, 0x67, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vega_commands_v1_signature_proto_rawDescOnce sync.Once
	file_vega_commands_v1_signature_proto_rawDescData = file_vega_commands_v1_signature_proto_rawDesc
)

func file_vega_commands_v1_signature_proto_rawDescGZIP() []byte {
	file_vega_commands_v1_signature_proto_rawDescOnce.Do(func() {
		file_vega_commands_v1_signature_proto_rawDescData = protoimpl.X.CompressGZIP(file_vega_commands_v1_signature_proto_rawDescData)
	})
	return file_vega_commands_v1_signature_proto_rawDescData
}

var file_vega_commands_v1_signature_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_vega_commands_v1_signature_proto_goTypes = []interface{}{
	(*Signature)(nil), // 0: vega.commands.v1.Signature
}
var file_vega_commands_v1_signature_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_vega_commands_v1_signature_proto_init() }
func file_vega_commands_v1_signature_proto_init() {
	if File_vega_commands_v1_signature_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vega_commands_v1_signature_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signature); i {
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
			RawDescriptor: file_vega_commands_v1_signature_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vega_commands_v1_signature_proto_goTypes,
		DependencyIndexes: file_vega_commands_v1_signature_proto_depIdxs,
		MessageInfos:      file_vega_commands_v1_signature_proto_msgTypes,
	}.Build()
	File_vega_commands_v1_signature_proto = out.File
	file_vega_commands_v1_signature_proto_rawDesc = nil
	file_vega_commands_v1_signature_proto_goTypes = nil
	file_vega_commands_v1_signature_proto_depIdxs = nil
}
