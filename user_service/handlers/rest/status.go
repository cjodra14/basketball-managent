package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
)

func GETStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tracer := otel.Tracer("status-handler")
        _, span := tracer.Start(ctx.Request.Context(), "Checking status")
        defer span.End()
		ctx.JSON(http.StatusOK, "online")
	}
}
