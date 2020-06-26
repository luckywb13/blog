package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/luckywb13/blog/middleware"
	"github.com/luckywb13/blog/routers/admin"
	"github.com/luckywb13/blog/routers/blog"
	"html/template"
	"time"
)

func InitRouter() (engine *gin.Engine) {
	engine = gin.Default()

	engine.Use(middleware.LogHandler())
	engine.SetFuncMap(template.FuncMap{
		"dateFormat":  DateFormat,
		"htmlHandler": HtmlHandler,
	})

	engine.LoadHTMLGlob("./views/**/**/*")
	engine.Static("/static", "./static")
	engine.Static("/images", "./images")

	engine.GET("/", blog.Home)
	engine.GET("home/index", blog.Home)
	engine.GET("about/index", blog.AboutIndex)
	engine.GET("talk/index", blog.TalkIndex)
	engine.GET("talk/getList", blog.GetTalkListForPage)


	engine.GET("article/detail", blog.GetArticleDetail)
	engine.GET("article/getList", blog.GetArticleListForPage)

	engine.GET("comment/index", blog.CommentIndex)
	engine.POST("comment/add", blog.AddComment)
	engine.GET("comment/getList", blog.GetCommentListForPage)

	adminGroup := engine.Group("/admin")
	adminGroup.GET("/login", admin.LoginIndex)
	adminGroup.POST("/login", admin.Login)
	adminGroup.GET("/logout", admin.Logout)

	adminGroup.Use(middleware.AuthHandler())
	{
		adminGroup.GET("/category/index", admin.CategoryIndex)
		adminGroup.POST("/category/add", admin.AddCategory)
		adminGroup.POST("/category/edit", admin.EditCategory)
		adminGroup.POST("/category/delete", admin.DelCategory)

		adminGroup.GET("/article/index", admin.ArticleIndex)
		adminGroup.GET("/article/getList", admin.GetArticleForPage)
		adminGroup.GET("/article/add", admin.AddArticleIndex)
		adminGroup.GET("/article/edit", admin.EditArticleIndex)
		adminGroup.POST("/article/add", admin.AddArticle)
		adminGroup.POST("/article/delete", admin.DelArticle)
		adminGroup.POST("/article/edit", admin.EditArticle)

		adminGroup.POST("/file/upload", admin.ImageUpload)

		adminGroup.GET("/talk/index", admin.TalkIndex)
		adminGroup.GET("/talk/add", admin.AddTalkIndex)
		adminGroup.GET("/talk/edit", admin.EditTalkIndex)
		adminGroup.GET("/talk/getList", admin.GetTalkListForPage)
		adminGroup.POST("/talk/add", admin.AddTalk)
		adminGroup.POST("/talk/delete", admin.DelTalk)
		adminGroup.POST("/talk/edit", admin.EditTalk)

	}

	return
}

func DateFormat(date time.Time, layout string) string {
	return date.Format(layout)
}

func HtmlHandler(content string) template.HTML {
	return template.HTML(content)
}
