package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Admin struct {
	Id       int    `gorm:"primary_key"`
	UserName string `gorm:"column:username"`
	PassWord string `gorm:"column:password"`
}

func GetAdminByUserName(userName string) (admin Admin, err error) {
	if err = dbObj.Where("username = ?", userName).First(&admin).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = NoData
		}
		return
	}
	return
}
