// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: EventService.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CalendarService_CreateEvent_FullMethodName    = "/event.CalendarService/CreateEvent"
	CalendarService_UpdateEvent_FullMethodName    = "/event.CalendarService/UpdateEvent"
	CalendarService_DeleteEvent_FullMethodName    = "/event.CalendarService/DeleteEvent"
	CalendarService_GetDayEvents_FullMethodName   = "/event.CalendarService/GetDayEvents"
	CalendarService_GetWeekEvents_FullMethodName  = "/event.CalendarService/GetWeekEvents"
	CalendarService_GetMonthEvents_FullMethodName = "/event.CalendarService/GetMonthEvents"
)

// CalendarServiceClient is the client API for CalendarService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalendarServiceClient interface {
	CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*CreateEventResponse, error)
	UpdateEvent(ctx context.Context, in *UpdateEventRequest, opts ...grpc.CallOption) (*UpdateEventResponse, error)
	DeleteEvent(ctx context.Context, in *DeleteEventRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetDayEvents(ctx context.Context, in *GetDayEventsRequest, opts ...grpc.CallOption) (*GetDayEventsResponse, error)
	GetWeekEvents(ctx context.Context, in *GetWeekEventsRequest, opts ...grpc.CallOption) (*GetWeekEventsResponse, error)
	GetMonthEvents(ctx context.Context, in *GetMonthEventsRequest, opts ...grpc.CallOption) (*GetMonthEventsResponse, error)
}

type calendarServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalendarServiceClient(cc grpc.ClientConnInterface) CalendarServiceClient {
	return &calendarServiceClient{cc}
}

func (c *calendarServiceClient) CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*CreateEventResponse, error) {
	out := new(CreateEventResponse)
	err := c.cc.Invoke(ctx, CalendarService_CreateEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarServiceClient) UpdateEvent(ctx context.Context, in *UpdateEventRequest, opts ...grpc.CallOption) (*UpdateEventResponse, error) {
	out := new(UpdateEventResponse)
	err := c.cc.Invoke(ctx, CalendarService_UpdateEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarServiceClient) DeleteEvent(ctx context.Context, in *DeleteEventRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CalendarService_DeleteEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarServiceClient) GetDayEvents(ctx context.Context, in *GetDayEventsRequest, opts ...grpc.CallOption) (*GetDayEventsResponse, error) {
	out := new(GetDayEventsResponse)
	err := c.cc.Invoke(ctx, CalendarService_GetDayEvents_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarServiceClient) GetWeekEvents(ctx context.Context, in *GetWeekEventsRequest, opts ...grpc.CallOption) (*GetWeekEventsResponse, error) {
	out := new(GetWeekEventsResponse)
	err := c.cc.Invoke(ctx, CalendarService_GetWeekEvents_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarServiceClient) GetMonthEvents(ctx context.Context, in *GetMonthEventsRequest, opts ...grpc.CallOption) (*GetMonthEventsResponse, error) {
	out := new(GetMonthEventsResponse)
	err := c.cc.Invoke(ctx, CalendarService_GetMonthEvents_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalendarServiceServer is the server API for CalendarService service.
// All implementations must embed UnimplementedCalendarServiceServer
// for forward compatibility
type CalendarServiceServer interface {
	CreateEvent(context.Context, *CreateEventRequest) (*CreateEventResponse, error)
	UpdateEvent(context.Context, *UpdateEventRequest) (*UpdateEventResponse, error)
	DeleteEvent(context.Context, *DeleteEventRequest) (*emptypb.Empty, error)
	GetDayEvents(context.Context, *GetDayEventsRequest) (*GetDayEventsResponse, error)
	GetWeekEvents(context.Context, *GetWeekEventsRequest) (*GetWeekEventsResponse, error)
	GetMonthEvents(context.Context, *GetMonthEventsRequest) (*GetMonthEventsResponse, error)
	mustEmbedUnimplementedCalendarServiceServer()
}

// UnimplementedCalendarServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCalendarServiceServer struct {
}

func (UnimplementedCalendarServiceServer) CreateEvent(context.Context, *CreateEventRequest) (*CreateEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}
func (UnimplementedCalendarServiceServer) UpdateEvent(context.Context, *UpdateEventRequest) (*UpdateEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEvent not implemented")
}
func (UnimplementedCalendarServiceServer) DeleteEvent(context.Context, *DeleteEventRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEvent not implemented")
}
func (UnimplementedCalendarServiceServer) GetDayEvents(context.Context, *GetDayEventsRequest) (*GetDayEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDayEvents not implemented")
}
func (UnimplementedCalendarServiceServer) GetWeekEvents(context.Context, *GetWeekEventsRequest) (*GetWeekEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWeekEvents not implemented")
}
func (UnimplementedCalendarServiceServer) GetMonthEvents(context.Context, *GetMonthEventsRequest) (*GetMonthEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMonthEvents not implemented")
}
func (UnimplementedCalendarServiceServer) mustEmbedUnimplementedCalendarServiceServer() {}

// UnsafeCalendarServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalendarServiceServer will
// result in compilation errors.
type UnsafeCalendarServiceServer interface {
	mustEmbedUnimplementedCalendarServiceServer()
}

func RegisterCalendarServiceServer(s grpc.ServiceRegistrar, srv CalendarServiceServer) {
	s.RegisterService(&CalendarService_ServiceDesc, srv)
}

func _CalendarService_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServiceServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalendarService_CreateEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServiceServer).CreateEvent(ctx, req.(*CreateEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarService_UpdateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServiceServer).UpdateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalendarService_UpdateEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServiceServer).UpdateEvent(ctx, req.(*UpdateEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarService_DeleteEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServiceServer).DeleteEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalendarService_DeleteEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServiceServer).DeleteEvent(ctx, req.(*DeleteEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarService_GetDayEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDayEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServiceServer).GetDayEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalendarService_GetDayEvents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServiceServer).GetDayEvents(ctx, req.(*GetDayEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarService_GetWeekEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWeekEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServiceServer).GetWeekEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalendarService_GetWeekEvents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServiceServer).GetWeekEvents(ctx, req.(*GetWeekEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarService_GetMonthEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMonthEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServiceServer).GetMonthEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalendarService_GetMonthEvents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServiceServer).GetMonthEvents(ctx, req.(*GetMonthEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CalendarService_ServiceDesc is the grpc.ServiceDesc for CalendarService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalendarService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "event.CalendarService",
	HandlerType: (*CalendarServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEvent",
			Handler:    _CalendarService_CreateEvent_Handler,
		},
		{
			MethodName: "UpdateEvent",
			Handler:    _CalendarService_UpdateEvent_Handler,
		},
		{
			MethodName: "DeleteEvent",
			Handler:    _CalendarService_DeleteEvent_Handler,
		},
		{
			MethodName: "GetDayEvents",
			Handler:    _CalendarService_GetDayEvents_Handler,
		},
		{
			MethodName: "GetWeekEvents",
			Handler:    _CalendarService_GetWeekEvents_Handler,
		},
		{
			MethodName: "GetMonthEvents",
			Handler:    _CalendarService_GetMonthEvents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "EventService.proto",
}
