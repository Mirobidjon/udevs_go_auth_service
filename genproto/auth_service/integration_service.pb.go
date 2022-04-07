// Code generated by protoc-gen-go. DO NOT EDIT.
// source: integration_service.proto

package auth_service

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/struct"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CreateIntegrationRequest struct {
	ProjectId            string   `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ClientPlatformId     string   `protobuf:"bytes,2,opt,name=client_platform_id,json=clientPlatformId,proto3" json:"client_platform_id,omitempty"`
	ClientTypeId         string   `protobuf:"bytes,3,opt,name=client_type_id,json=clientTypeId,proto3" json:"client_type_id,omitempty"`
	RoleId               string   `protobuf:"bytes,4,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	SecretKey            string   `protobuf:"bytes,5,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	Active               int32    `protobuf:"varint,6,opt,name=active,proto3" json:"active,omitempty"`
	ExpiresAt            string   `protobuf:"bytes,7,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	Title                string   `protobuf:"bytes,8,opt,name=title,proto3" json:"title,omitempty"`
	IpWhitelist          string   `protobuf:"bytes,9,opt,name=ip_whitelist,json=ipWhitelist,proto3" json:"ip_whitelist,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateIntegrationRequest) Reset()         { *m = CreateIntegrationRequest{} }
func (m *CreateIntegrationRequest) String() string { return proto.CompactTextString(m) }
func (*CreateIntegrationRequest) ProtoMessage()    {}
func (*CreateIntegrationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b223b7acfa203af, []int{0}
}

func (m *CreateIntegrationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateIntegrationRequest.Unmarshal(m, b)
}
func (m *CreateIntegrationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateIntegrationRequest.Marshal(b, m, deterministic)
}
func (m *CreateIntegrationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateIntegrationRequest.Merge(m, src)
}
func (m *CreateIntegrationRequest) XXX_Size() int {
	return xxx_messageInfo_CreateIntegrationRequest.Size(m)
}
func (m *CreateIntegrationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateIntegrationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateIntegrationRequest proto.InternalMessageInfo

func (m *CreateIntegrationRequest) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *CreateIntegrationRequest) GetClientPlatformId() string {
	if m != nil {
		return m.ClientPlatformId
	}
	return ""
}

func (m *CreateIntegrationRequest) GetClientTypeId() string {
	if m != nil {
		return m.ClientTypeId
	}
	return ""
}

func (m *CreateIntegrationRequest) GetRoleId() string {
	if m != nil {
		return m.RoleId
	}
	return ""
}

func (m *CreateIntegrationRequest) GetSecretKey() string {
	if m != nil {
		return m.SecretKey
	}
	return ""
}

func (m *CreateIntegrationRequest) GetActive() int32 {
	if m != nil {
		return m.Active
	}
	return 0
}

func (m *CreateIntegrationRequest) GetExpiresAt() string {
	if m != nil {
		return m.ExpiresAt
	}
	return ""
}

func (m *CreateIntegrationRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CreateIntegrationRequest) GetIpWhitelist() string {
	if m != nil {
		return m.IpWhitelist
	}
	return ""
}

type IntegrationPrimaryKey struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntegrationPrimaryKey) Reset()         { *m = IntegrationPrimaryKey{} }
func (m *IntegrationPrimaryKey) String() string { return proto.CompactTextString(m) }
func (*IntegrationPrimaryKey) ProtoMessage()    {}
func (*IntegrationPrimaryKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b223b7acfa203af, []int{1}
}

func (m *IntegrationPrimaryKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntegrationPrimaryKey.Unmarshal(m, b)
}
func (m *IntegrationPrimaryKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntegrationPrimaryKey.Marshal(b, m, deterministic)
}
func (m *IntegrationPrimaryKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntegrationPrimaryKey.Merge(m, src)
}
func (m *IntegrationPrimaryKey) XXX_Size() int {
	return xxx_messageInfo_IntegrationPrimaryKey.Size(m)
}
func (m *IntegrationPrimaryKey) XXX_DiscardUnknown() {
	xxx_messageInfo_IntegrationPrimaryKey.DiscardUnknown(m)
}

var xxx_messageInfo_IntegrationPrimaryKey proto.InternalMessageInfo

func (m *IntegrationPrimaryKey) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type IntegrationPrimaryKeyList struct {
	Ids                  []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntegrationPrimaryKeyList) Reset()         { *m = IntegrationPrimaryKeyList{} }
