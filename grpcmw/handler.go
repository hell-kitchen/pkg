package grpcmw

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type requestIDKey struct{}

func UnaryServerInterceptor(opt ...Option) grpc.UnaryServerInterceptor {
	var opts options
	opts.validator = defaultReqeustIDValidator
	for _, o := range opt {
		o.apply(&opts)
	}

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		var requestID string
		if opts.chainRequestID {
			requestID = HandleRequestIDChain(ctx, opts.validator)
		} else {
			requestID = HandleRequestID(ctx, opts.validator)
		}
		ctx = context.WithValue(ctx, requestIDKey{}, requestID)
		return handler(ctx, req)
	}
}

func FromContext(ctx context.Context) string {
	id, ok := ctx.Value(requestIDKey{}).(string)
	if !ok {
		return ""
	}
	return id
}
