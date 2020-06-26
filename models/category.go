package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Category struct {
	BaseModel
	Name   string `gorm:"name" json:"Name"`
	Status int8   `gorm:"status" json:"Status"`
}


type CST int8

const (
	_ CST = iota - 1
	CategoryAll
	CategoryNORMAL
	CategoryDELETE
)

func AddCategory(name string) error {
	category := Category{Name: name, Status: 1}
	return dbObj.Model(Category{}).Create(&category).Error
}

func GetCategoryById(id int) (category Category, err error) {
	if err = dbObj.Model(Category{}).Where("id = ?", id).First(&category).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = NoData
		}
		return
	}
	return
}

func DeleteCategory(id int) (err error) {
	if err = dbObj.Model(&Category{}).Where("id = ?", id).Update(map[string]int{"status": 2}).Error; err != nil {
		return
	}
	return
}

func EditCategory(id int, name string) (err error) {

	var categoryModel Category
	if err = dbObj.Model(Category{}).Where("id = ?", id).First(&categoryModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = NoData
		}
		return
	}

	categoryModel.Name = name
	err = dbObj.Model(&categoryModel).Update(&categoryModel).Error
	return
}


func GetCategoryList(status CST) (categoryList []Category, err error) {
	if status != CategoryAll {
		err = dbObj.Model(&Category{}).Where("status = ?", status).Order("created_at desc").Find(&categoryList).Error
	} else {
		err = dbObj.Model(&Category{}).Order("created_at desc").Find(&categoryList).Error
	}
	return

}
