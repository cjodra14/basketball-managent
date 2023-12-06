package main

import (
	"github.com/cjodra14/basketball-management/user_service/configuration"
	"github.com/cjodra14/basketball-management/user_service/server/rest"
	"fmt"

	"github.com/gin-gonic/gin"
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

	router := gin.Default()

	address := fmt.Sprintf("%s:%s", config.RESTServer.Address, config.RESTServer.Port)

	if err := router.Run(address); err != nil {
		logrus.Fatal(err)
	}

	var group run.Group

	group.Add(func() error {

		return rest.Init(config.RESTServer)
	}, func(e error) {
		logrus.Fatal(e)
	})

	// group.Add(func() error {
	// 	return grpcServer.InitTelemetryServiceServer(conf.GRPCServer, telemetryService)
	// }, func(e error) {
	// 	log.Fatal(e)
	// })

	if err := group.Run(); err != nil {
		logrus.Fatal(err)
	}
}
