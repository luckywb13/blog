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
	"math"
	"net/http"
	"strconv"
)

func Home(c *gin.Context) {

	var err error
	categoryList := make([]models.Category, 0)
	articleList := make([]models.Article, 0)
	topArticle := models.Article{}
	var num int
	var total float64
	if categoryList, err = models.GetCategoryList(models.CategoryNORMAL); err != nil {
		log.Error(err.Error())
		goto result
	}

	if num, articleList, err = models.GetArticleListForPage(1, 5, 0); err != nil {
		log.Error(err.Error())
		goto result
	}

	if topArticle, err = models.GetTopArticle(); err != nil {
		if err != models.NoData {
			log.Error(err.Error())
			goto result
		}
	}

	total = math.Ceil(float64(num) / 5.0)

result:
	c.HTML(http.StatusOK, "blog/home/index.html", gin.H{"title": conf.WebName, "highlight": "home", "categoryList": categoryList, "top": topArticle, "total": total, "articles": articleList})
}

func GetArticleDetail(c *gin.Context) {
	tmpId := c.Query("id")
	articleId, err := strconv.Atoi(tmpId)
	if err != nil {
		log.Error(err.Error())
		c.Redirect(http.StatusFound, "/home/index")
		return
	}

	article, err := models.GetArticleById(articleId)
	if err != nil {
		log.Error(err.Error())
		c.Redirect(http.StatusFound, "/home/index")
		return
	}

	article.Content = html.UnescapeString(article.Content)
	num, err := models.GetNumByArticleId(article.ID)
	if err != nil {
		log.Error(err.Error())
	} else {
		article.ReadNum = num.ReadNum
		article.CommentNum = num.CommentNum
	}
	_, articles, err := models.GetArticleListForPage(1, 5, 0)
	if err != nil {
		log.Error(err.Error())
	}

	tmp, err := models.GetCommentTotalByArticle(articleId)
	if err != nil {
		log.Error(err.Error())
	}

	total := math.Ceil(float64(tmp) / float64(5))
	twos, err := models.GetArticleByTwo(articleId)

	if err != nil {
		log.Error(err.Error())
	}

	if err := models.UpdateReadNum(articleId); err != nil {
		log.Error(err.Error())
	}

	c.HTML(http.StatusOK, "blog/article/detail.html", gin.H{"title": conf.WebName, "highlight": "home", "article": article, "articles": articles, "twos": twos, "total": total})
}

func GetArticleListForPage(c *gin.Context) {
	var args struct {
		PageIndex  int `form:"page_index" binding:"min=1"`
		PageSize   int `form:"page_size" binding:"min=5"`
		CategoryId int `form:"category_id" binding:"-"`
	}
	if err := c.ShouldBindQuery(&args); err != nil {
		resp.GetResponse(c, e.INVALID_PARAMS, util.Translate(err.(validator.ValidationErrors)), nil)
		return
	}

	_, articles, err := models.GetArticleListForPage(args.PageIndex, args.PageSize, args.CategoryId)
	if err != nil {
		log.Error(err.Error())
		resp.GetResponse(c, e.ERROR, "", nil)
		return
	}

	ids := make([]int, 0)
	for i := 0; i < len(articles); i++ {
		ids = append(ids, articles[i].ID)
	}

	nums, err := models.GetNumByArticleIds(ids)
	if err != nil {
		log.Error(err.Error())
		resp.GetResponse(c, e.ERROR, "", nil)
		return
	}

	type T struct {
		ReadNum    int
		CommentNum int
	}

	tmp1 := make(map[int]T)
	for i := 0; i < len(nums); i++ {
		tmp1[nums[i].ArticleId] = T{ReadNum: nums[i].ReadNum, CommentNum: nums[i].CommentNum}
	}

	for i := 0; i < len(articles); i++ {
		articles[i].ReadNum = tmp1[articles[i].ID].ReadNum
		articles[i].CommentNum = tmp1[articles[i].ID].CommentNum
	}

	resp.GetResponse(c, e.SUCCESS, "", articles)
	return
}
