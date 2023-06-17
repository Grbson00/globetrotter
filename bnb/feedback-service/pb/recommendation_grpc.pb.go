// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.2
// source: recommendation.proto

package pb

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

// RecommendationServiceDBEventsClient is the client API for RecommendationServiceDBEvents service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecommendationServiceDBEventsClient interface {
	CreateAccommodation(ctx context.Context, in *GraphAccommodation, opts ...grpc.CallOption) (*GraphEmpty, error)
	CreateUser(ctx context.Context, in *GraphUser, opts ...grpc.CallOption) (*GraphEmpty, error)
	CreateReservation(ctx context.Context, in *GraphReservation, opts ...grpc.CallOption) (*GraphEmpty, error)
	CreateReview(ctx context.Context, in *GraphReview, opts ...grpc.CallOption) (*GraphEmpty, error)
	DeleteAccommodation(ctx context.Context, in *GraphAccommodation, opts ...grpc.CallOption) (*GraphEmpty, error)
	DeleteUser(ctx context.Context, in *GraphUser, opts ...grpc.CallOption) (*GraphEmpty, error)
	DeleteReservation(ctx context.Context, in *GraphReservation, opts ...grpc.CallOption) (*GraphEmpty, error)
	DeleteReview(ctx context.Context, in *GraphReview, opts ...grpc.CallOption) (*GraphEmpty, error)
	UpdateAccommodation(ctx context.Context, in *GraphAccommodation, opts ...grpc.CallOption) (*GraphEmpty, error)
	UpdateUser(ctx context.Context, in *GraphUser, opts ...grpc.CallOption) (*GraphEmpty, error)
	UpdateReservation(ctx context.Context, in *GraphReservation, opts ...grpc.CallOption) (*GraphEmpty, error)
	UpdateReview(ctx context.Context, in *GraphReview, opts ...grpc.CallOption) (*GraphEmpty, error)
}

type recommendationServiceDBEventsClient struct {
	cc grpc.ClientConnInterface
}

func NewRecommendationServiceDBEventsClient(cc grpc.ClientConnInterface) RecommendationServiceDBEventsClient {
	return &recommendationServiceDBEventsClient{cc}
}

