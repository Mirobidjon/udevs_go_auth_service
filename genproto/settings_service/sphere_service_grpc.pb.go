// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.11.2
// source: sphere_service.proto

package settings_service

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SphereServiceClient is the client API for SphereService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SphereServiceClient interface {
	Create(ctx context.Context, in *CreateSphereRequest, opts ...grpc.CallOption) (*Sphere, error)
	GetByID(ctx context.Context, in *SpherePrimaryKey, opts ...grpc.CallOption) (*Sphere, error)
	GetList(ctx context.Context, in *GetSphereListRequest, opts ...grpc.CallOption) (*GetSphereListResponse, error)
	Update(ctx context.Context, in *UpdateSphereRequest, opts ...grpc.CallOption) (*Sphere, error)
	Delete(ctx context.Context, in *SpherePrimaryKey, opts ...grpc.CallOption) (*empty.Empty, error)
}

type sphereServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSphereServiceClient(cc grpc.ClientConnInterface) SphereServiceClient {
	return &sphereServiceClient{cc}
}

func (c *sphereServiceClient) Create(ctx context.Context, in *CreateSphereRequest, opts ...grpc.CallOption) (*Sphere, error) {
	out := new(Sphere)
	err := c.cc.Invoke(ctx, "/settings_service.SphereService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sphereServiceClient) GetByID(ctx context.Context, in *SpherePrimaryKey, opts ...grpc.CallOption) (*Sphere, error) {
	out := new(Sphere)
	err := c.cc.Invoke(ctx, "/settings_service.SphereService/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sphereServiceClient) GetList(ctx context.Context, in *GetSphereListRequest, opts ...grpc.CallOption) (*GetSphereListResponse, error) {
	out := new(GetSphereListResponse)
	err := c.cc.Invoke(ctx, "/settings_service.SphereService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sphereServiceClient) Update(ctx context.Context, in *UpdateSphereRequest, opts ...grpc.CallOption) (*Sphere, error) {
	out := new(Sphere)
	err := c.cc.Invoke(ctx, "/settings_service.SphereService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sphereServiceClient) Delete(ctx context.Context, in *SpherePrimaryKey, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/settings_service.SphereService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SphereServiceServer is the server API for SphereService service.
// All implementations must embed UnimplementedSphereServiceServer
// for forward compatibility
type SphereServiceServer interface {
	Create(context.Context, *CreateSphereRequest) (*Sphere, error)
	GetByID(context.Context, *SpherePrimaryKey) (*Sphere, error)
	GetList(context.Context, *GetSphereListRequest) (*GetSphereListResponse, error)
	Update(context.Context, *UpdateSphereRequest) (*Sphere, error)
	Delete(context.Context, *SpherePrimaryKey) (*empty.Empty, error)
	mustEmbedUnimplementedSphereServiceServer()
}

// UnimplementedSphereServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSphereServiceServer struct {
}

func (UnimplementedSphereServiceServer) Create(context.Context, *CreateSphereRequest) (*Sphere, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSphereServiceServer) GetByID(context.Context, *SpherePrimaryKey) (*Sphere, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedSphereServiceServer) GetList(context.Context, *GetSphereListRequest) (*GetSphereListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedSphereServiceServer) Update(context.Context, *UpdateSphereRequest) (*Sphere, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSphereServiceServer) Delete(context.Context, *SpherePrimaryKey) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedSphereServiceServer) mustEmbedUnimplementedSphereServiceServer() {}

// UnsafeSphereServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SphereServiceServer will
// result in compilation errors.
type UnsafeSphereServiceServer interface {
	mustEmbedUnimplementedSphereServiceServer()
}

func RegisterSphereServiceServer(s grpc.ServiceRegistrar, srv SphereServiceServer) {
	s.RegisterService(&SphereService_ServiceDesc, srv)
}

func _SphereService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSphereRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SphereServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/settings_service.SphereService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SphereServiceServer).Create(ctx, req.(*CreateSphereRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SphereService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpherePrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SphereServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/settings_service.SphereService/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SphereServiceServer).GetByID(ctx, req.(*SpherePrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _SphereService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSphereListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SphereServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/settings_service.SphereService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SphereServiceServer).GetList(ctx, req.(*GetSphereListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SphereService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSphereRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SphereServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/settings_service.SphereService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SphereServiceServer).Update(ctx, req.(*UpdateSphereRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SphereService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpherePrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SphereServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/settings_service.SphereService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SphereServiceServer).Delete(ctx, req.(*SpherePrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

// SphereService_ServiceDesc is the grpc.ServiceDesc for SphereService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SphereService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "settings_service.SphereService",
	HandlerType: (*SphereServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _SphereService_Create_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _SphereService_GetByID_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _SphereService_GetList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SphereService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SphereService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sphere_service.proto",
}
