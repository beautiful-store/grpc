// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: test/member/member.proto

package member

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type IDs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids string `protobuf:"bytes,1,opt,name=ids,proto3" json:"ids,omitempty"`
}

func (x *IDs) Reset() {
	*x = IDs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_member_member_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDs) ProtoMessage() {}

func (x *IDs) ProtoReflect() protoreflect.Message {
	mi := &file_test_member_member_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDs.ProtoReflect.Descriptor instead.
func (*IDs) Descriptor() ([]byte, []int) {
	return file_test_member_member_proto_rawDescGZIP(), []int{0}
}

func (x *IDs) GetIds() string {
	if x != nil {
		return x.Ids
	}
	return ""
}

type MemberList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberList []*Member `protobuf:"bytes,1,rep,name=member_list,json=memberList,proto3" json:"member_list,omitempty"`
}

func (x *MemberList) Reset() {
	*x = MemberList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_member_member_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberList) ProtoMessage() {}

func (x *MemberList) ProtoReflect() protoreflect.Message {
	mi := &file_test_member_member_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberList.ProtoReflect.Descriptor instead.
func (*MemberList) Descriptor() ([]byte, []int) {
	return file_test_member_member_proto_rawDescGZIP(), []int{1}
}

func (x *MemberList) GetMemberList() []*Member {
	if x != nil {
		return x.MemberList
	}
	return nil
}

type Member struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberID     int64  `protobuf:"varint,1,opt,name=memberID,proto3" json:"memberID,omitempty"`
	MemberName   string `protobuf:"bytes,2,opt,name=memberName,proto3" json:"memberName,omitempty"`
	Mobile       string `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	MaskedMobile string `protobuf:"bytes,4,opt,name=maskedMobile,proto3" json:"maskedMobile,omitempty"`
	Email        string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	OrgID        int64  `protobuf:"varint,6,opt,name=orgID,proto3" json:"orgID,omitempty"`
}

func (x *Member) Reset() {
	*x = Member{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_member_member_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Member) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Member) ProtoMessage() {}

func (x *Member) ProtoReflect() protoreflect.Message {
	mi := &file_test_member_member_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Member.ProtoReflect.Descriptor instead.
func (*Member) Descriptor() ([]byte, []int) {
	return file_test_member_member_proto_rawDescGZIP(), []int{2}
}

func (x *Member) GetMemberID() int64 {
	if x != nil {
		return x.MemberID
	}
	return 0
}

func (x *Member) GetMemberName() string {
	if x != nil {
		return x.MemberName
	}
	return ""
}

func (x *Member) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *Member) GetMaskedMobile() string {
	if x != nil {
		return x.MaskedMobile
	}
	return ""
}

func (x *Member) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Member) GetOrgID() int64 {
	if x != nil {
		return x.OrgID
	}
	return 0
}

var File_test_member_member_proto protoreflect.FileDescriptor

var file_test_member_member_proto_rawDesc = []byte{
	0x0a, 0x18, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2f, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x17, 0x0a, 0x03, 0x49, 0x44, 0x73, 0x12, 0x10, 0x0a, 0x03,
	0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x42,
	0x0a, 0x0a, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x0b,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x0a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x22, 0xac, 0x01, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x73, 0x6b, 0x65, 0x64, 0x4d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61, 0x73, 0x6b, 0x65, 0x64, 0x4d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x6f,
	0x72, 0x67, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49,
	0x44, 0x32, 0x64, 0x0a, 0x0c, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x53, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x54, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x10, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x49, 0x44, 0x73, 0x1a, 0x17,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12,
	0x16, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2f, 0x6f,
	0x72, 0x67, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x61, 0x75, 0x74, 0x69, 0x66, 0x75, 0x6c, 0x2d,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_member_member_proto_rawDescOnce sync.Once
	file_test_member_member_proto_rawDescData = file_test_member_member_proto_rawDesc
)

func file_test_member_member_proto_rawDescGZIP() []byte {
	file_test_member_member_proto_rawDescOnce.Do(func() {
		file_test_member_member_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_member_member_proto_rawDescData)
	})
	return file_test_member_member_proto_rawDescData
}

var file_test_member_member_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_test_member_member_proto_goTypes = []interface{}{
	(*IDs)(nil),        // 0: test.member.IDs
	(*MemberList)(nil), // 1: test.member.MemberList
	(*Member)(nil),     // 2: test.member.Member
}
var file_test_member_member_proto_depIdxs = []int32{
	2, // 0: test.member.MemberList.member_list:type_name -> test.member.Member
	0, // 1: test.member.MemberSevice.GetList:input_type -> test.member.IDs
	1, // 2: test.member.MemberSevice.GetList:output_type -> test.member.MemberList
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_test_member_member_proto_init() }
func file_test_member_member_proto_init() {
	if File_test_member_member_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_member_member_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDs); i {
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
		file_test_member_member_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberList); i {
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
		file_test_member_member_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Member); i {
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
			RawDescriptor: file_test_member_member_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_test_member_member_proto_goTypes,
		DependencyIndexes: file_test_member_member_proto_depIdxs,
		MessageInfos:      file_test_member_member_proto_msgTypes,
	}.Build()
	File_test_member_member_proto = out.File
	file_test_member_member_proto_rawDesc = nil
	file_test_member_member_proto_goTypes = nil
	file_test_member_member_proto_depIdxs = nil
}
