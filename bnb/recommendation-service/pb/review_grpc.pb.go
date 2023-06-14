// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.0
// source: review.proto

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

// FeedbackServiceClient is the client API for FeedbackService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeedbackServiceClient interface {
	GetHostReviewById(ctx context.Context, in *RequestReviewById, opts ...grpc.CallOption) (*HostReview, error)
	GetHostReviewsByUserId(ctx context.Context, in *RequestReviewsByUserId, opts ...grpc.CallOption) (FeedbackService_GetHostReviewsByUserIdClient, error)
	GetHostReviewsByHostId(ctx context.Context, in *RequestReviewsByHostId, opts ...grpc.CallOption) (FeedbackService_GetHostReviewsByHostIdClient, error)
	CalcAvgRatingForHost(ctx context.Context, in *RequestAvgRating, opts ...grpc.CallOption) (*AvgRatingResponse, error)
	GetAllAccommodationReviews(ctx context.Context, in *EmptyReviewMsg, opts ...grpc.CallOption) (FeedbackService_GetAllAccommodationReviewsClient, error)
	GetAccommodationReviewById(ctx context.Context, in *RequestReviewById, opts ...grpc.CallOption) (*AccommodationReview, error)
	GetAccommodationReviewsByUserId(ctx context.Context, in *RequestReviewsByUserId, opts ...grpc.CallOption) (FeedbackService_GetAccommodationReviewsByUserIdClient, error)
	GetAccommodationReviewsByAccommodationId(ctx context.Context, in *RequestReviewsByAccommodationId, opts ...grpc.CallOption) (FeedbackService_GetAccommodationReviewsByAccommodationIdClient, error)
	CalcAvgRatingForAccommodation(ctx context.Context, in *RequestAvgRating, opts ...grpc.CallOption) (*AvgRatingResponse, error)
}

type feedbackServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFeedbackServiceClient(cc grpc.ClientConnInterface) FeedbackServiceClient {
	return &feedbackServiceClient{cc}
}

