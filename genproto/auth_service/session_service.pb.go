// Code generated by protoc-gen-go. DO NOT EDIT.
// source: session_service.proto

package auth_service

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
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

type LoginRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff25770e79caedc0, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	UserFound            bool            `protobuf:"varint,1,opt,name=user_found,json=userFound,proto3" json:"user_found,omitempty"`
	ClientPlatform       *ClientPlatform `protobuf:"bytes,2,opt,name=client_platform,json=clientPlatform,proto3" json:"client_platform,omitempty"`
	ClientType           *ClientType     `protobuf:"bytes,3,opt,name=client_type,json=clientType,proto3" json:"client_type,omitempty"`
	User                 *User           `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	Role                 *Role           `protobuf:"bytes,5,opt,name=role,proto3" json:"role,omitempty"`
	Token                *Token          `protobuf:"bytes,6,opt,name=token,proto3" json:"token,omitempty"`
	Permissions          []*Permission   `protobuf:"bytes,7,rep,name=permissions,proto3" json:"permissions,omitempty"`
	Sessions             []*Session      `protobuf:"bytes,8,rep,name=sessions,proto3" json:"sessions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff25770e79caedc0, []int{1}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetUserFound() bool {
	if m != nil {
		return m.UserFound
	}
	return false
}

func (m *LoginResponse) GetClientPlatform() *ClientPlatform {
	if m != nil {
		return m.ClientPlatform
	}
	return nil
}

func (m *LoginResponse) GetClientType() *ClientType {
	if m != nil {
		return m.ClientType
	}
	return nil
}

func (m *LoginResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *LoginResponse) GetRole() *Role {
	if m != nil {
		return m.Role
	}
	return nil
}

func (m *LoginResponse) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *LoginResponse) GetPermissions() []*Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *LoginResponse) GetSessions() []*Session {
	if m != nil {
		return m.Sessions
	}
	return nil
}

type LogoutRequest struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutRequest) Reset()         { *m = LogoutRequest{} }
func (m *LogoutRequest) String() string { return proto.CompactTextString(m) }
func (*LogoutRequest) ProtoMessage()    {}
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff25770e79caedc0, []int{2}
}

