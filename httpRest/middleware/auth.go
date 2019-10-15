package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"goweb/httpRest"
	"goweb/mem"
	"goweb/util"
)

func checkId(id string) bool {
	return true
}

func checkUserName(userName string) bool {
	return true
}

func verifyToken(c *gin.Context) {
	if !viper.GetBool("filterToken") {
		c.Next()
	} else {
		token := c.GetHeader("Authorization")
		if token == "" {
			httpRest.SendResponse(c, mem.NotOK, nil)
			logrus.Info(mem.TokenNotExit)
			c.AbortWithStatus(200)
		}

		id, userName, err := util.ParseJwt(token, "")
		if err != nil {
			logrus.Info(err)
			httpRest.SendResponse(c, mem.NotOK, nil)
			c.AbortWithStatus(200)
		}

		if checkId(id) && checkUserName(userName) {
			c.Next()
		} else {
			logrus.Info(mem.UserInfoNotValid)
			httpRest.SendResponse(c, mem.NotOK, nil)
			c.AbortWithStatus(200)
		}
	}
}