// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: headscale/v1/acls.proto

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

type ACLPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Groups    map[string]*Group     `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Hosts     map[string]string     `protobuf:"bytes,2,rep,name=hosts,proto3" json:"hosts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	TagOwners map[string]*TagOwners `protobuf:"bytes,3,rep,name=tag_owners,json=tagOwners,proto3" json:"tag_owners,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Acls      []*ACL                `protobuf:"bytes,4,rep,name=acls,proto3" json:"acls,omitempty"`
	AclTest   []*ACLTest            `protobuf:"bytes,5,rep,name=acl_test,json=aclTest,proto3" json:"acl_test,omitempty"`
}

func (x *ACLPolicy) Reset() {
	*x = ACLPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_headscale_v1_acls_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ACLPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ACLPolicy) ProtoMessage() {}

func (x *ACLPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_headscale_v1_acls_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ACLPolicy.ProtoReflect.Descriptor instead.
func (*ACLPolicy) Descriptor() ([]byte, []int) {
	return file_headscale_v1_acls_proto_rawDescGZIP(), []int{0}
}

func (x *ACLPolicy) GetGroups() map[string]*Group {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *ACLPolicy) GetHosts() map[string]string {
	if x != nil {
		return x.Hosts
	}
	return nil
}

func (x *ACLPolicy) GetTagOwners() map[string]*TagOwners {
	if x != nil {
		return x.TagOwners
	}
	return nil
}

func (x *ACLPolicy) GetAcls() []*ACL {
	if x != nil {
		return x.Acls
	}
	return nil
}

func (x *ACLPolicy) GetAclTest() []*ACLTest {
	if x != nil {
		return x.AclTest
	}
	return nil
}

type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group []string `protobuf:"bytes,1,rep,name=group,proto3" json:"group,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_headscale_v1_acls_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_headscale_v1_acls_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_headscale_v1_acls_proto_rawDescGZIP(), []int{1}
}

func (x *Group) GetGroup() []string {
	if x != nil {
		return x.Group
	}
	return nil
}

type TagOwners struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TagOwners []string `protobuf:"bytes,1,rep,name=tag_owners,json=tagOwners,proto3" json:"tag_owners,omitempty"`
}

func (x *TagOwners) Reset() {
	*x = TagOwners{}
	if protoimpl.UnsafeEnabled {
		mi := &file_headscale_v1_acls_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagOwners) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagOwners) ProtoMessage() {}

func (x *TagOwners) ProtoReflect() protoreflect.Message {
	mi := &file_headscale_v1_acls_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagOwners.ProtoReflect.Descriptor instead.
func (*TagOwners) Descriptor() ([]byte, []int) {
	return file_headscale_v1_acls_proto_rawDescGZIP(), []int{2}
}

func (x *TagOwners) GetTagOwners() []string {
	if x != nil {
		return x.TagOwners
	}
	return nil
}

type ACL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action       string   `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	Protocol     string   `protobuf:"bytes,2,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Sources      []string `protobuf:"bytes,3,rep,name=sources,proto3" json:"sources,omitempty"`
	Destinations []string `protobuf:"bytes,4,rep,name=destinations,proto3" json:"destinations,omitempty"`
}

func (x *ACL) Reset() {
	*x = ACL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_headscale_v1_acls_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ACL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ACL) ProtoMessage() {}

func (x *ACL) ProtoReflect() protoreflect.Message {
	mi := &file_headscale_v1_acls_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ACL.ProtoReflect.Descriptor instead.
func (*ACL) Descriptor() ([]byte, []int) {
	return file_headscale_v1_acls_proto_rawDescGZIP(), []int{3}
}

func (x *ACL) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *ACL) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *ACL) GetSources() []string {
	if x != nil {
		return x.Sources
	}
	return nil
}

func (x *ACL) GetDestinations() []string {
	if x != nil {
		return x.Destinations
	}
	return nil
}

type ACLTest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source string   `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Accept []string `protobuf:"bytes,2,rep,name=accept,proto3" json:"accept,omitempty"`
	Deny   []string `protobuf:"bytes,3,rep,name=deny,proto3" json:"deny,omitempty"`
}

