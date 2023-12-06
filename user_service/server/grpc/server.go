package server

import (
	// "context"

	// grpcAPI "github.com/cjodra14/basketball-management/api/client/grpc"
	// apiModels "github.com/cjodra14/telemetry_backend/api/models"
	// "github.com/cjodra14/telemetry_backend/services"
	// grpcConfig "github.com/cjodra14/telemetry_commons/configuration/grpc"
	// commonsGrpc "github.com/cjodra14/telemetry_commons/services/grpc"
	"github.com/pkg/errors"
)


var (
	ErrInternalCellGrpcHandler       = errors.New("internal error")
	ErrCellGrpcHandlerCreatingServer = errors.New("error creating grpc server")
)

type UserServer struct {
	// grpcAPI.
	// telemetryService services.TelemetryService
}

// func newTelemetryServerGrpc() grpcAPI.TelemetryServiceServer {
// 	return &TelemetryServer{
// 		telemetryService: telemetryService,
// 	}
// }

// func InitTelemetryServiceServer(configuration grpcConfig.GRPCServerConfiguration, telemetryService services.TelemetryService) error {
// 	grpcServer, err := commonsGrpc.Server(configuration, commonsGrpc.NewGrpcServerInterceptor())
// 	if err != nil {
// 		return err
// 	}

// 	grpcService := newTelemetryServerGrpc(telemetryService)
// 	grpcServer.RegisterService(&grpcAPI.TelemetryService_ServiceDesc, grpcService)

// 	err = grpcServer.Serve()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return nil
// }

// func (telemetryServer *TelemetryServer) SaveTelemetry(ctx context.Context, in *grpcAPI.Telemetry) (*grpcAPI.TelemetryResponse, error) {
// 	err := telemetryServer.telemetryService.SaveTelemetry(ctx, apiModels.Telemetry{
// 		Timestamp: in.Timestamp,
// 		UserID:    in.UserId,
// 		GPS: apiModels.GPS{
// 			Latitude:            in.Gps.Latitude,
// 			Longitude:           in.Gps.Longitude,
// 			Speed:               in.Gps.Speed,
// 			Direction:           in.Gps.Direction,
// 			ConnectedSatellites: in.Gps.ConnectedSatellites,
// 		},
// 		Accelerometer: apiModels.Accelerometer{
// 			AngleX: in.Accelerometer.AngleX,
// 			AngleY: in.Accelerometer.AngleY,
// 			AngleZ: in.Accelerometer.AngleZ,
// 			GForce: in.Accelerometer.GForce,
// 		},
// 	})
// 	if err != nil {
// 		return &grpcAPI.TelemetryResponse{}, err
// 	}

// 	return &grpcAPI.TelemetryResponse{}, nil
// }

// func (telemetryServer *TelemetryServer) GetUserTelemetry(ctx context.Context, in *grpcAPI.UserTelemetryRequest) (*grpcAPI.UserTelemetryResponse, error) {
// 	telemetries, err := telemetryServer.telemetryService.GetuserTelemetries(ctx, in.UserId)
// 	if err != nil {
// 		return &grpcAPI.UserTelemetryResponse{}, err
// 	}

// 	grpcTelemetries := []*grpcAPI.Telemetry{}

// 	for _, telemetry := range telemetries {
// 		grpcTelemetries = append(grpcTelemetries, &grpcAPI.Telemetry{
// 			Timestamp: telemetry.Timestamp,
// 			UserId:    telemetry.UserID,
// 			Gps: &grpcAPI.GPS{
// 				Latitude:            telemetry.GPS.Latitude,
// 				Longitude:           telemetry.GPS.Longitude,
// 				Speed:               telemetry.GPS.Speed,
// 				Direction:           telemetry.GPS.Direction,
// 				ConnectedSatellites: telemetry.GPS.ConnectedSatellites,
// 			},
// 			Accelerometer: &grpcAPI.Accelerometer{
// 				AngleX: telemetry.Accelerometer.AngleX,
// 				AngleY: telemetry.Accelerometer.AngleY,
// 				AngleZ: telemetry.Accelerometer.AngleZ,
// 				GForce: telemetry.Accelerometer.GForce,
// 			},
// 		})
// 	}

// 	return &grpcAPI.UserTelemetryResponse{
// 		Telemetries: grpcTelemetries,
// 	}, nil
// }