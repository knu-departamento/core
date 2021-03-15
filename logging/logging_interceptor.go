package logging

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
	Log logrus.FieldLogger
}

func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Internal, "Could not retrieve metadata from the request")
	}
	service := path.Dir(info.FullMethod)[1:]
	method := path.Base(info.FullMethod)
	traceId := "null"
	if len(md.Get("trace_id")) > 0 {
		traceId = md.Get("trace_id")[0]
	}
	log := logrus.WithFields(logrus.Fields{
		"grpc.service": service,
		"grpc.method":  method,
		"trace_id":     traceId,
	})

	log.Infof("Received %s request", method)
	newCtx := context.WithValue(ctx, CtxLoggerMarker{}, CtxLogger{log})
	res, err := handler(newCtx, req)
	if err == nil {
		log.Infof("Request %s finished with code %d", method, codes.OK)
	} else {
		if code, ok := status.FromError(err); ok {
			log.Errorf("Request %s finished with code %d: %s", method, code.Code(), code.Message())
		}
	}
	return res, err
}