func (c *feedbackServiceClient) GetHostReviewById(ctx context.Context, in *RequestReviewById, opts ...grpc.CallOption) (*HostReview, error) {
	out := new(HostReview)
	err := c.cc.Invoke(ctx, "/pb.FeedbackService/GetHostReviewById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedbackServiceClient) GetHostReviewsByUserId(ctx context.Context, in *RequestReviewsByUserId, opts ...grpc.CallOption) (FeedbackService_GetHostReviewsByUserIdClient, error) {
	stream, err := c.cc.NewStream(ctx, &FeedbackService_ServiceDesc.Streams[0], "/pb.FeedbackService/GetHostReviewsByUserId", opts...)
	if err != nil {
		return nil, err
	}
	x := &feedbackServiceGetHostReviewsByUserIdClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FeedbackService_GetHostReviewsByUserIdClient interface {
	Recv() (*HostReview, error)
	grpc.ClientStream
}

type feedbackServiceGetHostReviewsByUserIdClient struct {
	grpc.ClientStream
}

func (x *feedbackServiceGetHostReviewsByUserIdClient) Recv() (*HostReview, error) {
	m := new(HostReview)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *feedbackServiceClient) GetHostReviewsByHostId(ctx context.Context, in *RequestReviewsByHostId, opts ...grpc.CallOption) (FeedbackService_GetHostReviewsByHostIdClient, error) {
	stream, err := c.cc.NewStream(ctx, &FeedbackService_ServiceDesc.Streams[1], "/pb.FeedbackService/GetHostReviewsByHostId", opts...)
	if err != nil {
		return nil, err
	}
	x := &feedbackServiceGetHostReviewsByHostIdClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FeedbackService_GetHostReviewsByHostIdClient interface {
	Recv() (*HostReview, error)
	grpc.ClientStream
}

type feedbackServiceGetHostReviewsByHostIdClient struct {
	grpc.ClientStream
}

func (x *feedbackServiceGetHostReviewsByHostIdClient) Recv() (*HostReview, error) {
	m := new(HostReview)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *feedbackServiceClient) CalcAvgRatingForHost(ctx context.Context, in *RequestAvgRating, opts ...grpc.CallOption) (*AvgRatingResponse, error) {
	out := new(AvgRatingResponse)
	err := c.cc.Invoke(ctx, "/pb.FeedbackService/CalcAvgRatingForHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedbackServiceClient) GetAllAccommodationReviews(ctx context.Context, in *EmptyReviewMsg, opts ...grpc.CallOption) (FeedbackService_GetAllAccommodationReviewsClient, error) {
	stream, err := c.cc.NewStream(ctx, &FeedbackService_ServiceDesc.Streams[2], "/pb.FeedbackService/GetAllAccommodationReviews", opts...)
	if err != nil {
		return nil, err
	}
	x := &feedbackServiceGetAllAccommodationReviewsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FeedbackService_GetAllAccommodationReviewsClient interface {
	Recv() (*AccommodationReview, error)
	grpc.ClientStream
}

type feedbackServiceGetAllAccommodationReviewsClient struct {
	grpc.ClientStream
}

func (x *feedbackServiceGetAllAccommodationReviewsClient) Recv() (*AccommodationReview, error) {
	m := new(AccommodationReview)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *feedbackServiceClient) GetAccommodationReviewById(ctx context.Context, in *RequestReviewById, opts ...grpc.CallOption) (*AccommodationReview, error) {
	out := new(AccommodationReview)
	err := c.cc.Invoke(ctx, "/pb.FeedbackService/GetAccommodationReviewById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedbackServiceClient) GetAccommodationReviewsByUserId(ctx context.Context, in *RequestReviewsByUserId, opts ...grpc.CallOption) (FeedbackService_GetAccommodationReviewsByUserIdClient, error) {
	stream, err := c.cc.NewStream(ctx, &FeedbackService_ServiceDesc.Streams[3], "/pb.FeedbackService/GetAccommodationReviewsByUserId", opts...)
	if err != nil {
		return nil, err
	}
	x := &feedbackServiceGetAccommodationReviewsByUserIdClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FeedbackService_GetAccommodationReviewsByUserIdClient interface {
	Recv() (*AccommodationReview, error)
	grpc.ClientStream
}

type feedbackServiceGetAccommodationReviewsByUserIdClient struct {
	grpc.ClientStream
}

func (x *feedbackServiceGetAccommodationReviewsByUserIdClient) Recv() (*AccommodationReview, error) {
	m := new(AccommodationReview)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *feedbackServiceClient) GetAccommodationReviewsByAccommodationId(ctx context.Context, in *RequestReviewsByAccommodationId, opts ...grpc.CallOption) (FeedbackService_GetAccommodationReviewsByAccommodationIdClient, error) {
	stream, err := c.cc.NewStream(ctx, &FeedbackService_ServiceDesc.Streams[4], "/pb.FeedbackService/GetAccommodationReviewsByAccommodationId", opts...)
	if err != nil {
		return nil, err
	}
	x := &feedbackServiceGetAccommodationReviewsByAccommodationIdClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FeedbackService_GetAccommodationReviewsByAccommodationIdClient interface {
	Recv() (*AccommodationReview, error)
	grpc.ClientStream
}

type feedbackServiceGetAccommodationReviewsByAccommodationIdClient struct {
	grpc.ClientStream
}

func (x *feedbackServiceGetAccommodationReviewsByAccommodationIdClient) Recv() (*AccommodationReview, error) {
	m := new(AccommodationReview)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *feedbackServiceClient) CalcAvgRatingForAccommodation(ctx context.Context, in *RequestAvgRating, opts ...grpc.CallOption) (*AvgRatingResponse, error) {
	out := new(AvgRatingResponse)
	err := c.cc.Invoke(ctx, "/pb.FeedbackService/CalcAvgRatingForAccommodation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedbackServiceServer is the server API for FeedbackService service.
// All implementations must embed UnimplementedFeedbackServiceServer
// for forward compatibility
type FeedbackServiceServer interface {
	GetHostReviewById(context.Context, *RequestReviewById) (*HostReview, error)
	GetHostReviewsByUserId(*RequestReviewsByUserId, FeedbackService_GetHostReviewsByUserIdServer) error
	GetHostReviewsByHostId(*RequestReviewsByHostId, FeedbackService_GetHostReviewsByHostIdServer) error
	CalcAvgRatingForHost(context.Context, *RequestAvgRating) (*AvgRatingResponse, error)
	GetAllAccommodationReviews(*EmptyReviewMsg, FeedbackService_GetAllAccommodationReviewsServer) error
	GetAccommodationReviewById(context.Context, *RequestReviewById) (*AccommodationReview, error)
	GetAccommodationReviewsByUserId(*RequestReviewsByUserId, FeedbackService_GetAccommodationReviewsByUserIdServer) error
	GetAccommodationReviewsByAccommodationId(*RequestReviewsByAccommodationId, FeedbackService_GetAccommodationReviewsByAccommodationIdServer) error
	CalcAvgRatingForAccommodation(context.Context, *RequestAvgRating) (*AvgRatingResponse, error)
	mustEmbedUnimplementedFeedbackServiceServer()
}

// UnimplementedFeedbackServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFeedbackServiceServer struct {
}

func (UnimplementedFeedbackServiceServer) GetHostReviewById(context.Context, *RequestReviewById) (*HostReview, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHostReviewById not implemented")
}
func (UnimplementedFeedbackServiceServer) GetHostReviewsByUserId(*RequestReviewsByUserId, FeedbackService_GetHostReviewsByUserIdServer) error {
	return status.Errorf(codes.Unimplemented, "method GetHostReviewsByUserId not implemented")
}
func (UnimplementedFeedbackServiceServer) GetHostReviewsByHostId(*RequestReviewsByHostId, FeedbackService_GetHostReviewsByHostIdServer) error {
	return status.Errorf(codes.Unimplemented, "method GetHostReviewsByHostId not implemented")
}
func (UnimplementedFeedbackServiceServer) CalcAvgRatingForHost(context.Context, *RequestAvgRating) (*AvgRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalcAvgRatingForHost not implemented")
}
func (UnimplementedFeedbackServiceServer) GetAllAccommodationReviews(*EmptyReviewMsg, FeedbackService_GetAllAccommodationReviewsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllAccommodationReviews not implemented")
}
func (UnimplementedFeedbackServiceServer) GetAccommodationReviewById(context.Context, *RequestReviewById) (*AccommodationReview, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccommodationReviewById not implemented")
}
func (UnimplementedFeedbackServiceServer) GetAccommodationReviewsByUserId(*RequestReviewsByUserId, FeedbackService_GetAccommodationReviewsByUserIdServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAccommodationReviewsByUserId not implemented")
}
func (UnimplementedFeedbackServiceServer) GetAccommodationReviewsByAccommodationId(*RequestReviewsByAccommodationId, FeedbackService_GetAccommodationReviewsByAccommodationIdServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAccommodationReviewsByAccommodationId not implemented")
}
func (UnimplementedFeedbackServiceServer) CalcAvgRatingForAccommodation(context.Context, *RequestAvgRating) (*AvgRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalcAvgRatingForAccommodation not implemented")
}
func (UnimplementedFeedbackServiceServer) mustEmbedUnimplementedFeedbackServiceServer() {}

// UnsafeFeedbackServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeedbackServiceServer will
// result in compilation errors.
type UnsafeFeedbackServiceServer interface {
	mustEmbedUnimplementedFeedbackServiceServer()
}

func RegisterFeedbackServiceServer(s grpc.ServiceRegistrar, srv FeedbackServiceServer) {
	s.RegisterService(&FeedbackService_ServiceDesc, srv)
}

func _FeedbackService_GetHostReviewById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestReviewById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedbackServiceServer).GetHostReviewById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FeedbackService/GetHostReviewById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedbackServiceServer).GetHostReviewById(ctx, req.(*RequestReviewById))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedbackService_GetHostReviewsByUserId_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestReviewsByUserId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FeedbackServiceServer).GetHostReviewsByUserId(m, &feedbackServiceGetHostReviewsByUserIdServer{stream})
}

type FeedbackService_GetHostReviewsByUserIdServer interface {
	Send(*HostReview) error
	grpc.ServerStream
}

type feedbackServiceGetHostReviewsByUserIdServer struct {
	grpc.ServerStream
}

func (x *feedbackServiceGetHostReviewsByUserIdServer) Send(m *HostReview) error {
	return x.ServerStream.SendMsg(m)
}

func _FeedbackService_GetHostReviewsByHostId_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestReviewsByHostId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FeedbackServiceServer).GetHostReviewsByHostId(m, &feedbackServiceGetHostReviewsByHostIdServer{stream})
}

type FeedbackService_GetHostReviewsByHostIdServer interface {
	Send(*HostReview) error
	grpc.ServerStream
}

type feedbackServiceGetHostReviewsByHostIdServer struct {
	grpc.ServerStream
}

func (x *feedbackServiceGetHostReviewsByHostIdServer) Send(m *HostReview) error {
	return x.ServerStream.SendMsg(m)
}

func _FeedbackService_CalcAvgRatingForHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestAvgRating)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedbackServiceServer).CalcAvgRatingForHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FeedbackService/CalcAvgRatingForHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedbackServiceServer).CalcAvgRatingForHost(ctx, req.(*RequestAvgRating))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedbackService_GetAllAccommodationReviews_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EmptyReviewMsg)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FeedbackServiceServer).GetAllAccommodationReviews(m, &feedbackServiceGetAllAccommodationReviewsServer{stream})
}

