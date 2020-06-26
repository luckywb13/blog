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
	"strconv"
)

func TalkIndex(c *gin.Context) {
	userName, _ := c.Get("username")
	c.HTML(http.StatusOK, "admin/talk/index.html", gin.H{"title": "说说管理", "highlight": "talk", "username": userName})
}

func AddTalkIndex(c *gin.Context) {
	userName, _ := c.Get("username")
	c.HTML(http.StatusOK, "admin/talk/add.html", gin.H{"title": "说说管理", "username": userName, "highlight": "talk"})
}

func EditTalkIndex(c *gin.Context) {
	userName, _ := c.Get("username")
	tid := c.Query("id")
	id, err := strconv.Atoi(tid)
	if err != nil || id <= 0 {
		c.Redirect(http.StatusFound, "/admin/talk/index")
		return
	}

	talk, err := models.GetTalkById(id)
	if err != nil {
		log.Error(err.Error())
		c.Redirect(http.StatusFound, "/admin/talk/index")
		return
	}

	c.HTML(http.StatusOK, "admin/talk/edit.html", gin.H{"title": "说说管理", "highlight": "talk", "username": userName, "talk": talk})
	return

}

func AddTalk(c *gin.Context) {
	var args struct {
		Title   string `form:"title" binding:"min=5" label:"标题"`
		Content string `form:"content" binding:"min=10" label:"内容"`
		MyImage string `form:"myimage" binding:"-" label:"图片"`
	}

	if err := c.ShouldBind(&args); err != nil {
		resp.GetResponse(c, e.INVALID_PARAMS, util.Translate(err.(validator.ValidationErrors)), nil)
		return
	}

	err := models.AddTalk(args.Title, args.Content, args.MyImage)
	if err != nil {
		log.Error(err.Error())
		resp.GetResponse(c, e.ERROR, "", nil)
		return
	}

	resp.GetResponse(c, e.SUCCESS, "", nil)
	return
}

func EditTalk(c *gin.Context) {
	var args struct {
		TalkId  int    `form:"talk_id" binding:"min=1"`
		Title   string `form:"title" binding:"required"`
		Content string `form:"content" binding:"required"`
		MyImage string `form:"myimage" binding:"-"`
	}

	if err := c.ShouldBind(&args); err != nil {
		resp.GetResponse(c, e.INVALID_PARAMS, util.Translate(err.(validator.ValidationErrors)), nil)
		return
	}

	talk, err := models.GetTalkById(args.TalkId)
	if err != nil {
		if err == models.NoData {
			resp.GetResponse(c, e.NOTALK, "", nil)
			return
		}
		log.Error(err.Error())
		resp.GetResponse(c, e.ERROR, "", nil)
		return
	}

	talk.Title = args.Title
	talk.Content = args.Content
	talk.Icon = args.MyImage
	if err = models.EditTalk(talk); err != nil {
		log.Error(err.Error())
		resp.GetResponse(c, e.ERROR, "", nil)
		return
	}

	resp.GetResponse(c, e.SUCCESS, "", nil)
	return
}

func DelTalk(c *gin.Context) {
	var args struct {
		Id int `form:"id" binding:"min=1"`
	}

	if err := c.ShouldBind(&args); err != nil {
		resp.GetResponse(c, e.INVALID_PARAMS, util.Translate(err.(validator.ValidationErrors)), nil)
		return
	}

	exists, err := models.ExistsTalkById(args.Id)
	if err != nil {
		log.Error(err.Error())
		resp.GetResponse(c, e.ERROR, "", nil)
		return
	}

	if !exists {
		resp.GetResponse(c, e.NOTALK, "", nil)
		return
	}

	if err := models.DeleteTalk(args.Id); err != nil {
		log.Error(err.Error())
		resp.GetResponse(c, e.ERROR, "", nil)
		return
	}

	resp.GetResponse(c, e.SUCCESS, "", nil)
	return
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

	total, talkList, err := models.GetTalkListForPage(args.PageIndex, args.PageSize)
	if err != nil {
		log.Error(err.Error())
		resp.GetResponse(c, e.ERROR, "", nil)
		return
	}

	resp.GetResponse(c, e.SUCCESS, "", map[string]interface{}{"list": talkList, "total": total})
	return
}
