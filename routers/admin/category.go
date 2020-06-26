package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/luckywb13/blog/models"
	"github.com/luckywb13/blog/pkg/e"
	"github.com/luckywb13/blog/pkg/log"
	"github.com/luckywb13/blog/pkg/resp"
	"github.com/luckywb13/blog/pkg/util"
	"net/http"
)

func CategoryIndex(c *gin.Context) {
	userName, _ := c.Get("username")
	categoryList, err := models.GetCategoryList(models.CategoryAll)
	if err != nil {
		log.Error(err.Error())
	}
	c.HTML(http.StatusOK, "admin/category/index.html", gin.H{"title": "分类管理", "highlight": "category", "username": userName, "data": categoryList})
}

func AddCategory(c *gin.Context) {
	var args struct {
		Name string `form:"name" binding:"min=5,max=50" label:"分类名称"`
	}

	if err := c.ShouldBind(&args); err != nil {
		resp.GetResponse(c, e.INVALID_PARAMS, util.Translate(err.(validator.ValidationErrors)), nil)
		return
	}

	err := models.AddCategory(args.Name)
	if err != nil {
		log.Error(err.Error())
		resp.GetResponse(c, e.ERROR, "", nil)
		return
	}
	resp.GetResponse(c, e.SUCCESS, "", nil)
	return
}

func EditCategory(c *gin.Context) {

	var args struct {
		Name string `form:"name" binding:"min=5,max=50" label:"分类名称"`
		Id   int    `form:"id" binding:"min=1"`
	}

	if err := c.ShouldBind(&args); err != nil {
		resp.GetResponse(c, e.INVALID_PARAMS, util.Translate(err.(validator.ValidationErrors)), nil)
		return
	}

	if err := models.EditCategory(args.Id, args.Name); err != nil {
		if err == models.NoData {
			resp.GetResponse(c, e.NOCATEGORY, "", nil)
			return
		}
		log.Error(err.Error())
		resp.GetResponse(c, e.ERROR, "", nil)
		return
	}

	resp.GetResponse(c, e.SUCCESS, "", nil)
	return
}

func DelCategory(c *gin.Context) {
	var args struct {
		Id int `form:"id" binding:"min=1"`
	}

	if err := c.ShouldBind(&args); err != nil {
		resp.GetResponse(c, e.INVALID_PARAMS, util.Translate(err.(validator.ValidationErrors)), nil)
		return
	}

	category, err := models.GetCategoryById(args.Id)
	if err != nil {
		if err != models.NoData {
			resp.GetResponse(c, e.NOCATEGORY, "", nil)
			return
		}
		log.Error(err.Error())
		resp.GetResponse(c, e.ERROR, "", nil)
		return
	}

	if category.Status == 2 {
		resp.GetResponse(c, e.NOCATEGORY, "", nil)
		return
	}

	count, err := models.GetArticleCountByCategoryId(args.Id)
	if err != nil {
		log.Error(err.Error())
		resp.GetResponse(c, e.SUCCESS, "", nil)
		return
	}

	if count > 0 {
		resp.GetResponse(c, e.EXISTSARTICLE, "", nil)
		return
	}

	err = models.DeleteCategory(args.Id)
	if err != nil {
		log.Error(err.Error())
		resp.GetResponse(c, e.ERROR, "", nil)
		return
	}

	resp.GetResponse(c, e.SUCCESS, "", nil)
	return
}
