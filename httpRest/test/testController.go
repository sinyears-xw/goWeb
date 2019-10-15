package test

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	. "goweb/httpRest"
	"goweb/httpRest/middleware"
	"goweb/mem"
)

type out1 struct {
	Name string `json:"name"`
	Code int    `json:"code"`
}

type in1 struct {
	Args1 []int `json:"args1" binding:"required"`
	Args2 int   `json:"args2"`
}

func init() {
	if mem.G != nil {
		rg := mem.G.Group("/test")
		middleware.UseDefault(rg)
		middleware.UseAuth(rg)
		{
			rg.GET("/a", a)
			rg.POST("/b", b)
			rg.GET("/c/:name", c1)
		}
	}
}

func a(c *gin.Context) {
	SendResponse(c, mem.OK, out1{
		Name: "xxx",
		Code: 0,
	})
}

func b(c *gin.Context) {
	var params in1

	err := c.ShouldBindJSON(&params)
	if err != nil {
		logrus.Error(err)
		SendResponse(c, mem.NotOK, nil)
		return
	}
	SendResponse(c, mem.OK, s2(params.Args1))
}

func c1(c *gin.Context) {
	name := c.Param("name")
	nameBytes := s1(name)
	SendResponse(c, mem.OK, nameBytes)
}
