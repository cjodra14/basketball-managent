package rest

import (
	"fmt"

	"github.com/cjodra14/basketball-management/user_service/configuration"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init(config configuration.RESTServer) error {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/healthz"},
	}))

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	router.Use(cors.New(corsConfig))

	usersGroup := router.Group("/users")
	UsersRouter(usersGroup)

	infoGroup := router.Group("/")
	InfoRouter(infoGroup)

	serverAddress := fmt.Sprintf("%s:%d", config.Address, config.Port)

	return router.Run(serverAddress)
}
