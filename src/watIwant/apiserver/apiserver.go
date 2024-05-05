package apiserver

import (
	"github.com/gin-gonic/gin"
)

func Handlers() *gin.Engine {
	engine := gin.Default()
	V1Routes(engine)
	return engine
}
