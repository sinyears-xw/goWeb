package router

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	_ "goweb/httpRest/test"
	"goweb/mem"
	"net/http"
)

func init()  {
	if viper.GetBool("http") {
		port := viper.GetString("httpPort")
		if port == "" {
			panic(mem.HttpPortNotValid)
		}

		if mem.G == nil {
			panic(mem.GinNotValid)
		}

		logrus.Print("HttpRest", port)
		mem.HttpServer = &http.Server{Addr: port, Handler: mem.G}
		go mem.HttpServer.ListenAndServe()
	}

}