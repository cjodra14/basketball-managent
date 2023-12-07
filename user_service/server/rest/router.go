package rest

import (
	"github.com/cjodra14/basketball-management/user_service/handlers/rest"
	"github.com/cjodra14/basketball-management/user_service/services"
	"github.com/gin-gonic/gin"
)

func InfoRouter(router *gin.RouterGroup) {
	router.GET("/status", rest.GETStatus())
}

func UsersRouter(router *gin.RouterGroup, userService services.UserService) {
	router.POST("/register", rest.RegisterUser(userService))
	router.POST("/login", rest.Login(userService))
	router.GET("/:id", rest.GetUser(userService))
}
