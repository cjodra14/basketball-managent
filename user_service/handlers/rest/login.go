package rest

import (
	"net/http"

	"github.com/cjodra14/basketball-management/user_service/api/models"
	"github.com/cjodra14/basketball-management/user_service/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
)

func Login(userService services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		tracer := otel.Tracer("login-handler")
		_, span := tracer.Start(c.Request.Context(), "login user")
		defer span.End()

		user := models.UserLogin{}
		if err := c.BindJSON(&user); err != nil {
			if err := c.AbortWithError(http.StatusBadRequest, err); err != nil {
				logrus.Debug(err)
			}

			return
		}

		if err := userService.Login(c, user); err != nil {
			if err := c.AbortWithError(http.StatusBadRequest, err); err != nil {
				logrus.Debug(err)
			}

			return
		}

		c.Status(http.StatusOK)
	}
}
