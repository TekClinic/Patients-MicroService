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
	GetPatient(ctx context.Context, in *GetPatientRequest, opts ...grpc.CallOption) (*GetPatientResponse, error)
	GetPatientsIDs(ctx context.Context, in *GetPatientsIDsRequest, opts ...grpc.CallOption) (*GetPatientsIDsResponse, error)
	CreatePatient(ctx context.Context, in *CreatePatientRequest, opts ...grpc.CallOption) (*CreatePatientResponse, error)
	DeletePatient(ctx context.Context, in *DeletePatientRequest, opts ...grpc.CallOption) (*DeletePatientResponse, error)
	UpdatePatient(ctx context.Context, in *UpdatePatientRequest, opts ...grpc.CallOption) (*UpdatePatientResponse, error)
}

type patientsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPatientsServiceClient(cc grpc.ClientConnInterface) PatientsServiceClient {
	return &patientsServiceClient{cc}
}

func (c *patientsServiceClient) GetPatient(ctx context.Context, in *GetPatientRequest, opts ...grpc.CallOption) (*GetPatientResponse, error) {
	out := new(GetPatientResponse)
	err := c.cc.Invoke(ctx, "/patients.PatientsService/GetPatient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *patientsServiceClient) GetPatientsIDs(ctx context.Context, in *GetPatientsIDsRequest, opts ...grpc.CallOption) (*GetPatientsIDsResponse, error) {
	out := new(GetPatientsIDsResponse)
	err := c.cc.Invoke(ctx, "/patients.PatientsService/GetPatientsIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *patientsServiceClient) CreatePatient(ctx context.Context, in *CreatePatientRequest, opts ...grpc.CallOption) (*CreatePatientResponse, error) {
	out := new(CreatePatientResponse)
	err := c.cc.Invoke(ctx, "/patients.PatientsService/CreatePatient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *patientsServiceClient) DeletePatient(ctx context.Context, in *DeletePatientRequest, opts ...grpc.CallOption) (*DeletePatientResponse, error) {
	out := new(DeletePatientResponse)
	err := c.cc.Invoke(ctx, "/patients.PatientsService/DeletePatient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *patientsServiceClient) UpdatePatient(ctx context.Context, in *UpdatePatientRequest, opts ...grpc.CallOption) (*UpdatePatientResponse, error) {
	out := new(UpdatePatientResponse)
	err := c.cc.Invoke(ctx, "/patients.PatientsService/UpdatePatient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PatientsServiceServer is the server API for PatientsService service.
// All implementations must embed UnimplementedPatientsServiceServer
// for forward compatibility
type PatientsServiceServer interface {
	GetPatient(context.Context, *GetPatientRequest) (*GetPatientResponse, error)
	GetPatientsIDs(context.Context, *GetPatientsIDsRequest) (*GetPatientsIDsResponse, error)
	CreatePatient(context.Context, *CreatePatientRequest) (*CreatePatientResponse, error)
	DeletePatient(context.Context, *DeletePatientRequest) (*DeletePatientResponse, error)
	UpdatePatient(context.Context, *UpdatePatientRequest) (*UpdatePatientResponse, error)
	mustEmbedUnimplementedPatientsServiceServer()
}

// UnimplementedPatientsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPatientsServiceServer struct {
}

func (UnimplementedPatientsServiceServer) GetPatient(context.Context, *GetPatientRequest) (*GetPatientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPatient not implemented")
}
func (UnimplementedPatientsServiceServer) GetPatientsIDs(context.Context, *GetPatientsIDsRequest) (*GetPatientsIDsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPatientsIDs not implemented")
}
func (UnimplementedPatientsServiceServer) CreatePatient(context.Context, *CreatePatientRequest) (*CreatePatientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePatient not implemented")
}
func (UnimplementedPatientsServiceServer) DeletePatient(context.Context, *DeletePatientRequest) (*DeletePatientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePatient not implemented")
}
func (UnimplementedPatientsServiceServer) UpdatePatient(context.Context, *UpdatePatientRequest) (*UpdatePatientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePatient not implemented")
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
	in := new(GetPatientRequest)
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
		return srv.(PatientsServiceServer).GetPatient(ctx, req.(*GetPatientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PatientsService_GetPatientsIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPatientsIDsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientsServiceServer).GetPatientsIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/patients.PatientsService/GetPatientsIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientsServiceServer).GetPatientsIDs(ctx, req.(*GetPatientsIDsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PatientsService_CreatePatient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePatientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientsServiceServer).CreatePatient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/patients.PatientsService/CreatePatient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientsServiceServer).CreatePatient(ctx, req.(*CreatePatientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PatientsService_DeletePatient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePatientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientsServiceServer).DeletePatient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/patients.PatientsService/DeletePatient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientsServiceServer).DeletePatient(ctx, req.(*DeletePatientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PatientsService_UpdatePatient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePatientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientsServiceServer).UpdatePatient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/patients.PatientsService/UpdatePatient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientsServiceServer).UpdatePatient(ctx, req.(*UpdatePatientRequest))
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
			MethodName: "GetPatientsIDs",
			Handler:    _PatientsService_GetPatientsIDs_Handler,
		},
		{
			MethodName: "CreatePatient",
			Handler:    _PatientsService_CreatePatient_Handler,
		},
		{
			MethodName: "DeletePatient",
			Handler:    _PatientsService_DeletePatient_Handler,
		},
		{
			MethodName: "UpdatePatient",
			Handler:    _PatientsService_UpdatePatient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "patients_service.proto",
}
