package resp

import (
	"github.com/gin-gonic/gin"
	"github.com/luckywb13/blog/pkg/e"
	"net/http"
)

type Response struct {
	Code     int         `json:"code"`
	Message  string      `json:"msg"`
	Data     interface{} `json:"data,omitempty"`
}

func GetResponse(c *gin.Context, code int, msg string, data interface{}){
	resp := Response{Code:code, Data:data}

	if len(msg) > 0 {
		resp.Message = msg
	} else {
		if msg, ok := e.MsgFlags[code]; ok {
			resp.Message = msg
		} else{
			resp.Message = e.MsgFlags[e.ERROR]
		}
	}

	c.Set(e.ContextCode, code)
	c.JSON(http.StatusOK, resp)
	return
}