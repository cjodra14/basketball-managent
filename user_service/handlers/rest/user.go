package rest

import (
	"net/http"

	"github.com/cjodra14/basketball-management/user_service/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
)

func GetUser(userService services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		tracer := otel.Tracer("get-user-handler")
        _, span := tracer.Start(c.Request.Context(), "getting user")
        defer span.End()
		userID := c.Param("id")
		user, err := userService.Get(c, userID)
		if err != nil {
			if err := c.AbortWithError(http.StatusInternalServerError, err); err != nil {
				logrus.Debug(err)
			}

			return
		}

		c.JSON(http.StatusOK, user)
	}
}
