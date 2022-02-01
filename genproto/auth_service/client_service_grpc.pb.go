// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: client_service.proto

package auth_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ClientServiceClient is the client API for ClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientServiceClient interface {
	CreateClientPlatform(ctx context.Context, in *CreateClientPlatformRequest, opts ...grpc.CallOption) (*ClientPlatform, error)
	GetClientPlatformByID(ctx context.Context, in *ClientPlatformPrimaryKey, opts ...grpc.CallOption) (*ClientPlatform, error)
	GetClientPlatformList(ctx context.Context, in *GetClientPlatformListRequest, opts ...grpc.CallOption) (*GetClientPlatformListResponse, error)
	UpdateClientPlatform(ctx context.Context, in *UpdateClientPlatformRequest, opts ...grpc.CallOption) (*ClientPlatform, error)
	DeleteClientPlatform(ctx context.Context, in *ClientPlatformPrimaryKey, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateClientType(ctx context.Context, in *CreateClientTypeRequest, opts ...grpc.CallOption) (*ClientType, error)
	GetClientTypeByID(ctx context.Context, in *ClientTypePrimaryKey, opts ...grpc.CallOption) (*CompleteClientType, error)
	GetClientTypeList(ctx context.Context, in *GetClientTypeListRequest, opts ...grpc.CallOption) (*GetClientTypeListResponse, error)
	UpdateClientType(ctx context.Context, in *UpdateClientTypeRequest, opts ...grpc.CallOption) (*ClientType, error)
	DeleteClientType(ctx context.Context, in *ClientTypePrimaryKey, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AddClient(ctx context.Context, in *AddClientRequest, opts ...grpc.CallOption) (*Client, error)
	UpdateClient(ctx context.Context, in *UpdateClientRequest, opts ...grpc.CallOption) (*Client, error)
	RemoveClient(ctx context.Context, in *ClientPrimaryKey, opts ...grpc.CallOption) (*Client, error)
	GetClientList(ctx context.Context, in *GetClientListRequest, opts ...grpc.CallOption) (*GetClientListResponse, error)
	GetClientMatrix(ctx context.Context, in *GetClientMatrixRequest, opts ...grpc.CallOption) (*GetClientMatrixResponse, error)
	AddRelation(ctx context.Context, in *AddRelationRequest, opts ...grpc.CallOption) (*Relation, error)
	UpdateRelation(ctx context.Context, in *UpdateRelationRequest, opts ...grpc.CallOption) (*Relation, error)
	RemoveRelation(ctx context.Context, in *RelationPrimaryKey, opts ...grpc.CallOption) (*Relation, error)
	AddUserInfoField(ctx context.Context, in *AddUserInfoFieldRequest, opts ...grpc.CallOption) (*UserInfoField, error)
	UpdateUserInfoField(ctx context.Context, in *UpdateUserInfoFieldRequest, opts ...grpc.CallOption) (*UserInfoField, error)
	RemoveUserInfoField(ctx context.Context, in *UserInfoFieldPrimaryKey, opts ...grpc.CallOption) (*UserInfoField, error)
}

type clientServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientServiceClient(cc grpc.ClientConnInterface) ClientServiceClient {
	return &clientServiceClient{cc}
}

func (c *clientServiceClient) CreateClientPlatform(ctx context.Context, in *CreateClientPlatformRequest, opts ...grpc.CallOption) (*ClientPlatform, error) {
	out := new(ClientPlatform)
	err := c.cc.Invoke(ctx, "/auth_service.ClientService/CreateClientPlatform", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) GetClientPlatformByID(ctx context.Context, in *ClientPlatformPrimaryKey, opts ...grpc.CallOption) (*ClientPlatform, error) {
	out := new(ClientPlatform)
	err := c.cc.Invoke(ctx, "/auth_service.ClientService/GetClientPlatformByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) GetClientPlatformList(ctx context.Context, in *GetClientPlatformListRequest, opts ...grpc.CallOption) (*GetClientPlatformListResponse, error) {
	out := new(GetClientPlatformListResponse)
	err := c.cc.Invoke(ctx, "/auth_service.ClientService/GetClientPlatformList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) UpdateClientPlatform(ctx context.Context, in *UpdateClientPlatformRequest, opts ...grpc.CallOption) (*ClientPlatform, error) {
	out := new(ClientPlatform)
	err := c.cc.Invoke(ctx, "/auth_service.ClientService/UpdateClientPlatform", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) DeleteClientPlatform(ctx context.Context, in *ClientPlatformPrimaryKey, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/auth_service.ClientService/DeleteClientPlatform", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) CreateClientType(ctx context.Context, in *CreateClientTypeRequest, opts ...grpc.CallOption) (*ClientType, error) {
	out := new(ClientType)
	err := c.cc.Invoke(ctx, "/auth_service.ClientService/CreateClientType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) GetClientTypeByID(ctx context.Context, in *ClientTypePrimaryKey, opts ...grpc.CallOption) (*CompleteClientType, error) {
	out := new(CompleteClientType)
	err := c.cc.Invoke(ctx, "/auth_service.ClientService/GetClientTypeByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) GetClientTypeList(ctx context.Context, in *GetClientTypeListRequest, opts ...grpc.CallOption) (*GetClientTypeListResponse, error) {
	out := new(GetClientTypeListResponse)
	err := c.cc.Invoke(ctx, "/auth_service.ClientService/GetClientTypeList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) UpdateClientType(ctx context.Context, in *UpdateClientTypeRequest, opts ...grpc.CallOption) (*ClientType, error) {
	out := new(ClientType)
	err := c.cc.Invoke(ctx, "/auth_service.ClientService/UpdateClientType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) DeleteClientType(ctx context.Context, in *ClientTypePrimaryKey, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/auth_service.ClientService/DeleteClientType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) AddClient(ctx context.Context, in *AddClientRequest, opts ...grpc.CallOption) (*Client, error) {
	out := new(Client)
	err := c.cc.Invoke(ctx, "/auth_service.ClientService/AddClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) UpdateClient(ctx context.Context, in *UpdateClientRequest, opts ...grpc.CallOption) (*Client, error) {
	out := new(Client)
	err := c.cc.Invoke(ctx, "/auth_service.ClientService/UpdateClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) RemoveClient(ctx context.Context, in *ClientPrimaryKey, opts ...grpc.CallOption) (*Client, error) {
	out := new(Client)
	err := c.cc.Invoke(ctx, "/auth_service.ClientService/RemoveClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) GetClientList(ctx context.Context, in *GetClientListRequest, opts ...grpc.CallOption) (*GetClientListResponse, error) {
	out := new(GetClientListResponse)
	err := c.cc.Invoke(ctx, "/auth_service.ClientService/GetClientList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) GetClientMatrix(ctx context.Context, in *GetClientMatrixRequest, opts ...grpc.CallOption) (*GetClientMatrixResponse, error) {
	out := new(GetClientMatrixResponse)
	err := c.cc.Invoke(ctx, "/auth_service.ClientService/GetClientMatrix", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) AddRelation(ctx context.Context, in *AddRelationRequest, opts ...grpc.CallOption) (*Relation, error) {
	out := new(Relation)
	err := c.cc.Invoke(ctx, "/auth_service.ClientService/AddRelation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) UpdateRelation(ctx context.Context, in *UpdateRelationRequest, opts ...grpc.CallOption) (*Relation, error) {
	out := new(Relation)
	err := c.cc.Invoke(ctx, "/auth_service.ClientService/UpdateRelation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) RemoveRelation(ctx context.Context, in *RelationPrimaryKey, opts ...grpc.CallOption) (*Relation, error) {
	out := new(Relation)
	err := c.cc.Invoke(ctx, "/auth_service.ClientService/RemoveRelation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) AddUserInfoField(ctx context.Context, in *AddUserInfoFieldRequest, opts ...grpc.CallOption) (*UserInfoField, error) {
	out := new(UserInfoField)
	err := c.cc.Invoke(ctx, "/auth_service.ClientService/AddUserInfoField", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) UpdateUserInfoField(ctx context.Context, in *UpdateUserInfoFieldRequest, opts ...grpc.CallOption) (*UserInfoField, error) {
	out := new(UserInfoField)
	err := c.cc.Invoke(ctx, "/auth_service.ClientService/UpdateUserInfoField", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) RemoveUserInfoField(ctx context.Context, in *UserInfoFieldPrimaryKey, opts ...grpc.CallOption) (*UserInfoField, error) {
	out := new(UserInfoField)
	err := c.cc.Invoke(ctx, "/auth_service.ClientService/RemoveUserInfoField", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientServiceServer is the server API for ClientService service.
// All implementations must embed UnimplementedClientServiceServer
// for forward compatibility
type ClientServiceServer interface {
	CreateClientPlatform(context.Context, *CreateClientPlatformRequest) (*ClientPlatform, error)
	GetClientPlatformByID(context.Context, *ClientPlatformPrimaryKey) (*ClientPlatform, error)
	GetClientPlatformList(context.Context, *GetClientPlatformListRequest) (*GetClientPlatformListResponse, error)
	UpdateClientPlatform(context.Context, *UpdateClientPlatformRequest) (*ClientPlatform, error)
	DeleteClientPlatform(context.Context, *ClientPlatformPrimaryKey) (*emptypb.Empty, error)
	CreateClientType(context.Context, *CreateClientTypeRequest) (*ClientType, error)
	GetClientTypeByID(context.Context, *ClientTypePrimaryKey) (*CompleteClientType, error)
	GetClientTypeList(context.Context, *GetClientTypeListRequest) (*GetClientTypeListResponse, error)
	UpdateClientType(context.Context, *UpdateClientTypeRequest) (*ClientType, error)
	DeleteClientType(context.Context, *ClientTypePrimaryKey) (*emptypb.Empty, error)
	AddClient(context.Context, *AddClientRequest) (*Client, error)
	UpdateClient(context.Context, *UpdateClientRequest) (*Client, error)
	RemoveClient(context.Context, *ClientPrimaryKey) (*Client, error)
	GetClientList(context.Context, *GetClientListRequest) (*GetClientListResponse, error)
	GetClientMatrix(context.Context, *GetClientMatrixRequest) (*GetClientMatrixResponse, error)
	AddRelation(context.Context, *AddRelationRequest) (*Relation, error)
	UpdateRelation(context.Context, *UpdateRelationRequest) (*Relation, error)
	RemoveRelation(context.Context, *RelationPrimaryKey) (*Relation, error)
	AddUserInfoField(context.Context, *AddUserInfoFieldRequest) (*UserInfoField, error)
	UpdateUserInfoField(context.Context, *UpdateUserInfoFieldRequest) (*UserInfoField, error)
	RemoveUserInfoField(context.Context, *UserInfoFieldPrimaryKey) (*UserInfoField, error)
	mustEmbedUnimplementedClientServiceServer()
}

// UnimplementedClientServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClientServiceServer struct {
}

func (UnimplementedClientServiceServer) CreateClientPlatform(context.Context, *CreateClientPlatformRequest) (*ClientPlatform, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClientPlatform not implemented")
}
func (UnimplementedClientServiceServer) GetClientPlatformByID(context.Context, *ClientPlatformPrimaryKey) (*ClientPlatform, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientPlatformByID not implemented")
}
func (UnimplementedClientServiceServer) GetClientPlatformList(context.Context, *GetClientPlatformListRequest) (*GetClientPlatformListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientPlatformList not implemented")
}
func (UnimplementedClientServiceServer) UpdateClientPlatform(context.Context, *UpdateClientPlatformRequest) (*ClientPlatform, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClientPlatform not implemented")
}
func (UnimplementedClientServiceServer) DeleteClientPlatform(context.Context, *ClientPlatformPrimaryKey) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteClientPlatform not implemented")
}
func (UnimplementedClientServiceServer) CreateClientType(context.Context, *CreateClientTypeRequest) (*ClientType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClientType not implemented")
}
func (UnimplementedClientServiceServer) GetClientTypeByID(context.Context, *ClientTypePrimaryKey) (*CompleteClientType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientTypeByID not implemented")
}
func (UnimplementedClientServiceServer) GetClientTypeList(context.Context, *GetClientTypeListRequest) (*GetClientTypeListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientTypeList not implemented")
}
func (UnimplementedClientServiceServer) UpdateClientType(context.Context, *UpdateClientTypeRequest) (*ClientType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClientType not implemented")
}
func (UnimplementedClientServiceServer) DeleteClientType(context.Context, *ClientTypePrimaryKey) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteClientType not implemented")
}
func (UnimplementedClientServiceServer) AddClient(context.Context, *AddClientRequest) (*Client, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddClient not implemented")
}
func (UnimplementedClientServiceServer) UpdateClient(context.Context, *UpdateClientRequest) (*Client, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClient not implemented")
}
func (UnimplementedClientServiceServer) RemoveClient(context.Context, *ClientPrimaryKey) (*Client, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveClient not implemented")
}
func (UnimplementedClientServiceServer) GetClientList(context.Context, *GetClientListRequest) (*GetClientListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientList not implemented")
}
func (UnimplementedClientServiceServer) GetClientMatrix(context.Context, *GetClientMatrixRequest) (*GetClientMatrixResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientMatrix not implemented")
}
func (UnimplementedClientServiceServer) AddRelation(context.Context, *AddRelationRequest) (*Relation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRelation not implemented")
}
func (UnimplementedClientServiceServer) UpdateRelation(context.Context, *UpdateRelationRequest) (*Relation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRelation not implemented")
}
func (UnimplementedClientServiceServer) RemoveRelation(context.Context, *RelationPrimaryKey) (*Relation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveRelation not implemented")
}
func (UnimplementedClientServiceServer) AddUserInfoField(context.Context, *AddUserInfoFieldRequest) (*UserInfoField, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserInfoField not implemented")
}
func (UnimplementedClientServiceServer) UpdateUserInfoField(context.Context, *UpdateUserInfoFieldRequest) (*UserInfoField, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserInfoField not implemented")
}
func (UnimplementedClientServiceServer) RemoveUserInfoField(context.Context, *UserInfoFieldPrimaryKey) (*UserInfoField, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUserInfoField not implemented")
}
func (UnimplementedClientServiceServer) mustEmbedUnimplementedClientServiceServer() {}

// UnsafeClientServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientServiceServer will
// result in compilation errors.
type UnsafeClientServiceServer interface {
	mustEmbedUnimplementedClientServiceServer()
}

func RegisterClientServiceServer(s grpc.ServiceRegistrar, srv ClientServiceServer) {
	s.RegisterService(&ClientService_ServiceDesc, srv)
}

func _ClientService_CreateClientPlatform_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClientPlatformRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).CreateClientPlatform(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.ClientService/CreateClientPlatform",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).CreateClientPlatform(ctx, req.(*CreateClientPlatformRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_GetClientPlatformByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientPlatformPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).GetClientPlatformByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.ClientService/GetClientPlatformByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).GetClientPlatformByID(ctx, req.(*ClientPlatformPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_GetClientPlatformList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClientPlatformListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).GetClientPlatformList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.ClientService/GetClientPlatformList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).GetClientPlatformList(ctx, req.(*GetClientPlatformListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_UpdateClientPlatform_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClientPlatformRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).UpdateClientPlatform(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.ClientService/UpdateClientPlatform",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).UpdateClientPlatform(ctx, req.(*UpdateClientPlatformRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_DeleteClientPlatform_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientPlatformPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).DeleteClientPlatform(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.ClientService/DeleteClientPlatform",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).DeleteClientPlatform(ctx, req.(*ClientPlatformPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_CreateClientType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClientTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).CreateClientType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.ClientService/CreateClientType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).CreateClientType(ctx, req.(*CreateClientTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_GetClientTypeByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientTypePrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).GetClientTypeByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.ClientService/GetClientTypeByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).GetClientTypeByID(ctx, req.(*ClientTypePrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_GetClientTypeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClientTypeListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).GetClientTypeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.ClientService/GetClientTypeList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).GetClientTypeList(ctx, req.(*GetClientTypeListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_UpdateClientType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClientTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).UpdateClientType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.ClientService/UpdateClientType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).UpdateClientType(ctx, req.(*UpdateClientTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_DeleteClientType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientTypePrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).DeleteClientType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.ClientService/DeleteClientType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).DeleteClientType(ctx, req.(*ClientTypePrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_AddClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).AddClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.ClientService/AddClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).AddClient(ctx, req.(*AddClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_UpdateClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).UpdateClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.ClientService/UpdateClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).UpdateClient(ctx, req.(*UpdateClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_RemoveClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).RemoveClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.ClientService/RemoveClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).RemoveClient(ctx, req.(*ClientPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_GetClientList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClientListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).GetClientList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.ClientService/GetClientList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).GetClientList(ctx, req.(*GetClientListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_GetClientMatrix_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClientMatrixRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).GetClientMatrix(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.ClientService/GetClientMatrix",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).GetClientMatrix(ctx, req.(*GetClientMatrixRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_AddRelation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRelationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).AddRelation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.ClientService/AddRelation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).AddRelation(ctx, req.(*AddRelationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_UpdateRelation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRelationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).UpdateRelation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.ClientService/UpdateRelation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).UpdateRelation(ctx, req.(*UpdateRelationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_RemoveRelation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RelationPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).RemoveRelation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.ClientService/RemoveRelation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).RemoveRelation(ctx, req.(*RelationPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_AddUserInfoField_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserInfoFieldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).AddUserInfoField(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.ClientService/AddUserInfoField",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).AddUserInfoField(ctx, req.(*AddUserInfoFieldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_UpdateUserInfoField_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserInfoFieldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).UpdateUserInfoField(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.ClientService/UpdateUserInfoField",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).UpdateUserInfoField(ctx, req.(*UpdateUserInfoFieldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_RemoveUserInfoField_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoFieldPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).RemoveUserInfoField(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.ClientService/RemoveUserInfoField",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).RemoveUserInfoField(ctx, req.(*UserInfoFieldPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientService_ServiceDesc is the grpc.ServiceDesc for ClientService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth_service.ClientService",
	HandlerType: (*ClientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateClientPlatform",
			Handler:    _ClientService_CreateClientPlatform_Handler,
		},
		{
			MethodName: "GetClientPlatformByID",
			Handler:    _ClientService_GetClientPlatformByID_Handler,
		},
		{
			MethodName: "GetClientPlatformList",
			Handler:    _ClientService_GetClientPlatformList_Handler,
		},
		{
			MethodName: "UpdateClientPlatform",
			Handler:    _ClientService_UpdateClientPlatform_Handler,
		},
		{
			MethodName: "DeleteClientPlatform",
			Handler:    _ClientService_DeleteClientPlatform_Handler,
		},
		{
			MethodName: "CreateClientType",
			Handler:    _ClientService_CreateClientType_Handler,
		},
		{
			MethodName: "GetClientTypeByID",
			Handler:    _ClientService_GetClientTypeByID_Handler,
		},
		{
			MethodName: "GetClientTypeList",
			Handler:    _ClientService_GetClientTypeList_Handler,
		},
		{
			MethodName: "UpdateClientType",
			Handler:    _ClientService_UpdateClientType_Handler,
		},
		{
			MethodName: "DeleteClientType",
			Handler:    _ClientService_DeleteClientType_Handler,
		},
		{
			MethodName: "AddClient",
			Handler:    _ClientService_AddClient_Handler,
		},
		{
			MethodName: "UpdateClient",
			Handler:    _ClientService_UpdateClient_Handler,
		},
		{
			MethodName: "RemoveClient",
			Handler:    _ClientService_RemoveClient_Handler,
		},
		{
			MethodName: "GetClientList",
			Handler:    _ClientService_GetClientList_Handler,
		},
		{
			MethodName: "GetClientMatrix",
			Handler:    _ClientService_GetClientMatrix_Handler,
		},
		{
			MethodName: "AddRelation",
			Handler:    _ClientService_AddRelation_Handler,
		},
		{
			MethodName: "UpdateRelation",
			Handler:    _ClientService_UpdateRelation_Handler,
		},
		{
			MethodName: "RemoveRelation",
			Handler:    _ClientService_RemoveRelation_Handler,
		},
		{
			MethodName: "AddUserInfoField",
			Handler:    _ClientService_AddUserInfoField_Handler,
		},
		{
			MethodName: "UpdateUserInfoField",
			Handler:    _ClientService_UpdateUserInfoField_Handler,
		},
		{
			MethodName: "RemoveUserInfoField",
			Handler:    _ClientService_RemoveUserInfoField_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "client_service.proto",
}
