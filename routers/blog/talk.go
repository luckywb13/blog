package blog

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/luckywb13/blog/conf"
	"github.com/luckywb13/blog/models"
	"github.com/luckywb13/blog/pkg/e"
	"github.com/luckywb13/blog/pkg/log"
	"github.com/luckywb13/blog/pkg/resp"
	"github.com/luckywb13/blog/pkg/util"
	"math"
	"net/http"
)

var defaultPageNum = 10

func TalkIndex(c *gin.Context) {
	tmp, err := models.GetTalkTotal()
	if err != nil {
		log.Error(err.Error())
	}

	total := math.Ceil(float64(tmp) / float64(defaultPageNum))
	c.HTML(http.StatusOK, "blog/talk/index.html", gin.H{"title": conf.WebName, "highlight": "talk", "total": total})
}

func GetTalkListForPage(c *gin.Context) {
	var args struct {
		PageIndex int `form:"page_index" binding:"min=1"`
		PageSize  int `form:"page_size" binding:"min=10"`
	}
	if err := c.ShouldBindQuery(&args); err != nil {
		resp.GetResponse(c, e.INVALID_PARAMS, util.Translate(err.(validator.ValidationErrors)), nil)
		return
	}

	_, talks, err := models.GetTalkListForPage(args.PageIndex, args.PageSize)
	if err != nil {
		log.Error(err.Error())
		resp.GetResponse(c, e.ERROR, "", talks)
		return
	}

	resp.GetResponse(c, e.SUCCESS, "", talks)
	return
}
