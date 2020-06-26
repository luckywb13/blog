package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/luckywb13/blog/models"
	"github.com/luckywb13/blog/pkg/e"
	"github.com/luckywb13/blog/pkg/log"
	"github.com/luckywb13/blog/pkg/resp"
	"github.com/luckywb13/blog/pkg/util"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func LoginIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login/index.html", gin.H{"title": "登录"})
}

func Login(c *gin.Context) {
	var args struct {
		UserName string `form:"username" binding:"email"`
		Password string `form:"password" binding:"required"`
	}

	if err := c.ShouldBind(&args); err != nil {
		resp.GetResponse(c, e.INVALID_PARAMS, util.Translate(err.(validator.ValidationErrors)), nil)
		return
	}

	admin, err := models.GetAdminByUserName(args.UserName)
	if err != nil {
		if err != models.NoData {
			resp.GetResponse(c, e.USERERROR, "", nil)
			return
		}
		log.Error(err.Error())
		resp.GetResponse(c, e.ERROR, "", nil)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(admin.PassWord), []byte(args.Password)); err != nil {
		resp.GetResponse(c, e.USERERROR, "", nil)
		return
	}

	var token string
	if token, err = util.GenerateToken(admin.Id, admin.UserName, admin.PassWord); err != nil {
		resp.GetResponse(c, e.USERERROR, "", nil)
		return
	}

	cookie := &http.Cookie{Name: "authorization", Value: token, Path: "/", HttpOnly: true}
	http.SetCookie(c.Writer, cookie)

	resp.GetResponse(c, e.SUCCESS, "", nil)
	return
}

func Logout(c *gin.Context) {
	c.SetCookie("authorization", "", -1, "/", "*.", false, true)
	c.Redirect(http.StatusFound, "/admin/login")

}
