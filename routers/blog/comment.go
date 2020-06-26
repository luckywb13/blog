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
	"html"
	"net/http"
)

func CommentIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "blog/comment/index.html", gin.H{"title": conf.WebName, "highlight": "comment"})
}

func AddComment(c *gin.Context) {
	var args struct {
		Email     string `form:"email" binding:"email" label:"邮箱"`
		Content   string `form:"content" binding:"min=10" label:"评论内容"`
		Name      string `form:"name" binding:"required" label:"姓名"`
		ArticleId int    `form:"article_id" binding:"min=1" label:"文章编号"`
	}

	if err := c.ShouldBind(&args); err != nil {
		resp.GetResponse(c, e.INVALID_PARAMS, util.Translate(err.(validator.ValidationErrors)), nil)
		return
	}

	err := models.AddComment(args.ArticleId, args.Email, args.Name, html.EscapeString(args.Content), c.Request.RemoteAddr)
	if err != nil {
		log.Error(err.Error())
		resp.GetResponse(c, e.ERROR, "", nil)
		return
	}

	resp.GetResponse(c, e.SUCCESS, "", nil)
	return
}

func GetCommentListForPage(c *gin.Context) {
	var args struct {
		PageIndex int `form:"page_index" binding:"min=1"`
		PageSize  int `form:"page_size" binding:"eq=5"`
		ArticleId int `form:"article_id" binding:"min=1"`
	}

	if err := c.ShouldBindQuery(&args); err != nil {
		resp.GetResponse(c, e.INVALID_PARAMS, util.Translate(err.(validator.ValidationErrors)), nil)
		return
	}

	comments, err := models.GetCommentListForPage(args.PageIndex, args.PageSize, args.ArticleId)
	if err != nil {
		log.Error(err.Error())
		resp.GetResponse(c, e.ERROR, "", nil)
		return
	}
	resp.GetResponse(c, e.SUCCESS, "", comments)
	return
}
