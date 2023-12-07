package grpc

import (
	"fmt"
	"net"

	"github.com/cjodra14/basketball-management/user_service/configuration"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type grpcServer struct {
	server   *grpc.Server
	listener net.Listener
}

func Server(config configuration.GRPCServer, grpcServerInterceptor GrpcServerInterceptor) (*grpcServer, error) {
	unaryServerInterceptors := grpcServerInterceptor.GetServerUnaryInterceptors()
	streamServerInterceptors := grpcServerInterceptor.GetServerStreamInterceptors()

	grpcCreds := insecure.NewCredentials()

	options := []grpc.ServerOption{
		grpc.Creds(grpcCreds),
		grpc_middleware.WithUnaryServerChain(unaryServerInterceptors...),
		grpc_middleware.WithStreamServerChain(streamServerInterceptors...),
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Address, config.Port))
	if err != nil {
		log.Error("grpcClientServerRunner: unable to create a tcp listener for the control server", err)

		return nil, err
	}

	return &grpcServer{
		server:   grpc.NewServer(options...),
		listener: listener,
	}, nil
}

func (gs *grpcServer) RegisterService(desc *grpc.ServiceDesc, srv interface{}) {
	gs.server.RegisterService(desc, srv)
}

func (gs *grpcServer) Serve() error {
	return gs.server.Serve(gs.listener)
}
