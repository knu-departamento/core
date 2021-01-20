package interceptors

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"path"
)

type CtxLoggerMarker struct{}

type CtxLogger struct {
	Log *logrus.Entry
}

func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Internal, "Could not retrieve metadata from the request")
	}
	service := path.Dir(info.FullMethod)[1:]
	method := path.Base(info.FullMethod)
	log := logrus.WithFields(logrus.Fields{
		"grpc.service": service,
		"grpc.method":  method,
		"trace_id":     md.Get("trace_id")[0],
	})

	newCtx := context.WithValue(ctx, CtxLoggerMarker{}, CtxLogger{log})
	return handler(newCtx, req)
}
