// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: user_service.proto

package auth_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UpsertUserInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string           `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Data   *structpb.Struct `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UpsertUserInfoRequest) Reset() {
	*x = UpsertUserInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertUserInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertUserInfoRequest) ProtoMessage() {}

func (x *UpsertUserInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertUserInfoRequest.ProtoReflect.Descriptor instead.
func (*UpsertUserInfoRequest) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{0}
}

func (x *UpsertUserInfoRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpsertUserInfoRequest) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId        string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ClientPlatformId string `protobuf:"bytes,3,opt,name=client_platform_id,json=clientPlatformId,proto3" json:"client_platform_id,omitempty"`
	ClientTypeId     string `protobuf:"bytes,4,opt,name=client_type_id,json=clientTypeId,proto3" json:"client_type_id,omitempty"`
	RoleId           string `protobuf:"bytes,5,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	Phone            string `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"`
	Email            string `protobuf:"bytes,7,opt,name=email,proto3" json:"email,omitempty"`
	Login            string `protobuf:"bytes,8,opt,name=login,proto3" json:"login,omitempty"`
	Password         string `protobuf:"bytes,9,opt,name=password,proto3" json:"password,omitempty"`
	Active           int32  `protobuf:"varint,10,opt,name=active,proto3" json:"active,omitempty"`
	ExpiresAt        string `protobuf:"bytes,11,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *CreateUserRequest) GetClientPlatformId() string {
	if x != nil {
		return x.ClientPlatformId
	}
	return ""
}

func (x *CreateUserRequest) GetClientTypeId() string {
	if x != nil {
		return x.ClientTypeId
	}
	return ""
}

func (x *CreateUserRequest) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *CreateUserRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUserRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *CreateUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateUserRequest) GetActive() int32 {
	if x != nil {
		return x.Active
	}
	return 0
}

func (x *CreateUserRequest) GetExpiresAt() string {
	if x != nil {
		return x.ExpiresAt
	}
	return ""
}

type UserPrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserPrimaryKey) Reset() {
	*x = UserPrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPrimaryKey) ProtoMessage() {}

func (x *UserPrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPrimaryKey.ProtoReflect.Descriptor instead.
func (*UserPrimaryKey) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{2}
}

func (x *UserPrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetUserListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *GetUserListRequest) Reset() {
	*x = GetUserListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserListRequest) ProtoMessage() {}

func (x *GetUserListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserListRequest.ProtoReflect.Descriptor instead.
func (*GetUserListRequest) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserListRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetUserListRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetUserListRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type GetUserListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32   `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Users []*User `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *GetUserListResponse) Reset() {
	*x = GetUserListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserListResponse) ProtoMessage() {}

func (x *GetUserListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserListResponse.ProtoReflect.Descriptor instead.
func (*GetUserListResponse) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserListResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetUserListResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type UpdateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProjectId        string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ClientPlatformId string `protobuf:"bytes,3,opt,name=client_platform_id,json=clientPlatformId,proto3" json:"client_platform_id,omitempty"`
	ClientTypeId     string `protobuf:"bytes,4,opt,name=client_type_id,json=clientTypeId,proto3" json:"client_type_id,omitempty"`
	RoleId           string `protobuf:"bytes,5,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	Phone            string `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"`
	Email            string `protobuf:"bytes,7,opt,name=email,proto3" json:"email,omitempty"`
	Login            string `protobuf:"bytes,8,opt,name=login,proto3" json:"login,omitempty"`
	Active           int32  `protobuf:"varint,10,opt,name=active,proto3" json:"active,omitempty"`
	ExpiresAt        string `protobuf:"bytes,11,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateUserRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateUserRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *UpdateUserRequest) GetClientPlatformId() string {
	if x != nil {
		return x.ClientPlatformId
	}
	return ""
}

func (x *UpdateUserRequest) GetClientTypeId() string {
	if x != nil {
		return x.ClientTypeId
	}
	return ""
}

func (x *UpdateUserRequest) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *UpdateUserRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UpdateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateUserRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *UpdateUserRequest) GetActive() int32 {
	if x != nil {
		return x.Active
	}
	return 0
}

func (x *UpdateUserRequest) GetExpiresAt() string {
	if x != nil {
		return x.ExpiresAt
	}
	return ""
}

type AddUserRelationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RelationId string `protobuf:"bytes,2,opt,name=relation_id,json=relationId,proto3" json:"relation_id,omitempty"`
}

func (x *AddUserRelationRequest) Reset() {
	*x = AddUserRelationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserRelationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserRelationRequest) ProtoMessage() {}

func (x *AddUserRelationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserRelationRequest.ProtoReflect.Descriptor instead.
func (*AddUserRelationRequest) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{6}
}

func (x *AddUserRelationRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddUserRelationRequest) GetRelationId() string {
	if x != nil {
		return x.RelationId
	}
	return ""
}

type UserRelationPrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RelationId string `protobuf:"bytes,2,opt,name=relation_id,json=relationId,proto3" json:"relation_id,omitempty"`
}

func (x *UserRelationPrimaryKey) Reset() {
	*x = UserRelationPrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRelationPrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRelationPrimaryKey) ProtoMessage() {}

func (x *UserRelationPrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRelationPrimaryKey.ProtoReflect.Descriptor instead.
func (*UserRelationPrimaryKey) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{7}
}

func (x *UserRelationPrimaryKey) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserRelationPrimaryKey) GetRelationId() string {
	if x != nil {
		return x.RelationId
	}
	return ""
}

type UserInfoPrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserInfoPrimaryKey) Reset() {
	*x = UserInfoPrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoPrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoPrimaryKey) ProtoMessage() {}

func (x *UserInfoPrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoPrimaryKey.ProtoReflect.Descriptor instead.
func (*UserInfoPrimaryKey) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{8}
}

func (x *UserInfoPrimaryKey) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_user_service_proto protoreflect.FileDescriptor

var file_user_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x1a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x15, 0x55, 0x70, 0x73,
	0x65, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xb4, 0x02, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x2c, 0x0a,
	0x12, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x22,
	0x20, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x5a, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x55, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x05, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x22, 0xa8, 0x02, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x22,
	0x52, 0x0a, 0x16, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x22, 0x52, 0x0a, 0x16, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x32, 0xf8, 0x04, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1c, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x12, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x54,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x1f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0a, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x55, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x24, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b,
	0x65, 0x79, 0x1a, 0x1a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00,
	0x12, 0x4f, 0x0a, 0x0e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x23, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x00, 0x42, 0x17, 0x5a, 0x15, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_user_service_proto_rawDescOnce sync.Once
	file_user_service_proto_rawDescData = file_user_service_proto_rawDesc
)

func file_user_service_proto_rawDescGZIP() []byte {
	file_user_service_proto_rawDescOnce.Do(func() {
		file_user_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_service_proto_rawDescData)
	})
	return file_user_service_proto_rawDescData
}

var file_user_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_user_service_proto_goTypes = []interface{}{
	(*UpsertUserInfoRequest)(nil),  // 0: auth_service.UpsertUserInfoRequest
	(*CreateUserRequest)(nil),      // 1: auth_service.CreateUserRequest
	(*UserPrimaryKey)(nil),         // 2: auth_service.UserPrimaryKey
	(*GetUserListRequest)(nil),     // 3: auth_service.GetUserListRequest
	(*GetUserListResponse)(nil),    // 4: auth_service.GetUserListResponse
	(*UpdateUserRequest)(nil),      // 5: auth_service.UpdateUserRequest
	(*AddUserRelationRequest)(nil), // 6: auth_service.AddUserRelationRequest
	(*UserRelationPrimaryKey)(nil), // 7: auth_service.UserRelationPrimaryKey
	(*UserInfoPrimaryKey)(nil),     // 8: auth_service.UserInfoPrimaryKey
	(*structpb.Struct)(nil),        // 9: google.protobuf.Struct
	(*User)(nil),                   // 10: auth_service.User
	(*emptypb.Empty)(nil),          // 11: google.protobuf.Empty
	(*UserRelation)(nil),           // 12: auth_service.UserRelation
	(*UserInfo)(nil),               // 13: auth_service.UserInfo
}
var file_user_service_proto_depIdxs = []int32{
	9,  // 0: auth_service.UpsertUserInfoRequest.data:type_name -> google.protobuf.Struct
	10, // 1: auth_service.GetUserListResponse.users:type_name -> auth_service.User
	1,  // 2: auth_service.UserService.CreateUser:input_type -> auth_service.CreateUserRequest
	2,  // 3: auth_service.UserService.GetUserByID:input_type -> auth_service.UserPrimaryKey
	3,  // 4: auth_service.UserService.GetUserList:input_type -> auth_service.GetUserListRequest
	5,  // 5: auth_service.UserService.UpdateUser:input_type -> auth_service.UpdateUserRequest
	2,  // 6: auth_service.UserService.DeleteUser:input_type -> auth_service.UserPrimaryKey
	6,  // 7: auth_service.UserService.AddUserRelation:input_type -> auth_service.AddUserRelationRequest
	7,  // 8: auth_service.UserService.RemoveUserRelation:input_type -> auth_service.UserRelationPrimaryKey
	0,  // 9: auth_service.UserService.UpsertUserInfo:input_type -> auth_service.UpsertUserInfoRequest
	10, // 10: auth_service.UserService.CreateUser:output_type -> auth_service.User
	10, // 11: auth_service.UserService.GetUserByID:output_type -> auth_service.User
	4,  // 12: auth_service.UserService.GetUserList:output_type -> auth_service.GetUserListResponse
	10, // 13: auth_service.UserService.UpdateUser:output_type -> auth_service.User
	11, // 14: auth_service.UserService.DeleteUser:output_type -> google.protobuf.Empty
	12, // 15: auth_service.UserService.AddUserRelation:output_type -> auth_service.UserRelation
	12, // 16: auth_service.UserService.RemoveUserRelation:output_type -> auth_service.UserRelation
	13, // 17: auth_service.UserService.UpsertUserInfo:output_type -> auth_service.UserInfo
	10, // [10:18] is the sub-list for method output_type
	2,  // [2:10] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_user_service_proto_init() }
func file_user_service_proto_init() {
	if File_user_service_proto != nil {
		return
	}
	file_auth_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_user_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertUserInfoRequest); i {
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
		file_user_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_user_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPrimaryKey); i {
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
		file_user_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserListRequest); i {
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
		file_user_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserListResponse); i {
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
		file_user_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserRequest); i {
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
		file_user_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserRelationRequest); i {
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
		file_user_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRelationPrimaryKey); i {
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
		file_user_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoPrimaryKey); i {
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
			RawDescriptor: file_user_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_service_proto_goTypes,
		DependencyIndexes: file_user_service_proto_depIdxs,
		MessageInfos:      file_user_service_proto_msgTypes,
	}.Build()
	File_user_service_proto = out.File
	file_user_service_proto_rawDesc = nil
	file_user_service_proto_goTypes = nil
	file_user_service_proto_depIdxs = nil
}
