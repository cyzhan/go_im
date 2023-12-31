// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: vendors.proto

package vendors

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// VendorsServiceClient is the client API for VendorsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VendorsServiceClient interface {
	UpdateSensitiveWord(ctx context.Context, in *UpdateSensitiveWordReq, opts ...grpc.CallOption) (*UpdateSensitiveWordResp, error)
	GetSensitiveWordPage(ctx context.Context, in *GetSensitiveWordPageReq, opts ...grpc.CallOption) (*GetSensitiveWordPageResp, error)
	GetSensitiveWordList(ctx context.Context, in *GetSensitiveWordListReq, opts ...grpc.CallOption) (*GetSensitiveWordListResp, error)
	GetConfig(ctx context.Context, in *GetConfigReq, opts ...grpc.CallOption) (*GetConfigResp, error)
}

type vendorsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVendorsServiceClient(cc grpc.ClientConnInterface) VendorsServiceClient {
	return &vendorsServiceClient{cc}
}

func (c *vendorsServiceClient) UpdateSensitiveWord(ctx context.Context, in *UpdateSensitiveWordReq, opts ...grpc.CallOption) (*UpdateSensitiveWordResp, error) {
	out := new(UpdateSensitiveWordResp)
	err := c.cc.Invoke(ctx, "/vendors.vendorsService/UpdateSensitiveWord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendorsServiceClient) GetSensitiveWordPage(ctx context.Context, in *GetSensitiveWordPageReq, opts ...grpc.CallOption) (*GetSensitiveWordPageResp, error) {
	out := new(GetSensitiveWordPageResp)
	err := c.cc.Invoke(ctx, "/vendors.vendorsService/GetSensitiveWordPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendorsServiceClient) GetSensitiveWordList(ctx context.Context, in *GetSensitiveWordListReq, opts ...grpc.CallOption) (*GetSensitiveWordListResp, error) {
	out := new(GetSensitiveWordListResp)
	err := c.cc.Invoke(ctx, "/vendors.vendorsService/GetSensitiveWordList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendorsServiceClient) GetConfig(ctx context.Context, in *GetConfigReq, opts ...grpc.CallOption) (*GetConfigResp, error) {
	out := new(GetConfigResp)
	err := c.cc.Invoke(ctx, "/vendors.vendorsService/GetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VendorsServiceServer is the server API for VendorsService service.
// All implementations should embed UnimplementedVendorsServiceServer
// for forward compatibility
type VendorsServiceServer interface {
	UpdateSensitiveWord(context.Context, *UpdateSensitiveWordReq) (*UpdateSensitiveWordResp, error)
	GetSensitiveWordPage(context.Context, *GetSensitiveWordPageReq) (*GetSensitiveWordPageResp, error)
	GetSensitiveWordList(context.Context, *GetSensitiveWordListReq) (*GetSensitiveWordListResp, error)
	GetConfig(context.Context, *GetConfigReq) (*GetConfigResp, error)
}

// UnimplementedVendorsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedVendorsServiceServer struct {
}

func (UnimplementedVendorsServiceServer) UpdateSensitiveWord(context.Context, *UpdateSensitiveWordReq) (*UpdateSensitiveWordResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSensitiveWord not implemented")
}
func (UnimplementedVendorsServiceServer) GetSensitiveWordPage(context.Context, *GetSensitiveWordPageReq) (*GetSensitiveWordPageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSensitiveWordPage not implemented")
}
func (UnimplementedVendorsServiceServer) GetSensitiveWordList(context.Context, *GetSensitiveWordListReq) (*GetSensitiveWordListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSensitiveWordList not implemented")
}
func (UnimplementedVendorsServiceServer) GetConfig(context.Context, *GetConfigReq) (*GetConfigResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}

// UnsafeVendorsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VendorsServiceServer will
// result in compilation errors.
type UnsafeVendorsServiceServer interface {
	mustEmbedUnimplementedVendorsServiceServer()
}

func RegisterVendorsServiceServer(s grpc.ServiceRegistrar, srv VendorsServiceServer) {
	s.RegisterService(&VendorsService_ServiceDesc, srv)
}

func _VendorsService_UpdateSensitiveWord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSensitiveWordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VendorsServiceServer).UpdateSensitiveWord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vendors.vendorsService/UpdateSensitiveWord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VendorsServiceServer).UpdateSensitiveWord(ctx, req.(*UpdateSensitiveWordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VendorsService_GetSensitiveWordPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSensitiveWordPageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VendorsServiceServer).GetSensitiveWordPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vendors.vendorsService/GetSensitiveWordPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VendorsServiceServer).GetSensitiveWordPage(ctx, req.(*GetSensitiveWordPageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VendorsService_GetSensitiveWordList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSensitiveWordListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VendorsServiceServer).GetSensitiveWordList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vendors.vendorsService/GetSensitiveWordList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VendorsServiceServer).GetSensitiveWordList(ctx, req.(*GetSensitiveWordListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VendorsService_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VendorsServiceServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vendors.vendorsService/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VendorsServiceServer).GetConfig(ctx, req.(*GetConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

// VendorsService_ServiceDesc is the grpc.ServiceDesc for VendorsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VendorsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vendors.vendorsService",
	HandlerType: (*VendorsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateSensitiveWord",
			Handler:    _VendorsService_UpdateSensitiveWord_Handler,
		},
		{
			MethodName: "GetSensitiveWordPage",
			Handler:    _VendorsService_GetSensitiveWordPage_Handler,
		},
		{
			MethodName: "GetSensitiveWordList",
			Handler:    _VendorsService_GetSensitiveWordList_Handler,
		},
		{
			MethodName: "GetConfig",
			Handler:    _VendorsService_GetConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vendors.proto",
}
