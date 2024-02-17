// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: payment-processing.proto

package api

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
	UserTransactionsService_GetUserMonthAverage_FullMethodName    = "/user.UserTransactionsService/GetUserMonthAverage"
	UserTransactionsService_GetLastUserTransaction_FullMethodName = "/user.UserTransactionsService/GetLastUserTransaction"
)

// UserTransactionsServiceClient is the client API for UserTransactionsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserTransactionsServiceClient interface {
	// Gets the user's month average
	GetUserMonthAverage(ctx context.Context, in *UserMonthAverageRequest, opts ...grpc.CallOption) (*UserMonthAverageResponse, error)
	// Gets last user transaction
	GetLastUserTransaction(ctx context.Context, in *LastUserTransactionRequest, opts ...grpc.CallOption) (*LastUserTransactionResponse, error)
}

type userTransactionsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserTransactionsServiceClient(cc grpc.ClientConnInterface) UserTransactionsServiceClient {
	return &userTransactionsServiceClient{cc}
}

func (c *userTransactionsServiceClient) GetUserMonthAverage(ctx context.Context, in *UserMonthAverageRequest, opts ...grpc.CallOption) (*UserMonthAverageResponse, error) {
	out := new(UserMonthAverageResponse)
	err := c.cc.Invoke(ctx, UserTransactionsService_GetUserMonthAverage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userTransactionsServiceClient) GetLastUserTransaction(ctx context.Context, in *LastUserTransactionRequest, opts ...grpc.CallOption) (*LastUserTransactionResponse, error) {
	out := new(LastUserTransactionResponse)
	err := c.cc.Invoke(ctx, UserTransactionsService_GetLastUserTransaction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserTransactionsServiceServer is the server API for UserTransactionsService service.
// All implementations must embed UnimplementedUserTransactionsServiceServer
// for forward compatibility
type UserTransactionsServiceServer interface {
	// Gets the user's month average
	GetUserMonthAverage(context.Context, *UserMonthAverageRequest) (*UserMonthAverageResponse, error)
	// Gets last user transaction
	GetLastUserTransaction(context.Context, *LastUserTransactionRequest) (*LastUserTransactionResponse, error)
	mustEmbedUnimplementedUserTransactionsServiceServer()
}

// UnimplementedUserTransactionsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserTransactionsServiceServer struct {
}

func (UnimplementedUserTransactionsServiceServer) GetUserMonthAverage(context.Context, *UserMonthAverageRequest) (*UserMonthAverageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserMonthAverage not implemented")
}
func (UnimplementedUserTransactionsServiceServer) GetLastUserTransaction(context.Context, *LastUserTransactionRequest) (*LastUserTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLastUserTransaction not implemented")
}
func (UnimplementedUserTransactionsServiceServer) mustEmbedUnimplementedUserTransactionsServiceServer() {
}

// UnsafeUserTransactionsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserTransactionsServiceServer will
// result in compilation errors.
type UnsafeUserTransactionsServiceServer interface {
	mustEmbedUnimplementedUserTransactionsServiceServer()
}

func RegisterUserTransactionsServiceServer(s grpc.ServiceRegistrar, srv UserTransactionsServiceServer) {
	s.RegisterService(&UserTransactionsService_ServiceDesc, srv)
}

func _UserTransactionsService_GetUserMonthAverage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserMonthAverageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserTransactionsServiceServer).GetUserMonthAverage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserTransactionsService_GetUserMonthAverage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserTransactionsServiceServer).GetUserMonthAverage(ctx, req.(*UserMonthAverageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserTransactionsService_GetLastUserTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LastUserTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserTransactionsServiceServer).GetLastUserTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserTransactionsService_GetLastUserTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserTransactionsServiceServer).GetLastUserTransaction(ctx, req.(*LastUserTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserTransactionsService_ServiceDesc is the grpc.ServiceDesc for UserTransactionsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserTransactionsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserTransactionsService",
	HandlerType: (*UserTransactionsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserMonthAverage",
			Handler:    _UserTransactionsService_GetUserMonthAverage_Handler,
		},
		{
			MethodName: "GetLastUserTransaction",
			Handler:    _UserTransactionsService_GetLastUserTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "payment-processing.proto",
}
