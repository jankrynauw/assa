// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package assa

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

// MembersClient is the client API for Members service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MembersClient interface {
	// Creates a new Member resource
	CreateMember(ctx context.Context, in *CreateMemberRequest, opts ...grpc.CallOption) (*Member, error)
	// Get a Member resource
	GetMember(ctx context.Context, in *GetMemberRequest, opts ...grpc.CallOption) (*Member, error)
}

type membersClient struct {
	cc grpc.ClientConnInterface
}

func NewMembersClient(cc grpc.ClientConnInterface) MembersClient {
	return &membersClient{cc}
}

func (c *membersClient) CreateMember(ctx context.Context, in *CreateMemberRequest, opts ...grpc.CallOption) (*Member, error) {
	out := new(Member)
	err := c.cc.Invoke(ctx, "/assa.v1.Members/CreateMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *membersClient) GetMember(ctx context.Context, in *GetMemberRequest, opts ...grpc.CallOption) (*Member, error) {
	out := new(Member)
	err := c.cc.Invoke(ctx, "/assa.v1.Members/GetMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MembersServer is the server API for Members service.
// All implementations must embed UnimplementedMembersServer
// for forward compatibility
type MembersServer interface {
	// Creates a new Member resource
	CreateMember(context.Context, *CreateMemberRequest) (*Member, error)
	// Get a Member resource
	GetMember(context.Context, *GetMemberRequest) (*Member, error)
	mustEmbedUnimplementedMembersServer()
}

// UnimplementedMembersServer must be embedded to have forward compatible implementations.
type UnimplementedMembersServer struct {
}

func (UnimplementedMembersServer) CreateMember(context.Context, *CreateMemberRequest) (*Member, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMember not implemented")
}
func (UnimplementedMembersServer) GetMember(context.Context, *GetMemberRequest) (*Member, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMember not implemented")
}
func (UnimplementedMembersServer) mustEmbedUnimplementedMembersServer() {}

// UnsafeMembersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MembersServer will
// result in compilation errors.
type UnsafeMembersServer interface {
	mustEmbedUnimplementedMembersServer()
}

func RegisterMembersServer(s grpc.ServiceRegistrar, srv MembersServer) {
	s.RegisterService(&Members_ServiceDesc, srv)
}

func _Members_CreateMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MembersServer).CreateMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/assa.v1.Members/CreateMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MembersServer).CreateMember(ctx, req.(*CreateMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Members_GetMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MembersServer).GetMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/assa.v1.Members/GetMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MembersServer).GetMember(ctx, req.(*GetMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Members_ServiceDesc is the grpc.ServiceDesc for Members service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Members_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "assa.v1.Members",
	HandlerType: (*MembersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMember",
			Handler:    _Members_CreateMember_Handler,
		},
		{
			MethodName: "GetMember",
			Handler:    _Members_GetMember_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "assa.proto",
}
