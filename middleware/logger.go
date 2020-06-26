package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/luckywb13/blog/pkg/e"
	"github.com/luckywb13/blog/pkg/log"
	"time"
)

func LogHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		method := c.Request.Method

		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		statusCode := c.Writer.Status()
		code := c.GetInt(e.ContextCode)
		clientIP := c.ClientIP()
		if raw != "" {
			path = path + "?" + raw
		}
		log.InfoF("METHOD:%s PATH:%s CODE:%d IP:%s TIME:%d CODE:%d", method, path, statusCode, clientIP, latency/time.Millisecond, code)
	}

}
