// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: kessel/inventory/v1beta2/check.proto

package v1beta2

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

type CheckResponse_Allowed int32

const (
	CheckResponse_ALLOWED_UNSPECIFIED CheckResponse_Allowed = 0
	CheckResponse_ALLOWED_TRUE        CheckResponse_Allowed = 1
	CheckResponse_ALLOWED_FALSE       CheckResponse_Allowed = 2 // e.g.  ALLOWED_CONDITIONAL = 3;
)

// Enum value maps for CheckResponse_Allowed.
var (
	CheckResponse_Allowed_name = map[int32]string{
		0: "ALLOWED_UNSPECIFIED",
		1: "ALLOWED_TRUE",
		2: "ALLOWED_FALSE",
	}
	CheckResponse_Allowed_value = map[string]int32{
		"ALLOWED_UNSPECIFIED": 0,
		"ALLOWED_TRUE":        1,
		"ALLOWED_FALSE":       2,
	}
)

func (x CheckResponse_Allowed) Enum() *CheckResponse_Allowed {
	p := new(CheckResponse_Allowed)
	*p = x
	return p
}

func (x CheckResponse_Allowed) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CheckResponse_Allowed) Descriptor() protoreflect.EnumDescriptor {
	return file_kessel_inventory_v1beta2_check_proto_enumTypes[0].Descriptor()
}

func (CheckResponse_Allowed) Type() protoreflect.EnumType {
	return &file_kessel_inventory_v1beta2_check_proto_enumTypes[0]
}

func (x CheckResponse_Allowed) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CheckResponse_Allowed.Descriptor instead.
func (CheckResponse_Allowed) EnumDescriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta2_check_proto_rawDescGZIP(), []int{1, 0}
}

type CheckForUpdateResponse_Allowed int32

const (
	CheckForUpdateResponse_ALLOWED_UNSPECIFIED CheckForUpdateResponse_Allowed = 0
	CheckForUpdateResponse_ALLOWED_TRUE        CheckForUpdateResponse_Allowed = 1
	CheckForUpdateResponse_ALLOWED_FALSE       CheckForUpdateResponse_Allowed = 2 // e.g.  ALLOWED_CONDITIONAL = 3;
)

// Enum value maps for CheckForUpdateResponse_Allowed.
var (
	CheckForUpdateResponse_Allowed_name = map[int32]string{
		0: "ALLOWED_UNSPECIFIED",
		1: "ALLOWED_TRUE",
		2: "ALLOWED_FALSE",
	}
	CheckForUpdateResponse_Allowed_value = map[string]int32{
		"ALLOWED_UNSPECIFIED": 0,
		"ALLOWED_TRUE":        1,
		"ALLOWED_FALSE":       2,
	}
)

func (x CheckForUpdateResponse_Allowed) Enum() *CheckForUpdateResponse_Allowed {
	p := new(CheckForUpdateResponse_Allowed)
	*p = x
	return p
}

func (x CheckForUpdateResponse_Allowed) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CheckForUpdateResponse_Allowed) Descriptor() protoreflect.EnumDescriptor {
	return file_kessel_inventory_v1beta2_check_proto_enumTypes[1].Descriptor()
}

func (CheckForUpdateResponse_Allowed) Type() protoreflect.EnumType {
	return &file_kessel_inventory_v1beta2_check_proto_enumTypes[1]
}

func (x CheckForUpdateResponse_Allowed) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CheckForUpdateResponse_Allowed.Descriptor instead.
func (CheckForUpdateResponse_Allowed) EnumDescriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta2_check_proto_rawDescGZIP(), []int{3, 0}
}

type CheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required parameters (from an authz perspective)
	// - resource type and id
	// - permission (cannot be derived from the type as a type may have multiple 'read' permissions. Ex: https://github.com/RedHatInsights/rbac-config/blob/master/configs/prod/schemas/src/notifications.ksl#L31)
	// - user (possibly derived from an identity token)
	Parent   *ObjectReference  `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	Relation string            `protobuf:"bytes,2,opt,name=relation,proto3" json:"relation,omitempty"`
	Subject  *SubjectReference `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"` // Consistency consistency = 4;
}

func (x *CheckRequest) Reset() {
	*x = CheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1beta2_check_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckRequest) ProtoMessage() {}

func (x *CheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1beta2_check_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckRequest.ProtoReflect.Descriptor instead.
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta2_check_proto_rawDescGZIP(), []int{0}
}

func (x *CheckRequest) GetParent() *ObjectReference {
	if x != nil {
		return x.Parent
	}
	return nil
}

func (x *CheckRequest) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

func (x *CheckRequest) GetSubject() *SubjectReference {
	if x != nil {
		return x.Subject
	}
	return nil
}

type CheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Allowed CheckResponse_Allowed `protobuf:"varint,1,opt,name=allowed,proto3,enum=kessel.inventory.v1beta2.CheckResponse_Allowed" json:"allowed,omitempty"`
}

