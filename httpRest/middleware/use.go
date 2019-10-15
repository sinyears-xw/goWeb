package middleware

import (
	"github.com/gin-gonic/gin"
)

func UseDefault(rg *gin.RouterGroup) {
	rg.Use(noCache)
	rg.Use(options)
	rg.Use(secure)
	rg.Use(gin.Recovery())
}

func UseAuth(rg *gin.RouterGroup) {
	rg.Use(verifyToken)
}