func (c *recommendationServiceDBEventsClient) CreateAccommodation(ctx context.Context, in *GraphAccommodation, opts ...grpc.CallOption) (*GraphEmpty, error) {
	out := new(GraphEmpty)
	err := c.cc.Invoke(ctx, "/pb.RecommendationServiceDBEvents/CreateAccommodation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recommendationServiceDBEventsClient) CreateUser(ctx context.Context, in *GraphUser, opts ...grpc.CallOption) (*GraphEmpty, error) {
	out := new(GraphEmpty)
	err := c.cc.Invoke(ctx, "/pb.RecommendationServiceDBEvents/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recommendationServiceDBEventsClient) CreateReservation(ctx context.Context, in *GraphReservation, opts ...grpc.CallOption) (*GraphEmpty, error) {
	out := new(GraphEmpty)
	err := c.cc.Invoke(ctx, "/pb.RecommendationServiceDBEvents/CreateReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recommendationServiceDBEventsClient) CreateReview(ctx context.Context, in *GraphReview, opts ...grpc.CallOption) (*GraphEmpty, error) {
	out := new(GraphEmpty)
	err := c.cc.Invoke(ctx, "/pb.RecommendationServiceDBEvents/CreateReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recommendationServiceDBEventsClient) DeleteAccommodation(ctx context.Context, in *GraphAccommodation, opts ...grpc.CallOption) (*GraphEmpty, error) {
	out := new(GraphEmpty)
	err := c.cc.Invoke(ctx, "/pb.RecommendationServiceDBEvents/DeleteAccommodation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recommendationServiceDBEventsClient) DeleteUser(ctx context.Context, in *GraphUser, opts ...grpc.CallOption) (*GraphEmpty, error) {
	out := new(GraphEmpty)
	err := c.cc.Invoke(ctx, "/pb.RecommendationServiceDBEvents/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recommendationServiceDBEventsClient) DeleteReservation(ctx context.Context, in *GraphReservation, opts ...grpc.CallOption) (*GraphEmpty, error) {
	out := new(GraphEmpty)
	err := c.cc.Invoke(ctx, "/pb.RecommendationServiceDBEvents/DeleteReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recommendationServiceDBEventsClient) DeleteReview(ctx context.Context, in *GraphReview, opts ...grpc.CallOption) (*GraphEmpty, error) {
	out := new(GraphEmpty)
	err := c.cc.Invoke(ctx, "/pb.RecommendationServiceDBEvents/DeleteReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recommendationServiceDBEventsClient) UpdateAccommodation(ctx context.Context, in *GraphAccommodation, opts ...grpc.CallOption) (*GraphEmpty, error) {
	out := new(GraphEmpty)
	err := c.cc.Invoke(ctx, "/pb.RecommendationServiceDBEvents/UpdateAccommodation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recommendationServiceDBEventsClient) UpdateUser(ctx context.Context, in *GraphUser, opts ...grpc.CallOption) (*GraphEmpty, error) {
	out := new(GraphEmpty)
	err := c.cc.Invoke(ctx, "/pb.RecommendationServiceDBEvents/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recommendationServiceDBEventsClient) UpdateReservation(ctx context.Context, in *GraphReservation, opts ...grpc.CallOption) (*GraphEmpty, error) {
	out := new(GraphEmpty)
	err := c.cc.Invoke(ctx, "/pb.RecommendationServiceDBEvents/UpdateReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recommendationServiceDBEventsClient) UpdateReview(ctx context.Context, in *GraphReview, opts ...grpc.CallOption) (*GraphEmpty, error) {
	out := new(GraphEmpty)
	err := c.cc.Invoke(ctx, "/pb.RecommendationServiceDBEvents/UpdateReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecommendationServiceDBEventsServer is the server API for RecommendationServiceDBEvents service.
// All implementations must embed UnimplementedRecommendationServiceDBEventsServer
// for forward compatibility
type RecommendationServiceDBEventsServer interface {
	CreateAccommodation(context.Context, *GraphAccommodation) (*GraphEmpty, error)
	CreateUser(context.Context, *GraphUser) (*GraphEmpty, error)
	CreateReservation(context.Context, *GraphReservation) (*GraphEmpty, error)
	CreateReview(context.Context, *GraphReview) (*GraphEmpty, error)
	DeleteAccommodation(context.Context, *GraphAccommodation) (*GraphEmpty, error)
	DeleteUser(context.Context, *GraphUser) (*GraphEmpty, error)
	DeleteReservation(context.Context, *GraphReservation) (*GraphEmpty, error)
	DeleteReview(context.Context, *GraphReview) (*GraphEmpty, error)
	UpdateAccommodation(context.Context, *GraphAccommodation) (*GraphEmpty, error)
	UpdateUser(context.Context, *GraphUser) (*GraphEmpty, error)
	UpdateReservation(context.Context, *GraphReservation) (*GraphEmpty, error)
	UpdateReview(context.Context, *GraphReview) (*GraphEmpty, error)
	mustEmbedUnimplementedRecommendationServiceDBEventsServer()
}

// UnimplementedRecommendationServiceDBEventsServer must be embedded to have forward compatible implementations.
type UnimplementedRecommendationServiceDBEventsServer struct {
}

func (UnimplementedRecommendationServiceDBEventsServer) CreateAccommodation(context.Context, *GraphAccommodation) (*GraphEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccommodation not implemented")
}
func (UnimplementedRecommendationServiceDBEventsServer) CreateUser(context.Context, *GraphUser) (*GraphEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedRecommendationServiceDBEventsServer) CreateReservation(context.Context, *GraphReservation) (*GraphEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReservation not implemented")
}
func (UnimplementedRecommendationServiceDBEventsServer) CreateReview(context.Context, *GraphReview) (*GraphEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReview not implemented")
}
func (UnimplementedRecommendationServiceDBEventsServer) DeleteAccommodation(context.Context, *GraphAccommodation) (*GraphEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccommodation not implemented")
}
func (UnimplementedRecommendationServiceDBEventsServer) DeleteUser(context.Context, *GraphUser) (*GraphEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedRecommendationServiceDBEventsServer) DeleteReservation(context.Context, *GraphReservation) (*GraphEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteReservation not implemented")
}
func (UnimplementedRecommendationServiceDBEventsServer) DeleteReview(context.Context, *GraphReview) (*GraphEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteReview not implemented")
}
func (UnimplementedRecommendationServiceDBEventsServer) UpdateAccommodation(context.Context, *GraphAccommodation) (*GraphEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccommodation not implemented")
}
func (UnimplementedRecommendationServiceDBEventsServer) UpdateUser(context.Context, *GraphUser) (*GraphEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedRecommendationServiceDBEventsServer) UpdateReservation(context.Context, *GraphReservation) (*GraphEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReservation not implemented")
}
func (UnimplementedRecommendationServiceDBEventsServer) UpdateReview(context.Context, *GraphReview) (*GraphEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReview not implemented")
}
func (UnimplementedRecommendationServiceDBEventsServer) mustEmbedUnimplementedRecommendationServiceDBEventsServer() {
}

// UnsafeRecommendationServiceDBEventsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecommendationServiceDBEventsServer will
// result in compilation errors.
type UnsafeRecommendationServiceDBEventsServer interface {
	mustEmbedUnimplementedRecommendationServiceDBEventsServer()
}

func RegisterRecommendationServiceDBEventsServer(s grpc.ServiceRegistrar, srv RecommendationServiceDBEventsServer) {
	s.RegisterService(&RecommendationServiceDBEvents_ServiceDesc, srv)
}

func _RecommendationServiceDBEvents_CreateAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GraphAccommodation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommendationServiceDBEventsServer).CreateAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RecommendationServiceDBEvents/CreateAccommodation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendationServiceDBEventsServer).CreateAccommodation(ctx, req.(*GraphAccommodation))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecommendationServiceDBEvents_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GraphUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommendationServiceDBEventsServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RecommendationServiceDBEvents/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendationServiceDBEventsServer).CreateUser(ctx, req.(*GraphUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecommendationServiceDBEvents_CreateReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GraphReservation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommendationServiceDBEventsServer).CreateReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RecommendationServiceDBEvents/CreateReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendationServiceDBEventsServer).CreateReservation(ctx, req.(*GraphReservation))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecommendationServiceDBEvents_CreateReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GraphReview)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommendationServiceDBEventsServer).CreateReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RecommendationServiceDBEvents/CreateReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendationServiceDBEventsServer).CreateReview(ctx, req.(*GraphReview))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecommendationServiceDBEvents_DeleteAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GraphAccommodation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommendationServiceDBEventsServer).DeleteAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RecommendationServiceDBEvents/DeleteAccommodation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendationServiceDBEventsServer).DeleteAccommodation(ctx, req.(*GraphAccommodation))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecommendationServiceDBEvents_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GraphUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommendationServiceDBEventsServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RecommendationServiceDBEvents/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendationServiceDBEventsServer).DeleteUser(ctx, req.(*GraphUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecommendationServiceDBEvents_DeleteReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GraphReservation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommendationServiceDBEventsServer).DeleteReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RecommendationServiceDBEvents/DeleteReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendationServiceDBEventsServer).DeleteReservation(ctx, req.(*GraphReservation))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecommendationServiceDBEvents_DeleteReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GraphReview)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommendationServiceDBEventsServer).DeleteReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RecommendationServiceDBEvents/DeleteReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendationServiceDBEventsServer).DeleteReview(ctx, req.(*GraphReview))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecommendationServiceDBEvents_UpdateAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GraphAccommodation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommendationServiceDBEventsServer).UpdateAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RecommendationServiceDBEvents/UpdateAccommodation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendationServiceDBEventsServer).UpdateAccommodation(ctx, req.(*GraphAccommodation))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecommendationServiceDBEvents_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GraphUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommendationServiceDBEventsServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RecommendationServiceDBEvents/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendationServiceDBEventsServer).UpdateUser(ctx, req.(*GraphUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecommendationServiceDBEvents_UpdateReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GraphReservation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommendationServiceDBEventsServer).UpdateReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RecommendationServiceDBEvents/UpdateReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendationServiceDBEventsServer).UpdateReservation(ctx, req.(*GraphReservation))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecommendationServiceDBEvents_UpdateReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GraphReview)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommendationServiceDBEventsServer).UpdateReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RecommendationServiceDBEvents/UpdateReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendationServiceDBEventsServer).UpdateReview(ctx, req.(*GraphReview))
	}
	return interceptor(ctx, in, info, handler)
}