func (m *LogoutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutRequest.Unmarshal(m, b)
}
func (m *LogoutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutRequest.Marshal(b, m, deterministic)
}
func (m *LogoutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutRequest.Merge(m, src)
}
func (m *LogoutRequest) XXX_Size() int {
	return xxx_messageInfo_LogoutRequest.Size(m)
}
func (m *LogoutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutRequest proto.InternalMessageInfo

func (m *LogoutRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type RefreshTokenRequest struct {
	RefreshToken         string   `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RefreshTokenRequest) Reset()         { *m = RefreshTokenRequest{} }
func (m *RefreshTokenRequest) String() string { return proto.CompactTextString(m) }
func (*RefreshTokenRequest) ProtoMessage()    {}
func (*RefreshTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff25770e79caedc0, []int{3}
}

func (m *RefreshTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshTokenRequest.Unmarshal(m, b)
}
func (m *RefreshTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshTokenRequest.Marshal(b, m, deterministic)
}
func (m *RefreshTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshTokenRequest.Merge(m, src)
}
func (m *RefreshTokenRequest) XXX_Size() int {
	return xxx_messageInfo_RefreshTokenRequest.Size(m)
}
func (m *RefreshTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshTokenRequest proto.InternalMessageInfo

func (m *RefreshTokenRequest) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

type RefreshTokenResponse struct {
	Token                *Token   `protobuf:"bytes,6,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RefreshTokenResponse) Reset()         { *m = RefreshTokenResponse{} }
func (m *RefreshTokenResponse) String() string { return proto.CompactTextString(m) }
func (*RefreshTokenResponse) ProtoMessage()    {}
func (*RefreshTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff25770e79caedc0, []int{4}
}

func (m *RefreshTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshTokenResponse.Unmarshal(m, b)
}
func (m *RefreshTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshTokenResponse.Marshal(b, m, deterministic)
}
func (m *RefreshTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshTokenResponse.Merge(m, src)
}
func (m *RefreshTokenResponse) XXX_Size() int {
	return xxx_messageInfo_RefreshTokenResponse.Size(m)
}
func (m *RefreshTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshTokenResponse proto.InternalMessageInfo

func (m *RefreshTokenResponse) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type HasAccessRequest struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	ProjectId            string   `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ClientPlatformId     string   `protobuf:"bytes,3,opt,name=client_platform_id,json=clientPlatformId,proto3" json:"client_platform_id,omitempty"`
	Path                 string   `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Method               string   `protobuf:"bytes,5,opt,name=method,proto3" json:"method,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HasAccessRequest) Reset()         { *m = HasAccessRequest{} }
func (m *HasAccessRequest) String() string { return proto.CompactTextString(m) }
func (*HasAccessRequest) ProtoMessage()    {}
func (*HasAccessRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff25770e79caedc0, []int{5}
}

func (m *HasAccessRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HasAccessRequest.Unmarshal(m, b)
}
func (m *HasAccessRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HasAccessRequest.Marshal(b, m, deterministic)
}
func (m *HasAccessRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HasAccessRequest.Merge(m, src)
}
func (m *HasAccessRequest) XXX_Size() int {
	return xxx_messageInfo_HasAccessRequest.Size(m)
}
func (m *HasAccessRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HasAccessRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HasAccessRequest proto.InternalMessageInfo

func (m *HasAccessRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *HasAccessRequest) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *HasAccessRequest) GetClientPlatformId() string {
	if m != nil {
		return m.ClientPlatformId
	}
	return ""
}

func (m *HasAccessRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *HasAccessRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

type HasAccessResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProjectId            string   `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ClientPlatformId     string   `protobuf:"bytes,3,opt,name=client_platform_id,json=clientPlatformId,proto3" json:"client_platform_id,omitempty"`
	ClientTypeId         string   `protobuf:"bytes,4,opt,name=client_type_id,json=clientTypeId,proto3" json:"client_type_id,omitempty"`
	UserId               string   `protobuf:"bytes,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RoleId               string   `protobuf:"bytes,6,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	Ip                   string   `protobuf:"bytes,7,opt,name=ip,proto3" json:"ip,omitempty"`
	Data                 string   `protobuf:"bytes,8,opt,name=data,proto3" json:"data,omitempty"`
	ExpiresAt            string   `protobuf:"bytes,9,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	CreatedAt            string   `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HasAccessResponse) Reset()         { *m = HasAccessResponse{} }
func (m *HasAccessResponse) String() string { return proto.CompactTextString(m) }
func (*HasAccessResponse) ProtoMessage()    {}
func (*HasAccessResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff25770e79caedc0, []int{6}
}

func (m *HasAccessResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HasAccessResponse.Unmarshal(m, b)
}
func (m *HasAccessResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HasAccessResponse.Marshal(b, m, deterministic)
}
func (m *HasAccessResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HasAccessResponse.Merge(m, src)
}
func (m *HasAccessResponse) XXX_Size() int {
	return xxx_messageInfo_HasAccessResponse.Size(m)
}
func (m *HasAccessResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HasAccessResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HasAccessResponse proto.InternalMessageInfo

func (m *HasAccessResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *HasAccessResponse) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *HasAccessResponse) GetClientPlatformId() string {
	if m != nil {
		return m.ClientPlatformId
	}
	return ""
}

func (m *HasAccessResponse) GetClientTypeId() string {
	if m != nil {
		return m.ClientTypeId
	}
	return ""
}

func (m *HasAccessResponse) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *HasAccessResponse) GetRoleId() string {
	if m != nil {
		return m.RoleId
	}
	return ""
}

func (m *HasAccessResponse) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *HasAccessResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *HasAccessResponse) GetExpiresAt() string {
	if m != nil {
		return m.ExpiresAt
	}
	return ""
}

func (m *HasAccessResponse) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *HasAccessResponse) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type CreateSessionRequest struct {
	ProjectId            string   `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ClientPlatformId     string   `protobuf:"bytes,3,opt,name=client_platform_id,json=clientPlatformId,proto3" json:"client_platform_id,omitempty"`
	ClientTypeId         string   `protobuf:"bytes,4,opt,name=client_type_id,json=clientTypeId,proto3" json:"client_type_id,omitempty"`
	UserId               string   `protobuf:"bytes,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RoleId               string   `protobuf:"bytes,6,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	Ip                   string   `protobuf:"bytes,7,opt,name=ip,proto3" json:"ip,omitempty"`
	Data                 string   `protobuf:"bytes,8,opt,name=data,proto3" json:"data,omitempty"`
	ExpiresAt            string   `protobuf:"bytes,9,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSessionRequest) Reset()         { *m = CreateSessionRequest{} }
func (m *CreateSessionRequest) String() string { return proto.CompactTextString(m) }
func (*CreateSessionRequest) ProtoMessage()    {}
func (*CreateSessionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff25770e79caedc0, []int{7}
}

func (m *CreateSessionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSessionRequest.Unmarshal(m, b)
}
func (m *CreateSessionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSessionRequest.Marshal(b, m, deterministic)
}
func (m *CreateSessionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSessionRequest.Merge(m, src)
}
func (m *CreateSessionRequest) XXX_Size() int {
	return xxx_messageInfo_CreateSessionRequest.Size(m)
}
func (m *CreateSessionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSessionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSessionRequest proto.InternalMessageInfo

func (m *CreateSessionRequest) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *CreateSessionRequest) GetClientPlatformId() string {
	if m != nil {
		return m.ClientPlatformId
	}
	return ""
}

func (m *CreateSessionRequest) GetClientTypeId() string {
	if m != nil {
		return m.ClientTypeId
	}
	return ""
}

func (m *CreateSessionRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *CreateSessionRequest) GetRoleId() string {
	if m != nil {
		return m.RoleId
	}
	return ""
}

func (m *CreateSessionRequest) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *CreateSessionRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *CreateSessionRequest) GetExpiresAt() string {
	if m != nil {
		return m.ExpiresAt
	}
	return ""
}

type UpdateSessionRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProjectId            string   `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ClientPlatformId     string   `protobuf:"bytes,3,opt,name=client_platform_id,json=clientPlatformId,proto3" json:"client_platform_id,omitempty"`
	ClientTypeId         string   `protobuf:"bytes,4,opt,name=client_type_id,json=clientTypeId,proto3" json:"client_type_id,omitempty"`
	UserId               string   `protobuf:"bytes,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RoleId               string   `protobuf:"bytes,6,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	Ip                   string   `protobuf:"bytes,7,opt,name=ip,proto3" json:"ip,omitempty"`
	Data                 string   `protobuf:"bytes,8,opt,name=data,proto3" json:"data,omitempty"`
	ExpiresAt            string   `protobuf:"bytes,9,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateSessionRequest) Reset()         { *m = UpdateSessionRequest{} }
func (m *UpdateSessionRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateSessionRequest) ProtoMessage()    {}
func (*UpdateSessionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff25770e79caedc0, []int{8}
}

func (m *UpdateSessionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateSessionRequest.Unmarshal(m, b)
}
func (m *UpdateSessionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateSessionRequest.Marshal(b, m, deterministic)
}
func (m *UpdateSessionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateSessionRequest.Merge(m, src)
}
func (m *UpdateSessionRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateSessionRequest.Size(m)
}
func (m *UpdateSessionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateSessionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateSessionRequest proto.InternalMessageInfo

func (m *UpdateSessionRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateSessionRequest) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *UpdateSessionRequest) GetClientPlatformId() string {
	if m != nil {
		return m.ClientPlatformId
	}
	return ""
}

func (m *UpdateSessionRequest) GetClientTypeId() string {
	if m != nil {
		return m.ClientTypeId
	}
	return ""
}

func (m *UpdateSessionRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UpdateSessionRequest) GetRoleId() string {
	if m != nil {
		return m.RoleId
	}
	return ""
}

func (m *UpdateSessionRequest) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *UpdateSessionRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *UpdateSessionRequest) GetExpiresAt() string {
	if m != nil {
		return m.ExpiresAt
	}
	return ""
}

type SessionPrimaryKey struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionPrimaryKey) Reset()         { *m = SessionPrimaryKey{} }
func (m *SessionPrimaryKey) String() string { return proto.CompactTextString(m) }
func (*SessionPrimaryKey) ProtoMessage()    {}
func (*SessionPrimaryKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff25770e79caedc0, []int{9}
}

func (m *SessionPrimaryKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionPrimaryKey.Unmarshal(m, b)
}
func (m *SessionPrimaryKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionPrimaryKey.Marshal(b, m, deterministic)
}
func (m *SessionPrimaryKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionPrimaryKey.Merge(m, src)
}
func (m *SessionPrimaryKey) XXX_Size() int {
	return xxx_messageInfo_SessionPrimaryKey.Size(m)
}
func (m *SessionPrimaryKey) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionPrimaryKey.DiscardUnknown(m)
}

var xxx_messageInfo_SessionPrimaryKey proto.InternalMessageInfo

func (m *SessionPrimaryKey) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetSessionListRequest struct {
	Limit                int32    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int32    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Search               string   `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSessionListRequest) Reset()         { *m = GetSessionListRequest{} }
