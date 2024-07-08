// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: staffer-protos/resume.proto

package genprotos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	ResumeService_Create_FullMethodName = "/staff.ResumeService/Create"
	ResumeService_Get_FullMethodName    = "/staff.ResumeService/Get"
	ResumeService_Update_FullMethodName = "/staff.ResumeService/Update"
	ResumeService_Delete_FullMethodName = "/staff.ResumeService/Delete"
	ResumeService_GetAll_FullMethodName = "/staff.ResumeService/GetAll"
)

// ResumeServiceClient is the client API for ResumeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResumeServiceClient interface {
	Create(ctx context.Context, in *ResumeCreate, opts ...grpc.CallOption) (*Void, error)
	Get(ctx context.Context, in *Byid, opts ...grpc.CallOption) (*ResumeGetRes, error)
	Update(ctx context.Context, in *ResumeUpdate, opts ...grpc.CallOption) (*Void, error)
	Delete(ctx context.Context, in *Byid, opts ...grpc.CallOption) (*Void, error)
	GetAll(ctx context.Context, in *ResumeGetAll, opts ...grpc.CallOption) (*ResumeGetAllRes, error)
}

type resumeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewResumeServiceClient(cc grpc.ClientConnInterface) ResumeServiceClient {
	return &resumeServiceClient{cc}
}

func (c *resumeServiceClient) Create(ctx context.Context, in *ResumeCreate, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, ResumeService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resumeServiceClient) Get(ctx context.Context, in *Byid, opts ...grpc.CallOption) (*ResumeGetRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResumeGetRes)
	err := c.cc.Invoke(ctx, ResumeService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resumeServiceClient) Update(ctx context.Context, in *ResumeUpdate, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, ResumeService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resumeServiceClient) Delete(ctx context.Context, in *Byid, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, ResumeService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resumeServiceClient) GetAll(ctx context.Context, in *ResumeGetAll, opts ...grpc.CallOption) (*ResumeGetAllRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResumeGetAllRes)
	err := c.cc.Invoke(ctx, ResumeService_GetAll_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResumeServiceServer is the server API for ResumeService service.
// All implementations must embed UnimplementedResumeServiceServer
// for forward compatibility
type ResumeServiceServer interface {
	Create(context.Context, *ResumeCreate) (*Void, error)
	Get(context.Context, *Byid) (*ResumeGetRes, error)
	Update(context.Context, *ResumeUpdate) (*Void, error)
	Delete(context.Context, *Byid) (*Void, error)
	GetAll(context.Context, *ResumeGetAll) (*ResumeGetAllRes, error)
	mustEmbedUnimplementedResumeServiceServer()
}

// UnimplementedResumeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedResumeServiceServer struct {
}

func (UnimplementedResumeServiceServer) Create(context.Context, *ResumeCreate) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedResumeServiceServer) Get(context.Context, *Byid) (*ResumeGetRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedResumeServiceServer) Update(context.Context, *ResumeUpdate) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedResumeServiceServer) Delete(context.Context, *Byid) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedResumeServiceServer) GetAll(context.Context, *ResumeGetAll) (*ResumeGetAllRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedResumeServiceServer) mustEmbedUnimplementedResumeServiceServer() {}

// UnsafeResumeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResumeServiceServer will
// result in compilation errors.
type UnsafeResumeServiceServer interface {
	mustEmbedUnimplementedResumeServiceServer()
}

func RegisterResumeServiceServer(s grpc.ServiceRegistrar, srv ResumeServiceServer) {
	s.RegisterService(&ResumeService_ServiceDesc, srv)
}

func _ResumeService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResumeCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResumeServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResumeService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResumeServiceServer).Create(ctx, req.(*ResumeCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResumeService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Byid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResumeServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResumeService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResumeServiceServer).Get(ctx, req.(*Byid))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResumeService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResumeUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResumeServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResumeService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResumeServiceServer).Update(ctx, req.(*ResumeUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResumeService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Byid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResumeServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResumeService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResumeServiceServer).Delete(ctx, req.(*Byid))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResumeService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResumeGetAll)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResumeServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResumeService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResumeServiceServer).GetAll(ctx, req.(*ResumeGetAll))
	}
	return interceptor(ctx, in, info, handler)
}

// ResumeService_ServiceDesc is the grpc.ServiceDesc for ResumeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResumeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "staff.ResumeService",
	HandlerType: (*ResumeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ResumeService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ResumeService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ResumeService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ResumeService_Delete_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _ResumeService_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "staffer-protos/resume.proto",
}