// RecommendationServiceDBEvents_ServiceDesc is the grpc.ServiceDesc for RecommendationServiceDBEvents service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RecommendationServiceDBEvents_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.RecommendationServiceDBEvents",
	HandlerType: (*RecommendationServiceDBEventsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccommodation",
			Handler:    _RecommendationServiceDBEvents_CreateAccommodation_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _RecommendationServiceDBEvents_CreateUser_Handler,
		},
		{
			MethodName: "CreateReservation",
			Handler:    _RecommendationServiceDBEvents_CreateReservation_Handler,
		},
		{
			MethodName: "CreateReview",
			Handler:    _RecommendationServiceDBEvents_CreateReview_Handler,
		},
		{
			MethodName: "DeleteAccommodation",
			Handler:    _RecommendationServiceDBEvents_DeleteAccommodation_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _RecommendationServiceDBEvents_DeleteUser_Handler,
		},
		{
			MethodName: "DeleteReservation",
			Handler:    _RecommendationServiceDBEvents_DeleteReservation_Handler,
		},
		{
			MethodName: "DeleteReview",
			Handler:    _RecommendationServiceDBEvents_DeleteReview_Handler,
		},
		{
			MethodName: "UpdateAccommodation",
			Handler:    _RecommendationServiceDBEvents_UpdateAccommodation_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _RecommendationServiceDBEvents_UpdateUser_Handler,
		},
		{
			MethodName: "UpdateReservation",
			Handler:    _RecommendationServiceDBEvents_UpdateReservation_Handler,
		},
		{
			MethodName: "UpdateReview",
			Handler:    _RecommendationServiceDBEvents_UpdateReview_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "recommendation.proto",
}