func (m *IntegrationPrimaryKeyList) String() string { return proto.CompactTextString(m) }
func (*IntegrationPrimaryKeyList) ProtoMessage()    {}
func (*IntegrationPrimaryKeyList) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b223b7acfa203af, []int{2}
}

func (m *IntegrationPrimaryKeyList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntegrationPrimaryKeyList.Unmarshal(m, b)
}
func (m *IntegrationPrimaryKeyList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntegrationPrimaryKeyList.Marshal(b, m, deterministic)
}
func (m *IntegrationPrimaryKeyList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntegrationPrimaryKeyList.Merge(m, src)
}
func (m *IntegrationPrimaryKeyList) XXX_Size() int {
	return xxx_messageInfo_IntegrationPrimaryKeyList.Size(m)
}
func (m *IntegrationPrimaryKeyList) XXX_DiscardUnknown() {
	xxx_messageInfo_IntegrationPrimaryKeyList.DiscardUnknown(m)
}

var xxx_messageInfo_IntegrationPrimaryKeyList proto.InternalMessageInfo

func (m *IntegrationPrimaryKeyList) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type GetIntegrationListRequest struct {
	Limit                int32    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int32    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Search               string   `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
	ClientPlatformId     string   `protobuf:"bytes,4,opt,name=client_platform_id,json=clientPlatformId,proto3" json:"client_platform_id,omitempty"`
	ClientTypeId         string   `protobuf:"bytes,5,opt,name=client_type_id,json=clientTypeId,proto3" json:"client_type_id,omitempty"`
	ProjectId            string   `protobuf:"bytes,6,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetIntegrationListRequest) Reset()         { *m = GetIntegrationListRequest{} }
func (m *GetIntegrationListRequest) String() string { return proto.CompactTextString(m) }
func (*GetIntegrationListRequest) ProtoMessage()    {}
func (*GetIntegrationListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b223b7acfa203af, []int{3}
}

func (m *GetIntegrationListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetIntegrationListRequest.Unmarshal(m, b)
}
func (m *GetIntegrationListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetIntegrationListRequest.Marshal(b, m, deterministic)
}
func (m *GetIntegrationListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetIntegrationListRequest.Merge(m, src)
}
func (m *GetIntegrationListRequest) XXX_Size() int {
	return xxx_messageInfo_GetIntegrationListRequest.Size(m)
}
func (m *GetIntegrationListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetIntegrationListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetIntegrationListRequest proto.InternalMessageInfo

func (m *GetIntegrationListRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetIntegrationListRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *GetIntegrationListRequest) GetSearch() string {
	if m != nil {
		return m.Search
	}
	return ""
}

func (m *GetIntegrationListRequest) GetClientPlatformId() string {
	if m != nil {
		return m.ClientPlatformId
	}
	return ""
}

func (m *GetIntegrationListRequest) GetClientTypeId() string {
	if m != nil {
		return m.ClientTypeId
	}
	return ""
}

func (m *GetIntegrationListRequest) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

type GetIntegrationListResponse struct {
	Count                int32          `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Integrations         []*Integration `protobuf:"bytes,2,rep,name=integrations,proto3" json:"integrations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetIntegrationListResponse) Reset()         { *m = GetIntegrationListResponse{} }
func (m *GetIntegrationListResponse) String() string { return proto.CompactTextString(m) }
func (*GetIntegrationListResponse) ProtoMessage()    {}
func (*GetIntegrationListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b223b7acfa203af, []int{4}
}

func (m *GetIntegrationListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetIntegrationListResponse.Unmarshal(m, b)
}
func (m *GetIntegrationListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetIntegrationListResponse.Marshal(b, m, deterministic)
}
func (m *GetIntegrationListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetIntegrationListResponse.Merge(m, src)
}
func (m *GetIntegrationListResponse) XXX_Size() int {
	return xxx_messageInfo_GetIntegrationListResponse.Size(m)
}
func (m *GetIntegrationListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetIntegrationListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetIntegrationListResponse proto.InternalMessageInfo

func (m *GetIntegrationListResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *GetIntegrationListResponse) GetIntegrations() []*Integration {
	if m != nil {
		return m.Integrations
	}
	return nil
}

type UpdateIntegrationRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProjectId            string   `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ClientPlatformId     string   `protobuf:"bytes,3,opt,name=client_platform_id,json=clientPlatformId,proto3" json:"client_platform_id,omitempty"`
	ClientTypeId         string   `protobuf:"bytes,4,opt,name=client_type_id,json=clientTypeId,proto3" json:"client_type_id,omitempty"`
	RoleId               string   `protobuf:"bytes,5,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	Active               int32    `protobuf:"varint,6,opt,name=active,proto3" json:"active,omitempty"`
	ExpiresAt            string   `protobuf:"bytes,7,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	Title                string   `protobuf:"bytes,8,opt,name=title,proto3" json:"title,omitempty"`
	IpWhitelist          string   `protobuf:"bytes,9,opt,name=ip_whitelist,json=ipWhitelist,proto3" json:"ip_whitelist,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateIntegrationRequest) Reset()         { *m = UpdateIntegrationRequest{} }
func (m *UpdateIntegrationRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateIntegrationRequest) ProtoMessage()    {}
func (*UpdateIntegrationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b223b7acfa203af, []int{5}
}

func (m *UpdateIntegrationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateIntegrationRequest.Unmarshal(m, b)
}
func (m *UpdateIntegrationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateIntegrationRequest.Marshal(b, m, deterministic)
}
func (m *UpdateIntegrationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateIntegrationRequest.Merge(m, src)
}
func (m *UpdateIntegrationRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateIntegrationRequest.Size(m)
}
func (m *UpdateIntegrationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateIntegrationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateIntegrationRequest proto.InternalMessageInfo

func (m *UpdateIntegrationRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateIntegrationRequest) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *UpdateIntegrationRequest) GetClientPlatformId() string {
	if m != nil {
		return m.ClientPlatformId
	}
	return ""
}

func (m *UpdateIntegrationRequest) GetClientTypeId() string {
	if m != nil {
		return m.ClientTypeId
	}
	return ""
}

func (m *UpdateIntegrationRequest) GetRoleId() string {
	if m != nil {
		return m.RoleId
	}
	return ""
}

func (m *UpdateIntegrationRequest) GetActive() int32 {
	if m != nil {
		return m.Active
	}
	return 0
}

func (m *UpdateIntegrationRequest) GetExpiresAt() string {
	if m != nil {
		return m.ExpiresAt
	}
	return ""
}

func (m *UpdateIntegrationRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *UpdateIntegrationRequest) GetIpWhitelist() string {
	if m != nil {
		return m.IpWhitelist
	}
	return ""
}

type AddIntegrationRelationRequest struct {
	IntegrationId        string   `protobuf:"bytes,1,opt,name=integration_id,json=integrationId,proto3" json:"integration_id,omitempty"`
	RelationId           string   `protobuf:"bytes,2,opt,name=relation_id,json=relationId,proto3" json:"relation_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddIntegrationRelationRequest) Reset()         { *m = AddIntegrationRelationRequest{} }
func (m *AddIntegrationRelationRequest) String() string { return proto.CompactTextString(m) }
func (*AddIntegrationRelationRequest) ProtoMessage()    {}
func (*AddIntegrationRelationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b223b7acfa203af, []int{6}
}

func (m *AddIntegrationRelationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddIntegrationRelationRequest.Unmarshal(m, b)
}
func (m *AddIntegrationRelationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddIntegrationRelationRequest.Marshal(b, m, deterministic)
}
func (m *AddIntegrationRelationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddIntegrationRelationRequest.Merge(m, src)
}
func (m *AddIntegrationRelationRequest) XXX_Size() int {
	return xxx_messageInfo_AddIntegrationRelationRequest.Size(m)
}
func (m *AddIntegrationRelationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddIntegrationRelationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddIntegrationRelationRequest proto.InternalMessageInfo

func (m *AddIntegrationRelationRequest) GetIntegrationId() string {
	if m != nil {
		return m.IntegrationId
	}
	return ""
}

func (m *AddIntegrationRelationRequest) GetRelationId() string {
	if m != nil {
		return m.RelationId
	}
	return ""
}

type IntegrationRelationPrimaryKey struct {
	IntegrationId        string   `protobuf:"bytes,1,opt,name=integration_id,json=integrationId,proto3" json:"integration_id,omitempty"`
	RelationId           string   `protobuf:"bytes,2,opt,name=relation_id,json=relationId,proto3" json:"relation_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntegrationRelationPrimaryKey) Reset()         { *m = IntegrationRelationPrimaryKey{} }
func (m *IntegrationRelationPrimaryKey) String() string { return proto.CompactTextString(m) }
func (*IntegrationRelationPrimaryKey) ProtoMessage()    {}
func (*IntegrationRelationPrimaryKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b223b7acfa203af, []int{7}
}

func (m *IntegrationRelationPrimaryKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntegrationRelationPrimaryKey.Unmarshal(m, b)
}
func (m *IntegrationRelationPrimaryKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntegrationRelationPrimaryKey.Marshal(b, m, deterministic)
}
func (m *IntegrationRelationPrimaryKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntegrationRelationPrimaryKey.Merge(m, src)
}
func (m *IntegrationRelationPrimaryKey) XXX_Size() int {
	return xxx_messageInfo_IntegrationRelationPrimaryKey.Size(m)
}
func (m *IntegrationRelationPrimaryKey) XXX_DiscardUnknown() {
	xxx_messageInfo_IntegrationRelationPrimaryKey.DiscardUnknown(m)
}

var xxx_messageInfo_IntegrationRelationPrimaryKey proto.InternalMessageInfo

func (m *IntegrationRelationPrimaryKey) GetIntegrationId() string {
	if m != nil {
		return m.IntegrationId
	}
	return ""
}

func (m *IntegrationRelationPrimaryKey) GetRelationId() string {
	if m != nil {
		return m.RelationId
	}
	return ""
}

type GetIntegrationSessionsResponse struct {
	Sessions             []*Session `protobuf:"bytes,1,rep,name=sessions,proto3" json:"sessions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetIntegrationSessionsResponse) Reset()         { *m = GetIntegrationSessionsResponse{} }
func (m *GetIntegrationSessionsResponse) String() string { return proto.CompactTextString(m) }
func (*GetIntegrationSessionsResponse) ProtoMessage()    {}
func (*GetIntegrationSessionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b223b7acfa203af, []int{8}
}

func (m *GetIntegrationSessionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetIntegrationSessionsResponse.Unmarshal(m, b)
}
func (m *GetIntegrationSessionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetIntegrationSessionsResponse.Marshal(b, m, deterministic)
}
func (m *GetIntegrationSessionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetIntegrationSessionsResponse.Merge(m, src)
}
func (m *GetIntegrationSessionsResponse) XXX_Size() int {
	return xxx_messageInfo_GetIntegrationSessionsResponse.Size(m)
}
func (m *GetIntegrationSessionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetIntegrationSessionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetIntegrationSessionsResponse proto.InternalMessageInfo

func (m *GetIntegrationSessionsResponse) GetSessions() []*Session {
	if m != nil {
		return m.Sessions
	}
	return nil
}

type AddSessionToIntegrationRequest struct {
	IntegrationId        string   `protobuf:"bytes,1,opt,name=integration_id,json=integrationId,proto3" json:"integration_id,omitempty"`
	SecretKey            string   `protobuf:"bytes,2,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	ExpiresAt            string   `protobuf:"bytes,3,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	Date                 string   `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddSessionToIntegrationRequest) Reset()         { *m = AddSessionToIntegrationRequest{} }
func (m *AddSessionToIntegrationRequest) String() string { return proto.CompactTextString(m) }
func (*AddSessionToIntegrationRequest) ProtoMessage()    {}
func (*AddSessionToIntegrationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b223b7acfa203af, []int{9}
}

func (m *AddSessionToIntegrationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddSessionToIntegrationRequest.Unmarshal(m, b)
}
func (m *AddSessionToIntegrationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddSessionToIntegrationRequest.Marshal(b, m, deterministic)
}
func (m *AddSessionToIntegrationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddSessionToIntegrationRequest.Merge(m, src)
}
func (m *AddSessionToIntegrationRequest) XXX_Size() int {
	return xxx_messageInfo_AddSessionToIntegrationRequest.Size(m)
}
func (m *AddSessionToIntegrationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddSessionToIntegrationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddSessionToIntegrationRequest proto.InternalMessageInfo

func (m *AddSessionToIntegrationRequest) GetIntegrationId() string {
	if m != nil {
		return m.IntegrationId
	}
	return ""
}

func (m *AddSessionToIntegrationRequest) GetSecretKey() string {
	if m != nil {
		return m.SecretKey
	}
	return ""
}

func (m *AddSessionToIntegrationRequest) GetExpiresAt() string {
	if m != nil {
		return m.ExpiresAt
	}
	return ""
}

func (m *AddSessionToIntegrationRequest) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

type AddSessionToIntegrationResponse struct {
	IntegrationFound     bool            `protobuf:"varint,1,opt,name=integration_found,json=integrationFound,proto3" json:"integration_found,omitempty"`
	ClientPlatform       *ClientPlatform `protobuf:"bytes,2,opt,name=client_platform,json=clientPlatform,proto3" json:"client_platform,omitempty"`
	ClientType           *ClientType     `protobuf:"bytes,3,opt,name=client_type,json=clientType,proto3" json:"client_type,omitempty"`
	Integration          *Integration    `protobuf:"bytes,4,opt,name=integration,proto3" json:"integration,omitempty"`
	Token                *Token          `protobuf:"bytes,6,opt,name=token,proto3" json:"token,omitempty"`
	Permissions          []*Permission   `protobuf:"bytes,7,rep,name=permissions,proto3" json:"permissions,omitempty"`
	Session              *Session        `protobuf:"bytes,8,opt,name=session,proto3" json:"session,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AddSessionToIntegrationResponse) Reset()         { *m = AddSessionToIntegrationResponse{} }
func (m *AddSessionToIntegrationResponse) String() string { return proto.CompactTextString(m) }
func (*AddSessionToIntegrationResponse) ProtoMessage()    {}
func (*AddSessionToIntegrationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b223b7acfa203af, []int{10}
}

func (m *AddSessionToIntegrationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddSessionToIntegrationResponse.Unmarshal(m, b)
}
func (m *AddSessionToIntegrationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddSessionToIntegrationResponse.Marshal(b, m, deterministic)
}
func (m *AddSessionToIntegrationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddSessionToIntegrationResponse.Merge(m, src)
}
func (m *AddSessionToIntegrationResponse) XXX_Size() int {
	return xxx_messageInfo_AddSessionToIntegrationResponse.Size(m)
}
func (m *AddSessionToIntegrationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddSessionToIntegrationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddSessionToIntegrationResponse proto.InternalMessageInfo

func (m *AddSessionToIntegrationResponse) GetIntegrationFound() bool {
	if m != nil {
		return m.IntegrationFound
	}
	return false
}

func (m *AddSessionToIntegrationResponse) GetClientPlatform() *ClientPlatform {
	if m != nil {
		return m.ClientPlatform
	}
	return nil
}

func (m *AddSessionToIntegrationResponse) GetClientType() *ClientType {
	if m != nil {
		return m.ClientType
	}
	return nil
}

func (m *AddSessionToIntegrationResponse) GetIntegration() *Integration {
	if m != nil {
		return m.Integration
	}
	return nil
}

func (m *AddSessionToIntegrationResponse) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *AddSessionToIntegrationResponse) GetPermissions() []*Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *AddSessionToIntegrationResponse) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

type GetIntegrationTokenRequest struct {
	IntegrationId        string   `protobuf:"bytes,1,opt,name=integration_id,json=integrationId,proto3" json:"integration_id,omitempty"`
	SessionId            string   `protobuf:"bytes,2,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetIntegrationTokenRequest) Reset()         { *m = GetIntegrationTokenRequest{} }
func (m *GetIntegrationTokenRequest) String() string { return proto.CompactTextString(m) }
func (*GetIntegrationTokenRequest) ProtoMessage()    {}
func (*GetIntegrationTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b223b7acfa203af, []int{11}
}

func (m *GetIntegrationTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetIntegrationTokenRequest.Unmarshal(m, b)
}
func (m *GetIntegrationTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetIntegrationTokenRequest.Marshal(b, m, deterministic)
}
func (m *GetIntegrationTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetIntegrationTokenRequest.Merge(m, src)
}
func (m *GetIntegrationTokenRequest) XXX_Size() int {
	return xxx_messageInfo_GetIntegrationTokenRequest.Size(m)
}
func (m *GetIntegrationTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetIntegrationTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetIntegrationTokenRequest proto.InternalMessageInfo

func (m *GetIntegrationTokenRequest) GetIntegrationId() string {
	if m != nil {
		return m.IntegrationId
	}
	return ""
}

func (m *GetIntegrationTokenRequest) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateIntegrationRequest)(nil), "auth_service.CreateIntegrationRequest")
	proto.RegisterType((*IntegrationPrimaryKey)(nil), "auth_service.IntegrationPrimaryKey")
	proto.RegisterType((*IntegrationPrimaryKeyList)(nil), "auth_service.IntegrationPrimaryKeyList")
	proto.RegisterType((*GetIntegrationListRequest)(nil), "auth_service.GetIntegrationListRequest")
	proto.RegisterType((*GetIntegrationListResponse)(nil), "auth_service.GetIntegrationListResponse")
	proto.RegisterType((*UpdateIntegrationRequest)(nil), "auth_service.UpdateIntegrationRequest")
	proto.RegisterType((*AddIntegrationRelationRequest)(nil), "auth_service.AddIntegrationRelationRequest")
	proto.RegisterType((*IntegrationRelationPrimaryKey)(nil), "auth_service.IntegrationRelationPrimaryKey")
	proto.RegisterType((*GetIntegrationSessionsResponse)(nil), "auth_service.GetIntegrationSessionsResponse")
	proto.RegisterType((*AddSessionToIntegrationRequest)(nil), "auth_service.AddSessionToIntegrationRequest")
	proto.RegisterType((*AddSessionToIntegrationResponse)(nil), "auth_service.AddSessionToIntegrationResponse")
	proto.RegisterType((*GetIntegrationTokenRequest)(nil), "auth_service.GetIntegrationTokenRequest")
}

func init() { proto.RegisterFile("integration_service.proto", fileDescriptor_7b223b7acfa203af) }

var fileDescriptor_7b223b7acfa203af = []byte{
	// 942 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0x5d, 0x72, 0xe3, 0x44,
	0x10, 0xf6, 0x4f, 0x64, 0x27, 0xad, 0x10, 0x92, 0xd9, 0x4d, 0xa2, 0x98, 0xec, 0x6e, 0x10, 0x3f,
	0x31, 0xc5, 0xae, 0x53, 0x98, 0x27, 0xa0, 0x78, 0xc8, 0xfe, 0x51, 0x2a, 0x78, 0x48, 0x69, 0xc3,
	0xb2, 0xc5, 0x8b, 0xcb, 0xb1, 0xda, 0xce, 0xb0, 0xb2, 0xa4, 0xd5, 0x8c, 0x97, 0xf8, 0x28, 0x1c,
	0x00, 0x6e, 0xc0, 0x4d, 0x38, 0x0a, 0x07, 0xa0, 0x66, 0x34, 0xb6, 0x67, 0x14, 0xc9, 0x08, 0x8a,
	0xe2, 0x4d, 0xdd, 0xfd, 0x4d, 0x4f, 0xff, 0x7c, 0xdd, 0x23, 0x38, 0xa2, 0x11, 0xc7, 0x49, 0x3a,
	0xe4, 0x34, 0x8e, 0x06, 0x0c, 0xd3, 0xb7, 0x74, 0x84, 0xbd, 0x24, 0x8d, 0x79, 0x4c, 0xb6, 0x87,
	0x33, 0x7e, 0xbd, 0xd0, 0x75, 0x40, 0x48, 0x99, 0xa5, 0xf3, 0xde, 0x24, 0x8e, 0x27, 0x21, 0x9e,
	0x49, 0xe9, 0x6a, 0x36, 0x3e, 0xc3, 0x69, 0xc2, 0xe7, 0xca, 0x78, 0x9c, 0x37, 0x32, 0x9e, 0xce,
	0x46, 0x3c, 0xb3, 0xba, 0xbf, 0x37, 0xc0, 0x79, 0x92, 0xe2, 0x90, 0xa3, 0xb7, 0xba, 0xd8, 0xc7,
	0x37, 0x33, 0x64, 0x9c, 0xdc, 0x03, 0x48, 0xd2, 0xf8, 0x27, 0x1c, 0xf1, 0x01, 0x0d, 0x9c, 0xfa,
	0x49, 0xbd, 0xbb, 0xe5, 0x6f, 0x29, 0x8d, 0x17, 0x90, 0x87, 0x40, 0x46, 0x21, 0xc5, 0x88, 0x0f,
	0x92, 0x70, 0xc8, 0xc7, 0x71, 0x3a, 0x15, 0xb0, 0x86, 0x84, 0xed, 0x66, 0x96, 0x0b, 0x65, 0xf0,
	0x02, 0xf2, 0x21, 0xec, 0x28, 0x34, 0x9f, 0x27, 0x28, 0x90, 0x4d, 0x89, 0xdc, 0xce, 0xb4, 0x97,
	0xf3, 0x04, 0xbd, 0x80, 0x1c, 0x42, 0x3b, 0x8d, 0x43, 0x69, 0xde, 0x90, 0xe6, 0x96, 0x10, 0xbd,
	0x40, 0xc4, 0xc2, 0x70, 0x94, 0x22, 0x1f, 0xbc, 0xc6, 0xb9, 0x63, 0x65, 0xb1, 0x64, 0x9a, 0x6f,
	0x71, 0x4e, 0x0e, 0xa0, 0x35, 0x1c, 0x71, 0xfa, 0x16, 0x9d, 0xd6, 0x49, 0xbd, 0x6b, 0xf9, 0x4a,
	0x12, 0xc7, 0xf0, 0x26, 0xa1, 0x29, 0xb2, 0xc1, 0x90, 0x3b, 0xed, 0xec, 0x98, 0xd2, 0x9c, 0x73,
	0x72, 0x17, 0x2c, 0x4e, 0x79, 0x88, 0xce, 0xa6, 0xb4, 0x64, 0x02, 0x79, 0x1f, 0xb6, 0x69, 0x32,
	0xf8, 0xf9, 0x9a, 0x72, 0x0c, 0x29, 0xe3, 0xce, 0x96, 0x34, 0xda, 0x34, 0xf9, 0x61, 0xa1, 0x72,
	0x4f, 0x61, 0x5f, 0x2b, 0xd8, 0x45, 0x4a, 0xa7, 0xc3, 0x74, 0x2e, 0x02, 0xd9, 0x81, 0xc6, 0xb2,
	0x56, 0x0d, 0x1a, 0xb8, 0x8f, 0xe0, 0xa8, 0x10, 0xf8, 0x1d, 0x65, 0x9c, 0xec, 0x42, 0x93, 0x06,
	0xcc, 0xa9, 0x9f, 0x34, 0xbb, 0x5b, 0xbe, 0xf8, 0x74, 0xff, 0xa8, 0xc3, 0xd1, 0x37, 0xc8, 0xb5,
	0x23, 0x02, 0xb8, 0x68, 0xc8, 0x5d, 0xb0, 0x42, 0x3a, 0xa5, 0x5c, 0xfa, 0xb7, 0xfc, 0x4c, 0x10,
	0xb9, 0xc7, 0xe3, 0x31, 0x43, 0x2e, 0x6b, 0x6f, 0xf9, 0x4a, 0x12, 0x7a, 0x86, 0xc3, 0x74, 0x74,
	0xad, 0x2a, 0xad, 0xa4, 0x92, 0xbe, 0x6d, 0x54, 0xee, 0x9b, 0x55, 0xd0, 0x37, 0x93, 0x2a, 0xad,
	0x1c, 0x55, 0xdc, 0x37, 0xd0, 0x29, 0xca, 0x8a, 0x25, 0x71, 0xc4, 0x50, 0xa4, 0x35, 0x8a, 0x67,
	0xd1, 0x32, 0x2d, 0x29, 0x90, 0xaf, 0x61, 0x5b, 0x1b, 0x06, 0xe6, 0x34, 0x4e, 0x9a, 0x5d, 0xbb,
	0x7f, 0xd4, 0xd3, 0xc7, 0xa0, 0xa7, 0xb3, 0xd6, 0x80, 0xbb, 0xbf, 0x35, 0xc0, 0xf9, 0x3e, 0x09,
	0x8a, 0x99, 0x9d, 0xeb, 0x52, 0x2e, 0xfc, 0x46, 0x35, 0xa6, 0x37, 0x2b, 0x57, 0x6c, 0x63, 0x3d,
	0xd3, 0x2d, 0x83, 0xe9, 0xff, 0x37, 0x95, 0x27, 0x70, 0xef, 0x3c, 0x08, 0x8c, 0x22, 0x85, 0x46,
	0xb1, 0x3e, 0x82, 0x1d, 0x7d, 0x2b, 0x2d, 0x0b, 0xf7, 0x8e, 0xa6, 0xf5, 0x02, 0xf2, 0x00, 0xec,
	0x54, 0x9d, 0x5c, 0x15, 0x11, 0x16, 0x2a, 0x2f, 0x10, 0x17, 0x15, 0xdc, 0xa2, 0xcd, 0xce, 0x7f,
	0x75, 0xd1, 0x0b, 0xb8, 0x6f, 0xb2, 0xed, 0x05, 0x32, 0x26, 0x48, 0xb1, 0x64, 0xdc, 0x67, 0xb0,
	0xc9, 0x94, 0x4e, 0x4e, 0x9f, 0xdd, 0xdf, 0x37, 0x79, 0xa5, 0x4e, 0xf8, 0x4b, 0x98, 0xfb, 0x4b,
	0x1d, 0xee, 0x9f, 0x07, 0x81, 0x32, 0x5c, 0xc6, 0x05, 0xac, 0xaa, 0x18, 0xbf, 0xb9, 0xca, 0x1a,
	0xf9, 0x55, 0x66, 0xf6, 0xb9, 0x99, 0xef, 0x33, 0x81, 0x0d, 0x41, 0x6a, 0xc5, 0x29, 0xf9, 0xed,
	0xfe, 0xda, 0x84, 0x07, 0xa5, 0xb1, 0xa9, 0x94, 0x3f, 0x85, 0x3d, 0x3d, 0xb8, 0x71, 0x3c, 0x8b,
	0xb2, 0xf8, 0x36, 0xfd, 0x5d, 0xcd, 0xf0, 0x5c, 0xe8, 0xc9, 0x33, 0x78, 0x37, 0x47, 0x78, 0x19,
	0xa7, 0xdd, 0x3f, 0x36, 0xcb, 0xf4, 0xc4, 0xe0, 0xbe, 0xbf, 0x63, 0xce, 0x02, 0xf9, 0x02, 0x6c,
	0x6d, 0x12, 0x64, 0x2e, 0x76, 0xdf, 0x29, 0x72, 0x21, 0x86, 0xc2, 0x87, 0xd5, 0x80, 0x90, 0xaf,
	0xc0, 0xd6, 0xa2, 0x92, 0xd9, 0xae, 0x1d, 0x7e, 0x1d, 0x4d, 0x3e, 0x01, 0x8b, 0xc7, 0xaf, 0x31,
	0x92, 0x13, 0x64, 0xf7, 0xef, 0x98, 0xc7, 0x2e, 0x85, 0xc9, 0xcf, 0x10, 0xe4, 0x4b, 0xb0, 0x13,
	0x4c, 0xa7, 0x54, 0x91, 0xa1, 0x2d, 0xc9, 0x90, 0x0b, 0xf1, 0x62, 0x09, 0xf0, 0x75, 0x30, 0x39,
	0x83, 0xb6, 0xa2, 0x87, 0x1c, 0xba, 0x52, 0x12, 0x2d, 0x50, 0xee, 0x55, 0x7e, 0x0d, 0x66, 0xa1,
	0xfc, 0x0b, 0xfa, 0x48, 0x7f, 0xda, 0xae, 0x52, 0x1a, 0x2f, 0xe8, 0xff, 0xd9, 0x06, 0x62, 0x50,
	0x5f, 0xc6, 0x42, 0x5e, 0xc1, 0xde, 0xad, 0x77, 0x9e, 0x7c, 0x9c, 0x6b, 0x45, 0xc9, 0x8f, 0x40,
	0xa7, 0xbc, 0xee, 0x6e, 0x8d, 0xbc, 0x04, 0x62, 0x26, 0xf5, 0x78, 0xee, 0x3d, 0x25, 0x1f, 0x94,
	0x1e, 0x59, 0x0d, 0xfc, 0x7a, 0xbf, 0x11, 0x1c, 0xde, 0x7e, 0x33, 0x84, 0x6f, 0x46, 0x4e, 0x2b,
	0x38, 0x17, 0xe8, 0x4e, 0xd7, 0x04, 0x96, 0xbf, 0x41, 0x6e, 0x8d, 0xd0, 0x7c, 0x1e, 0xf2, 0x89,
	0x3e, 0xfd, 0x7b, 0x0f, 0x59, 0x8d, 0xfe, 0xc9, 0x55, 0xaf, 0x60, 0xef, 0xd6, 0xd3, 0x94, 0x6f,
	0x46, 0xd9, 0xdb, 0xb5, 0xbe, 0x68, 0x3e, 0xec, 0x3d, 0xc5, 0x10, 0x4d, 0xcf, 0x95, 0x7a, 0x71,
	0xd0, 0xcb, 0x7e, 0x14, 0x7b, 0x8b, 0x1f, 0xc5, 0xde, 0x33, 0xf1, 0x17, 0xe9, 0xd6, 0xc8, 0x14,
	0x0e, 0x8a, 0xd7, 0x69, 0x35, 0xc7, 0x0f, 0xd7, 0x15, 0x26, 0xbf, 0x99, 0xdd, 0x1a, 0xb9, 0x81,
	0xc3, 0x92, 0x5d, 0x46, 0x72, 0xae, 0xd6, 0xaf, 0xe3, 0xce, 0xa3, 0x8a, 0xe8, 0xe5, 0xcd, 0x2f,
	0xe1, 0x4e, 0xc1, 0x78, 0x92, 0xb5, 0x9d, 0xd5, 0x27, 0xb8, 0x53, 0xb4, 0x68, 0xdc, 0x1a, 0xb9,
	0x82, 0xe3, 0xac, 0x29, 0xea, 0xfe, 0xe7, 0x69, 0x3c, 0xd5, 0xd3, 0xaa, 0x7e, 0x41, 0x69, 0x93,
	0x1e, 0x1f, 0xfe, 0xb8, 0x3f, 0xc1, 0x48, 0xaa, 0xcf, 0x74, 0x6f, 0x57, 0x2d, 0xa9, 0xfb, 0xfc,
	0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2f, 0xf6, 0x74, 0x33, 0x5a, 0x0c, 0x00, 0x00,
}
