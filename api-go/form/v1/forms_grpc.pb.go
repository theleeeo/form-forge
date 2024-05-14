// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: form/v1/forms.proto

package v1

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

const (
	FormService_GetByID_FullMethodName = "/form.v1.FormService/GetByID"
	FormService_Create_FullMethodName  = "/form.v1.FormService/Create"
)

// FormServiceClient is the client API for FormService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FormServiceClient interface {
	GetByID(ctx context.Context, in *GetByIDRequest, opts ...grpc.CallOption) (*GetByIDResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
}

type formServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFormServiceClient(cc grpc.ClientConnInterface) FormServiceClient {
	return &formServiceClient{cc}
}

func (c *formServiceClient) GetByID(ctx context.Context, in *GetByIDRequest, opts ...grpc.CallOption) (*GetByIDResponse, error) {
	out := new(GetByIDResponse)
	err := c.cc.Invoke(ctx, FormService_GetByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *formServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, FormService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FormServiceServer is the server API for FormService service.
// All implementations should embed UnimplementedFormServiceServer
// for forward compatibility
type FormServiceServer interface {
	GetByID(context.Context, *GetByIDRequest) (*GetByIDResponse, error)
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
}

// UnimplementedFormServiceServer should be embedded to have forward compatible implementations.
type UnimplementedFormServiceServer struct {
}

func (UnimplementedFormServiceServer) GetByID(context.Context, *GetByIDRequest) (*GetByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedFormServiceServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

// UnsafeFormServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FormServiceServer will
// result in compilation errors.
type UnsafeFormServiceServer interface {
	mustEmbedUnimplementedFormServiceServer()
}

func RegisterFormServiceServer(s grpc.ServiceRegistrar, srv FormServiceServer) {
	s.RegisterService(&FormService_ServiceDesc, srv)
}

func _FormService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FormServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FormService_GetByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FormServiceServer).GetByID(ctx, req.(*GetByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FormService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FormServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FormService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FormServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FormService_ServiceDesc is the grpc.ServiceDesc for FormService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FormService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "form.v1.FormService",
	HandlerType: (*FormServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetByID",
			Handler:    _FormService_GetByID_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _FormService_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "form/v1/forms.proto",
}
