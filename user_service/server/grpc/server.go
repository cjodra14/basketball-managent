package grpc

import (
	// "context"

	"context"

	grpcAPI "github.com/cjodra14/basketball-management/user_service/api/client/gRPC"
	"github.com/cjodra14/basketball-management/user_service/configuration"
	"github.com/cjodra14/basketball-management/user_service/services"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var (
	ErrInternalCellGrpcHandler       = errors.New("internal error")
	ErrCellGrpcHandlerCreatingServer = errors.New("error creating grpc server")
)

type UserServer struct {
	grpcAPI.UnimplementedUserServiceServer
	userService services.UserService
}

func newUserServerGrpc(userService services.UserService) grpcAPI.UserServiceServer {
	return &UserServer{
		userService: userService,
	}
}

func InitUserServiceServer(configuration configuration.GRPCServer, userService services.UserService) error {
	grpcServer, err := Server(configuration, NewGrpcServerInterceptor())
	if err != nil {
		return err
	}

	grpcService := newUserServerGrpc(userService)
	grpcServer.RegisterService(&grpcAPI.UserService_ServiceDesc, grpcService)

	err = grpcServer.Serve()
	if err != nil {
		logrus.Fatal(err)
	}

	return nil
}

func (userServer *UserServer) GetUser(ctx context.Context, in *grpcAPI.GetUserRequest) (*grpcAPI.UserResponse, error) {
	user, err := userServer.userService.Get(ctx, in.Id)
	if err != nil {
		return &grpcAPI.UserResponse{}, err
	}

	return &grpcAPI.UserResponse{
		Id:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	}, nil
}

func (userServer *UserServer) CreateUser(ctx context.Context, in *grpcAPI.CreateUserRequest) (*grpcAPI.UserResponse, error) {
	return &grpcAPI.UserResponse{}, nil
}

func (userServer *UserServer) AuthenticateUser(ctx context.Context, in *grpcAPI.AuthenticateRequest) (*grpcAPI.TokenResponse, error) {
	return &grpcAPI.TokenResponse{}, nil
}
