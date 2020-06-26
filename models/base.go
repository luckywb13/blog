package models

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/luckywb13/blog/conf"
	"time"
)

var NoData = errors.New("NoData")

type BaseModel struct {
	ID        int       `gorm:"primary_key" json:"ID"`
	CreatedAt time.Time `gorm:"column:created_at;default:null" json:"CreatedAt"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:null" json:"UpdatedAt"`
}

var dbObj *gorm.DB

func Init(c *conf.DB) (err error) {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local", c.User, c.Password, c.Host, c.Port, c.Name)

	if dbObj, err = gorm.Open("mysql", dns); err != nil {
		return
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return c.Prefix + defaultTableName
	}

	dbObj.SingularTable(true)
	dbObj.LogMode(conf.Debug)
	dbObj.DB().SetMaxIdleConns(c.MaxIdle)
	dbObj.DB().SetMaxOpenConns(c.MaxOpen)

	return
}

func Close() {
	_ = dbObj.Close()
}
