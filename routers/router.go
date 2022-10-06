package routers

import (
	"github.com/gin-gonic/gin"
)

type RouterContract interface {
	Setup()
}

func NewDefaultRouter() *gin.Engine {
	return gin.Default()
}

// noinspection GoUnusedExportedFunction
func NewEmptyRouter() *gin.Engine {
	return gin.New()
}