type FeedbackService_GetAllAccommodationReviewsServer interface {
	Send(*AccommodationReview) error
	grpc.ServerStream
}

type feedbackServiceGetAllAccommodationReviewsServer struct {
	grpc.ServerStream
}

func (x *feedbackServiceGetAllAccommodationReviewsServer) Send(m *AccommodationReview) error {
	return x.ServerStream.SendMsg(m)
}

func _FeedbackService_GetAccommodationReviewById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestReviewById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedbackServiceServer).GetAccommodationReviewById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FeedbackService/GetAccommodationReviewById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedbackServiceServer).GetAccommodationReviewById(ctx, req.(*RequestReviewById))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedbackService_GetAccommodationReviewsByUserId_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestReviewsByUserId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FeedbackServiceServer).GetAccommodationReviewsByUserId(m, &feedbackServiceGetAccommodationReviewsByUserIdServer{stream})
}

type FeedbackService_GetAccommodationReviewsByUserIdServer interface {
	Send(*AccommodationReview) error
	grpc.ServerStream
}

type feedbackServiceGetAccommodationReviewsByUserIdServer struct {
	grpc.ServerStream
}

func (x *feedbackServiceGetAccommodationReviewsByUserIdServer) Send(m *AccommodationReview) error {
	return x.ServerStream.SendMsg(m)
}

