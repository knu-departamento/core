package interceptors

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AllowedCredentials map[string]string

func AuthInterceptor(allowedCredentials AllowedCredentials) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Internal, "Could not retrieve metadata from the request")
		}
		logger := logrus.WithField("trace_id", md.Get("trace_id")[0])

		clientIdCollection, foundId := md["client_id"]
		if !foundId {
			return nil, status.Error(codes.Unauthenticated, "Request does not contain client id as metadata")
		}
		clientId := clientIdCollection[0]

		clientKeyCollection, foundKey := md["client_key"]
		if !foundKey {
			logger.WithField("client_id", clientId).Error()
			return nil, status.Error(codes.Unauthenticated, "Request does not contain client key as metadata")
		}
		clientKey := clientKeyCollection[0]

		authenticated := allowedCredentials[clientId] == clientKey
		if !authenticated {
			logger.WithField("client_id", clientId).Warn("Invalid attempt to authenticate!")
			return nil, status.Error(codes.Unauthenticated, "Invalid credentials")
		}

		return handler(ctx, req)
	}
}