func (x *ACLTest) Reset() {
	*x = ACLTest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_headscale_v1_acls_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ACLTest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ACLTest) ProtoMessage() {}

func (x *ACLTest) ProtoReflect() protoreflect.Message {
	mi := &file_headscale_v1_acls_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ACLTest.ProtoReflect.Descriptor instead.
func (*ACLTest) Descriptor() ([]byte, []int) {
	return file_headscale_v1_acls_proto_rawDescGZIP(), []int{4}
}

func (x *ACLTest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *ACLTest) GetAccept() []string {
	if x != nil {
		return x.Accept
	}
	return nil
}

func (x *ACLTest) GetDeny() []string {
	if x != nil {
		return x.Deny
	}
	return nil
}

type ListACLPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListACLPolicyRequest) Reset() {
	*x = ListACLPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_headscale_v1_acls_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListACLPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListACLPolicyRequest) ProtoMessage() {}

func (x *ListACLPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_headscale_v1_acls_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListACLPolicyRequest.ProtoReflect.Descriptor instead.
func (*ListACLPolicyRequest) Descriptor() ([]byte, []int) {
	return file_headscale_v1_acls_proto_rawDescGZIP(), []int{5}
}

type ListACLPolicyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Policy *ACLPolicy `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
}

func (x *ListACLPolicyResponse) Reset() {
	*x = ListACLPolicyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_headscale_v1_acls_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListACLPolicyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListACLPolicyResponse) ProtoMessage() {}

func (x *ListACLPolicyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_headscale_v1_acls_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListACLPolicyResponse.ProtoReflect.Descriptor instead.
func (*ListACLPolicyResponse) Descriptor() ([]byte, []int) {
	return file_headscale_v1_acls_proto_rawDescGZIP(), []int{6}
}

func (x *ListACLPolicyResponse) GetPolicy() *ACLPolicy {
	if x != nil {
		return x.Policy
	}
	return nil
}

var File_headscale_v1_acls_proto protoreflect.FileDescriptor

var file_headscale_v1_acls_proto_rawDesc = []byte{
	0x0a, 0x17, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x63, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x68, 0x65, 0x61, 0x64, 0x73,
	0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x83, 0x04, 0x0a, 0x09, 0x41, 0x43, 0x4c, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x3b, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x43, 0x4c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x12, 0x38, 0x0a, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x43, 0x4c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x45, 0x0a, 0x0a,
	0x74, 0x61, 0x67, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x43, 0x4c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x54, 0x61, 0x67, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x74, 0x61, 0x67, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x73, 0x12, 0x25, 0x0a, 0x04, 0x61, 0x63, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x43, 0x4c, 0x52, 0x04, 0x61, 0x63, 0x6c, 0x73, 0x12, 0x30, 0x0a, 0x08, 0x61, 0x63,
	0x6c, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x68,
	0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x43, 0x4c, 0x54,
	0x65, 0x73, 0x74, 0x52, 0x07, 0x61, 0x63, 0x6c, 0x54, 0x65, 0x73, 0x74, 0x1a, 0x4e, 0x0a, 0x0b,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x29, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x68,
	0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x38, 0x0a, 0x0a,
	0x48, 0x6f, 0x73, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x55, 0x0a, 0x0e, 0x54, 0x61, 0x67, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x68, 0x65, 0x61, 0x64,
	0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x67, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x1d, 0x0a,
	0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x2a, 0x0a, 0x09,
	0x54, 0x61, 0x67, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x67,
	0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x61, 0x67, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x22, 0x77, 0x0a, 0x03, 0x41, 0x43, 0x4c, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x22, 0x0a,
	0x0c, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0x4d, 0x0a, 0x07, 0x41, 0x43, 0x4c, 0x54, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x65, 0x6e, 0x79, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x6e, 0x79,
	0x22, 0x16, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x43, 0x4c, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x48, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x43, 0x4c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2f, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x43, 0x4c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6a, 0x75, 0x61, 0x6e, 0x66, 0x6f, 0x6e, 0x74, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x73, 0x63,
	0x61, 0x6c, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_headscale_v1_acls_proto_rawDescOnce sync.Once
	file_headscale_v1_acls_proto_rawDescData = file_headscale_v1_acls_proto_rawDesc
)

func file_headscale_v1_acls_proto_rawDescGZIP() []byte {
	file_headscale_v1_acls_proto_rawDescOnce.Do(func() {
		file_headscale_v1_acls_proto_rawDescData = protoimpl.X.CompressGZIP(file_headscale_v1_acls_proto_rawDescData)
	})
	return file_headscale_v1_acls_proto_rawDescData
}

var file_headscale_v1_acls_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_headscale_v1_acls_proto_goTypes = []interface{}{
	(*ACLPolicy)(nil),             // 0: headscale.v1.ACLPolicy
	(*Group)(nil),                 // 1: headscale.v1.Group
	(*TagOwners)(nil),             // 2: headscale.v1.TagOwners
	(*ACL)(nil),                   // 3: headscale.v1.ACL
	(*ACLTest)(nil),               // 4: headscale.v1.ACLTest
	(*ListACLPolicyRequest)(nil),  // 5: headscale.v1.ListACLPolicyRequest
	(*ListACLPolicyResponse)(nil), // 6: headscale.v1.ListACLPolicyResponse
	nil,                           // 7: headscale.v1.ACLPolicy.GroupsEntry
	nil,                           // 8: headscale.v1.ACLPolicy.HostsEntry
	nil,                           // 9: headscale.v1.ACLPolicy.TagOwnersEntry
}
var file_headscale_v1_acls_proto_depIdxs = []int32{
	7, // 0: headscale.v1.ACLPolicy.groups:type_name -> headscale.v1.ACLPolicy.GroupsEntry
	8, // 1: headscale.v1.ACLPolicy.hosts:type_name -> headscale.v1.ACLPolicy.HostsEntry
	9, // 2: headscale.v1.ACLPolicy.tag_owners:type_name -> headscale.v1.ACLPolicy.TagOwnersEntry
	3, // 3: headscale.v1.ACLPolicy.acls:type_name -> headscale.v1.ACL
	4, // 4: headscale.v1.ACLPolicy.acl_test:type_name -> headscale.v1.ACLTest
	0, // 5: headscale.v1.ListACLPolicyResponse.policy:type_name -> headscale.v1.ACLPolicy
	1, // 6: headscale.v1.ACLPolicy.GroupsEntry.value:type_name -> headscale.v1.Group
	2, // 7: headscale.v1.ACLPolicy.TagOwnersEntry.value:type_name -> headscale.v1.TagOwners
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_headscale_v1_acls_proto_init() }
func file_headscale_v1_acls_proto_init() {
	if File_headscale_v1_acls_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_headscale_v1_acls_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ACLPolicy); i {
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
		file_headscale_v1_acls_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
		file_headscale_v1_acls_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagOwners); i {
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
		file_headscale_v1_acls_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ACL); i {
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
		file_headscale_v1_acls_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ACLTest); i {
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
		file_headscale_v1_acls_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListACLPolicyRequest); i {
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
		file_headscale_v1_acls_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListACLPolicyResponse); i {
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
			RawDescriptor: file_headscale_v1_acls_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_headscale_v1_acls_proto_goTypes,
		DependencyIndexes: file_headscale_v1_acls_proto_depIdxs,
		MessageInfos:      file_headscale_v1_acls_proto_msgTypes,
	}.Build()
	File_headscale_v1_acls_proto = out.File
	file_headscale_v1_acls_proto_rawDesc = nil
	file_headscale_v1_acls_proto_goTypes = nil
	file_headscale_v1_acls_proto_depIdxs = nil
}
