package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Talk struct {
	BaseModel
	Title   string `gorm:"column:title" json:"Title"`
	Content string `gorm:"column:content" json:"Content"`
	Icon    string `gorm:"Icon" json:"Icon"`
	Status  int8   `gorm:"status" json:"Status"`
}

func AddTalk(title string, content string, image string) (err error) {
	talk := Talk{Title: title, Content: content, Icon: image, Status: 1}
	err = dbObj.Model(Talk{}).Create(&talk).Error
	return
}

func EditTalk(talk Talk) (err error) {
	err = dbObj.Model(&talk).Updates(Talk{
		Title:   talk.Title,
		Content: talk.Content,
		Icon:    talk.Icon}).Error
	return
}

func DeleteTalk(id int) (err error) {
	err = dbObj.Model(&Talk{}).Where("id = ?", id).Update(map[string]int{"status": 2}).Error
	return
}

func GetTalkById(id int) (talk Talk, err error) {
	if err = dbObj.Model(Talk{}).First(&talk, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = NoData
		}
		return
	}
	return
}

func GetTalkListForPage(pageIndex, pageNum int) (total int,talks []Talk, err error) {

	talks = make([]Talk, 0)
	if err = dbObj.Model(Talk{}).Where("status=?", 1).Count(&total).Error; err != nil {
		return
	}

	if total <= 0 {
		return
	}

	err = dbObj.Model(Talk{}).Where("status=?", 1).Order("created_at desc", true).Limit(pageNum).Offset((pageIndex-1)*pageNum).Find(&talks).Error
	return
}

func GetTalkTotal() (count int, err error) {
	err = dbObj.Model(Talk{}).Where("status=?", 1).Count(&count).Error
	return
}


func ExistsTalkById(id int) (exists bool, err error) {
	var count int
	if err = dbObj.Model(&Talk{}).Where("id = ?", id).Count(&count).Error; err != nil {
		return
	}

	if count > 0 {
		exists = true
	}
	return
}