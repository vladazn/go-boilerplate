// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.13.0
// source: party.proto

package frontoffice

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

// PartyServiceClient is the client API for PartyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PartyServiceClient interface {
	CreateParty(ctx context.Context, in *CreatePartyRequest, opts ...grpc.CallOption) (*CreatePartyResponse, error)
}

type partyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPartyServiceClient(cc grpc.ClientConnInterface) PartyServiceClient {
	return &partyServiceClient{cc}
}

func (c *partyServiceClient) CreateParty(ctx context.Context, in *CreatePartyRequest, opts ...grpc.CallOption) (*CreatePartyResponse, error) {
	out := new(CreatePartyResponse)
	err := c.cc.Invoke(ctx, "/boilerplate.PartyService/CreateParty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PartyServiceServer is the server API for PartyService service.
// All implementations must embed UnimplementedPartyServiceServer
// for forward compatibility
type PartyServiceServer interface {
	CreateParty(context.Context, *CreatePartyRequest) (*CreatePartyResponse, error)
	mustEmbedUnimplementedPartyServiceServer()
}

// UnimplementedPartyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPartyServiceServer struct {
}

func (UnimplementedPartyServiceServer) CreateParty(context.Context, *CreatePartyRequest) (*CreatePartyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateParty not implemented")
}
func (UnimplementedPartyServiceServer) mustEmbedUnimplementedPartyServiceServer() {}

// UnsafePartyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PartyServiceServer will
// result in compilation errors.
type UnsafePartyServiceServer interface {
	mustEmbedUnimplementedPartyServiceServer()
}

func RegisterPartyServiceServer(s grpc.ServiceRegistrar, srv PartyServiceServer) {
	s.RegisterService(&PartyService_ServiceDesc, srv)
}

func _PartyService_CreateParty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).CreateParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/boilerplate.PartyService/CreateParty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).CreateParty(ctx, req.(*CreatePartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PartyService_ServiceDesc is the grpc.ServiceDesc for PartyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PartyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "boilerplate.PartyService",
	HandlerType: (*PartyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateParty",
			Handler:    _PartyService_CreateParty_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "party.proto",
}
