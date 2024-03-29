// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: patients_service.proto

package patients_protobuf

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

// PatientsServiceClient is the client API for PatientsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PatientsServiceClient interface {
	GetPatient(ctx context.Context, in *PatientRequest, opts ...grpc.CallOption) (*Patient, error)
	GetPatientsIds(ctx context.Context, in *PatientsRequest, opts ...grpc.CallOption) (*PaginatedResponse, error)
}

type patientsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPatientsServiceClient(cc grpc.ClientConnInterface) PatientsServiceClient {
	return &patientsServiceClient{cc}
}

func (c *patientsServiceClient) GetPatient(ctx context.Context, in *PatientRequest, opts ...grpc.CallOption) (*Patient, error) {
	out := new(Patient)
	err := c.cc.Invoke(ctx, "/patients.PatientsService/GetPatient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *patientsServiceClient) GetPatientsIds(ctx context.Context, in *PatientsRequest, opts ...grpc.CallOption) (*PaginatedResponse, error) {
	out := new(PaginatedResponse)
	err := c.cc.Invoke(ctx, "/patients.PatientsService/GetPatientsIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PatientsServiceServer is the server API for PatientsService service.
// All implementations must embed UnimplementedPatientsServiceServer
// for forward compatibility
type PatientsServiceServer interface {
	GetPatient(context.Context, *PatientRequest) (*Patient, error)
	GetPatientsIds(context.Context, *PatientsRequest) (*PaginatedResponse, error)
	mustEmbedUnimplementedPatientsServiceServer()
}

// UnimplementedPatientsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPatientsServiceServer struct {
}

func (UnimplementedPatientsServiceServer) GetPatient(context.Context, *PatientRequest) (*Patient, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPatient not implemented")
}
func (UnimplementedPatientsServiceServer) GetPatientsIds(context.Context, *PatientsRequest) (*PaginatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPatientsIds not implemented")
}
func (UnimplementedPatientsServiceServer) mustEmbedUnimplementedPatientsServiceServer() {}

// UnsafePatientsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PatientsServiceServer will
// result in compilation errors.
type UnsafePatientsServiceServer interface {
	mustEmbedUnimplementedPatientsServiceServer()
}

func RegisterPatientsServiceServer(s grpc.ServiceRegistrar, srv PatientsServiceServer) {
	s.RegisterService(&PatientsService_ServiceDesc, srv)
}

func _PatientsService_GetPatient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientsServiceServer).GetPatient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/patients.PatientsService/GetPatient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientsServiceServer).GetPatient(ctx, req.(*PatientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PatientsService_GetPatientsIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatientsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientsServiceServer).GetPatientsIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/patients.PatientsService/GetPatientsIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientsServiceServer).GetPatientsIds(ctx, req.(*PatientsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PatientsService_ServiceDesc is the grpc.ServiceDesc for PatientsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PatientsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "patients.PatientsService",
	HandlerType: (*PatientsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPatient",
			Handler:    _PatientsService_GetPatient_Handler,
		},
		{
			MethodName: "GetPatientsIds",
			Handler:    _PatientsService_GetPatientsIds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "patients_service.proto",
}