func (x *CheckResponse) Reset() {
	*x = CheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1beta2_check_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckResponse) ProtoMessage() {}

func (x *CheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1beta2_check_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckResponse.ProtoReflect.Descriptor instead.
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta2_check_proto_rawDescGZIP(), []int{1}
}

func (x *CheckResponse) GetAllowed() CheckResponse_Allowed {
	if x != nil {
		return x.Allowed
	}
	return CheckResponse_ALLOWED_UNSPECIFIED
}

type CheckForUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required parameters (from an authz perspective)
	// - resource type and id
	// - permission (cannot be derived from type as types may have multiple edit permissions Ex: https://github.com/RedHatInsights/rbac-config/blob/master/configs/prod/schemas/src/notifications.ksl#L37)
	// - user (possibly derived from an identity token)
	Parent   *ObjectReference  `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	Relation string            `protobuf:"bytes,2,opt,name=relation,proto3" json:"relation,omitempty"`
	Subject  *SubjectReference `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
}

func (x *CheckForUpdateRequest) Reset() {
	*x = CheckForUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1beta2_check_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckForUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckForUpdateRequest) ProtoMessage() {}

func (x *CheckForUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1beta2_check_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckForUpdateRequest.ProtoReflect.Descriptor instead.
func (*CheckForUpdateRequest) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta2_check_proto_rawDescGZIP(), []int{2}
}

func (x *CheckForUpdateRequest) GetParent() *ObjectReference {
	if x != nil {
		return x.Parent
	}
	return nil
}

func (x *CheckForUpdateRequest) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

func (x *CheckForUpdateRequest) GetSubject() *SubjectReference {
	if x != nil {
		return x.Subject
	}
	return nil
}

type CheckForUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Allowed CheckForUpdateResponse_Allowed `protobuf:"varint,1,opt,name=allowed,proto3,enum=kessel.inventory.v1beta2.CheckForUpdateResponse_Allowed" json:"allowed,omitempty"`
}

func (x *CheckForUpdateResponse) Reset() {
	*x = CheckForUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1beta2_check_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckForUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckForUpdateResponse) ProtoMessage() {}

func (x *CheckForUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1beta2_check_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckForUpdateResponse.ProtoReflect.Descriptor instead.
func (*CheckForUpdateResponse) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta2_check_proto_rawDescGZIP(), []int{3}
}

func (x *CheckForUpdateResponse) GetAllowed() CheckForUpdateResponse_Allowed {
	if x != nil {
		return x.Allowed
	}
	return CheckForUpdateResponse_ALLOWED_UNSPECIFIED
}

var File_kessel_inventory_v1beta2_check_proto protoreflect.FileDescriptor

var file_kessel_inventory_v1beta2_check_proto_rawDesc = []byte{
	0x0a, 0x24, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25,
	0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xcc, 0x01, 0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x49, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2e, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x06,
	0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x23,
	0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x4c, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2e,
	0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x22, 0xa3, 0x01, 0x0a, 0x0d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x07, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x41, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x64, 0x52, 0x07, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x22, 0x47,
	0x0a, 0x07, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x12, 0x17, 0x0a, 0x13, 0x41, 0x4c, 0x4c,
	0x4f, 0x57, 0x45, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x5f, 0x54, 0x52,
	0x55, 0x45, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x5f,
	0x46, 0x41, 0x4c, 0x53, 0x45, 0x10, 0x02, 0x22, 0xd5, 0x01, 0x0a, 0x15, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x46, 0x6f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x49, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x29, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2e, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x06, 0xba, 0x48,
	0x03, 0xc8, 0x01, 0x01, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x08,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x4c, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2e, 0x53, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x06,
	0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22,
	0xb5, 0x01, 0x0a, 0x16, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x46, 0x6f, 0x72, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x07, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x38, 0x2e, 0x6b, 0x65,
	0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x46, 0x6f, 0x72, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x41, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x64, 0x52, 0x07, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x22, 0x47,
	0x0a, 0x07, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x12, 0x17, 0x0a, 0x13, 0x41, 0x4c, 0x4c,
	0x4f, 0x57, 0x45, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x5f, 0x54, 0x52,
	0x55, 0x45, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x5f,
	0x46, 0x41, 0x4c, 0x53, 0x45, 0x10, 0x02, 0x32, 0xc0, 0x02, 0x0a, 0x12, 0x4b, 0x65, 0x73, 0x73,
	0x65, 0x6c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x81,
	0x01, 0x0a, 0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x26, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65,
	0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x32, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x27, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x21, 0x3a, 0x01, 0x2a, 0x22, 0x1c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2f, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x12, 0xa5, 0x01, 0x0a, 0x0e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x46, 0x6f, 0x72, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2f, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x46, 0x6f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x32, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x46, 0x6f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a,
	0x3a, 0x01, 0x2a, 0x22, 0x25, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2f, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x66, 0x6f, 0x72, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x72, 0x0a, 0x28, 0x6f, 0x72,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x6b, 0x65, 0x73,
	0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2d, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kessel_inventory_v1beta2_check_proto_rawDescOnce sync.Once
	file_kessel_inventory_v1beta2_check_proto_rawDescData = file_kessel_inventory_v1beta2_check_proto_rawDesc
)

