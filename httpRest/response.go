package httpRest

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"goweb/util"
	"io/ioutil"
	"net/http"
	"strings"
)
import "goweb/mem"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err *mem.Err, data interface{}) {
	code, message := err.DecodeErr()
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func GetResponse(method string, url string, content map[string]interface{}, auth ...string) (rp Response, err error) {
	rp = Response{}
	err = nil
	client := &http.Client{}
	url = viper.GetString("httpServer") + viper.GetString("httpPort") + url
	req, err := http.NewRequest(method, url, strings.NewReader(util.Map2Json(content)))

	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if auth != nil {
		req.Header.Set("Authorization", auth[0])
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &rp)
	if err != nil {
		return
	}
	return
}