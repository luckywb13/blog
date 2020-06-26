package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/luckywb13/blog/pkg/log"
	"github.com/luckywb13/blog/pkg/util"
	"net/http"
	"time"
)

func AuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var token *http.Cookie

		if token, err = c.Request.Cookie("authorization"); err == nil {
			value := token.Value

			claims, err := util.ParseToken(value)
			if err != nil {
				log.Error(err.Error())
				c.Redirect(http.StatusFound, "admin/login")
				return
			}

			if time.Now().Unix() > claims.ExpiresAt {
				c.Redirect(http.StatusFound, "admin/login")
				return
			}

			c.Set("id", claims.Id)
			c.Set("username", claims.Username)
			c.Next()
			return
		}
		c.Redirect(http.StatusFound, "/admin/login")
		return
	}
}
