package common

import (
	"context"

	"github.com/gin-gonic/gin"
)

func GetContext(ctx *gin.Context) context.Context {
	// todo: add logger, request-id here in context
	return context.Background()
}
