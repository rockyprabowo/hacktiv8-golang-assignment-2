package routers

import (
	"github.com/gin-gonic/gin"
)

// RouterContract defines a contract for the route definitions.
type RouterContract interface {
	Setup()
}

// NewDefaultRouter instantiates a default Gin engine.
func NewDefaultRouter() *gin.Engine {
	return gin.Default()
}

// NewEmptyRouter instantiates an empty Gin engine.
// TODO: Remove this maybe?
// noinspection GoUnusedExportedFunction
func NewEmptyRouter() *gin.Engine {
	return gin.New()
}
