package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/luckywb13/blog/models"
	"github.com/luckywb13/blog/pkg/e"
	"github.com/luckywb13/blog/pkg/log"
	"github.com/luckywb13/blog/pkg/resp"
	"github.com/luckywb13/blog/pkg/util"
	"html"
	"net/http"
	"strconv"
)

func ArticleIndex(c *gin.Context) {
	userName, _ := c.Get("username")
	c.HTML(http.StatusOK, "admin/article/index.html", gin.H{"title": "文章管理", "highlight": "article", "username": userName})
}

func AddArticleIndex(c *gin.Context) {
	userName, _ := c.Get("username")
	categoryList, err := models.GetCategoryList(models.CategoryNORMAL)

	if err != nil {
		log.Error(err.Error())
	}

	c.HTML(http.StatusOK, "admin/article/add.html", gin.H{"title": "文章管理", "highlight": "article", "username": userName, "categorys": categoryList})
}

func EditArticleIndex(c *gin.Context) {
	userName, _ := c.Get("username")

	tid := c.Query("id")
	id, err := strconv.Atoi(tid)
	if err != nil || id <= 0 {
		c.Redirect(http.StatusFound, "/admin/article/index")
		return
	}

	article, err := models.GetArticleById(id)
	if err != nil {
		log.Error(err.Error())
		c.Redirect(http.StatusFound, "/admin/article/index")
		return
	}

	categoryList, err := models.GetCategoryList(models.CategoryNORMAL)
	if err != nil {
		log.Error(err.Error())
	}

	article.Content = html.UnescapeString(article.Content)
	c.HTML(http.StatusOK, "admin/article/edit.html", gin.H{"title": "文章管理", "highlight": "article", "categorys": categoryList, "username": userName, "article": article})
	return

}

func GetArticleForPage(c *gin.Context) {
	var args struct {
		PageIndex int `form:"page_index" binding:"min=1"`
		PageSize  int `form:"page_size" binding:"min=10"`
	}

	if err := c.ShouldBindQuery(&args); err != nil {
		resp.GetResponse(c, e.INVALID_PARAMS, util.Translate(err.(validator.ValidationErrors)), nil)
		return
	}

	total, articles, err := models.GetArticleListForPage(args.PageIndex, args.PageSize, 0)
	if err != nil {
		log.Error(err.Error())
		resp.GetResponse(c, e.ERROR, "", nil)
		return
	}

	resp.GetResponse(c, e.SUCCESS, "", map[string]interface{}{"list": articles, "total": total})
	return
}

func AddArticle(c *gin.Context) {
	var args struct {
		CategoryId int    `form:"category" binding:"min=1"`
		Title      string `form:"title" binding:"min=10"`
		Content    string `form:"content" binding:"min=10"`
		Depiction  string `form:"depiction" binding:"min=10"`
		MyImage    string `form:"myimage" binding:"-"`
		Top        int8   `form:"top" binding:"min=1,max=2"`
	}

	if err := c.ShouldBind(&args); err != nil {
		resp.GetResponse(c, e.INVALID_PARAMS, util.Translate(err.(validator.ValidationErrors)), nil)
		return
	}

	if err := models.AddArticle(args.CategoryId, args.Title, html.EscapeString(args.Content), args.Top, args.Depiction, args.MyImage); err != nil {
		resp.GetResponse(c, e.ERROR, "", nil)
		return
	}
	resp.GetResponse(c, e.SUCCESS, "", nil)
	return
}

func EditArticle(c *gin.Context) {

	var args struct {
		ArticleId  int    `form:"articleId" binding:"min=1"`
		CategoryId int    `form:"category" binding:"min=1"`
		Title      string `form:"title" binding:"min=10"`
		Content    string `form:"content" binding:"min=10"`
		Depiction  string `form:"depiction" binding:"min=10"`
		MyImage    string `form:"myimage" binding:"-"`
		Top        int8   `form:"top" binding:"min=1,max=2"`
	}

	if err := c.ShouldBind(&args); err != nil {
		resp.GetResponse(c, e.INVALID_PARAMS, util.Translate(err.(validator.ValidationErrors)), nil)
		return
	}

	article, err := models.GetArticleById(args.ArticleId)
	if err != nil {
		if err == models.NoData {
			resp.GetResponse(c, e.NOARTICLE, "", nil)
			return
		}
		log.Error(err.Error())
		resp.GetResponse(c, e.ERROR, "", nil)
		return
	}

	article.CategoryId = args.CategoryId
	article.Title = args.Title
	article.Content = html.EscapeString(args.Content)
	article.Top = args.Top
	article.Depiction = args.Depiction
	article.Image = args.MyImage
	if err = models.EditArticle(article); err != nil {
		log.Error(err.Error())
		resp.GetResponse(c, e.ERROR, "", nil)
		return
	}

	resp.GetResponse(c, e.SUCCESS, "", nil)
	return
}

func DelArticle(c *gin.Context) {
	var args struct {
		ArticleId int `form:"id" binding:"min=1" `
	}

	if err := c.ShouldBind(&args); err != nil {
		resp.GetResponse(c, e.INVALID_PARAMS, util.Translate(err.(validator.ValidationErrors)), nil)
		return
	}

	exists, err := models.ExistsArticleById(args.ArticleId)
	if err != nil {
		log.Error(err.Error())
		resp.GetResponse(c, e.ERROR, "", nil)
		return
	}

	if !exists {
		resp.GetResponse(c, e.NOARTICLE, "", nil)
		return
	}

	if err := models.DeleteArticle(args.ArticleId); err != nil {
		resp.GetResponse(c, e.ERROR, "", nil)
		return
	}

	resp.GetResponse(c, e.SUCCESS, "", nil)
	return

}
