// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: testapi.proto

package testapi

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

type FetchJSON struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SampleCount uint32             `protobuf:"varint,1,opt,name=sampleCount,json=sample-count,proto3" json:"sampleCount,omitempty"`
	Disclaimer  string             `protobuf:"bytes,2,opt,name=disclaimer,proto3" json:"disclaimer,omitempty"`
	Results     map[string]*Result `protobuf:"bytes,3,rep,name=results,proto3" json:"results,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *FetchJSON) Reset() {
	*x = FetchJSON{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchJSON) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchJSON) ProtoMessage() {}

func (x *FetchJSON) ProtoReflect() protoreflect.Message {
	mi := &file_testapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchJSON.ProtoReflect.Descriptor instead.
func (*FetchJSON) Descriptor() ([]byte, []int) {
	return file_testapi_proto_rawDescGZIP(), []int{0}
}

func (x *FetchJSON) GetSampleCount() uint32 {
	if x != nil {
		return x.SampleCount
	}
	return 0
}

func (x *FetchJSON) GetDisclaimer() string {
	if x != nil {
		return x.Disclaimer
	}
	return ""
}

func (x *FetchJSON) GetResults() map[string]*Result {
	if x != nil {
		return x.Results
	}
	return nil
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *Metadata `protobuf:"bytes,1,opt,name=metadata,json=meta-data,proto3" json:"metadata,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_testapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_testapi_proto_rawDescGZIP(), []int{1}
}

func (x *Result) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DmpSampleId string `protobuf:"bytes,1,opt,name=dmp_sample_id,json=dmpSampleId,proto3" json:"dmp_sample_id,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testapi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_testapi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_testapi_proto_rawDescGZIP(), []int{2}
}

func (x *Metadata) GetDmpSampleId() string {
	if x != nil {
		return x.DmpSampleId
	}
	return ""
}

var File_testapi_proto protoreflect.FileDescriptor

var file_testapi_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x74, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc6, 0x01, 0x0a, 0x09, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x21, 0x0a,
	0x0b, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0c, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x73, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x69, 0x73, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x65, 0x72,
	0x12, 0x31, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4a, 0x53, 0x4f, 0x4e, 0x2e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x1a, 0x43, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x30, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x26, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52,
	0x09, 0x6d, 0x65, 0x74, 0x61, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2e, 0x0a, 0x08, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x0d, 0x64, 0x6d, 0x70, 0x5f, 0x73, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x6d, 0x70, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x64, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x76, 0x65, 0x72, 0x79, 0x6e, 0x69,
	0x63, 0x65, 0x64, 0x61, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x6d, 0x70, 0x61, 0x74, 0x68, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_testapi_proto_rawDescOnce sync.Once
	file_testapi_proto_rawDescData = file_testapi_proto_rawDesc
)

func file_testapi_proto_rawDescGZIP() []byte {
	file_testapi_proto_rawDescOnce.Do(func() {
		file_testapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_testapi_proto_rawDescData)
	})
	return file_testapi_proto_rawDescData
}

var file_testapi_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_testapi_proto_goTypes = []interface{}{
	(*FetchJSON)(nil), // 0: FetchJSON
	(*Result)(nil),    // 1: Result
	(*Metadata)(nil),  // 2: Metadata
	nil,               // 3: FetchJSON.ResultsEntry
}
var file_testapi_proto_depIdxs = []int32{
	3, // 0: FetchJSON.results:type_name -> FetchJSON.ResultsEntry
	2, // 1: Result.metadata:type_name -> Metadata
	1, // 2: FetchJSON.ResultsEntry.value:type_name -> Result
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_testapi_proto_init() }
func file_testapi_proto_init() {
	if File_testapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_testapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchJSON); i {
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
		file_testapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_testapi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
			RawDescriptor: file_testapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_testapi_proto_goTypes,
		DependencyIndexes: file_testapi_proto_depIdxs,
		MessageInfos:      file_testapi_proto_msgTypes,
	}.Build()
	File_testapi_proto = out.File
	file_testapi_proto_rawDesc = nil
	file_testapi_proto_goTypes = nil
	file_testapi_proto_depIdxs = nil
}
