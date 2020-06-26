package admin

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/luckywb13/blog/pkg/e"
	"github.com/luckywb13/blog/pkg/log"
	"github.com/luckywb13/blog/pkg/resp"
	"net/http"
	"time"
)

func ImageUpload(c *gin.Context) {
	uFile, err := c.FormFile("imgFile")
	if err != nil {
		log.Error(err)
		resp.GetResponse(c, e.INVALID_PARAMS, "", nil)
		return
	}

	h := md5.New()
	h.Write([]byte(uFile.Filename))
	cipherStr := h.Sum(nil)
	localName := fmt.Sprintf("%s/%d%s", "images", time.Now().Unix(), hex.EncodeToString(cipherStr))
	if err = c.SaveUploadedFile(uFile, localName); err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusOK, map[string]interface{}{"error": 1008, "msg": "保存文件失败"})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"error": 0, "url": "/" + localName})
	return

}
