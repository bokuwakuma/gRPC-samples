package interceptor

import (
	"context"

	"github.com/bokuwakuma/grpc-samples/front/support"
	"github.com/bokuwakuma/grpc-samples/shared/md"
	"google.golang.org/grpc"
)

func XTraceID(ctx context.Context,
	method string,
	req, reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption) error {
	traceID := support.GetTraceIDFromContext(ctx)
	ctx = md.AddTraceIDToContext(ctx, traceID)
	return invoker(ctx, method, req, reply, cc, opts...)
}

func XUserID(ctx context.Context,
	method string,
	req, reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption) error {
	user := support.GetUserFromContext(ctx)
	ctx = md.AddUserIDToContext(ctx, user.Id)
	return invoker(ctx, method, req, reply, cc, opts...)
}
