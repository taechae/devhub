// Copyright 2019 The Grafeas Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: grafeas/v1/build.proto

package grafeas

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Note holding the version of the provider's builder and the signature of the
// provenance message in the build details occurrence.
type BuildNote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Immutable. Version of the builder which produced this build.
	BuilderVersion string `protobuf:"bytes,1,opt,name=builder_version,json=builderVersion,proto3" json:"builder_version,omitempty"`
}

func (x *BuildNote) Reset() {
	*x = BuildNote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grafeas_v1_build_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildNote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildNote) ProtoMessage() {}

func (x *BuildNote) ProtoReflect() protoreflect.Message {
	mi := &file_grafeas_v1_build_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildNote.ProtoReflect.Descriptor instead.
func (*BuildNote) Descriptor() ([]byte, []int) {
	return file_grafeas_v1_build_proto_rawDescGZIP(), []int{0}
}

func (x *BuildNote) GetBuilderVersion() string {
	if x != nil {
		return x.BuilderVersion
	}
	return ""
}

// Details of a build occurrence.
type BuildOccurrence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The actual provenance for the build.
	Provenance *BuildProvenance `protobuf:"bytes,1,opt,name=provenance,proto3" json:"provenance,omitempty"`
	// Serialized JSON representation of the provenance, used in generating the
	// build signature in the corresponding build note. After verifying the
	// signature, `provenance_bytes` can be unmarshalled and compared to the
	// provenance to confirm that it is unchanged. A base64-encoded string
	// representation of the provenance bytes is used for the signature in order
	// to interoperate with openssl which expects this format for signature
	// verification.
	//
	// The serialized form is captured both to avoid ambiguity in how the
	// provenance is marshalled to json as well to prevent incompatibilities with
	// future changes.
	ProvenanceBytes string `protobuf:"bytes,2,opt,name=provenance_bytes,json=provenanceBytes,proto3" json:"provenance_bytes,omitempty"`
	// Deprecated. See InTotoStatement for the replacement.
	// In-toto Provenance representation as defined in spec.
	IntotoProvenance *InTotoProvenance `protobuf:"bytes,3,opt,name=intoto_provenance,json=intotoProvenance,proto3" json:"intoto_provenance,omitempty"`
	// In-toto Statement representation as defined in spec.
	// The intoto_statement can contain any type of provenance. The serialized
	// payload of the statement can be stored and signed in the Occurrence's
	// envelope.
	IntotoStatement *InTotoStatement `protobuf:"bytes,4,opt,name=intoto_statement,json=intotoStatement,proto3" json:"intoto_statement,omitempty"`
}

func (x *BuildOccurrence) Reset() {
	*x = BuildOccurrence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grafeas_v1_build_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildOccurrence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildOccurrence) ProtoMessage() {}

func (x *BuildOccurrence) ProtoReflect() protoreflect.Message {
	mi := &file_grafeas_v1_build_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildOccurrence.ProtoReflect.Descriptor instead.
func (*BuildOccurrence) Descriptor() ([]byte, []int) {
	return file_grafeas_v1_build_proto_rawDescGZIP(), []int{1}
}

func (x *BuildOccurrence) GetProvenance() *BuildProvenance {
	if x != nil {
		return x.Provenance
	}
	return nil
}

func (x *BuildOccurrence) GetProvenanceBytes() string {
	if x != nil {
		return x.ProvenanceBytes
	}
	return ""
}

func (x *BuildOccurrence) GetIntotoProvenance() *InTotoProvenance {
	if x != nil {
		return x.IntotoProvenance
	}
	return nil
}

func (x *BuildOccurrence) GetIntotoStatement() *InTotoStatement {
	if x != nil {
		return x.IntotoStatement
	}
	return nil
}

var File_grafeas_v1_build_proto protoreflect.FileDescriptor

var file_grafeas_v1_build_proto_rawDesc = []byte{
	0x0a, 0x16, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61,
	0x73, 0x2e, 0x76, 0x31, 0x1a, 0x22, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x69, 0x6e, 0x74, 0x6f, 0x74, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x74, 0x6f, 0x74, 0x6f, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x72, 0x61,
	0x66, 0x65, 0x61, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x09, 0x42, 0x75, 0x69, 0x6c,
	0x64, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x8c,
	0x02, 0x0a, 0x0f, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x4f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x29, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x76, 0x65,
	0x6e, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x49, 0x0a, 0x11, 0x69, 0x6e,
	0x74, 0x6f, 0x74, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x6e, 0x54, 0x6f, 0x74, 0x6f, 0x50, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x10, 0x69, 0x6e, 0x74, 0x6f, 0x74, 0x6f, 0x50, 0x72, 0x6f, 0x76, 0x65,
	0x6e, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x10, 0x69, 0x6e, 0x74, 0x6f, 0x74, 0x6f, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x54,
	0x6f, 0x74, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0f, 0x69, 0x6e,
	0x74, 0x6f, 0x74, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x51, 0x0a,
	0x0d, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x50, 0x01,
	0x5a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2f,
	0x76, 0x31, 0x3b, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x52, 0x41,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grafeas_v1_build_proto_rawDescOnce sync.Once
	file_grafeas_v1_build_proto_rawDescData = file_grafeas_v1_build_proto_rawDesc
)

func file_grafeas_v1_build_proto_rawDescGZIP() []byte {
	file_grafeas_v1_build_proto_rawDescOnce.Do(func() {
		file_grafeas_v1_build_proto_rawDescData = protoimpl.X.CompressGZIP(file_grafeas_v1_build_proto_rawDescData)
	})
	return file_grafeas_v1_build_proto_rawDescData
}

var file_grafeas_v1_build_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grafeas_v1_build_proto_goTypes = []interface{}{
	(*BuildNote)(nil),        // 0: grafeas.v1.BuildNote
	(*BuildOccurrence)(nil),  // 1: grafeas.v1.BuildOccurrence
	(*BuildProvenance)(nil),  // 2: grafeas.v1.BuildProvenance
	(*InTotoProvenance)(nil), // 3: grafeas.v1.InTotoProvenance
	(*InTotoStatement)(nil),  // 4: grafeas.v1.InTotoStatement
}
var file_grafeas_v1_build_proto_depIdxs = []int32{
	2, // 0: grafeas.v1.BuildOccurrence.provenance:type_name -> grafeas.v1.BuildProvenance
	3, // 1: grafeas.v1.BuildOccurrence.intoto_provenance:type_name -> grafeas.v1.InTotoProvenance
	4, // 2: grafeas.v1.BuildOccurrence.intoto_statement:type_name -> grafeas.v1.InTotoStatement
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_grafeas_v1_build_proto_init() }
func file_grafeas_v1_build_proto_init() {
	if File_grafeas_v1_build_proto != nil {
		return
	}
	file_grafeas_v1_intoto_provenance_proto_init()
	file_grafeas_v1_intoto_statement_proto_init()
	file_grafeas_v1_provenance_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_grafeas_v1_build_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildNote); i {
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
		file_grafeas_v1_build_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildOccurrence); i {
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
			RawDescriptor: file_grafeas_v1_build_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_grafeas_v1_build_proto_goTypes,
		DependencyIndexes: file_grafeas_v1_build_proto_depIdxs,
		MessageInfos:      file_grafeas_v1_build_proto_msgTypes,
	}.Build()
	File_grafeas_v1_build_proto = out.File
	file_grafeas_v1_build_proto_rawDesc = nil
	file_grafeas_v1_build_proto_goTypes = nil
	file_grafeas_v1_build_proto_depIdxs = nil
}
