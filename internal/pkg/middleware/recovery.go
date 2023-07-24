package middleware

import (
	"net/http"
	"xsis/assignment/internal/pkg/handler"
	"xsis/assignment/internal/pkg/reason"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			err := recover()
			if err != nil {
				logrus.Error("error message: ", err)
				handler.ResponseError(ctx, http.StatusInternalServerError, reason.InternalServerError)
			}
		}()

		ctx.Next()
	}
}