func _FeedbackService_GetAccommodationReviewsByAccommodationId_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestReviewsByAccommodationId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FeedbackServiceServer).GetAccommodationReviewsByAccommodationId(m, &feedbackServiceGetAccommodationReviewsByAccommodationIdServer{stream})
}

type FeedbackService_GetAccommodationReviewsByAccommodationIdServer interface {
	Send(*AccommodationReview) error
	grpc.ServerStream
}

type feedbackServiceGetAccommodationReviewsByAccommodationIdServer struct {
	grpc.ServerStream
}

func (x *feedbackServiceGetAccommodationReviewsByAccommodationIdServer) Send(m *AccommodationReview) error {
	return x.ServerStream.SendMsg(m)
}

func _FeedbackService_CalcAvgRatingForAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestAvgRating)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedbackServiceServer).CalcAvgRatingForAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FeedbackService/CalcAvgRatingForAccommodation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedbackServiceServer).CalcAvgRatingForAccommodation(ctx, req.(*RequestAvgRating))
	}
	return interceptor(ctx, in, info, handler)
}

// FeedbackService_ServiceDesc is the grpc.ServiceDesc for FeedbackService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FeedbackService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.FeedbackService",
	HandlerType: (*FeedbackServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHostReviewById",
			Handler:    _FeedbackService_GetHostReviewById_Handler,
		},
		{
			MethodName: "CalcAvgRatingForHost",
			Handler:    _FeedbackService_CalcAvgRatingForHost_Handler,
		},
		{
			MethodName: "GetAccommodationReviewById",
			Handler:    _FeedbackService_GetAccommodationReviewById_Handler,
		},
		{
			MethodName: "CalcAvgRatingForAccommodation",
			Handler:    _FeedbackService_CalcAvgRatingForAccommodation_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetHostReviewsByUserId",
			Handler:       _FeedbackService_GetHostReviewsByUserId_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetHostReviewsByHostId",
			Handler:       _FeedbackService_GetHostReviewsByHostId_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetAllAccommodationReviews",
			Handler:       _FeedbackService_GetAllAccommodationReviews_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetAccommodationReviewsByUserId",
			Handler:       _FeedbackService_GetAccommodationReviewsByUserId_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetAccommodationReviewsByAccommodationId",
			Handler:       _FeedbackService_GetAccommodationReviewsByAccommodationId_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "review.proto",
}
