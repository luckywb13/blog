package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Comment struct {
	BaseModel
	ArticleId int       `gorm:"column:article_id" json:"ArticleId"`
	Email     string    `gorm:"column:email" json:"Email"`
	Name      string    `gorm:"column:name" json:"Name"`
	Content   string    `gorm:"column:content" json:"Content"`
	Ip        string    `gorm:"column:ip" json:"-"`
}


func GetCommentListForPage(pageIndex, pageNum, articleId int) (comments []Comment, err error) {
	if err = dbObj.Model(Comment{}).Where("article_id = ?", articleId).Order("created_at desc").Limit(pageNum).Offset((pageIndex - 1)*pageNum).Find(&comments).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = NoData
		}
	}
	return
}

func GetCommentTotalByArticle(articleId int) (count int, err error) {
	err = dbObj.Model(Comment{}).Where("article_id = ?", articleId).Count(&count).Error
	return
}

func AddComment(articleId int, email, name, content, ip string) (err error) {

	tx := dbObj.Begin()
	comment := Comment{ArticleId: articleId, Email: email, Name: name, Content: content, Ip: ip}
	if err = dbObj.Model(Comment{}).Create(&comment).Error; err != nil {
		tx.Rollback()
		return
	}
	if err = UpdateCommentNum(articleId); err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	return
}