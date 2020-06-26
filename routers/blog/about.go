package blog

import (
	"github.com/gin-gonic/gin"
	"github.com/luckywb13/blog/conf"
	"net/http"
)

func AboutIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "blog/about/index.html", gin.H{"title": conf.WebName, "highlight": "about"})
}
