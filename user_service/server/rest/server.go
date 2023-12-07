package rest

import (
	"fmt"

	"github.com/cjodra14/basketball-management/user_service/configuration"
	"github.com/cjodra14/basketball-management/user_service/services"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init(config configuration.RESTServer, userService services.UserService) error {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(otelgin.Middleware("user-service"))
	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/status"},
	}))

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	router.Use(cors.New(corsConfig))

	usersGroup := router.Group("/users")
	UsersRouter(usersGroup, userService)

	infoGroup := router.Group("/")
	InfoRouter(infoGroup)

	serverAddress := fmt.Sprintf("%s:%d", config.Address, config.Port)

	return router.Run(serverAddress)
}
