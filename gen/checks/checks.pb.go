// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: checks.proto

package checks

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

type CheckType int32

const (
	CheckType_CheckTypeUnknown           CheckType = 0
	CheckType_CheckTypeVulnerability     CheckType = 1
	CheckType_CheckTypeMalware           CheckType = 2
	CheckType_CheckTypePopularity        CheckType = 3
	CheckType_CheckTypeMaintenance       CheckType = 4
	CheckType_CheckTypeSecurityScorecard CheckType = 5
	CheckType_CheckTypeLicense           CheckType = 6
	CheckType_CheckTypeOther             CheckType = 100
)

// Enum value maps for CheckType.
var (
	CheckType_name = map[int32]string{
		0:   "CheckTypeUnknown",
		1:   "CheckTypeVulnerability",
		2:   "CheckTypeMalware",
		3:   "CheckTypePopularity",
		4:   "CheckTypeMaintenance",
		5:   "CheckTypeSecurityScorecard",
		6:   "CheckTypeLicense",
		100: "CheckTypeOther",
	}
	CheckType_value = map[string]int32{
		"CheckTypeUnknown":           0,
		"CheckTypeVulnerability":     1,
		"CheckTypeMalware":           2,
		"CheckTypePopularity":        3,
		"CheckTypeMaintenance":       4,
		"CheckTypeSecurityScorecard": 5,
		"CheckTypeLicense":           6,
		"CheckTypeOther":             100,
	}
)

func (x CheckType) Enum() *CheckType {
	p := new(CheckType)
	*p = x
	return p
}

func (x CheckType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CheckType) Descriptor() protoreflect.EnumDescriptor {
	return file_checks_proto_enumTypes[0].Descriptor()
}

func (CheckType) Type() protoreflect.EnumType {
	return &file_checks_proto_enumTypes[0]
}

func (x CheckType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CheckType.Descriptor instead.
func (CheckType) EnumDescriptor() ([]byte, []int) {
	return file_checks_proto_rawDescGZIP(), []int{0}
}

var File_checks_proto protoreflect.FileDescriptor

var file_checks_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xd6,
	0x01, 0x0a, 0x09, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x56,
	0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x10, 0x01, 0x12, 0x14,
	0x0a, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x4d, 0x61, 0x6c, 0x77, 0x61,
	0x72, 0x65, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x79, 0x70,
	0x65, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x10, 0x03, 0x12, 0x18, 0x0a,
	0x14, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x63, 0x65, 0x10, 0x04, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x54, 0x79, 0x70, 0x65, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x63, 0x61, 0x72, 0x64, 0x10, 0x05, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x10, 0x06, 0x12, 0x12, 0x0a,
	0x0e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x10,
	0x64, 0x22, 0x04, 0x08, 0x07, 0x10, 0x63, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x66, 0x65, 0x64, 0x65, 0x70, 0x2f, 0x76, 0x65,
	0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_checks_proto_rawDescOnce sync.Once
	file_checks_proto_rawDescData = file_checks_proto_rawDesc
)

func file_checks_proto_rawDescGZIP() []byte {
	file_checks_proto_rawDescOnce.Do(func() {
		file_checks_proto_rawDescData = protoimpl.X.CompressGZIP(file_checks_proto_rawDescData)
	})
	return file_checks_proto_rawDescData
}

var file_checks_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_checks_proto_goTypes = []any{
	(CheckType)(0), // 0: CheckType
}
var file_checks_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_checks_proto_init() }
func file_checks_proto_init() {
	if File_checks_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_checks_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_checks_proto_goTypes,
		DependencyIndexes: file_checks_proto_depIdxs,
		EnumInfos:         file_checks_proto_enumTypes,
	}.Build()
	File_checks_proto = out.File
	file_checks_proto_rawDesc = nil
	file_checks_proto_goTypes = nil
	file_checks_proto_depIdxs = nil
}
