package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Num struct {
	BaseModel
	ArticleId  int `gorm:"primary_key" json:"ArticleId"`
	CommentNum int `gorm:"column:comment_num" json:"CommentNum"`
	ReadNum    int `gorm:"column:read_num" json:"ReadNum"`
}


func GetNumByArticleId(articleId int) (num Num, err error) {
	if err = dbObj.Model(Num{}).Where("article_id = ?", articleId).First(&num).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = NoData
		}
		return
	}
	return
}

func GetNumByArticleIds(articleIds []int) (nums []Num, err error) {
	if len(articleIds) <= 0 {
		return
	}

	err = dbObj.Model(Num{}).Where("article_id in (?)", articleIds).Find(&nums).Error
	return
}

func UpdateCommentNum(articleId int)(err error) {
	var num int
	if err = dbObj.Model(Num{}).Where("article_id = ?", articleId).Count(&num).Error; err != nil {
		return
	}

	if num > 0 {
		if err = dbObj.Model(Num{}).Where("article_id = ?", articleId).Update("comment_num", gorm.Expr("comment_num + ?", 1)).Error; err != nil {
			return
		}
		return
	}

	err = dbObj.Model(Num{}).Create(&Num{ArticleId:articleId, CommentNum:1, ReadNum:0}).Error
	return
}

func UpdateReadNum(articleId int) (err error) {
	var num int
	if err = dbObj.Model(Num{}).Where("article_id = ?", articleId).Count(&num).Error; err != nil {
		return
	}

	if num > 0 {
		if err = dbObj.Model(Num{}).Where("article_id = ?", articleId).Update("read_num", gorm.Expr("read_num + ?", 1)).Error; err != nil {
			return
		}
		return
	}

	err = dbObj.Model(Num{}).Create(&Num{ArticleId:articleId, ReadNum:1}).Error
	return
}