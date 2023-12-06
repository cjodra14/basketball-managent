package rest

import (
	"github.com/gin-gonic/gin"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
)

func InfoRouter(router *gin.RouterGroup) {
	// router.GET("/healthz", handlers.GETHealthz())

	// router.GET("/doc/", func(c *gin.Context) { c.Redirect(http.StatusSeeOther, "/doc-api/index.html") })
	// router.GET("/doc-api/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, func(c *ginSwagger.Config) {
	// 	c.InstanceName = "telemetry_backend"
	// }))
}

func UsersRouter(router *gin.RouterGroup) {
	router.POST("/" /*handlers.SaveTelemetry(telemetryService)*/)
	router.POST("/login" /*handlers.GetUserTelemetries(telemetryService)*/)
	router.GET("/:id" /*handlers.GetUserTelemetries(telemetryService)*/)
	router.DELETE("/:id" /*handlers.GetUserTelemetries(telemetryService)*/)
	router.PUT("/:id" /*handlers.GetUserTelemetries(telemetryService)*/)
}