func (m *GetSessionListRequest) String() string { return proto.CompactTextString(m) }
func (*GetSessionListRequest) ProtoMessage()    {}
func (*GetSessionListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff25770e79caedc0, []int{10}
}

func (m *GetSessionListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSessionListRequest.Unmarshal(m, b)
}
func (m *GetSessionListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSessionListRequest.Marshal(b, m, deterministic)
}
func (m *GetSessionListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSessionListRequest.Merge(m, src)
}
func (m *GetSessionListRequest) XXX_Size() int {
	return xxx_messageInfo_GetSessionListRequest.Size(m)
}
func (m *GetSessionListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSessionListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSessionListRequest proto.InternalMessageInfo

func (m *GetSessionListRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetSessionListRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *GetSessionListRequest) GetSearch() string {
	if m != nil {
		return m.Search
	}
	return ""
}

type GetSessionListResponse struct {
	Count                int32      `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Sessions             []*Session `protobuf:"bytes,2,rep,name=sessions,proto3" json:"sessions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetSessionListResponse) Reset()         { *m = GetSessionListResponse{} }
func (m *GetSessionListResponse) String() string { return proto.CompactTextString(m) }
func (*GetSessionListResponse) ProtoMessage()    {}
func (*GetSessionListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff25770e79caedc0, []int{11}
}

func (m *GetSessionListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSessionListResponse.Unmarshal(m, b)
}
func (m *GetSessionListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSessionListResponse.Marshal(b, m, deterministic)
}
func (m *GetSessionListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSessionListResponse.Merge(m, src)
}
func (m *GetSessionListResponse) XXX_Size() int {
	return xxx_messageInfo_GetSessionListResponse.Size(m)
}
func (m *GetSessionListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSessionListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetSessionListResponse proto.InternalMessageInfo

func (m *GetSessionListResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *GetSessionListResponse) GetSessions() []*Session {
	if m != nil {
		return m.Sessions
	}
	return nil
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "auth_service.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "auth_service.LoginResponse")
	proto.RegisterType((*LogoutRequest)(nil), "auth_service.LogoutRequest")
	proto.RegisterType((*RefreshTokenRequest)(nil), "auth_service.RefreshTokenRequest")
	proto.RegisterType((*RefreshTokenResponse)(nil), "auth_service.RefreshTokenResponse")
	proto.RegisterType((*HasAccessRequest)(nil), "auth_service.HasAccessRequest")
	proto.RegisterType((*HasAccessResponse)(nil), "auth_service.HasAccessResponse")
	proto.RegisterType((*CreateSessionRequest)(nil), "auth_service.CreateSessionRequest")
	proto.RegisterType((*UpdateSessionRequest)(nil), "auth_service.UpdateSessionRequest")
	proto.RegisterType((*SessionPrimaryKey)(nil), "auth_service.SessionPrimaryKey")
	proto.RegisterType((*GetSessionListRequest)(nil), "auth_service.GetSessionListRequest")
	proto.RegisterType((*GetSessionListResponse)(nil), "auth_service.GetSessionListResponse")
}

func init() { proto.RegisterFile("session_service.proto", fileDescriptor_ff25770e79caedc0) }

var fileDescriptor_ff25770e79caedc0 = []byte{
	// 788 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x56, 0xdd, 0x4e, 0xeb, 0x46,
	0x10, 0x26, 0x09, 0x09, 0xc9, 0x24, 0xa4, 0xb0, 0x24, 0x60, 0x99, 0xfe, 0x80, 0xa9, 0xaa, 0x56,
	0xaa, 0x82, 0x4a, 0xaf, 0xca, 0x4d, 0x15, 0x10, 0xb4, 0x51, 0x51, 0x85, 0x16, 0x50, 0xa5, 0x4a,
	0x55, 0x64, 0xec, 0x49, 0xe2, 0x36, 0xf1, 0xba, 0xbb, 0xeb, 0xb6, 0x79, 0x89, 0x3e, 0x40, 0x1f,
	0xa2, 0x8f, 0x70, 0xee, 0xcf, 0x4b, 0x1d, 0x1d, 0xed, 0x8f, 0x8d, 0x1d, 0x38, 0x47, 0x5c, 0x9c,
	0x3b, 0xee, 0x3c, 0xf3, 0x7d, 0x33, 0xfa, 0x3c, 0xfb, 0x79, 0xd6, 0xd0, 0x17, 0x28, 0x44, 0xc4,
	0xe2, 0xb1, 0x40, 0xfe, 0x57, 0x14, 0xe0, 0x20, 0xe1, 0x4c, 0x32, 0xd2, 0xf1, 0x53, 0x39, 0xcb,
	0x72, 0x2e, 0xa8, 0xc8, 0x20, 0xee, 0xfe, 0x94, 0xb1, 0xe9, 0x1c, 0x8f, 0x75, 0x74, 0x9f, 0x4e,
	0x8e, 0x71, 0x91, 0xc8, 0xa5, 0x01, 0xbd, 0x4b, 0xe8, 0x5c, 0xb1, 0x69, 0x14, 0x53, 0xfc, 0x33,
	0x45, 0x21, 0x89, 0x0b, 0xcd, 0x54, 0x20, 0x8f, 0xfd, 0x05, 0x3a, 0x95, 0x83, 0xca, 0x97, 0x2d,
	0x9a, 0xc7, 0x0a, 0x4b, 0x7c, 0x21, 0xfe, 0x66, 0x3c, 0x74, 0xaa, 0x06, 0xcb, 0x62, 0xef, 0xbf,
	0x1a, 0x6c, 0xda, 0x46, 0x22, 0x61, 0xb1, 0x40, 0xf2, 0x09, 0x80, 0xaa, 0x1c, 0x4f, 0x58, 0x1a,
	0x87, 0xba, 0x57, 0x93, 0xb6, 0x54, 0xe6, 0x52, 0x25, 0xc8, 0x05, 0x7c, 0x14, 0xcc, 0x23, 0x8c,
	0xe5, 0x38, 0x99, 0xfb, 0x72, 0xc2, 0xf8, 0x42, 0xf7, 0x6c, 0x9f, 0x7c, 0x3c, 0x28, 0xbe, 0xc9,
	0xe0, 0x5c, 0x93, 0xae, 0x2d, 0x87, 0x76, 0x83, 0x52, 0x4c, 0xbe, 0x83, 0xb6, 0x6d, 0x23, 0x97,
	0x09, 0x3a, 0x35, 0xdd, 0xc2, 0x79, 0xaa, 0xc5, 0xed, 0x32, 0x41, 0x0a, 0x41, 0xfe, 0x4c, 0xbe,
	0x80, 0x75, 0x25, 0xc7, 0x59, 0xd7, 0x35, 0xa4, 0x5c, 0x73, 0x27, 0x90, 0x53, 0x8d, 0x2b, 0x1e,
	0x67, 0x73, 0x74, 0xea, 0x4f, 0xf1, 0x28, 0x9b, 0x23, 0xd5, 0x38, 0xf9, 0x0a, 0xea, 0x92, 0xfd,
	0x81, 0xb1, 0xd3, 0xd0, 0xc4, 0x9d, 0x32, 0xf1, 0x56, 0x41, 0xd4, 0x30, 0xc8, 0x29, 0xb4, 0x13,
	0xe4, 0x8b, 0x48, 0x1f, 0xa4, 0x70, 0x36, 0x0e, 0x6a, 0x8f, 0x55, 0x5f, 0xe7, 0x04, 0x5a, 0x24,
	0x93, 0x6f, 0xa0, 0x69, 0x1d, 0x20, 0x9c, 0xa6, 0x2e, 0xec, 0x97, 0x0b, 0x6f, 0x0c, 0x4a, 0x73,
	0x9a, 0x77, 0xa2, 0xcf, 0x86, 0xa5, 0x32, 0x3b, 0xe5, 0x43, 0xe8, 0xf8, 0x41, 0x80, 0x42, 0x8c,
	0x8d, 0x62, 0x73, 0xd2, 0x6d, 0x93, 0xd3, 0x4a, 0xbd, 0x53, 0xd8, 0xa1, 0x38, 0xe1, 0x28, 0x66,
	0x46, 0xb9, 0xad, 0x3c, 0x82, 0x4d, 0x6e, 0xd2, 0xb6, 0xd4, 0x18, 0xa1, 0xc3, 0x0b, 0x5c, 0x6f,
	0x08, 0xbd, 0x72, 0xad, 0xb5, 0xc4, 0xf3, 0x27, 0xe4, 0xfd, 0x5f, 0x81, 0xad, 0x1f, 0x7d, 0x31,
	0xd4, 0x8a, 0x9e, 0x2f, 0x5b, 0xb9, 0x2e, 0xe1, 0xec, 0x77, 0x0c, 0xe4, 0x38, 0xca, 0x5c, 0xda,
	0xb2, 0x99, 0x51, 0x48, 0xbe, 0x06, 0xb2, 0xe2, 0x3a, 0x45, 0xab, 0x69, 0xda, 0x56, 0xd9, 0x5a,
	0xa3, 0x90, 0x10, 0x58, 0x4f, 0x7c, 0x39, 0xd3, 0x0e, 0x69, 0x51, 0xfd, 0x4c, 0x76, 0xa1, 0xb1,
	0x40, 0x39, 0x63, 0xa1, 0xf6, 0x43, 0x8b, 0xda, 0xc8, 0x7b, 0x5d, 0x85, 0xed, 0x82, 0x60, 0xfb,
	0xc6, 0x5d, 0xa8, 0x46, 0xa1, 0xd5, 0x59, 0x8d, 0xc2, 0x0f, 0x2b, 0xef, 0x73, 0xe8, 0x16, 0xbc,
	0xaf, 0x98, 0x46, 0x68, 0xe7, 0xc1, 0xe4, 0xa3, 0x90, 0xec, 0xc1, 0x86, 0xfe, 0x0e, 0xa3, 0x5c,
	0xb1, 0x0a, 0x0d, 0xa0, 0x7c, 0xab, 0x80, 0x86, 0x01, 0x54, 0x38, 0x0a, 0xb5, 0xe8, 0xc4, 0xd9,
	0xb0, 0xa2, 0x13, 0x35, 0x86, 0xd0, 0x97, 0xbe, 0xd3, 0x34, 0x63, 0x50, 0xcf, 0xea, 0x45, 0xf0,
	0x9f, 0x24, 0xe2, 0x28, 0xc6, 0xbe, 0x74, 0x5a, 0xe6, 0x45, 0x6c, 0x66, 0x28, 0x15, 0x1c, 0x70,
	0xf4, 0x25, 0x86, 0x0a, 0x06, 0x03, 0xdb, 0x8c, 0x81, 0xd3, 0x24, 0xcc, 0xe0, 0xb6, 0x81, 0x6d,
	0x66, 0x28, 0xbd, 0x37, 0x15, 0xe8, 0x9d, 0x6b, 0x72, 0xe6, 0x65, 0x6b, 0x80, 0x17, 0x32, 0x3e,
	0xef, 0xdf, 0x2a, 0xf4, 0xee, 0xf4, 0x38, 0x56, 0x06, 0xf0, 0x42, 0xfd, 0xe4, 0x1d, 0xc1, 0xb6,
	0x9d, 0xc4, 0x35, 0x8f, 0x16, 0x3e, 0x5f, 0xfe, 0x84, 0xcb, 0xd5, 0x61, 0x78, 0xbf, 0x41, 0xff,
	0x07, 0x94, 0x96, 0x77, 0x15, 0x89, 0x7c, 0xdd, 0xf5, 0xa0, 0x3e, 0x8f, 0x16, 0x91, 0xd4, 0xdc,
	0x3a, 0x35, 0x81, 0xfa, 0x92, 0xd9, 0x64, 0x22, 0x50, 0xea, 0xb9, 0xd5, 0xa9, 0x8d, 0x54, 0x5e,
	0xa0, 0xcf, 0x83, 0x99, 0x1d, 0x94, 0x8d, 0x3c, 0x1f, 0x76, 0x57, 0xdb, 0xdb, 0xaf, 0xbc, 0x07,
	0xf5, 0x80, 0xa5, 0x71, 0xde, 0x5f, 0x07, 0xa5, 0x45, 0x5d, 0x7d, 0xd6, 0xa2, 0x3e, 0x79, 0x55,
	0x85, 0xae, 0xcd, 0xde, 0x18, 0x12, 0x39, 0x83, 0xba, 0xbe, 0x57, 0x89, 0x5b, 0x2e, 0x2e, 0xde,
	0xda, 0xee, 0xfe, 0x93, 0x98, 0x51, 0xe7, 0xad, 0x91, 0xef, 0xa1, 0x61, 0xf6, 0x3f, 0x79, 0x4c,
	0x7c, 0xb8, 0x15, 0xdc, 0xdd, 0x81, 0xf9, 0x53, 0x18, 0x64, 0x7f, 0x0a, 0x83, 0x0b, 0xf5, 0xa7,
	0xe0, 0xad, 0x91, 0x5f, 0xa0, 0x53, 0x5c, 0xe8, 0xe4, 0x70, 0xe5, 0x12, 0x7c, 0x7c, 0x51, 0xb8,
	0xde, 0xfb, 0x28, 0xb9, 0xb2, 0x9f, 0xa1, 0x95, 0x2f, 0x4d, 0xf2, 0x69, 0xb9, 0x64, 0x75, 0xfd,
	0xbb, 0x9f, 0xbd, 0x13, 0xcf, 0xfa, 0x9d, 0xed, 0xfd, 0xda, 0x9f, 0x62, 0xac, 0xf5, 0x1f, 0x17,
	0xc9, 0xf7, 0x0d, 0x9d, 0xfb, 0xf6, 0x6d, 0x00, 0x00, 0x00, 0xff, 0xff, 0x95, 0x32, 0x43, 0x6c,
	0x3e, 0x09, 0x00, 0x00,
}