func file_kessel_inventory_v1beta2_check_proto_rawDescGZIP() []byte {
	file_kessel_inventory_v1beta2_check_proto_rawDescOnce.Do(func() {
		file_kessel_inventory_v1beta2_check_proto_rawDescData = protoimpl.X.CompressGZIP(file_kessel_inventory_v1beta2_check_proto_rawDescData)
	})
	return file_kessel_inventory_v1beta2_check_proto_rawDescData
}

var file_kessel_inventory_v1beta2_check_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_kessel_inventory_v1beta2_check_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_kessel_inventory_v1beta2_check_proto_goTypes = []any{
	(CheckResponse_Allowed)(0),          // 0: kessel.inventory.v1beta2.CheckResponse.Allowed
	(CheckForUpdateResponse_Allowed)(0), // 1: kessel.inventory.v1beta2.CheckForUpdateResponse.Allowed
	(*CheckRequest)(nil),                // 2: kessel.inventory.v1beta2.CheckRequest
	(*CheckResponse)(nil),               // 3: kessel.inventory.v1beta2.CheckResponse
	(*CheckForUpdateRequest)(nil),       // 4: kessel.inventory.v1beta2.CheckForUpdateRequest
	(*CheckForUpdateResponse)(nil),      // 5: kessel.inventory.v1beta2.CheckForUpdateResponse
	(*ObjectReference)(nil),             // 6: kessel.inventory.v1beta2.ObjectReference
	(*SubjectReference)(nil),            // 7: kessel.inventory.v1beta2.SubjectReference
}
var file_kessel_inventory_v1beta2_check_proto_depIdxs = []int32{
	6, // 0: kessel.inventory.v1beta2.CheckRequest.parent:type_name -> kessel.inventory.v1beta2.ObjectReference
	7, // 1: kessel.inventory.v1beta2.CheckRequest.subject:type_name -> kessel.inventory.v1beta2.SubjectReference
	0, // 2: kessel.inventory.v1beta2.CheckResponse.allowed:type_name -> kessel.inventory.v1beta2.CheckResponse.Allowed
	6, // 3: kessel.inventory.v1beta2.CheckForUpdateRequest.parent:type_name -> kessel.inventory.v1beta2.ObjectReference
	7, // 4: kessel.inventory.v1beta2.CheckForUpdateRequest.subject:type_name -> kessel.inventory.v1beta2.SubjectReference
	1, // 5: kessel.inventory.v1beta2.CheckForUpdateResponse.allowed:type_name -> kessel.inventory.v1beta2.CheckForUpdateResponse.Allowed
	2, // 6: kessel.inventory.v1beta2.KesselCheckService.Check:input_type -> kessel.inventory.v1beta2.CheckRequest
	4, // 7: kessel.inventory.v1beta2.KesselCheckService.CheckForUpdate:input_type -> kessel.inventory.v1beta2.CheckForUpdateRequest
	3, // 8: kessel.inventory.v1beta2.KesselCheckService.Check:output_type -> kessel.inventory.v1beta2.CheckResponse
	5, // 9: kessel.inventory.v1beta2.KesselCheckService.CheckForUpdate:output_type -> kessel.inventory.v1beta2.CheckForUpdateResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_kessel_inventory_v1beta2_check_proto_init() }
func file_kessel_inventory_v1beta2_check_proto_init() {
	if File_kessel_inventory_v1beta2_check_proto != nil {
		return
	}
	file_kessel_inventory_v1beta2_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_kessel_inventory_v1beta2_check_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CheckRequest); i {
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
		file_kessel_inventory_v1beta2_check_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CheckResponse); i {
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
		file_kessel_inventory_v1beta2_check_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CheckForUpdateRequest); i {
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
		file_kessel_inventory_v1beta2_check_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CheckForUpdateResponse); i {
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
			RawDescriptor: file_kessel_inventory_v1beta2_check_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kessel_inventory_v1beta2_check_proto_goTypes,
		DependencyIndexes: file_kessel_inventory_v1beta2_check_proto_depIdxs,
		EnumInfos:         file_kessel_inventory_v1beta2_check_proto_enumTypes,
		MessageInfos:      file_kessel_inventory_v1beta2_check_proto_msgTypes,
	}.Build()
	File_kessel_inventory_v1beta2_check_proto = out.File
	file_kessel_inventory_v1beta2_check_proto_rawDesc = nil
	file_kessel_inventory_v1beta2_check_proto_goTypes = nil
	file_kessel_inventory_v1beta2_check_proto_depIdxs = nil
}
