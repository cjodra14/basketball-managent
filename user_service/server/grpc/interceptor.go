package grpc

import (
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type GrpcServerInterceptor interface {
	AddServerUnaryInterceptors(interceptors []grpc.UnaryServerInterceptor)
	AddServerStreamInterceptors(interceptors []grpc.StreamServerInterceptor)
	GetServerUnaryInterceptors() []grpc.UnaryServerInterceptor
	GetServerStreamInterceptors() []grpc.StreamServerInterceptor
}

var _ GrpcServerInterceptor = (*serverInterceptor)(nil)

type serverInterceptor struct {
	grpcUnaryServerInterceptors  []grpc.UnaryServerInterceptor
	grpcStreamServerInterceptors []grpc.StreamServerInterceptor
}

func NewGrpcServerInterceptor() *serverInterceptor {
	entry := logrus.NewEntry(logrus.New())
	grpc_logrus.ReplaceGrpcLogger(entry)

	return &serverInterceptor{
		// grpcUnaryServerInterceptors: []grpc.UnaryServerInterceptor{
		// 	grpc_auth.UnaryServerInterceptor(grpcCustomMiddlewares.TraceabilityMiddleware),
		// 	grpcCustomMiddlewares.LogUnaryInterceptor(entry),
		// },
		// grpcStreamServerInterceptors: []grpc.StreamServerInterceptor{
		// 	grpc_auth.StreamServerInterceptor(grpcCustomMiddlewares.TraceabilityMiddleware),
		// 	grpcCustomMiddlewares.LogStreamInterceptor(entry),
		// },
	}
}

func (a *serverInterceptor) AddServerUnaryInterceptors(interceptors []grpc.UnaryServerInterceptor) {
	a.grpcUnaryServerInterceptors = append(a.grpcUnaryServerInterceptors, interceptors...)
}

func (a *serverInterceptor) AddServerStreamInterceptors(interceptors []grpc.StreamServerInterceptor) {
	a.grpcStreamServerInterceptors = append(a.grpcStreamServerInterceptors, interceptors...)
}

func (a *serverInterceptor) GetServerUnaryInterceptors() []grpc.UnaryServerInterceptor {
	return a.grpcUnaryServerInterceptors
}

func (a *serverInterceptor) GetServerStreamInterceptors() []grpc.StreamServerInterceptor {
	return a.grpcStreamServerInterceptors
}