package main

import (
	"github.com/cjodra14/basketball-management/user_service/configuration"
	grpc "github.com/cjodra14/basketball-management/user_service/server/grpc"
	"github.com/cjodra14/basketball-management/user_service/server/rest"
	"github.com/cjodra14/basketball-management/user_service/services/users"
	"github.com/cjodra14/basketball-management/user_service/storage/postgres"
	"github.com/cjodra14/basketball-management/user_service/tracing"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"

	"github.com/kelseyhightower/envconfig"
	"github.com/oklog/run"
	"github.com/sirupsen/logrus"
)

func main() {

	config := configuration.Configuration{}

	err := envconfig.Process("", &config)
	if err != nil {
		logrus.Fatal(err.Error())
	}

	logrus.SetLevel(configuration.SetLogLevel(config.Logs))

	logrus.Debug(config)

	tp, tpErr := tracing.JaegerTraceProvider()
	if tpErr != nil {
		logrus.Fatal(tpErr)
	}
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	storage, err := postgres.NewPostgresUserStorage(config.Storage.PostgresConfiguration)
	if err != nil {
		logrus.Fatal(err)
	}
	userService := users.New(storage)

	var group run.Group

	group.Add(func() error {

		return rest.Init(config.RESTServer, userService)
	}, func(e error) {
		logrus.Fatal(e)
	})

	group.Add(func() error {
		return grpc.InitUserServiceServer(config.GRPCServer, userService)
	}, func(e error) {
		logrus.Fatal(e)
	})

	if err := group.Run(); err != nil {
		logrus.Fatal(err)
	}
}
